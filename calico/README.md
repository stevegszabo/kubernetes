# calico

```
szabos@ubuntu:~/src/github/kubernetes/calico$ kc get ns -l zone --show-labels
NAME            STATUS   AGE    LABELS
kube-system     Active   537d   kubernetes.io/metadata.name=kube-system,zone=kube-system
nginx-ingress   Active   9h     kubernetes.io/metadata.name=nginx-ingress,zone=nginx-ingress
webapp          Active   9h     kubernetes.io/metadata.name=webapp,zone=webapp

szabos@ubuntu:~/src/github/kubernetes/calico$ calicoctl get ippool default-ipv4-ippool -o wide
NAME                  CIDR            NAT    IPIPMODE   VXLANMODE   DISABLED   SELECTOR
default-ipv4-ippool   10.244.0.0/16   true   Never      Always      false      all()

szabos@ubuntu:~/src/github/kubernetes/calico$ calicoctl ipam show
+----------+---------------+-----------+------------+--------------+
| GROUPING |     CIDR      | IPS TOTAL | IPS IN USE |   IPS FREE   |
+----------+---------------+-----------+------------+--------------+
| IP Pool  | 10.244.0.0/16 |     65536 | 52 (0%)    | 65484 (100%) |
+----------+---------------+-----------+------------+--------------+

szabos@ubuntu:~/src/github/kubernetes/calico$ calicoctl ipam show --show-blocks
+----------+-----------------+-----------+------------+--------------+
| GROUPING |      CIDR       | IPS TOTAL | IPS IN USE |   IPS FREE   |
+----------+-----------------+-----------+------------+--------------+
| IP Pool  | 10.244.0.0/16   |     65536 | 52 (0%)    | 65484 (100%) |
| Block    | 10.244.0.0/26   |        64 | 42 (66%)   | 22 (34%)     |
| Block    | 10.244.0.128/26 |        64 | 0 (0%)     | 64 (100%)    |
| Block    | 10.244.0.192/26 |        64 | 0 (0%)     | 64 (100%)    |
| Block    | 10.244.0.64/26  |        64 | 0 (0%)     | 64 (100%)    |
| Block    | 10.244.1.0/26   |        64 | 10 (16%)   | 54 (84%)     |
| Block    | 10.244.1.128/26 |        64 | 0 (0%)     | 64 (100%)    |
| Block    | 10.244.1.192/26 |        64 | 0 (0%)     | 64 (100%)    |
| Block    | 10.244.1.64/26  |        64 | 0 (0%)     | 64 (100%)    |
+----------+-----------------+-----------+------------+--------------+
```
