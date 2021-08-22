# kubernetes

```
kubeadm init --pod-network-cidr=192.168.0.0/16
kubeadm certs check-expiration

kubeadm config images list
kubeadm config images pull

kubeadm upgrade plan
```

![External control plane](images/kube-control-external.png)
![Stacked control plane](images/kube-control-stacked.png)
![Control plane certificates](images/kube-certs.png)
