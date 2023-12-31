# openAPI schema part0 for appengine.cnrm.cloud.google.com.v1alpha1.AppEngineServiceSplitTraffic

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
      migrateTraffic:
        description: If set to true traffic will be migrated to this version.
        type: boolean
      project:
        description: Immutable.
        type: string
      resourceID:
        description: Immutable. Optional. The service of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      split:
        description: Mapping that defines fractional HTTP traffic diversion to different
          versions within the service.
        properties:
          allocations:
            additionalProperties:
              type: string
            description: Mapping from version IDs within the service to fractional
              (0.000, 1] allocations of traffic for that version. Each version can
              be specified only once, but some versions in the service may not have
              any traffic allocation. Services that have traffic allocated cannot
              be deleted until either the service is deleted or their traffic allocation
              is removed. Allocations must sum to 1. Up to two decimal place precision
              is supported for IP-based splits and up to three decimal places is supported
              for cookie-based splits.
            type: object
          shardBy:
            description: 'Mechanism used to determine which version a request is sent
              to. The traffic selection algorithm will be stable for either type until
              allocations are changed. Possible values: ["UNSPECIFIED", "COOKIE",
              "IP", "RANDOM"].'
            type: string
        required:
        - allocations
        type: object
    required:
    - split
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
