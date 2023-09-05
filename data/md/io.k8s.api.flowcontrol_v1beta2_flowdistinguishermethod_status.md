# openAPI schema status for io.k8s.api.flowcontrol.v1beta2.FlowDistinguisherMethod

## description

FlowDistinguisherMethod specifies the method of a flow distinguisher.

## schema

```yaml
|
  description: FlowDistinguisherMethod specifies the method of a flow distinguisher.
  properties:
    type:
      description: '`type` is the type of flow distinguisher method The supported types
        are "ByUser" and "ByNamespace". Required.'
      type: string
  required:
  - type
  type: object

```
