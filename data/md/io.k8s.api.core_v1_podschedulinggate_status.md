# openAPI schema status for io.k8s.api.core.v1.PodSchedulingGate

## description

PodSchedulingGate is associated to a Pod to guard its scheduling.

## schema

```yaml
|
  description: PodSchedulingGate is associated to a Pod to guard its scheduling.
  properties:
    name:
      description: Name of the scheduling gate. Each scheduling gate must have a unique
        name field.
      type: string
  required:
  - name
  type: object

```
