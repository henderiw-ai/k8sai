# openAPI schema status for io.k8s.api.core.v1.ConfigMapVolumeSource

## description

Adapts a ConfigMap into a volume.

The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.

## schema

```yaml
|
  description: |-
    Adapts a ConfigMap into a volume.

    The contents of the target ConfigMap's Data field will be presented in a volume as files using the keys in the Data field as the file names, unless the items element is populated with specific mappings of keys to paths. ConfigMap volumes support ownership management and SELinux relabeling.
  properties:
    defaultMode:
      description: 'defaultMode is optional: mode bits used to set permissions on created
        files by default. Must be an octal value between 0000 and 0777 or a decimal
        value between 0 and 511. YAML accepts both octal and decimal values, JSON requires
        decimal values for mode bits. Defaults to 0644. Directories within the path
        are not affected by this setting. This might be in conflict with other options
        that affect the file mode, like fsGroup, and the result can be other mode bits
        set.'
      format: int32
      type: integer
    items:
      description: items if unspecified, each key-value pair in the Data field of the
        referenced ConfigMap will be projected into the volume as a file whose name
        is the key and content is the value. If specified, the listed keys will be projected
        into the specified paths, and unlisted keys will not be present. If a key is
        specified which is not present in the ConfigMap, the volume setup will error
        unless it is marked optional. Paths must be relative and may not contain the
        '..' path or start with '..'.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.KeyToPath'
      type: array
    name:
      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
      type: string
    optional:
      description: optional specify whether the ConfigMap or its keys must be defined
      type: boolean
  type: object

```