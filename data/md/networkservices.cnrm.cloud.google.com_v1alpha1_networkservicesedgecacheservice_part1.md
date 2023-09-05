# openAPI schema part1 for networkservices.cnrm.cloud.google.com.v1alpha1.NetworkServicesEdgeCacheService

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: A human-readable description of the resource.
        type: string
      disableHttp2:
        description: |-
          Disables HTTP/2.

          HTTP/2 (h2) is enabled by default and recommended for performance. HTTP/2 improves connection re-use and reduces connection setup overhead by sending multiple streams over the same connection.

          Some legacy HTTP clients may have issues with HTTP/2 connections due to broken HTTP/2 implementations. Setting this to true will prevent HTTP/2 from being advertised and negotiated.
        type: boolean
      disableQuic:
        description: HTTP/3 (IETF QUIC) and Google QUIC are enabled by default.
        type: boolean
      edgeSecurityPolicy:
        description: Resource URL that points at the Cloud Armor edge security policy
          that is applied on each request against the EdgeCacheService.
        type: string
      edgeSslCertificates:
        description: |-
          URLs to sslCertificate resources that are used to authenticate connections between users and the EdgeCacheService.

          Note that only "global" certificates with a "scope" of "EDGE_CACHE" can be attached to an EdgeCacheService.
        items:
          type: string
        type: array
      logConfig:
        description: Specifies the logging options for the traffic served by this
          service. If logging is enabled, logs will be exported to Cloud Logging.
        properties:
          enable:
            description: Specifies whether to enable logging for traffic served by
              this service.
            type: boolean
          sampleRate:
            description: |-
              Configures the sampling rate of requests, where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported. The default value is 1.0, and the value of the field must be in [0, 1].

              This field can only be specified if logging is enabled for this service.
            type: number
        type: object
      projectRef:
        description: The project that this resource belongs to.
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
            description: 'Allowed value: The `name` field of a `Project` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      requireTls:
        description: |-
          Require TLS (HTTPS) for all clients connecting to this service.

          Clients who connect over HTTP (port 80) will receive a HTTP 301 to the same URL over HTTPS (port 443).
          You must have at least one (1) edgeSslCertificate specified to enable this.
        type: boolean
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      spec:
        required:
        - projectRef
        - routing
        type: object
required:
- spec
type: object

```
