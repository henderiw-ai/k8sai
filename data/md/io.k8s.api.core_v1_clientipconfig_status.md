# openAPI schema status for io.k8s.api.core.v1.ClientIPConfig

## description

ClientIPConfig represents the configurations of Client IP based session affinity.

## schema

```yaml
|
  description: ClientIPConfig represents the configurations of Client IP based session
    affinity.
  properties:
    timeoutSeconds:
      description: timeoutSeconds specifies the seconds of ClientIP type session sticky
        time. The value must be >0 && <=86400(for 1 day) if ServiceAffinity == "ClientIP".
        Default value is 10800(for 3 hours).
      format: int32
      type: integer
  type: object

```
