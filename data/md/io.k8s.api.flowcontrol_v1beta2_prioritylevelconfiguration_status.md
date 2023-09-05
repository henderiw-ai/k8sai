# openAPI schema status for io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfiguration

## description

PriorityLevelConfiguration represents the configuration of a priority level.

## schema

```yaml
|
  description: PriorityLevelConfiguration represents the configuration of a priority
    level.
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
      description: '`metadata` is the standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
    spec:
      $ref: '#/definitions/io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationSpec'
      description: '`spec` is the specification of the desired behavior of a "request-priority".
        More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
    status:
      $ref: '#/definitions/io.k8s.api.flowcontrol.v1beta2.PriorityLevelConfigurationStatus'
      description: '`status` is the current status of a "request-priority". More info:
        https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
  type: object
  x-kubernetes-group-version-kind:
  - group: flowcontrol.apiserver.k8s.io
    kind: PriorityLevelConfiguration
    version: v1beta2

```