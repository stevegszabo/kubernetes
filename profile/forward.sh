#!/bin/bash

set -o errexit

LOG_FILE=/tmp/kubectl.stdout.log

(cat <<EOF
istio-system svc/prometheus           9090:9090
istio-system svc/grafana              3000:3000
istio-system svc/kiali                20001:20001
istio-system svc/istio-ingressgateway 8443:443
vault        svc/vault                8200:8200
argocd       svc/argocd-server        8080:443
EOF
) | while read NAMESPACE SERVICE PORTS
do
nohup kubectl -n $NAMESPACE port-forward --address 0.0.0.0 $SERVICE $PORTS >> $LOG_FILE 2>&1 &
done

exit 0
