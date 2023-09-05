# openAPI schema status for io.k8s.api.core.v1.Affinity

## description

Affinity is a group of affinity scheduling rules.

## schema

```yaml
|
  description: Affinity is a group of affinity scheduling rules.
  properties:
    nodeAffinity:
      $ref: '#/definitions/io.k8s.api.core.v1.NodeAffinity'
      description: Describes node affinity scheduling rules for the pod.
    podAffinity:
      $ref: '#/definitions/io.k8s.api.core.v1.PodAffinity'
      description: Describes pod affinity scheduling rules (e.g. co-locate this pod
        in the same node, zone, etc. as some other pod(s)).
    podAntiAffinity:
      $ref: '#/definitions/io.k8s.api.core.v1.PodAntiAffinity'
      description: Describes pod anti-affinity scheduling rules (e.g. avoid putting
        this pod in the same node, zone, etc. as some other pod(s)).
  type: object

```
