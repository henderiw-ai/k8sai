# openAPI schema status for io.k8s.api.flowcontrol.v1beta3.ServiceAccountSubject

## description

ServiceAccountSubject holds detailed information for service-account-kind subject.

## schema

```yaml
|
  description: ServiceAccountSubject holds detailed information for service-account-kind
    subject.
  properties:
    name:
      description: '`name` is the name of matching ServiceAccount objects, or "*" to
        match regardless of name. Required.'
      type: string
    namespace:
      description: '`namespace` is the namespace of matching ServiceAccount objects.
        Required.'
      type: string
  required:
  - namespace
  - name
  type: object

```
