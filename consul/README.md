# consul

```
INGRESS_GATEWAY=$(kc -n consul get -o json svc -l component=ingress-gateway | jq -r .items[].spec.clusterIP)

curl -v -H "Host: frontend.ingress.consul" http://$INGRESS_GATEWAY/
```
