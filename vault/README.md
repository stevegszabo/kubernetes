# vault

![VaultAgent](images/vault-agent.png)

```
Unseal Key 1: BVBM0zdWs2MUX8+Vw+DPFmmcYU5PBtD14DO2CndNfTXQ
Unseal Key 2: UBO28nH2KZrZriWRmZUIhraWnaMKbYKzMEOjnM2Swtwq
Unseal Key 3: /rAigPxz+DkXS6xKbmJOhQzQBbCemKBczeS6iPos9Qz5
Unseal Key 4: HJxoTs5qtErZbwCOIKFVesxFcwx1Apog3mN0Wj86zrZ5
Unseal Key 5: v1NjX3BAEFHPzPwd3Kr0F5QXoJcQ6RNmy6Kb+mTAL1/q
Initial Root Token: s.Ni4R8qIv87976TU816bRbzTe
```

```
VAULT_ADDR=http://127.0.0.1:8200/
VAULT_TOKEN=s.Ni4R8qIv87976TU816bRbzTe
export VAULT_ADDR VAULT_TOKEN
```

```
kc -n vault port-forward --address 0.0.0.0 svc/vault 8200:8200

kc -n vault exec -it vault-0 -- vault operator init
kc -n vault exec -it vault-0 -- vault status

# on each vault-N pod
kc -n vault exec -it vault-0 -- vault operator unseal BVBM0zdWs2MUX8+Vw+DPFmmcYU5PBtD14DO2CndNfTXQ
kc -n vault exec -it vault-0 -- vault operator unseal UBO28nH2KZrZriWRmZUIhraWnaMKbYKzMEOjnM2Swtwq
kc -n vault exec -it vault-0 -- vault operator unseal HJxoTs5qtErZbwCOIKFVesxFcwx1Apog3mN0Wj86zrZ5
```

```
kc -n argo-demo create sa vault-auth
kc -n argo-demo apply -f vault-cluster-role-binding.yaml

VAULT_SA_NAME=$(kc -n argo-demo get -o json sa vault-auth | jq -r .secrets[].name)
VAULT_SA_JWT_TOKEN=$(kc -n argo-demo get -o json secret $VAULT_SA_NAME | jq -r .data.token | base64 -d)

VAULT_SA_CA_CRT=$(kc config view --raw --minify --flatten -o json | jq -r '.clusters[0].cluster["certificate-authority-data"]' | base64 -d)
VAULT_K8S_HOST=$(kc config view --raw --minify --flatten -o json | jq -r .clusters[0].cluster.server)

vault login

vault auth list
vault auth enable kubernetes
vault write auth/kubernetes/config \
      token_reviewer_jwt="$VAULT_SA_JWT_TOKEN" \
      kubernetes_host="$VAULT_K8S_HOST" \
      kubernetes_ca_cert="$VAULT_SA_CA_CRT" \
      issuer="https://kubernetes.default.svc.cluster.local"

vault secrets list
vault secrets enable -path=secret kv-v2

vault kv list secret
vault kv list secret/webapp
vault kv get secret/webapp/config
vault kv put secret/webapp/config username="szabo" password="p@ssw0rd" ttl="30s"

vault policy list
vault policy read webapp
vault policy write webapp - <<EOF
path "secret/data/webapp/config" {
    capabilities = ["read", "list"]
}
EOF

vault list auth/kubernetes/role
vault read auth/kubernetes/role/webapp
vault write auth/kubernetes/role/webapp \
      bound_service_account_names=vault-auth \
      bound_service_account_namespaces=argo-demo \
      policies=webapp \
      ttl=24h
```
