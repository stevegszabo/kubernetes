# calico

```
szabos@ubuntu:~/src/github/kubernetes/calico$ kc get ns -l zone --show-labels
NAME            STATUS   AGE    LABELS
kube-system     Active   537d   kubernetes.io/metadata.name=kube-system,zone=kube-system
nginx-ingress   Active   9h     kubernetes.io/metadata.name=nginx-ingress,zone=nginx-ingress
webapp          Active   9h     kubernetes.io/metadata.name=webapp,zone=webapp
```
