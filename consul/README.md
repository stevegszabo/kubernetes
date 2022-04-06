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

curl -v -k -XPUT -d @config-terminating-gw-datadog-node-01.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -XPUT -d @config-terminating-gw-datadog-node-02.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -XPUT -d @config-terminating-gw-datadog-node-03.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -XPUT -d @config-terminating-gw-datadog-node-04.json $CONSUL_HTTP_ADDR/v1/catalog/register

curl -s -k $CONSUL_HTTP_ADDR/v1/catalog/service/datadog | jq -r .

curl -v -k -XPUT -d @config-delete.json $CONSUL_HTTP_ADDR/v1/catalog/deregister
```

```
DEMO_NAMESPACE=consul-demo
DEMO_LABEL="service=frontend"
DEMO_PATH=.items[0].metadata.name
DEMO_POD=$(kc -n $DEMO_NAMESPACE get -o json po -l $DEMO_LABEL | jq -r $DEMO_PATH)

kc -n $DEMO_NAMESPACE exec -it -c frontend $DEMO_POD -- curl -v http://localhost:2000/
```
