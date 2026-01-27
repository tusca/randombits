# proxy


resources/service.yaml
```yaml
apiVersion: v1
kind: Service
metadata:
  name: TARGET-ingress-svc
  namespace: ingress
spec:
  type: ExternalName
  externalName: TARGETDNS.com
  ports:
  - name: https
    port: 443
```

resources/ingress.yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: TARGET
  namespace: ingress
  annotations:
    nginx.ingress.kubernetes.io/backend-protocol: "HTTPS"
    nginx.ingress.kubernetes.io/upstream-vhost: "TARGETDNS"
    nginx.ingress.kubernetes.io/proxy-ssl-server-name: "true"
spec:
  ingressClassName: nginx
  rules:
  - host: DNSNAME
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: TARGET-ingress-svc
            port:
              number: 443
  tls:
  - hosts:
    - k8s-charts.flexpond.com
    secretName: wildcard-tls
```


kustomization.yaml
```yaml
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: ingress

resources:
  - resources/ingress.yaml
  - resources/service.yaml
```
