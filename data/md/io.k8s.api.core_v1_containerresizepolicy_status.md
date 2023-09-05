# openAPI schema status for io.k8s.api.core.v1.ContainerResizePolicy

## description

ContainerResizePolicy represents resource resize policy for the container.

## schema

```yaml
|
  description: ContainerResizePolicy represents resource resize policy for the container.
  properties:
    resourceName:
      description: 'Name of the resource to which this resource resize policy applies.
        Supported values: cpu, memory.'
      type: string
    restartPolicy:
      description: Restart policy to apply when specified resource is resized. If not
        specified, it defaults to NotRequired.
      type: string
  required:
  - resourceName
  - restartPolicy
  type: object

```