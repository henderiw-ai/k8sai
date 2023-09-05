# openAPI schema status for io.k8s.api.core.v1.Binding

## description

Binding ties one object to another; for example, a pod is bound to a node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods instead.

## schema

```yaml
|
  description: Binding ties one object to another; for example, a pod is bound to a
    node by a scheduler. Deprecated in 1.7, please use the bindings subresource of pods
    instead.
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
    target:
      $ref: '#/definitions/io.k8s.api.core.v1.ObjectReference'
      description: The target object that you want to bind to the standard object.
  required:
  - target
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: Binding
    version: v1

```