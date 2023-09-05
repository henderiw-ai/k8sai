# openAPI schema part5 for accesscontextmanager.cnrm.cloud.google.com.v1beta1.AccessContextManagerServicePerimeter

## schema

```yaml
properties:
  spec:
    properties:
      status:
        properties:
          vpcAccessibleServices:
            properties:
              vpcAccessibleServices:
                description: |-
                  Specifies how APIs are allowed to communicate within the Service
                  Perimeter.
                properties:
                  allowedServices:
                    description: |-
                      The list of APIs usable within the Service Perimeter.
                      Must be empty unless 'enableRestriction' is True.
                    items:
                      type: string
                    type: array
                  enableRestriction:
                    description: |-
                      Whether to restrict API calls within the Service Perimeter to the
                      list of APIs specified in 'allowedServices'.
                    type: boolean
                type: object
      title:
        description: Human readable title. Must be unique within the Policy.
        type: string
      useExplicitDryRunSpec:
        description: |-
          Use explicit dry run spec flag. Ordinarily, a dry-run spec implicitly exists
          for all Service Perimeters, and that spec is identical to the status for those
          Service Perimeters. When this flag is set, it inhibits the generation of the
          implicit spec, thereby allowing the user to explicitly provide a
          configuration ("spec") to use in a dry-run version of the Service Perimeter.
          This allows the user to test changes to the enforced config ("status") without
          actually enforcing them. This testing is done through analyzing the differences
          between currently enforced and suggested restrictions. useExplicitDryRunSpec must
          bet set to True if any of the fields in the spec are set to non-default values.
        type: boolean
  status:
    properties:
      conditions:
        description: Conditions represent the latest available observation of the
          resource's current state.
        items:
          properties:
            lastTransitionTime:
              description: Last time the condition transitioned from one status to
                another.
              type: string
            message:
              description: Human-readable message indicating details about last transition.
              type: string
            reason:
              description: Unique, one-word, CamelCase reason for the condition's
                last transition.
              type: string
            status:
              description: Status is the status of the condition. Can be True, False,
                Unknown.
              type: string
            type:
              description: Type is the type of the condition.
              type: string
          type: object
        type: array
      createTime:
        description: Time the AccessPolicy was created in UTC.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: Time the AccessPolicy was updated in UTC.
        type: string
    type: object
required:
- spec
type: object

```
