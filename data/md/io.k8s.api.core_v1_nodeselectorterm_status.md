# openAPI schema status for io.k8s.api.core.v1.NodeSelectorTerm

## description

A null or empty node selector term matches no objects. The requirements of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.

## schema

```yaml
|
  description: A null or empty node selector term matches no objects. The requirements
    of them are ANDed. The TopologySelectorTerm type implements a subset of the NodeSelectorTerm.
  properties:
    matchExpressions:
      description: A list of node selector requirements by node's labels.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.NodeSelectorRequirement'
      type: array
    matchFields:
      description: A list of node selector requirements by node's fields.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.NodeSelectorRequirement'
      type: array
  type: object
  x-kubernetes-map-type: atomic

```
