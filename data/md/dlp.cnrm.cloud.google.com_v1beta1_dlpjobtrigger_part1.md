# openAPI schema part1 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: User provided description (max 256 chars)
        type: string
      displayName:
        description: Display name (max 100 chars)
        type: string
      spec:
        required:
        - inspectJob
        - projectRef
        - status
        - triggers
        type: object
required:
- spec
type: object

```
