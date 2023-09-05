# node srlinux example with defaualt parameters

## description

This is a example to deploy a srlinux node using the network node operator, with the default paremeters:
- model: ixr3dl
- image: latest

We deploy a certificate, s secret with credentials and a node using the `srlinux.nokia.com` provider in the default namespace.

## example yaml resources

add a certificate, which aligns with the node name that will be used later in the example. The certificate uses a self-signed certificate issuer that was setup earlier.

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata: 
  name: leaf1
spec:
  dnsNames:
  - leaf1.default.svc
  - leaf1.default.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: selfsigned-issuer
  secretName: leaf1
```

add a secret using the provider name.

```bash
kubectl create secret generic -n default srlinux.nokia.com \
  --from-literal username=admin \
  --from-literal password=NokiaSrl1! \
  --type kubernetes.io/basic-auth
```

deploy a srlinux node wiht the default parameters

```yaml
apiVersion: inv.nephio.org/v1alpha1
kind: Node
metadata:
  name: leaf1
  labels:
    topo.nephio.org/position: leaf
    topo.nephio.org/rack: rack1
    topo.nephio.org/rack-index: "0"
spec:
  provider: srlinux.nokia.com
```
