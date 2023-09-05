# openAPI schema part6 for monitoring.cnrm.cloud.google.com.v1beta1.MonitoringDashboard

## schema

```yaml
properties:
  spec:
    properties:
      mosaicLayout:
        properties:
          columns:
            description: The number of columns in the mosaic grid.
            format: int64
            type: integer
          mosaicLayout:
            description: The content is arranged as a grid of tiles, with each content
              widget occupying one or more tiles.
            type: object
required:
- spec
type: object

```
