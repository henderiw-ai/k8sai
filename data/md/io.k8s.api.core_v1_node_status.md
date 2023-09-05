# openAPI schema status for io.k8s.api.core.v1.Node

## description

Node is a worker node in Kubernetes. Each node will have a unique identifier in the cache (i.e. in etcd).

## schema

```yaml
|
  description: Node is a worker node in Kubernetes. Each node will have a unique identifier
    in the cache (i.e. in etcd).
  properties:
    apiVersion:
      description: 'APIVersion defines the versioned schema of this representation of
        an object. Servers should convert recognized schemas to the latest internal
        value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
      type: string
    kind:
      description: 'Kind is a string value representing the REST resource this object
        represents. Servers may infer this from the endpoint the client submits requests
        to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
      type: string
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta'
      description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
    spec:
      $ref: '#/definitions/io.k8s.api.core.v1.NodeSpec'
      description: Spec defines the behavior of a node. https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
    status:
      $ref: '#/definitions/io.k8s.api.core.v1.NodeStatus'
      description: 'Most recently observed status of the node. Populated by the system.
        Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: Node
    version: v1

```