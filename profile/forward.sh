#!/bin/bash

set -o errexit

LOG_FILE=/tmp/kubectl.stdout.log

(cat <<EOF
argocd                     svc/argocd-server           8443:443
hashi-vault                svc/r1-vault                8200:8200
hashi-consul               svc/r1-consul-server        8500:8500
hashi-consul               svc/r1-consul-ui            8080:80
EOF
) | while read NAMESPACE SERVICE PORTS
do
nohup kubectl -n $NAMESPACE port-forward --address 0.0.0.0 $SERVICE $PORTS >> $LOG_FILE 2>&1 &
done

exit 0
