# openAPI schema status for io.k8s.api.networking.v1.ServiceBackendPort

## description

ServiceBackendPort is the service port being referenced.

## schema

```yaml
|
  description: ServiceBackendPort is the service port being referenced.
  properties:
    name:
      description: name is the name of the port on the Service. This is a mutually exclusive
        setting with "Number".
      type: string
    number:
      description: number is the numerical port number (e.g. 80) on the Service. This
        is a mutually exclusive setting with "Name".
      format: int32
      type: integer
  type: object

```
