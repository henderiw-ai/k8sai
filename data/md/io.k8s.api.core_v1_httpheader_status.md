# openAPI schema status for io.k8s.api.core.v1.HTTPHeader

## description

HTTPHeader describes a custom header to be used in HTTP probes

## schema

```yaml
|
  description: HTTPHeader describes a custom header to be used in HTTP probes
  properties:
    name:
      description: The header field name. This will be canonicalized upon output, so
        case-variant names will be understood as the same header.
      type: string
    value:
      description: The header field value
      type: string
  required:
  - name
  - value
  type: object

```
