CONTAINER_RUNTIME_ENDPOINT=unix:///run/containerd/containerd.sock
IMAGE_SERVICE_ENDPOINT=$CONTAINER_RUNTIME_ENDPOINT
export CONTAINER_RUNTIME_ENDPOINT IMAGE_SERVICE_ENDPOINT

VAULT_ADDR=http://127.0.0.1:8200/
VAULT_TOKEN=XXXXXXXXXXXXXXXXXXX
export VAULT_ADDR VAULT_TOKEN

CONSUL_HTTP_ADDR=http://localhost:8500
CONSUL_HTTP_TOKEN=aaaaaaaaaaaaaaaaaaaa
CONSUL_HEADER="X-Consul-Token: $CONSUL_HTTP_TOKEN"
export CONSUL_HTTP_ADDR CONSUL_HTTP_TOKEN CONSUL_HEADER

SOPS_AGE_KEY_FILE=$HOME/.age/age.key
SOPS_AGE_RECIPIENTS=XXXXXXXXXXXXXXXXXX
export SOPS_AGE_KEY_FILE SOPS_AGE_RECIPIENTS

DATASTORE_TYPE=kubernetes
KUBECONFIG=~/.kube/config
export DATASTORE_TYPE KUBECONFIG

PATH=$PATH:$HOME/bin
PATH=$PATH:$HOME/.krew/bin
PATH=$PATH:$HOME/.istioctl/bin
export PATH

alias ls='ls -ACF --color'
alias ll='ls -l'

alias kc='kubectl'
alias kpp='watch -n 1 kubectl get po -A'
alias kps='watch -n 1 kubectl get po,svc,ep,deploy,no'

alias cdk='cd ~/src/kubernetes/'
alias cda='cd ~/src/argocd-example-apps/'
