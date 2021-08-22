# kubernetes

```
kubeadm init --pod-network-cidr=192.168.0.0/16
kubeadm certs check-expiration
kubeadm certs renew all

kubeadm config images list
kubeadm config images pull

kubeadm upgrade plan
kubeadm upgrade apply v1.22.1 --dry-run --yes
```

![External control plane](images/kube-control-external.png)
![Stacked control plane](images/kube-control-stacked.png)
![Control plane certificates](images/kube-certs.png)
