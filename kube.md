# kube

Random bits and piece losely related to kubernetes

## list pods with image and status as a table

```
kubectl get pods -o json | jq -r  '.items[] as $pod | "\($pod.metadata.name) \($pod.spec.containers[0].image) \($pod.status.phase)"' | column -t -s ' '
```
