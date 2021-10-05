package rollout

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	patchtypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/pointer"

	"github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	"github.com/argoproj/argo-rollouts/utils/annotations"
	"github.com/argoproj/argo-rollouts/utils/aws"
	"github.com/argoproj/argo-rollouts/utils/conditions"
	"github.com/argoproj/argo-rollouts/utils/defaults"
	logutil "github.com/argoproj/argo-rollouts/utils/log"
	"github.com/argoproj/argo-rollouts/utils/record"
	replicasetutil "github.com/argoproj/argo-rollouts/utils/replicaset"
	serviceutil "github.com/argoproj/argo-rollouts/utils/service"
)

const (
	switchSelectorPatch = `{
	"spec": {
		"selector": {
			"` + v1alpha1.DefaultRolloutUniqueLabelKey + `": "%s"
		}
	}
}`
	switchSelectorAndAddManagedByPatch = `{
	"metadata": {
		"annotations": {
			"` + v1alpha1.ManagedByRolloutsKey + `": "%s"
		}
	},
	"spec": {
		"selector": {
			"` + v1alpha1.DefaultRolloutUniqueLabelKey + `": "%s"
		}
	}
}`
)

func generatePatch(service *corev1.Service, newRolloutUniqueLabelValue string, r *v1alpha1.Rollout) string {
	if _, ok := service.Annotations[v1alpha1.ManagedByRolloutsKey]; !ok {
		return fmt.Sprintf(switchSelectorAndAddManagedByPatch, r.Name, newRolloutUniqueLabelValue)
	}
	return fmt.Sprintf(switchSelectorPatch, newRolloutUniqueLabelValue)
}

// switchSelector switch the selector on an existing service to a new value
func (c rolloutContext) switchServiceSelector(service *corev1.Service, newRolloutUniqueLabelValue string, r *v1alpha1.Rollout) error {
	ctx := context.TODO()
	if service.Spec.Selector == nil {
		service.Spec.Selector = make(map[string]string)
	}
	_, hasManagedRollout := serviceutil.HasManagedByAnnotation(service)
	oldPodHash, ok := service.Spec.Selector[v1alpha1.DefaultRolloutUniqueLabelKey]
	if ok && oldPodHash == newRolloutUniqueLabelValue && hasManagedRollout {
		return nil
	}
	patch := generatePatch(service, newRolloutUniqueLabelValue, r)
	_, err := c.kubeclientset.CoreV1().Services(service.Namespace).Patch(ctx, service.Name, patchtypes.StrategicMergePatchType, []byte(patch), metav1.PatchOptions{})
	if err != nil {
		return err
	}
	msg := fmt.Sprintf("Switched selector for service '%s' from '%s' to '%s'", service.Name, oldPodHash, newRolloutUniqueLabelValue)
	c.recorder.Eventf(r, record.EventOptions{EventReason: "SwitchService"}, msg)
	service.Spec.Selector[v1alpha1.DefaultRolloutUniqueLabelKey] = newRolloutUniqueLabelValue
	return err
}

func (c *rolloutContext) reconcilePreviewService(previewSvc *corev1.Service) error {
	if previewSvc == nil {
		return nil
	}
	newPodHash := c.newRS.Labels[v1alpha1.DefaultRolloutUniqueLabelKey]
	err := c.switchServiceSelector(previewSvc, newPodHash, c.rollout)
	if err != nil {
		return err
	}

	return nil
}

func (c *rolloutContext) reconcileActiveService(activeSvc *corev1.Service) error {
	if !replicasetutil.ReadyForPause(c.rollout, c.newRS, c.allRSs) || !annotations.IsSaturated(c.rollout, c.newRS) {
		c.log.Infof("skipping active service switch: New RS '%s' is not fully saturated", c.newRS.Name)
		return nil
	}

	newPodHash := activeSvc.Spec.Selector[v1alpha1.DefaultRolloutUniqueLabelKey]
	if c.isBlueGreenFastTracked(activeSvc) {
		newPodHash = c.newRS.Labels[v1alpha1.DefaultRolloutUniqueLabelKey]
	}
	if c.pauseContext.CompletedBlueGreenPause() && c.completedPrePromotionAnalysis() {
		newPodHash = c.newRS.Labels[v1alpha1.DefaultRolloutUniqueLabelKey]
	}

	if c.rollout.Status.Abort {
		newPodHash = c.rollout.Status.StableRS
	}

	err := c.switchServiceSelector(activeSvc, newPodHash, c.rollout)
	if err != nil {
		return err
	}
	return nil
}

// areTargetsVerified is a convenience to determine if the pod targets have been verified with
// underlying load balancer. If check was not performed or unnecessary, returns true.
func (c *rolloutContext) areTargetsVerified() bool {
	return c.targetsVerified == nil || *c.targetsVerified
}

