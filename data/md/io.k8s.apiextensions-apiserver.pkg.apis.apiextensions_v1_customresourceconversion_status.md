# openAPI schema status for io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceConversion

## description

CustomResourceConversion describes how to convert different versions of a CR.

## schema

```yaml
|
  description: CustomResourceConversion describes how to convert different versions
    of a CR.
  properties:
    strategy:
      description: |-
        strategy specifies how custom resources are converted between versions. Allowed values are: - `"None"`: The converter only change the apiVersion and would not touch any other field in the custom resource. - `"Webhook"`: API Server will call to an external webhook to do the conversion. Additional information
          is needed for this option. This requires spec.preserveUnknownFields to be false, and spec.conversion.webhook to be set.
      type: string
    webhook:
      $ref: '#/definitions/io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.WebhookConversion'
      description: webhook describes how to call the conversion webhook. Required when
        `strategy` is set to `"Webhook"`.
  required:
  - strategy
  type: object

```