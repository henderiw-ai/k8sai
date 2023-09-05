# openAPI schema status for io.k8s.api.autoscaling.v2.HorizontalPodAutoscaler

## description

HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler, which automatically manages the replica count of any resource implementing the scale subresource based on the metrics specified.

## schema

```yaml
|
  description: HorizontalPodAutoscaler is the configuration for a horizontal pod autoscaler,
    which automatically manages the replica count of any resource implementing the scale
    subresource based on the metrics specified.
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
      description: 'metadata is the standard object metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
    spec:
      $ref: '#/definitions/io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerSpec'
      description: 'spec is the specification for the behaviour of the autoscaler. More
        info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status.'
    status:
      $ref: '#/definitions/io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerStatus'
      description: status is the current information about the autoscaler.
  type: object
  x-kubernetes-group-version-kind:
  - group: autoscaling
    kind: HorizontalPodAutoscaler
    version: v2

```