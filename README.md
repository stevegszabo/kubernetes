# kubernetes playground

```
aws eks list-clusters | jq -r .
aws eks list-nodegroups --cluster-name=eng-cluster-01 | jq -r .

eksctl get cluster --name=eng-cluster-01
eksctl get cluster --name=eng-cluster-01 -o json | jq -r .

eksctl get nodegroup --cluster=eng-cluster-01 --name=eng-group-01
eksctl get nodegroup --cluster=eng-cluster-01 --name=eng-group-01 -o json | jq -r .

eksctl create cluster --dry-run -f eng-cluster-01.yaml
eksctl delete cluster --region=ca-central-1 --name=eng-cluster-01

eksctl utils describe-stacks --region=ca-central-1 --cluster=eng-cluster-01
```

```
kubeadm init --apiserver-advertise-address=192.168.56.101 --pod-network-cidr=10.0.0.0/16
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
