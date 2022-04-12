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

kc -n $CONSUL_NAMESPACE get -o json secret r1-consul-bootstrap-acl-token | jq -r .data.token | base64 -d

kc -n $CONSUL_NAMESPACE port-forward --address 0.0.0.0 $CONSUL_POD 8500:8500

CONSUL_HTTP_ADDR=http://localhost:8500
CONSUL_HTTP_TOKEN=aaaaaaaaaaaaaaaaaaaa
CONSUL_HEADER="X-Consul-Token: $CONSUL_HTTP_TOKEN"

curl -v -k -H "$CONSUL_HEADER" -XPUT -d @config-terminating-gw-elastic-node-01.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -H "$CONSUL_HEADER" -XPUT -d @config-terminating-gw-elastic-node-02.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -H "$CONSUL_HEADER" -XPUT -d @config-terminating-gw-elastic-node-03.json $CONSUL_HTTP_ADDR/v1/catalog/register
curl -v -k -H "$CONSUL_HEADER" -XPUT -d @config-terminating-gw-elastic-node-04.json $CONSUL_HTTP_ADDR/v1/catalog/register

curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/services | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/service/elastic | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/service/ingress-gateway | jq -r .[].ServiceID

curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/nodes | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/node/elastic-01 | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/node/elastic-02 | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/node/elastic-03 | jq -r .
curl -s -k -H "$CONSUL_HEADER" $CONSUL_HTTP_ADDR/v1/catalog/node/elastic-04 | jq -r .

curl -v -k -H "$CONSUL_HEADER" -XPUT -d @config-delete-node.json $CONSUL_HTTP_ADDR/v1/catalog/deregister
```

```
consul members
consul operator raft list-peers
consul operator autopilot state
consul operator autopilot get-config

consul catalog nodes
consul catalog services

consul config list -kind proxy-defaults
consul config read -kind proxy-defaults -name global

consul config list -kind service-defaults
consul config read -kind service-defaults -name frontend

consul services deregister -id=r1-consul-ingress-gateway-7775f8456c-k2s2g
```

```
DEMO_NAMESPACE=consul-demo
DEMO_LABEL="service=frontend"
DEMO_PATH=.items[0].metadata.name
DEMO_POD=$(kc -n $DEMO_NAMESPACE get -o json po -l $DEMO_LABEL | jq -r $DEMO_PATH)

kc -n $DEMO_NAMESPACE exec -it -c frontend $DEMO_POD -- curl -v http://localhost:2000/
```
