# openAPI schema status for io.k8s.api.networking.v1.IngressPortStatus

## description

IngressPortStatus represents the error condition of a service port

## schema

```yaml
|
  description: IngressPortStatus represents the error condition of a service port
  properties:
    error:
      description: |-
        error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use
          CamelCase names
        - cloud provider specific error values must have names that comply with the
          format foo.example.com/CamelCase.
      type: string
    port:
      description: port is the port number of the ingress port.
      format: int32
      type: integer
    protocol:
      description: 'protocol is the protocol of the ingress port. The supported values
        are: "TCP", "UDP", "SCTP"'
      type: string
  required:
  - port
  - protocol
  type: object

```
