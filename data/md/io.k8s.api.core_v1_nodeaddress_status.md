# openAPI schema status for io.k8s.api.core.v1.NodeAddress

## description

NodeAddress contains information for the node's address.

## schema

```yaml
|
  description: NodeAddress contains information for the node's address.
  properties:
    address:
      description: The node address.
      type: string
    type:
      description: Node address type, one of Hostname, ExternalIP or InternalIP.
      type: string
  required:
  - type
  - address
  type: object

```
