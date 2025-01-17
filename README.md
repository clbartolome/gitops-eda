# gitops-eda

Repository for fully automated installation and configuration of the necessary environment to run GitOps/EDA on OpenShift.

> [!IMPORTANT]  
> Tested versions: 
> - OpenShift: 4.17
> - OpenShift GitOps: 1.14.0

## Install

- Open a terminal
- Login into OpenShift
- Run installation:

```sh
SERVICENOW_HOST=https://your-servicenow-developer-instance-hostname
SERVICENOW_USER=your-servicenow-developer-instance-username
SERVICENOW_PASS=your-servicenow-developer-instance-password
CLUSTER_DOMAIN=$(oc whoami --show-server | sed 's~https://api\.~~' | sed 's~:.*~~')
ansible-playbook installation/install.yaml -e "ocp_host=$CLUSTER_DOMAIN"
```

## Unnstall

- Open a terminal
- Login into OpenShift
- Run installation:
```sh
CLUSTER_DOMAIN=$(oc whoami --show-server | sed 's~https://api\.~~' | sed 's~:.*~~')
ansible-playbook installation/uninstall.yaml -e "ocp_host=$CLUSTER_DOMAIN"
```