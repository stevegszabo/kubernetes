# kubernetes playground

```
aws eks list-clusters | jq -r .
aws eks describe-cluster --name=eks-cluster-01 | jq -r .
aws eks list-nodegroups --cluster-name=eks-cluster-01 | jq -r .

aws eks update-kubeconfig --region ca-central-1 --name eks-cluster-01

NODE_GROUP=$(aws eks list-nodegroups --cluster-name=eks-cluster-01 | jq -r .nodegroups[0])
aws eks describe-nodegroup --cluster-name eks-cluster-01 --nodegroup-name $NODE_GROUP | jq -r .

eksctl get cluster --name=eks-cluster-01
eksctl get cluster --name=eks-cluster-01 -o json | jq -r .

eksctl get nodegroup --cluster=eks-cluster-01 --name=$NODE_GROUP
eksctl get nodegroup --cluster=eks-cluster-01 --name=$NODE_GROUP -o json | jq -r .
```

```
kubeadm init --apiserver-advertise-address=192.168.56.101 --pod-network-cidr=10.0.0.0/16
kubeadm init --apiserver-advertise-address=192.168.56.201 --pod-network-cidr=10.0.0.0/16

kubeadm certs check-expiration
kubeadm certs renew all

kubeadm config images list
kubeadm config images pull

kubeadm token list
kubeadm token create --print-join-command

kubeadm upgrade plan
kubeadm upgrade apply v1.22.1 --dry-run --yes
```

![Stacked control plane](images/kube-control-stacked.png)
![Control plane certificates](images/kube-certs.png)
