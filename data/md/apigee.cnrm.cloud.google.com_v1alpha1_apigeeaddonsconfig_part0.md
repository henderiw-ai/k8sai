# openAPI schema part0 for apigee.cnrm.cloud.google.com.v1alpha1.ApigeeAddonsConfig

## schema

```yaml
properties:
  apiVersion:
    description: 'apiVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
    type: string
  kind:
    description: 'kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    properties:
      addonsConfig:
        description: Addon configurations of the Apigee organization.
        properties:
          advancedApiOpsConfig:
            description: Configuration for the Monetization add-on.
            properties:
              enabled:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: boolean
            type: object
          apiSecurityConfig:
            description: Configuration for the Monetization add-on.
            properties:
              enabled:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: boolean
              expiresAt:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: string
            type: object
          connectorsPlatformConfig:
            description: Configuration for the Monetization add-on.
            properties:
              enabled:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: boolean
              expiresAt:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: string
            type: object
          integrationConfig:
            description: Configuration for the Monetization add-on.
            properties:
              enabled:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: boolean
            type: object
          monetizationConfig:
            description: Configuration for the Monetization add-on.
            properties:
              enabled:
                description: Flag that specifies whether the Advanced API Ops add-on
                  is enabled.
                type: boolean
            type: object
        type: object
      org:
        description: Immutable. Name of the Apigee organization.
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
    required:
    - org
    type: object
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
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
    type: object
required:
- spec
type: object

```
