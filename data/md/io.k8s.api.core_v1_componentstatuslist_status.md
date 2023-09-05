# openAPI schema status for io.k8s.api.core.v1.ComponentStatusList

## description

Status of all the conditions for the component as a list of ComponentStatus objects. Deprecated: This API is deprecated in v1.19+

## schema

```yaml
|
  description: 'Status of all the conditions for the component as a list of ComponentStatus
    objects. Deprecated: This API is deprecated in v1.19+'
  properties:
    apiVersion:
      description: 'APIVersion defines the versioned schema of this representation of
        an object. Servers should convert recognized schemas to the latest internal
        value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
      type: string
    items:
      description: List of ComponentStatus objects.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.ComponentStatus'
      type: array
    kind:
      description: 'Kind is a string value representing the REST resource this object
        represents. Servers may infer this from the endpoint the client submits requests
        to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
      type: string
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta'
      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
  required:
  - items
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: ComponentStatusList
    version: v1

```