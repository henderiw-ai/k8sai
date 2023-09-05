# node srlinux example for ixrd2l using a specific image

## description

This is a example to deploy a srlinux node using the network node operator, with the following paremeters:
- model: ixrdl
- image: ghcr.io/nokia/srlinux:23.7.1

We deploy a certificate, s secret with credentials and a node using the `srlinux.nokia.com` provider in the default namespace.

## example yaml resources

Add a certificate, which aligns with the node name used later in the example. The certificate uses a self-signed certificate issuer that was setup earlier.

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata: 
  name: leaf2
spec:
  dnsNames:
  - leaf2.default.svc
  - leaf2.default.svc.cluster.local
  issuerRef:
    kind: Issuer
    name: selfsigned-issuer
  secretName: leaf2
```

add a secret using the provider `srlinux.nokia.com` name.

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
  name: leaf2
  labels:
    topo.nephio.org/position: leaf
    topo.nephio.org/rack: rack2
    topo.nephio.org/rack-index: "0"
spec:
  provider: srlinux.nokia.com
  parametersRef:
    apiVersion: inv.nephio.org/v1alpha1
    kind: NodeConfig
    name: leaf2
```

The NodeConfig details the specific parameters such as the model and image.

```yaml
apiVersion: inv.nephio.org/v1alpha1
kind: NodeConfig
metadata:
  name: leaf2
spec:
  provider: srlinux.nokia.com
  model: ixrd2l
  image: ghcr.io/nokia/srlinux:23.7.1
```