// awsVerifyTargetGroups examines a Service and verifies that the underlying AWS TargetGroup has all
// of the Service's Endpoint IPs and ports registered. Only valid for services which are reachable
// by an ALB Ingress, which can be determined if there exists a TargetGroupBinding object in the
// namespace that references the given service
func (c *rolloutContext) awsVerifyTargetGroups(svc *corev1.Service) error {
	if !c.shouldVerifyTargetGroup(svc) {
		return nil
	}
	logCtx := c.log.WithField(logutil.ServiceKey, svc.Name)
	logCtx.Infof("Verifying target group")

	ctx := context.TODO()
	// find all TargetGroupBindings in the namespace which reference the service name + port
	tgBindings, err := aws.GetTargetGroupBindingsByService(ctx, c.dynamicclientset, *svc)
	if err != nil {
		return err
	}
	if len(tgBindings) == 0 {
		// no TargetGroupBinding for the service found (e.g. it is in-cluster blue-green). nothing to verify
		return nil
	}

	c.targetsVerified = pointer.BoolPtr(false)

	// get endpoints of service
	endpoints, err := c.kubeclientset.CoreV1().Endpoints(svc.Namespace).Get(ctx, svc.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	awsClnt, err := aws.NewClient()
	if err != nil {
		return err
	}

	for _, tgb := range tgBindings {
		verifyRes, err := aws.VerifyTargetGroupBinding(ctx, c.log, awsClnt, tgb, endpoints, svc)
		if err != nil {
			c.recorder.Warnf(c.rollout, record.EventOptions{EventReason: conditions.TargetGroupVerifyErrorReason}, conditions.TargetGroupVerifyErrorMessage, svc.Name, tgb.Spec.TargetGroupARN, err)
			return err
		}
		if verifyRes == nil {
			// verification not applicable
			continue
		}
		if !verifyRes.Verified {
			c.recorder.Warnf(c.rollout, record.EventOptions{EventReason: conditions.TargetGroupUnverifiedReason}, conditions.TargetGroupUnverifiedRegistrationMessage, svc.Name, tgb.Spec.TargetGroupARN, verifyRes.EndpointsRegistered, verifyRes.EndpointsTotal)
			c.enqueueRolloutAfter(c.rollout, defaults.GetRolloutVerifyRetryInterval())
			return nil
		}
		c.recorder.Eventf(c.rollout, record.EventOptions{EventReason: conditions.TargetGroupVerifiedReason}, conditions.TargetGroupVerifiedRegistrationMessage, svc.Name, tgb.Spec.TargetGroupARN, verifyRes.EndpointsRegistered)
	}
	c.targetsVerified = pointer.BoolPtr(true)
	return nil
}

// shouldVerifyTargetGroup returns whether or not we should verify the target group
func (c *rolloutContext) shouldVerifyTargetGroup(svc *corev1.Service) bool {
	if !defaults.VerifyTargetGroup() {
		// feature is disabled
		return false
	}
	desiredPodHash := c.newRS.Labels[v1alpha1.DefaultRolloutUniqueLabelKey]
	if c.rollout.Spec.Strategy.BlueGreen != nil {
		if c.rollout.Status.StableRS == desiredPodHash {
			// for blue-green, we only verify targets right after switching active service. So if
			// we are fully promoted, then there is no need to verify targets.
			// NOTE: this is the opposite of canary, where we only verify targets if stable == desired
			return false
		}
		svcPodHash := svc.Spec.Selector[v1alpha1.DefaultRolloutUniqueLabelKey]
		if svcPodHash != desiredPodHash {
			// we have not yet switched service selector
			return false
		}
		if c.rollout.Status.BlueGreen.PostPromotionAnalysisRunStatus != nil {
			// we already started post-promotion analysis, so verification already occurred
			return false
		}
		return true
	} else if c.rollout.Spec.Strategy.Canary != nil {
		if c.rollout.Spec.Strategy.Canary.TrafficRouting == nil || c.rollout.Spec.Strategy.Canary.TrafficRouting.ALB == nil {
			// not ALB canary, so no need to verify targets
			return false
		}
		if c.rollout.Status.StableRS != desiredPodHash {
			// for canary, we only verify targets right after switching stable service, which happens
			// after the update. So if stable != desired, we are still in the middle of an update
			// and there is no need to verify targets.
			// NOTE: this is the opposite of blue-green, where we only verify targets if stable != active
			return false
		}
		return true
	}
	// should not get here
	return false
}

func (c *rolloutContext) getPreviewAndActiveServices() (*corev1.Service, *corev1.Service, error) {
	var previewSvc *corev1.Service
	var activeSvc *corev1.Service
	var err error

	if c.rollout.Spec.Strategy.BlueGreen.PreviewService != "" {
		previewSvc, err = c.servicesLister.Services(c.rollout.Namespace).Get(c.rollout.Spec.Strategy.BlueGreen.PreviewService)
		if err != nil {
			return nil, nil, err
		}
	}
	activeSvc, err = c.servicesLister.Services(c.rollout.Namespace).Get(c.rollout.Spec.Strategy.BlueGreen.ActiveService)
	if err != nil {
		return nil, nil, err
	}
	return previewSvc, activeSvc, nil
}

func (c *rolloutContext) reconcileStableAndCanaryService() error {
	if c.rollout.Spec.Strategy.Canary == nil {
		return nil
	}
	err := c.ensureSVCTargets(c.rollout.Spec.Strategy.Canary.StableService, c.stableRS)
	if err != nil {
		return err
	}

	if replicasetutil.IsReplicaSetReady(c.newRS) {
		err = c.ensureSVCTargets(c.rollout.Spec.Strategy.Canary.CanaryService, c.newRS)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *rolloutContext) ensureSVCTargets(svcName string, rs *appsv1.ReplicaSet) error {
	if rs == nil || svcName == "" {
		return nil
	}
	svc, err := c.servicesLister.Services(c.rollout.Namespace).Get(svcName)
	if err != nil {
		return err
	}
	if svc.Spec.Selector[v1alpha1.DefaultRolloutUniqueLabelKey] != rs.Labels[v1alpha1.DefaultRolloutUniqueLabelKey] {
		err = c.switchServiceSelector(svc, rs.Labels[v1alpha1.DefaultRolloutUniqueLabelKey], c.rollout)
		if err != nil {
			return err
		}
	}
	return nil
}