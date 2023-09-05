# openAPI schema part0 for vertexai.cnrm.cloud.google.com.v1alpha1.VertexAIFeaturestoreEntityTypeFeature

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
      description:
        description: Description of the feature.
        type: string
      entitytype:
        description: Immutable. The name of the Featurestore to use, in the format
          projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      valueType:
        description: Immutable. Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType.
        type: string
    required:
    - entitytype
    - valueType
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
      createTime:
        description: The timestamp of when the entity type was created in RFC3339
          UTC "Zulu" format, with nanosecond resolution and up to nine fractional
          digits.
        type: string
      etag:
        description: Used to perform consistent read-modify-write updates.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      region:
        description: The region of the feature.
        type: string
      updateTime:
        description: The timestamp when the entity type was most recently updated
          in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
          fractional digits.
        type: string
    type: object
required:
- spec
type: object

```