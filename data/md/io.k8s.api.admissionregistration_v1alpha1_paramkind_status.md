# openAPI schema status for io.k8s.api.admissionregistration.v1alpha1.ParamKind

## description

ParamKind is a tuple of Group Kind and Version.

## schema

```yaml
|
  description: ParamKind is a tuple of Group Kind and Version.
  properties:
    apiVersion:
      description: APIVersion is the API group version the resources belong to. In format
        of "group/version". Required.
      type: string
    kind:
      description: Kind is the API kind the resources belong to. Required.
      type: string
  type: object
  x-kubernetes-map-type: atomic

```
