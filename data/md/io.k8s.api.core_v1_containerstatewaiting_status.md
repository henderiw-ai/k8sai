# openAPI schema status for io.k8s.api.core.v1.ContainerStateWaiting

## description

ContainerStateWaiting is a waiting state of a container.

## schema

```yaml
|
  description: ContainerStateWaiting is a waiting state of a container.
  properties:
    message:
      description: Message regarding why the container is not yet running.
      type: string
    reason:
      description: (brief) reason the container is not yet running.
      type: string
  type: object

```
