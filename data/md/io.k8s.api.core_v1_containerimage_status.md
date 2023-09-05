# openAPI schema status for io.k8s.api.core.v1.ContainerImage

## description

Describe a container image

## schema

```yaml
|
  description: Describe a container image
  properties:
    names:
      description: Names by which this image is known. e.g. ["kubernetes.example/hyperkube:v1.0.7",
        "cloud-vendor.registry.example/cloud-vendor/hyperkube:v1.0.7"]
      items:
        type: string
      type: array
    sizeBytes:
      description: The size of the image in bytes.
      format: int64
      type: integer
  type: object

```
