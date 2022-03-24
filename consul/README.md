# consul

```
INGRESS_NAMESPACE=consul
INGRESS_LABEL="component=ingress-gateway"
INGRESS_PATH=.items[].spec.clusterIP
INGRESS_GATEWAY=$(kc -n $INGRESS_NAMESPACE get -o json svc -l $INGRESS_LABEL | jq -r $INGRESS_PATH)

curl -v -H "Host: frontend.ingress.consul" http://$INGRESS_GATEWAY/
```

```
CONSUL_NAMESPACE=consul
CONSUL_LABEL="component=server"
CONSUL_PATH=.items[0].metadata.name
CONSUL_POD=$(kc -n $CONSUL_NAMESPACE get po -o json -l $CONSUL_LABEL | jq -r $CONSUL_PATH)

kc -n $CONSUL_NAMESPACE port-forward --address 0.0.0.0 $CONSUL_POD 8500:8500

CONSUL_HTTP_ADDR=http://localhost:8500
export CONSUL_HTTP_ADDR

curl --request PUT --data @config-external.json --insecure $CONSUL_HTTP_ADDR/v1/catalog/register

curl -v -H "Host: ifconfig.me" http://localhost:1234/
```
