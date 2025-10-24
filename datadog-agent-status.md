#!/bin/bash

target_pod_name=$1

if [ -z "$target_pod_name" ]; then
  echo "usage: datadog-agent-status PODNAME"
  exit 1
fi

node_name=$(kubectl get pod $target_pod_name -o=jsonpath='{.spec.nodeName}')

echo "Node is $node_name"

pods_on_node=$(kubectl get pod -n monitoring --field-selector spec.nodeName=$node_name -o=jsonpath='{.items[*].metadata.name}')

datadog_pod_name=""
for pod in $pods_on_node; do
echo $pod
done

for pod in $pods_on_node; do
  if [[ $pod == datadog-* && ${#pod} -eq 13 ]]; then
    datadog_pod_name=$pod
    break
  fi
done

if [ -z "$datadog_pod_name" ]; then
  echo "No Datadog pod found on the same node."
  exit 1
fi

echo "Found datadog pod on that node : $datadog_pod_name"

kubectl exec -n monitoring -it $datadog_pod_name -- /bin/bash -c "agent status"


