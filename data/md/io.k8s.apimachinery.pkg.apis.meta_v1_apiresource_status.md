# openAPI schema status for io.k8s.apimachinery.pkg.apis.meta.v1.APIResource

## description

APIResource specifies the name of a resource and whether it is namespaced.

## schema

```yaml
|
  description: APIResource specifies the name of a resource and whether it is namespaced.
  properties:
    categories:
      description: categories is a list of the grouped resources this resource belongs
        to (e.g. 'all')
      items:
        type: string
      type: array
    group:
      description: 'group is the preferred group of the resource.  Empty implies the
        group of the containing resource list. For subresources, this may have a different
        value, for example: Scale".'
      type: string
    kind:
      description: kind is the kind for the resource (e.g. 'Foo' is the kind for a resource
        'foo')
      type: string
    name:
      description: name is the plural name of the resource.
      type: string
    namespaced:
      description: namespaced indicates if a resource is namespaced or not.
      type: boolean
    shortNames:
      description: shortNames is a list of suggested short names of the resource.
      items:
        type: string
      type: array
    singularName:
      description: singularName is the singular name of the resource.  This allows clients
        to handle plural and singular opaquely. The singularName is more correct for
        reporting status on a single item and both singular and plural are allowed from
        the kubectl CLI interface.
      type: string
    storageVersionHash:
      description: The hash value of the storage version, the version this resource
        is converted to when written to the data store. Value must be treated as opaque
        by clients. Only equality comparison on the value is valid. This is an alpha
        feature and may change or be removed in the future. The field is populated by
        the apiserver only if the StorageVersionHash feature gate is enabled. This field
        will remain optional even if it graduates.
      type: string
    verbs:
      description: verbs is a list of supported kube verbs (this includes get, list,
        watch, create, update, patch, delete, deletecollection, and proxy)
      items:
        type: string
      type: array
    version:
      description: 'version is the preferred version of the resource.  Empty implies
        the version of the containing resource list For subresources, this may have
        a different value, for example: v1 (while inside a v1beta1 version of the core
        resource''s group)".'
      type: string
  required:
  - name
  - singularName
  - namespaced
  - kind
  - verbs
  type: object

```