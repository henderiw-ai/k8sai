# openAPI schema status for io.k8s.apimachinery.pkg.apis.meta.v1.Preconditions

## description

Preconditions must be fulfilled before an operation (update, delete, etc.) is carried out.

## schema

```yaml
|
  description: Preconditions must be fulfilled before an operation (update, delete,
    etc.) is carried out.
  properties:
    resourceVersion:
      description: Specifies the target ResourceVersion
      type: string
    uid:
      description: Specifies the target UID.
      type: string
  type: object

```
