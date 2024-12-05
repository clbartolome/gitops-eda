# gitops-eda

Repository for fully automated installation and configuration of the necessary environment to run GitOps/EDA on OpenShift.

> [!IMPORTANT]  
> Tested versions: 
> - OpenShift: 4.17
> - OpenShift GitOps: 1.14.0

## Install

- Open a terminal
- Login into OpenShift
- Access installation->ansible-navigator: `cd installation/ansible-navigator`
- Run installation:
```sh
OPENSHIFT_TOKEN=$(oc whoami --show-token)
CLUSTER_DOMAIN=$(oc whoami --show-server | sed 's~https://api\.~~' | sed 's~:.*~~')
ansible-navigator run ../install.yaml -m stdout \
    -e "ocp_host=$CLUSTER_DOMAIN" \
    -e "api_token=$OPENSHIFT_TOKEN"
```