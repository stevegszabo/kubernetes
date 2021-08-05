#!/bin/bash

HASHI_HOST=hashi-demo.example.com
HASHI_CERT=example.com.crt

while true
do
	curl -v -HHost:$HASHI_HOST --resolve $HASHI_HOST:8443:127.0.0.1 --cacert $HASHI_CERT https://$HASHI_HOST:8443/
	date
	sleep 1
done
