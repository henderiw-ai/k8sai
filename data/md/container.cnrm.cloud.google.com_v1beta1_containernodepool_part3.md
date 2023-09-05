# openAPI schema part3 for container.cnrm.cloud.google.com.v1beta1.ContainerNodePool

## schema

```yaml
properties:
  status:
    properties:
      status:
        properties:
          conditions:
            description: Conditions represent the latest available observation of
              the resource's current state.
            items:
              properties:
                lastTransitionTime:
                  description: Last time the condition transitioned from one status
                    to another.
                  type: string
                message:
                  description: Human-readable message indicating details about last
                    transition.
                  type: string
                reason:
                  description: Unique, one-word, CamelCase reason for the condition's
                    last transition.
                  type: string
                status:
                  description: Status is the status of the condition. Can be True,
                    False, Unknown.
                  type: string
                type:
                  description: Type is the type of the condition.
                  type: string
              type: object
            type: array
          instanceGroupUrls:
            description: The resource URLs of the managed instance groups associated
              with this node pool.
            items:
              type: string
            type: array
          managedInstanceGroupUrls:
            description: List of instance group URLs which have been assigned to this
              node pool.
            items:
              type: string
            type: array
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          operation:
            type: string
        type: object
required:
- spec
type: object

```
