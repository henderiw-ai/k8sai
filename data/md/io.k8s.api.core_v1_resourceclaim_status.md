# openAPI schema status for io.k8s.api.core.v1.ResourceClaim

## description

ResourceClaim references one entry in PodSpec.ResourceClaims.

## schema

```yaml
|
  description: ResourceClaim references one entry in PodSpec.ResourceClaims.
  properties:
    name:
      description: Name must match the name of one entry in pod.spec.resourceClaims
        of the Pod where this field is used. It makes that resource available inside
        a container.
      type: string
  required:
  - name
  type: object

```
