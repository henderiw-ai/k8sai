# openAPI schema status for io.k8s.api.resource.v1alpha2.ResourceClaimSchedulingStatus

## description

ResourceClaimSchedulingStatus contains information about one particular ResourceClaim with "WaitForFirstConsumer" allocation mode.

## schema

```yaml
|
  description: ResourceClaimSchedulingStatus contains information about one particular
    ResourceClaim with "WaitForFirstConsumer" allocation mode.
  properties:
    name:
      description: Name matches the pod.spec.resourceClaims[*].Name field.
      type: string
    unsuitableNodes:
      description: |-
        UnsuitableNodes lists nodes that the ResourceClaim cannot be allocated for.

        The size of this field is limited to 128, the same as for PodSchedulingSpec.PotentialNodes. This may get increased in the future, but not reduced.
      items:
        type: string
      type: array
      x-kubernetes-list-type: atomic
  type: object

```
