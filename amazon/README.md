# amazon

```
echo | openssl s_client -connect oidc.eks.ca-central-1.amazonaws.com:443 2>&1 | openssl x509 -fingerprint -noout
```
