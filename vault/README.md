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
VAULT_APPLICATION_NS=argo-demo
VAULT_APPLICATION_SA=vault-auth

VAULT_CLUSTER_NS=default
VAULT_CLUSTER_SA=vault-cluster-auth

VAULT_CLUSTER_SA_TOKEN=$(kc -n $VAULT_CLUSTER_NS get -o json sa $VAULT_CLUSTER_SA | jq -r .secrets[0].name)
VAULT_CLUSTER_SA_JWT_TOKEN=$(kc -n $VAULT_CLUSTER_NS get -o json secret $VAULT_CLUSTER_SA_TOKEN | jq -r .data.token | base64 -d)

VAULT_CLUSTER_HOST=$(kc config view --raw --minify --flatten -o json | jq -r .clusters[0].cluster.server)
VAULT_CLUSTER_CA=$(kc config view --raw --minify --flatten -o json | jq -r '.clusters[0].cluster["certificate-authority-data"]' | base64 -d)

vault login

vault auth list
vault auth enable -path=dev-cluster -description "dev k8s" kubernetes
vault auth enable -path=pat-cluster -description "pat k8s" kubernetes
vault auth enable -path=prd-cluster -description "prd k8s" kubernetes
vault auth disable dev-cluster
vault auth disable pat-cluster
vault auth disable prd-cluster

vault read auth/dev-cluster/config
vault write auth/dev-cluster/config \
      token_reviewer_jwt="$VAULT_CLUSTER_SA_JWT_TOKEN" \
      kubernetes_host="$VAULT_CLUSTER_HOST" \
      kubernetes_ca_cert="$VAULT_CLUSTER_CA" \
      issuer="https://kubernetes.default.svc.cluster.local"

vault secrets list
vault secrets enable -path=secret kv-v2

vault kv list secret
vault kv list secret/webapp
vault kv get secret/webapp/config
vault kv put secret/webapp/config username="myuser" password="mypass101" ttl="30s"
vault kv metadata get secret/webapp/config
vault kv metadata delete secret/webapp/config

vault policy list
vault policy read webapp
vault policy delete webapp
vault policy write webapp - <<EOF
path "secret/data/webapp/config" {
    capabilities = ["read", "list"]
}
EOF

vault list auth/dev-cluster/role
vault read auth/dev-cluster/role/webapp
vault delete auth/dev-cluster/role/webapp
vault write auth/dev-cluster/role/webapp \
      bound_service_account_names=$VAULT_APPLICATION_SA \
      bound_service_account_namespaces=$VAULT_APPLICATION_NS \
      policies=webapp \
      ttl=24h
```

```
annotations:
  traffic.sidecar.istio.io/excludeOutboundPorts: "8200"
  vault.hashicorp.com/agent-init-first: "true"
  vault.hashicorp.com/agent-inject: "true"
  vault.hashicorp.com/agent-inject-secret-config: "secret/webapp/config"
  vault.hashicorp.com/agent-inject-command-config: "id"
  vault.hashicorp.com/agent-inject-template-config: |
    {{- with secret "secret/webapp/config" -}}
    MYSQL_USERNAME={{ .Data.data.username }}
    MYSQL_PASSWORD={{ .Data.data.password }}
    export MYSQL_USERNAME MYSQL_PASSWORD
    {{- end }}
  vault.hashicorp.com/auth-path: "auth/dev-cluster"
  vault.hashicorp.com/role: "webapp"
  vault.hashicorp.com/tls-skip-verify: "false"
```
