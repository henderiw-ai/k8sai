# openAPI schema part0 for datacatalog.cnrm.cloud.google.com.v1alpha1.DataCatalogTag

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
      column:
        description: |-
          Resources like Entry can have schemas associated with them. This scope allows users to attach tags to an
          individual column based on that schema.

          For attaching a tag to a nested column, use '.' to separate the column names. Example:
          'outer_column.inner_column'.
        type: string
      fields:
        description: |-
          This maps the ID of a tag field to the value of and additional information about that field.
          Valid field IDs are defined by the tag's template. A tag must have at least 1 field and at most 500 fields.
        items:
          properties:
            boolValue:
              description: Holds the value for a tag field with boolean type.
              type: boolean
            displayName:
              description: The display name of this field.
              type: string
            doubleValue:
              description: Holds the value for a tag field with double type.
              type: number
            enumValue:
              description: The display name of the enum value.
              type: string
            fieldName:
              type: string
            order:
              description: |-
                The order of this field with respect to other fields in this tag. For example, a higher value can indicate
                a more important field. The value can be negative. Multiple fields can have the same order, and field orders
                within a tag do not have to be sequential.
              type: integer
            stringValue:
              description: Holds the value for a tag field with string type.
              type: string
            timestampValue:
              description: Holds the value for a tag field with timestamp type.
              type: string
          required:
          - fieldName
          type: object
        type: array
      parent:
        description: |-
          Immutable. The name of the parent this tag is attached to. This can be the name of an entry or an entry group. If an entry group, the tag will be attached to
          all entries in that group.
        type: string
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
      template:
        description: |-
          Immutable. The resource name of the tag template that this tag uses. Example:
          projects/{project_id}/locations/{location}/tagTemplates/{tagTemplateId}
          This field cannot be modified after creation.
        type: string
    required:
    - fields
    - template
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
      name:
        description: |-
          The resource name of the tag in URL format. Example:
          projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/entries/{entryId}/tags/{tag_id} or
          projects/{project_id}/locations/{location}/entrygroups/{entryGroupId}/tags/{tag_id}
          where tag_id is a system-generated identifier. Note that this Tag may not actually be stored in the location in this name.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      templateDisplayname:
        description: The display name of the tag template.
        type: string
    type: object
required:
- spec
type: object

```