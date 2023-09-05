# openAPI schema part2 for dlp.cnrm.cloud.google.com.v1beta1.DLPInspectTemplate

## schema

```yaml
properties:
  spec:
    properties:
      organizationRef:
        properties:
          organizationRef:
            description: Immutable. The Organization that this resource belongs to.
              Only one of [organizationRef, projectRef] may be specified.
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
                description: 'Allowed value: The Google Cloud resource name of a Google
                  Cloud Organization (format: `organizations/{{name}}`).'
                type: string
              name:
                description: |-
                  [WARNING] Organization not yet supported in Config Connector, use 'external' field to reference existing resources.
                  Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
      projectRef:
        description: Immutable. The Project that this resource belongs to. Only one
          of [organizationRef, projectRef] may be specified.
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
            description: 'Allowed value: The Google Cloud resource name of a `Project`
              resource (format: `projects/{{name}}`).'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
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
        description: Output only. The creation timestamp of an inspectTemplate.
        format: date-time
        type: string
      locationId:
        description: Output only. The geographic location where this resource is stored.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: Output only. The last update timestamp of an inspectTemplate.
        format: date-time
        type: string
    type: object
type: object

```
