VAULT_ADDR=http://127.0.0.1:8200/
VAULT_TOKEN=s.1kC4RkeURylB1hYYToEQp94S
export VAULT_ADDR VAULT_TOKEN

SOPS_AGE_KEY_FILE=$HOME/.age/age.key
SOPS_AGE_RECIPIENTS=age18affr3vq66ehdmp49qhlpv9lclw2rjl4w4qdxwasz8szdvyqwczsqfjkca
export SOPS_AGE_KEY_FILE SOPS_AGE_RECIPIENTS

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
