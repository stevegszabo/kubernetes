# cert-manager

```
CERT_NAMESPACE=app-webapp-demo
CERT_SECRET=webapp-demo

kc -n $CERT_NAMESPACE get -o json secret $CERT_SECRET | \
jq -r '.data["tls.crt"]' | \
base64 -d | \
openssl x509 -text -in -
```
