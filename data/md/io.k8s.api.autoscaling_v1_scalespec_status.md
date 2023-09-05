# openAPI schema status for io.k8s.api.autoscaling.v1.ScaleSpec

## description

ScaleSpec describes the attributes of a scale subresource.

## schema

```yaml
|
  description: ScaleSpec describes the attributes of a scale subresource.
  properties:
    replicas:
      description: replicas is the desired number of instances for the scaled object.
      format: int32
      type: integer
  type: object

```
