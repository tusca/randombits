# kube

Preface: freelens is an amazing tool that can avoid quite some cli, but I do love CLI so I use ... both ;-)

Random bits and piece losely related to kubernetes

## list pods with image and status as a table

```
kubectl get pods -o json | jq -r  '.items[] as $pod | "\($pod.metadata.name) \($pod.spec.containers[0].image) \($pod.status.phase)"' | column -t -s ' '
```

## refresh an external secret without deleting the secret

```
kubectl annotate externalsecret NAME force-sync=$(date +%s) --overwrite
```

## show pods with ips and nodes

```
kubectl get pods -o wide
```

## use the force to drain the node

```
kubectl drain NODENAME  --ignore-daemonsets --delete-emptydir-data --force
```

## drain node of pod

```
kubectl drain $(kubectl get pod PODNAME -o=json | jq -r .spec.nodeName) --ignore-daemonsets --delete-local-data --force
```

## show all http proxies

```
kubectl get httpproxy --all-namespaces -o json | jq -r '.items | sort_by(.spec.ingressClassName, .metadata.namespace, .metadata.name) [] | [.spec.ingressClassName, .metadata.namespace, .metadata.name, .spec.virtualhost.fqdn] | @tsv' | column -t
```


# Scripts

## Drain Node on Which POD runs

```
export POD=$1
if [[ "$POD" == "" ]]; then
  echo "Usage: drainnodeofpod PODID"
  exit 0
fi
 
kubectl drain $(kubectl get pod $POD -o=json | jq -r .spec.nodeName) --ignore-daemonsets --delete-local-data $2

```

## View kubernetes secret

```
#!/bin/bash

SECRET_NAME=$1
SECRET_KEY=$2

if [ -z "$SECRET_NAME" ]; then
    kubectl get secret
    exit 0
fi

if [ -z "$SECRET_KEY" ]; then
    kubectl get secret "$SECRET_NAME" -o yaml | yq '.data |= with_entries(.value |= @base64d)'
    exit 0
fi

kubectl get secret "$SECRET_NAME" -o yaml | yq '.data."'"$SECRET_KEY"'"' | base64 --decode

```

## list all pods running on nodes in a specific karpenter nodepool

```

#!/bin/bash
# Check if POOL parameter is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <POOL>"
  exit 1
fi

POOL=$1
# Get nodes with the specified nodepool label
NODES=$(kubectl get nodes -l karpenter.sh/nodepool=$POOL -o jsonpath='{.items[*].metadata.name}')

# Loop through each node
for NODE in $NODES; do
  echo "Node: $NODE"
  echo "Pods:"
  kubectl get pods --all-namespaces --field-selector spec.nodeName=$NODE -o custom-columns='NAMESPACE:.metadata.namespace,POD:.metadata.name' --no-headers
  echo ""
done



```
