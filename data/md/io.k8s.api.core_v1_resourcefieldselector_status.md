# openAPI schema status for io.k8s.api.core.v1.ResourceFieldSelector

## description

ResourceFieldSelector represents container resources (cpu, memory) and their output format

## schema

```yaml
|
  description: ResourceFieldSelector represents container resources (cpu, memory) and
    their output format
  properties:
    containerName:
      description: 'Container name: required for volumes, optional for env vars'
      type: string
    divisor:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.api.resource.Quantity'
      description: Specifies the output format of the exposed resources, defaults to
        "1"
    resource:
      description: 'Required: resource to select'
      type: string
  required:
  - resource
  type: object
  x-kubernetes-map-type: atomic

```
