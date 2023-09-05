# openAPI schema status for io.k8s.apimachinery.pkg.apis.meta.v1.GroupVersionForDiscovery

## description

GroupVersion contains the "group/version" and "version" string of a version. It is made a struct to keep extensibility.

## schema

```yaml
|
  description: GroupVersion contains the "group/version" and "version" string of a version.
    It is made a struct to keep extensibility.
  properties:
    groupVersion:
      description: groupVersion specifies the API group and version in the form "group/version"
      type: string
    version:
      description: version specifies the version in the form of "version". This is to
        save the clients the trouble of splitting the GroupVersion.
      type: string
  required:
  - groupVersion
  - version
  type: object

```