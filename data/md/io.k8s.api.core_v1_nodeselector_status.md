# openAPI schema status for io.k8s.api.core.v1.NodeSelector

## description

A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.

## schema

```yaml
|
  description: A node selector represents the union of the results of one or more label
    queries over a set of nodes; that is, it represents the OR of the selectors represented
    by the node selector terms.
  properties:
    nodeSelectorTerms:
      description: Required. A list of node selector terms. The terms are ORed.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.NodeSelectorTerm'
      type: array
  required:
  - nodeSelectorTerms
  type: object
  x-kubernetes-map-type: atomic

```