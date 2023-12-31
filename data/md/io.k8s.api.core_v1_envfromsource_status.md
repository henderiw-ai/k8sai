# openAPI schema status for io.k8s.api.core.v1.EnvFromSource

## description

EnvFromSource represents the source of a set of ConfigMaps

## schema

```yaml
|
  description: EnvFromSource represents the source of a set of ConfigMaps
  properties:
    configMapRef:
      $ref: '#/definitions/io.k8s.api.core.v1.ConfigMapEnvSource'
      description: The ConfigMap to select from
    prefix:
      description: An optional identifier to prepend to each key in the ConfigMap. Must
        be a C_IDENTIFIER.
      type: string
    secretRef:
      $ref: '#/definitions/io.k8s.api.core.v1.SecretEnvSource'
      description: The Secret to select from
  type: object

```
