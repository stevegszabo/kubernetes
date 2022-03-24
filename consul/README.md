# consul

```
INGRESS_NAMESPACE=consul
INGRESS_LABEL="component=ingress-gateway"
INGRESS_PATH=.items[].spec.clusterIP
INGRESS_GATEWAY=$(kc -n $INGRESS_NAMESPACE get -o json svc -l $INGRESS_LABEL | jq -r $INGRESS_PATH)

curl -v -H "Host: frontend.ingress.consul" http://$INGRESS_GATEWAY/
```
