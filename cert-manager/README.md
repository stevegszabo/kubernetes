# cert-manager

```
$ kc get clusterissuer
NAME          READY   AGE
eng-cluster   True    10m

$ kc get issuer
NAME          READY   AGE
webapp-demo   True    9m50s

$ kc get cert
NAME             READY   SECRET           AGE
webapp-demo      True    webapp-demo      9m52s
webapp-demo-ca   True    webapp-demo-ca   10m

$ kc get crs
NAME                   APPROVED   DENIED   READY   ISSUER        REQUESTOR                                         AGE
webapp-demo-ca-rvxk5   True                True    eng-cluster   system:serviceaccount:cert-manager:cert-manager   10m
webapp-demo-ngcqx      True                True    webapp-demo   system:serviceaccount:cert-manager:cert-manager   9m57s
```

```
CERT_NAMESPACE=app-webapp-demo
CERT_SECRET=webapp-demo

kc -n $CERT_NAMESPACE get -o json secret $CERT_SECRET | \
jq -r '.data["ca.crt"]' | \
base64 -d | \
openssl x509 -text -in -

kc -n $CERT_NAMESPACE get -o json secret $CERT_SECRET | \
jq -r '.data["tls.crt"]' | \
base64 -d | \
openssl x509 -text -in -
```

```
        Issuer: CN = webapp-demo-ca
        Validity
            Not Before: Feb  4 00:02:09 2022 GMT
            Not After : May  5 00:02:09 2022 GMT
        Subject: CN = www.webapp.com
        Subject Public Key Info:
            Public Key Algorithm: id-ecPublicKey
                Public-Key: (256 bit)
                pub:
                    04:9f:fd:51:e5:e3:e6:6c:45:20:b3:5c:e0:d0:92:
                    e1:9d:90:1b:c2:d6:3f:96:1f:2a:c3:a7:66:a6:e8:
                    7f:96:24:20:96:de:68:5c:be:a8:0e:cd:90:04:db:
                    dd:24:5a:38:ea:ac:1a:72:05:4a:34:15:1d:72:d5:
                    03:91:63:69:c2
                ASN1 OID: prime256v1
                NIST CURVE: P-256
        X509v3 extensions:
            X509v3 Key Usage: critical
                Digital Signature, Key Encipherment
            X509v3 Basic Constraints: critical
                CA:FALSE
            X509v3 Authority Key Identifier:
                keyid:C9:6C:F8:B1:98:52:AC:B8:53:EC:BA:8F:B8:C2:26:51:69:6B:BC:26

            X509v3 Subject Alternative Name:
                DNS:www.webapp.ca, DNS:www.webapp.net
```
