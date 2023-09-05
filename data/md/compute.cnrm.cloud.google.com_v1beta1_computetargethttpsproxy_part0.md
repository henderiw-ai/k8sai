# openAPI schema part0 for compute.cnrm.cloud.google.com.v1beta1.ComputeTargetHTTPSProxy

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
      certificateMapRef:
        description: |-
          Only the `external` field is supported to configure the reference.

          A reference to the CertificateMap resource uri that identifies a
          certificate map associated with the given target proxy. This field
          can only be set for global target proxies.
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
            description: 'Allowed value: string of the format `//certificatemanager.googleapis.com/projects/{{project}}/locations/{{location}}/certificateMaps/{{value}}`,
              where {{value}} is the `name` field of a `CertificateManagerCertificateMap`
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
        description: Immutable. An optional description of this resource.
        type: string
      httpKeepAliveTimeoutSec:
        description: |-
          Immutable. Specifies how long to keep a connection open, after completing a response,
          while there is no matching traffic (in seconds). If an HTTP keepalive is
          not specified, a default value (610 seconds) will be used. For Global
          external HTTP(S) load balancer, the minimum allowed value is 5 seconds and
          the maximum allowed value is 1200 seconds. For Global external HTTP(S)
          load balancer (classic), this option is not available publicly.
        type: integer
      location:
        description: 'Location represents the geographical location of the ComputeTargetHTTPSProxy.
          Specify a region name or "global" for global resources. Reference: GCP definition
          of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)'
        type: string
      proxyBind:
        description: |-
          Immutable. This field only applies when the forwarding rule that references
          this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        type: boolean
      quicOverride:
        description: |-
          Specifies the QUIC override policy for this resource. This determines
          whether the load balancer will attempt to negotiate QUIC with clients
          or not. Can specify one of NONE, ENABLE, or DISABLE. If NONE is
          specified, Google manages whether QUIC is used. Default value: "NONE" Possible values: ["NONE", "ENABLE", "DISABLE"].
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      sslCertificates:
        items:
          description: |-
            A list of ComputeSSLCertificate resources that are used to
            authenticate connections between users and the load balancer. At
            least one SSL certificate must be specified.
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
              description: 'Allowed value: The `selfLink` field of a `ComputeSSLCertificate`
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
      sslPolicyRef:
        description: |-
          A reference to the ComputeSSLPolicy resource that will be
          associated with the ComputeTargetHTTPSProxy resource. If not set,
          the ComputeTargetHTTPSProxy resource will not have any SSL policy
          configured.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeSSLPolicy`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      urlMapRef:
        description: |-
          A reference to the ComputeURLMap resource that defines the mapping
          from URL to the BackendService.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeURLMap`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
    required:
    - location
    - urlMapRef
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
      creationTimestamp:
        description: Creation timestamp in RFC3339 text format.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      proxyId:
        description: The unique identifier for the resource.
        type: integer
      selfLink:
        type: string
    type: object
required:
- spec
type: object

```