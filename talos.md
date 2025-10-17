# Talos

## Links

- https://www.talos.dev/
- https://factory.talos.dev/ - iso with extensions and in place upgrades
- https://github.com/siderolabs/talos/releases - talos releases
- https://github.com/siderolabs/kubelet/releases - kubernetes releases

## Commands

Re-apply worker config (with upgraded images for instance)
```
talosctl apply-config --nodes NODE_IP -e CONTROLPLANE_IP --file worker.yaml  --talosconfig=./talosconfig
```

Same for control plane
```
talosctl apply-config --nodes CONTROLPLANE_IP -e CONTROLPLANE_IP --file controlplane.yaml --talosconfig=./talosconfig
```


Upgrade Talos (replace image link from result from factory)
```
talosctl upgrade --nodes NODE_IP --image factory.talos.dev/metal-installer/* -e CONTROLPLANE_IP --talosconfig=./talosconfig
```

