# calico

```
szabos@ubuntu:~/src/github/kubernetes/calico$ kc -n calico-system get po
NAME                                       READY   STATUS    RESTARTS   AGE
calico-kube-controllers-7f58dbcbbd-p4jcp   1/1     Running   0          32m
calico-node-bspjz                          1/1     Running   0          32m
calico-typha-5fdf48d8c8-xqt2d              1/1     Running   0          32m
```

* calico-node: Calico-node runs on every Kubernetes cluster node as a DaemonSet. It is responsible for enforcing network policy, setting up routes on the nodes, plus managing any virtual interfaces for IPIP, VXLAN, or WireGuard
* calico-typha: Typha is as a stateful proxy for the Kubernetes API server. It's used by every calico-node pod to query and watch Kubernetes resources without putting excessive load on the Kubernetes API server.  The Tigera Operator automatically scales the number of Typha instances as the cluster size grows
* calico-kube-controllers: Runs a variety of Calico specific controllers that automate synchronization of resources. For example, when a Kubernetes node is deleted, it tidies up any IP addresses or other Calico resources associated with the node

```
szabos@calico:~$ kc get tigerastatus
NAME                            AVAILABLE   PROGRESSING   DEGRADED   SINCE
apiserver                       True        False         False      9h
calico                          True        False         False      9h
compliance                      True        False         False      5d7h
intrusion-detection             True        False         False      2d22h
log-collector                   True        False         False      9h
management-cluster-connection   True        False         False      5d7h
monitor                         True        False         False      5d7h

szabos@calico:~$ kc get -o json clusterinformations default | jq -r .spec
{
  "calicoVersion": "v3.21.0",
  "clusterGUID": "c74d861cec9e4726a7f0a742af02384b",
  "clusterType": "typha,kdd,k8s,operator,bgp,kubeadm",
  "cnxVersion": "v3.11.1",
  "datastoreReady": true
}

szabos@calico:~$ kc get -o json felixconfiguration default | jq -r .spec
{
  "logSeverityScreen": "Info",
  "reportingInterval": "0s",
  "tproxyMode": "Disabled",
  "vxlanEnabled": true
}

szabos@calico:~$ kc get -o json ippools default-ipv4-ippool | jq -r .spec
{
  "allowedUses": [
    "Workload",
    "Tunnel"
  ],
  "blockSize": 26,
  "cidr": "10.0.0.0/16",
  "ipipMode": "Never",
  "natOutgoing": true,
  "nodeSelector": "all()",
  "vxlanMode": "CrossSubnet"
}

szabos@calico:~$ kc get globalreporttype
NAME             CREATED AT
cis-benchmark    2022-01-08T16:32:34Z
inventory        2022-01-08T16:32:33Z
network-access   2022-01-08T16:32:33Z
policy-audit     2022-01-08T16:32:33Z

szabos@calico:~$ calicoctl get ippool default-ipv4-ippool -o wide
NAME                  CIDR          NAT    IPIPMODE   VXLANMODE     DISABLED   DISABLEBGPEXPORT   SELECTOR
default-ipv4-ippool   10.0.0.0/16   true   Never      CrossSubnet   false      false              all()

szabos@calico:~$ calicoctl ipam show
+----------+-------------+-----------+------------+--------------+
| GROUPING |    CIDR     | IPS TOTAL | IPS IN USE |   IPS FREE   |
+----------+-------------+-----------+------------+--------------+
| IP Pool  | 10.0.0.0/16 |     65536 | 28 (0%)    | 65508 (100%) |
+----------+-------------+-----------+------------+--------------+

szabos@calico:~$ calicoctl ipam show --show-blocks
+----------+----------------+-----------+------------+--------------+
| GROUPING |      CIDR      | IPS TOTAL | IPS IN USE |   IPS FREE   |
+----------+----------------+-----------+------------+--------------+
| IP Pool  | 10.0.0.0/16    |     65536 | 28 (0%)    | 65508 (100%) |
| Block    | 10.0.15.192/26 |        64 | 28 (44%)   | 36 (56%)     |
+----------+----------------+-----------+------------+--------------+
```
