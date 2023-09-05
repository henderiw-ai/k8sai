# openAPI schema status for io.k8s.api.core.v1.NodeDaemonEndpoints

## description

NodeDaemonEndpoints lists ports opened by daemons running on the Node.

## schema

```yaml
|
  description: NodeDaemonEndpoints lists ports opened by daemons running on the Node.
  properties:
    kubeletEndpoint:
      $ref: '#/definitions/io.k8s.api.core.v1.DaemonEndpoint'
      description: Endpoint on which Kubelet is listening.
  type: object

```
