# openAPI schema status for io.k8s.api.core.v1.Capabilities

## description

Adds and removes POSIX capabilities from running containers.

## schema

```yaml
|
  description: Adds and removes POSIX capabilities from running containers.
  properties:
    add:
      description: Added capabilities
      items:
        type: string
      type: array
    drop:
      description: Removed capabilities
      items:
        type: string
      type: array
  type: object

```
