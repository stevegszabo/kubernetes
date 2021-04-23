![Istio Kiali](images/istio.kiali.png)

```
kubectl -n istio-system port-forward --address 0.0.0.0 svc/kiali 20001:20001
kubectl -n istio-system port-forward --address 0.0.0.0 svc/istio-ingressgateway 8443:443

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 \
-subj '/O=example Inc./CN=example.com' -keyout example.com.key -out example.com.crt

openssl req -out httpbin.example.com.csr -newkey rsa:2048 -nodes \
-keyout httpbin.example.com.key -subj "/CN=httpbin.example.com/O=httpbin organization"

openssl x509 -req -days 365 -CA example.com.crt -CAkey example.com.key \
-set_serial 0 -in httpbin.example.com.csr -out httpbin.example.com.crt

kubectl -n istio-system create secret tls httpbin-credential \
--key=httpbin.example.com.key --cert=httpbin.example.com.crt

curl -s -HHost:httpbin.example.com \
--resolve "httpbin.example.com:8443:127.0.0.1" \
--cacert example.com.crt "https://httpbin.example.com:8443/webapp"
```
