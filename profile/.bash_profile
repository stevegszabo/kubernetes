DOCKER_REGISTRY=notebook.local:5000
DOCKER_IMAGE=webapp
DOCKER_TAG=latest
DOCKER_FILE=Dockerfile.webapp
DOCKER_URL=http://$DOCKER_REGISTRY/v2/$DOCKER_IMAGE/manifests/$DOCKER_TAG
export DOCKER_REGISTRY DOCKER_IMAGE DOCKER_TAG DOCKER_FILE DOCKER_URL

VAULT_ADDR=http://127.0.0.1:8200/
VAULT_TOKEN=s.Ni4R8qIv87976TU816bRbzTe
export VAULT_ADDR VAULT_TOKEN

SOPS_AGE_KEY_FILE=$HOME/.age/age.key
SOPS_AGE_RECIPIENTS=age18affr3vq66ehdmp49qhlpv9lclw2rjl4w4qdxwasz8szdvyqwczsqfjkca
export SOPS_AGE_KEY_FILE SOPS_AGE_RECIPIENTS

PATH=$PATH:$HOME/.istioctl/bin
PATH=$PATH:$HOME/bin
PATH=$PATH:$HOME/.krew/bin
export PATH

alias ls='ls -ACF --color'
alias ll='ls -l'

alias kc='kubectl'
alias kps='watch -n 1 kubectl get po,svc,ep,deploy,no'

alias cdk='cd ~/src/kubernetes/'
