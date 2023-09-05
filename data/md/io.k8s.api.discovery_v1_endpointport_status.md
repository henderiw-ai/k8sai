# openAPI schema status for io.k8s.api.discovery.v1.EndpointPort

## description

EndpointPort represents a Port used by an EndpointSlice

## schema

```yaml
|
  description: EndpointPort represents a Port used by an EndpointSlice
  properties:
    appProtocol:
      description: |-
        The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:

        * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).

        * Kubernetes-defined prefixed names:
          * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540
          * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455
          * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455

        * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.
      type: string
    name:
      description: 'name represents the name of this port. All ports in an EndpointSlice
        must have a unique name. If the EndpointSlice is dervied from a Kubernetes service,
        this corresponds to the Service.ports[].name. Name must either be an empty string
        or pass DNS_LABEL validation: * must be no more than 63 characters long. * must
        consist of lower case alphanumeric characters or ''-''. * must start and end
        with an alphanumeric character. Default is empty string.'
      type: string
    port:
      description: port represents the port number of the endpoint. If this is not specified,
        ports are not restricted and must be interpreted in the context of the specific
        consumer.
      format: int32
      type: integer
    protocol:
      description: protocol represents the IP protocol for this port. Must be UDP, TCP,
        or SCTP. Default is TCP.
      type: string
  type: object
  x-kubernetes-map-type: atomic

```