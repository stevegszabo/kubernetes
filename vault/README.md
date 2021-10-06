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
kubectl exec -it vault-0 -- vault operator init
kubectl exec -it vault-0 -- vault operator unseal BVBM0zdWs2MUX8+Vw+DPFmmcYU5PBtD14DO2CndNfTXQ
kubectl exec -it vault-0 -- vault operator unseal UBO28nH2KZrZriWRmZUIhraWnaMKbYKzMEOjnM2Swtwq
kubectl exec -it vault-0 -- vault operator unseal HJxoTs5qtErZbwCOIKFVesxFcwx1Apog3mN0Wj86zrZ5
```
