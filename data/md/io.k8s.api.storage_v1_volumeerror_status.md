# openAPI schema status for io.k8s.api.storage.v1.VolumeError

## description

VolumeError captures an error encountered during a volume operation.

## schema

```yaml
|
  description: VolumeError captures an error encountered during a volume operation.
  properties:
    message:
      description: message represents the error encountered during Attach or Detach
        operation. This string may be logged, so it should not contain sensitive information.
      type: string
    time:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: time represents the time the error was encountered.
  type: object

```
