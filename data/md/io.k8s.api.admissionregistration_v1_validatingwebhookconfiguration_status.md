# openAPI schema status for io.k8s.api.admissionregistration.v1.ValidatingWebhookConfiguration

## description

ValidatingWebhookConfiguration describes the configuration of and admission webhook that accept or reject and object without changing it.

## schema

```yaml
|
  description: ValidatingWebhookConfiguration describes the configuration of and admission
    webhook that accept or reject and object without changing it.
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
      description: 'Standard object metadata; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.'
    webhooks:
      description: Webhooks is a list of webhooks and the affected resources and operations.
      items:
        $ref: '#/definitions/io.k8s.api.admissionregistration.v1.ValidatingWebhook'
      type: array
      x-kubernetes-patch-merge-key: name
      x-kubernetes-patch-strategy: merge
  type: object
  x-kubernetes-group-version-kind:
  - group: admissionregistration.k8s.io
    kind: ValidatingWebhookConfiguration
    version: v1

```