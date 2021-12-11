# kubernetes playground

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
