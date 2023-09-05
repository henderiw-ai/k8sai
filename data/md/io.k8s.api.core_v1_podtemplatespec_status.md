# openAPI schema status for io.k8s.api.core.v1.PodTemplateSpec

## description

PodTemplateSpec describes the data a pod should have when created from a template

## schema

```yaml
|
  description: PodTemplateSpec describes the data a pod should have when created from
    a template
  properties:
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta'
      description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
    spec:
      $ref: '#/definitions/io.k8s.api.core.v1.PodSpec'
      description: 'Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
  type: object

```