# openAPI schema part0 for networksecurity.cnrm.cloud.google.com.v1beta1.NetworkSecurityClientTLSPolicy

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
      clientCertificate:
        description: Optional. Defines a mechanism to provision client identity (public
          and private keys) for peer to peer authentication. The presence of this
          dictates mTLS.
        properties:
          certificateProviderInstance:
            description: The certificate provider instance specification that will
              be passed to the data plane, which will be used to load necessary credential
              information.
            properties:
              pluginInstance:
                description: Required. Plugin instance name, used to locate and load
                  CertificateProvider instance configuration. Set to "google_cloud_private_spiffe"
                  to use Certificate Authority Service certificate provider instance.
                type: string
            required:
            - pluginInstance
            type: object
          grpcEndpoint:
            description: gRPC specific configuration to access the gRPC server to
              obtain the cert and private key.
            properties:
              targetUri:
                description: Required. The target URI of the gRPC endpoint. Only UDS
                  path is supported, and should start with “unix:”.
                type: string
            required:
            - targetUri
            type: object
        type: object
      description:
        description: Optional. Free-text description of the resource.
        type: string
      location:
        description: Immutable. The location for the resource
        type: string
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
      serverValidationCa:
        description: Required. Defines the mechanism to obtain the Certificate Authority
          certificate to validate the server certificate.
        items:
          properties:
            certificateProviderInstance:
              description: The certificate provider instance specification that will
                be passed to the data plane, which will be used to load necessary
                credential information.
              properties:
                pluginInstance:
                  description: Required. Plugin instance name, used to locate and
                    load CertificateProvider instance configuration. Set to "google_cloud_private_spiffe"
                    to use Certificate Authority Service certificate provider instance.
                  type: string
              required:
              - pluginInstance
              type: object
            grpcEndpoint:
              description: gRPC specific configuration to access the gRPC server to
                obtain the CA certificate.
              properties:
                targetUri:
                  description: Required. The target URI of the gRPC endpoint. Only
                    UDS path is supported, and should start with “unix:”.
                  type: string
              required:
              - targetUri
              type: object
          type: object
        type: array
      sni:
        description: 'Optional. Server Name Indication string to present to the server
          during TLS handshake. E.g: "secure.example.com".'
        type: string
    required:
    - location
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
        description: Output only. The timestamp when the resource was created.
        format: date-time
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: Output only. The timestamp when the resource was updated.
        format: date-time
        type: string
    type: object
required:
- spec
type: object

```
