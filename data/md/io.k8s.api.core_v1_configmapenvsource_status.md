# openAPI schema status for io.k8s.api.core.v1.ConfigMapEnvSource

## description

ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.

The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.

## schema

```yaml
|
  description: |-
    ConfigMapEnvSource selects a ConfigMap to populate the environment variables with.

    The contents of the target ConfigMap's Data field will represent the key-value pairs as environment variables.
  properties:
    name:
      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
      type: string
    optional:
      description: Specify whether the ConfigMap must be defined
      type: boolean
  type: object

```
