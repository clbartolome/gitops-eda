# gitops-eda

Repository for fully automated installation and configuration of the necessary environment to run GitOps/EDA on OpenShift.

> [!IMPORTANT]  
> Tested versions: 
> - OpenShift: 4.17
> - OpenShift GitOps: 1.14.0
> - Openshift Custom Metrics Autoscaler: 2.15.1-6

## Install

- Open a terminal
- Login into OpenShift
- Run installation:

```sh
export SERVICENOW_HOST=https://your-servicenow-developer-instance-hostname
export SERVICENOW_USER=your-servicenow-developer-instance-username
export SERVICENOW_PASS=your-servicenow-developer-instance-password
export CLUSTER_DOMAIN=$(oc whoami --show-server | sed 's~https://api\.~~' | sed 's~:.*~~')
ansible-playbook installation/install.yaml -e "ocp_host=$CLUSTER_DOMAIN"
```

### Extra variables

`-e "install_collections=true"` Install required Ansible collections

`-e "install_operators=true"` Install required Openshift operators

> **IMPORTANT NOTE** Change the storage class in pvc.yaml according to your storage classes.

## Uninstall

- Open a terminal
- Login into OpenShift
- Run installation:
```sh
CLUSTER_DOMAIN=$(oc whoami --show-server | sed 's~https://api\.~~' | sed 's~:.*~~')
ansible-playbook installation/uninstall.yaml -e "ocp_host=$CLUSTER_DOMAIN"
```

## Keda Demo: Run load test

In order to generate load and simulate/trigger this demo we've used K6. Review official documentation [here](https://grafana.com/docs/k6/latest/)

Install **k6** locally:

```sh
sudo dnf install https://dl.k6.io/rpm/repo.rpm
sudo dnf install k6
```

Currently, Keda is scaling based on a metric exposed by the payment application named http_requests_total. This metric returns the cumulative number of requests per pod.

The query we use `rate(http_requests_total{job="payment"}[1m])`, calculates the average rate of requests per second for each pod (or instance).

KEDA will evaluate the result independently for each instance, and if at least one pod exceeds the threshold of 1 request per second, KEDA will increase the total number of pods.

kube_horizontalpodautoscaler_status_current_replicas{horizontalpodautoscaler="keda-hpa-payment-scaler", namespace="payment"}/kube_horizontalpodautoscaler_spec_max_replicas{horizontalpodautoscaler="keda-hpa-payment-scaler", namespace="payment"}*100

To force the keda scaling use the folloing command:

```
k6 run script.js
```

When the alert is triggered, show the servicenow incident and accept the Approvals in the AAP console.
The replicas should be modified in Gitea and changes should be applied by ArgoCD


## PVC Demo: Run remediation

In the Openshift console, open the terminal of the Payment Application pod.

Execute the following command to fill the disk

```
dd if=/dev/zero of=/mnt/file bs=$((1024*1024)) count=$((10*1024))
```

When the alert is triggered, show the servicenow incident and accept the Approvals in the AAP console.
The PVC size should be increased in Gitea and changes should be applied by ArgoCD
