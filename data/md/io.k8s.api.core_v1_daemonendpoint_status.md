# openAPI schema status for io.k8s.api.core.v1.DaemonEndpoint

## description

DaemonEndpoint contains information about a single Daemon endpoint.

## schema

```yaml
|
  description: DaemonEndpoint contains information about a single Daemon endpoint.
  properties:
    Port:
      description: Port number of the given endpoint.
      format: int32
      type: integer
  required:
  - Port
  type: object

```
