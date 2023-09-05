# openAPI schema part1 for networkservices.cnrm.cloud.google.com.v1beta1.NetworkServicesHTTPRoute

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: Optional. A free-text description of the resource. Max length
          1024 characters.
        type: string
      gateways:
        items:
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
              description: 'Allowed value: The `selfLink` field of a `NetworkServicesGateway`
                resource.'
              type: string
            name:
              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
              type: string
            namespace:
              description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
              type: string
          type: object
        type: array
      hostnames:
        description: Required. Hostnames define a set of hosts that should match against
          the HTTP host header to select a HttpRoute to process the request. Hostname
          is the fully qualified domain name of a network host, as defined by RFC
          1123 with the exception that ip addresses are not allowed. Wildcard hosts
          are supported as "*" (no prefix or suffix allowed).
        items:
          type: string
        type: array
      location:
        description: Immutable. The location for the resource
        type: string
      meshes:
        items:
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
              description: 'Allowed value: The `selfLink` field of a `NetworkServicesMesh`
                resource.'
              type: string
            name:
              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
              type: string
            namespace:
              description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
              type: string
          type: object
        type: array
      projectRef:
        description: Immutable. The Project that this resource belongs to.
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
            description: |-
              The project for the resource

              Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      spec:
        required:
        - hostnames
        - location
        - projectRef
        - rules
        type: object
required:
- spec
type: object

```
