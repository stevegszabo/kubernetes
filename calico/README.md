# calico

* calico-node: Calico-node runs on every Kubernetes cluster node as a DaemonSet. It is responsible for enforcing network policy, setting up routes on the nodes, plus managing any virtual interfaces for IPIP, VXLAN, or WireGuard
* calico-typha: Typha is as a stateful proxy for the Kubernetes API server. It's used by every calico-node pod to query and watch Kubernetes resources without putting excessive load on the Kubernetes API server.  The Tigera Operator automatically scales the number of Typha instances as the cluster size grows
* calico-kube-controllers: Runs a variety of Calico specific controllers that automate synchronization of resources. For example, when a Kubernetes node is deleted, it tidies up any IP addresses or other Calico resources associated with the node

```
szabos@ubuntu:~/src/github/kubernetes/calico$ kc get ns -l zone --show-labels
NAME            STATUS   AGE    LABELS
kube-system     Active   537d   kubernetes.io/metadata.name=kube-system,zone=kube-system
nginx-ingress   Active   9h     kubernetes.io/metadata.name=nginx-ingress,zone=nginx-ingress
webapp          Active   9h     kubernetes.io/metadata.name=webapp,zone=webapp

szabos@ubuntu:~/src/github/kubernetes/calico$ kc get tigerastatus calico
NAME     AVAILABLE   PROGRESSING   DEGRADED   SINCE
calico   True        False         False      16m

szabos@ubuntu:~/src/github/kubernetes/calico$ kc get -o json clusterinformations default | jq -r .spec
{
  "calicoVersion": "v3.19.1",
  "clusterGUID": "777501f84e7a4fe9b8d55f49be5fc39e",
  "clusterType": "k8s,kdd,bgp,kubeadm",
  "datastoreReady": true
}

szabos@ubuntu:~/src/github/kubernetes/calico$ kc get -o json felixconfiguration default | jq -r .spec
{
  "bpfLogLevel": "",
  "logSeverityScreen": "Info",
  "reportingInterval": "0s",
  "vxlanEnabled": true,
  "vxlanMTU": 1450,
  "vxlanPort": 8472,
  "vxlanVNI": 1
}

szabos@ubuntu:~/src/github/kubernetes/calico$ kc get -o json ippools default-ipv4-ippool | jq -r .spec
{
  "blockSize": 26,
  "cidr": "10.244.0.0/16",
  "ipipMode": "Never",
  "natOutgoing": true,
  "nodeSelector": "all()",
  "vxlanMode": "Always"
}

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
