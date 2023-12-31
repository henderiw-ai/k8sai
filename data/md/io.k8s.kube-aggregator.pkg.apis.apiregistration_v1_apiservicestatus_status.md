# openAPI schema status for io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIServiceStatus

## description

APIServiceStatus contains derived information about an API server

## schema

```yaml
|
  description: APIServiceStatus contains derived information about an API server
  properties:
    conditions:
      description: Current service state of apiService.
      items:
        $ref: '#/definitions/io.k8s.kube-aggregator.pkg.apis.apiregistration.v1.APIServiceCondition'
      type: array
      x-kubernetes-list-map-keys:
      - type
      x-kubernetes-list-type: map
      x-kubernetes-patch-merge-key: type
      x-kubernetes-patch-strategy: merge
  type: object

```
