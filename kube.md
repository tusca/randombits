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


# Scripts

## TBD
