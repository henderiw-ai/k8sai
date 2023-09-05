# openAPI schema part1 for accesscontextmanager.cnrm.cloud.google.com.v1beta1.AccessContextManagerServicePerimeter

## schema

```yaml
properties:
  spec:
    properties:
      accessPolicyRef:
        description: |-
          The AccessContextManagerAccessPolicy this
          AccessContextManagerServicePerimeter lives in.
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
            description: 'Allowed value: string of the format `accessPolicies/{{value}}`,
              where {{value}} is the `name` field of an `AccessContextManagerAccessPolicy`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      description:
        description: |-
          Description of the ServicePerimeter and its use. Does not affect
          behavior.
        type: string
      perimeterType:
        description: |-
          Immutable. Specifies the type of the Perimeter. There are two types: regular and
          bridge. Regular Service Perimeter contains resources, access levels,
          and restricted services. Every resource can be in at most
          ONE regular Service Perimeter.

          In addition to being in a regular service perimeter, a resource can also
          be in zero or more perimeter bridges. A perimeter bridge only contains
          resources. Cross project operations are permitted if all effected
          resources share some perimeter (whether bridge or regular). Perimeter
          Bridge does not contain access levels or services: those are governed
          entirely by the regular perimeter that resource is in.

          Perimeter Bridges are typically useful when building more complex
          topologies with many independent perimeters that need to share some data
          with a common perimeter, but should not be able to share data among
          themselves. Default value: "PERIMETER_TYPE_REGULAR" Possible values: ["PERIMETER_TYPE_REGULAR", "PERIMETER_TYPE_BRIDGE"].
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      spec:
        required:
        - accessPolicyRef
        - title
        type: object
required:
- spec
type: object

```
