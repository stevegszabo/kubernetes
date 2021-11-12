# gatekeeper

```
szabos@master:~$ kc get -o json MybankNamespace mybanknamespace | jq -r .status.violations
[
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "argo-rollouts"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "argocd"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "calico-apiserver"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "calico-system"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "default"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "gatekeeper"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "hashi-demo"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "istio-operator"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "istio-system"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "kube-node-lease"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "kube-public"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "kube-system"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "tigera-operator"
  },
  {
    "enforcementAction": "deny",
    "kind": "Namespace",
    "message": "Namespace resources must include labels: {\"application\", \"environment\"}",
    "name": "vault"
  }
]
```
