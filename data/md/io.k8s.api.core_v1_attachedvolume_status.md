# openAPI schema status for io.k8s.api.core.v1.AttachedVolume

## description

AttachedVolume describes a volume attached to a node

## schema

```yaml
|
  description: AttachedVolume describes a volume attached to a node
  properties:
    devicePath:
      description: DevicePath represents the device path where the volume should be
        available
      type: string
    name:
      description: Name of the attached volume
      type: string
  required:
  - name
  - devicePath
  type: object

```
