```
kubectl -n istio-system port-forward --address 0.0.0.0 svc/kiali 20001:20001
kubectl -n istio-system port-forward --address 0.0.0.0 svc/istio-ingressgateway 8443:443

curl -s -HHost:httpbin.example.com \
--resolve "httpbin.example.com:8443:127.0.0.1" \
--cacert example.com.crt "https://httpbin.example.com:8443/webapp"
```
