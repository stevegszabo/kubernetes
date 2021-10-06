# vault

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
