# openAPI schema status for io.k8s.api.core.v1.Pod

## description

Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.

## schema

```yaml
|
  description: Pod is a collection of containers that can run on a host. This resource
    is created by clients and scheduled onto hosts.
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
      $ref: '#/definitions/io.k8s.api.core.v1.PodSpec'
      description: 'Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
    status:
      $ref: '#/definitions/io.k8s.api.core.v1.PodStatus'
      description: 'Most recently observed status of the pod. This data may not be up
        to date. Populated by the system. Read-only. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: Pod
    version: v1

```