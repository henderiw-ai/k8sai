# openAPI schema status for io.k8s.api.core.v1.ContainerStateRunning

## description

ContainerStateRunning is a running state of a container.

## schema

```yaml
|
  description: ContainerStateRunning is a running state of a container.
  properties:
    startedAt:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: Time at which the container was last (re-)started
  type: object

```
