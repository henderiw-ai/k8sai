# openAPI schema status for io.k8s.api.core.v1.Sysctl

## description

Sysctl defines a kernel parameter to be set

## schema

```yaml
|
  description: Sysctl defines a kernel parameter to be set
  properties:
    name:
      description: Name of a property to set
      type: string
    value:
      description: Value of a property to set
      type: string
  required:
  - name
  - value
  type: object

```
