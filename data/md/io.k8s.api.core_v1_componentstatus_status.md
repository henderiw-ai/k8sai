# openAPI schema status for io.k8s.api.core.v1.ComponentStatus

## description

ComponentStatus (and ComponentStatusList) holds the cluster validation info. Deprecated: This API is deprecated in v1.19+

## schema

```yaml
|
  description: 'ComponentStatus (and ComponentStatusList) holds the cluster validation
    info. Deprecated: This API is deprecated in v1.19+'
  properties:
    apiVersion:
      description: 'APIVersion defines the versioned schema of this representation of
        an object. Servers should convert recognized schemas to the latest internal
        value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
      type: string
    conditions:
      description: List of component conditions observed
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.ComponentCondition'
      type: array
      x-kubernetes-patch-merge-key: type
      x-kubernetes-patch-strategy: merge
    kind:
      description: 'Kind is a string value representing the REST resource this object
        represents. Servers may infer this from the endpoint the client submits requests
        to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
      type: string
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta'
      description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: ComponentStatus
    version: v1

```