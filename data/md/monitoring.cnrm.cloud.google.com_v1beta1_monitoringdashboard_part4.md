# openAPI schema part4 for monitoring.cnrm.cloud.google.com.v1beta1.MonitoringDashboard

## schema

```yaml
properties:
  spec:
    properties:
      gridLayout:
        properties:
          columns:
            description: The number of columns into which the view's width is divided.
              If omitted or set to zero, a system default will be used while rendering.
            format: int64
            type: integer
          gridLayout:
            description: Content is arranged with a basic layout that re-flows a simple
              list of informational elements like widgets or tiles.
            type: object
required:
- spec
type: object

```
