# openAPI schema status for io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.ServiceReference

## description

ServiceReference holds a reference to Service.legacy.k8s.io

## schema

```yaml
|
  description: ServiceReference holds a reference to Service.legacy.k8s.io
  properties:
    name:
      description: Name is the name of the service
      type: string
    namespace:
      description: Namespace is the namespace of the service
      type: string
    port:
      description: If specified, the port on the service that hosting webhook. Default
        to 443 for backward compatibility. `port` should be a valid port number (1-65535,
        inclusive).
      format: int32
      type: integer
  type: object

```
