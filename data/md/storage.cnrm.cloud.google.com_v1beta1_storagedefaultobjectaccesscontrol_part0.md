# openAPI schema part0 for storage.cnrm.cloud.google.com.v1beta1.StorageDefaultObjectAccessControl

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
      bucketRef:
        description: Reference to the bucket.
        oneOf:
        - not:
            required:
            - external
          required:
          - name
        - not:
            anyOf:
            - required:
              - name
            - required:
              - namespace
          required:
          - external
        properties:
          external:
            description: 'Allowed value: The `name` field of a `StorageBucket` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      entity:
        description: |-
          The entity holding the permission, in one of the following forms:
            * user-{{userId}}
            * user-{{email}} (such as "user-liz@example.com")
            * group-{{groupId}}
            * group-{{email}} (such as "group-example@googlegroups.com")
            * domain-{{domain}} (such as "domain-example.com")
            * project-team-{{projectId}}
            * allUsers
            * allAuthenticatedUsers.
        type: string
      object:
        description: The name of the object, if applied to an object.
        type: string
      role:
        description: 'The access permission for the entity. Possible values: ["OWNER",
          "READER"].'
        type: string
    required:
    - bucketRef
    - entity
    - role
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
      domain:
        description: The domain associated with the entity.
        type: string
      email:
        description: The email address associated with the entity.
        type: string
      entityId:
        description: The ID for the entity.
        type: string
      generation:
        description: The content generation of the object, if applied to an object.
        type: integer
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      projectTeam:
        description: The project team associated with the entity.
        properties:
          projectNumber:
            description: The project team associated with the entity.
            type: string
          team:
            description: 'The team. Possible values: ["editors", "owners", "viewers"].'
            type: string
        type: object
    type: object
required:
- spec
type: object

```