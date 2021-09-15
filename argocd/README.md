# argocd

```
kc -n argocd get secret -o json argocd-initial-admin-secret | jq -r .data.password | base64 -d
kc -n argocd port-forward svc/argocd-server 8080:443

argocd login --insecure --username admin localhost:8080
argocd account update-password
argocd admin dashboard --port 9090
```
