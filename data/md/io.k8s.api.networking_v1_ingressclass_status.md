# openAPI schema status for io.k8s.api.networking.v1.IngressClass

## description

IngressClass represents the class of the Ingress, referenced by the Ingress Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to indicate that an IngressClass should be considered default. When a single IngressClass resource has this annotation set to true, new Ingress resources without a class specified will be assigned this default class.

## schema

```yaml
|
  description: IngressClass represents the class of the Ingress, referenced by the Ingress
    Spec. The `ingressclass.kubernetes.io/is-default-class` annotation can be used to
    indicate that an IngressClass should be considered default. When a single IngressClass
    resource has this annotation set to true, new Ingress resources without a class
    specified will be assigned this default class.
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
      $ref: '#/definitions/io.k8s.api.networking.v1.IngressClassSpec'
      description: 'spec is the desired state of the IngressClass. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status'
  type: object
  x-kubernetes-group-version-kind:
  - group: networking.k8s.io
    kind: IngressClass
    version: v1

```