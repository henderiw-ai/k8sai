# openAPI schema status for io.k8s.api.core.v1.PortStatus

## schema

```yaml
|
  properties:
    error:
      description: |-
        Error is to record the problem with the service port The format of the error shall comply with the following rules: - built-in error values shall be specified in this file and those shall use
          CamelCase names
        - cloud provider specific error values must have names that comply with the
          format foo.example.com/CamelCase.
      type: string
    port:
      description: Port is the port number of the service port of which status is recorded
        here
      format: int32
      type: integer
    protocol:
      description: 'Protocol is the protocol of the service port of which status is
        recorded here The supported values are: "TCP", "UDP", "SCTP"'
      type: string
  required:
  - port
  - protocol
  type: object

```