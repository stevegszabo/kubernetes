# argocd

```
kc -n argocd get secret -o json argocd-initial-admin-secret | jq -r .data.password | base64 -d
kc -n argocd port-forward --address 0.0.0.0 svc/argocd-server 8080:443
kc config get-contexts -o name

argocd login --insecure --username admin localhost:8080
argocd account update-password

argocd cluster list
argocd cluster add kubernetes-admin@kubernetes

argocd proj create argo-demo \
--description "Argo demo project" \
--src https://github.com/argoproj/argocd-example-apps.git \
--dest https://kubernetes.default.svc,argo-demo \
--allow-cluster-resource "*/*"

argocd proj list
argocd proj get argo-demo
argocd proj delete argo-demo

argocd app create guestbook \
--repo https://github.com/argoproj/argocd-example-apps.git \
--path guestbook \
--dest-server https://kubernetes.default.svc \
--dest-namespace argo-demo \
--project argo-demo

argocd app list
argocd app get guestbook
argocd app sync guestbook
argocd app delete -y guestbook
```

```
szabos@ubuntu:~$ argocd app get guestbook
Name:               guestbook
Project:            argo-demo
Server:             https://kubernetes.default.svc
Namespace:          argo-demo
URL:                https://localhost:8080/applications/guestbook
Repo:               https://github.com/argoproj/argocd-example-apps.git
Target:
Path:               guestbook
SyncWindow:         Sync Allowed
Sync Policy:        <none>
Sync Status:        Synced to  (53e28ff)
Health Status:      Healthy

GROUP  KIND        NAMESPACE  NAME          STATUS  HEALTH   HOOK  MESSAGE
       Service     argo-demo  guestbook-ui  Synced  Healthy        service/guestbook-ui unchanged
apps   Deployment  argo-demo  guestbook-ui  Synced  Healthy        deployment.apps/guestbook-ui unchanged
```

![ArgoCD](images/argo-demo.png)
