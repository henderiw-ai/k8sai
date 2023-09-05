# openAPI schema status for io.k8s.api.batch.v1.UncountedTerminatedPods

## description

UncountedTerminatedPods holds UIDs of Pods that have terminated but haven't been accounted in Job status counters.

## schema

```yaml
|
  description: UncountedTerminatedPods holds UIDs of Pods that have terminated but haven't
    been accounted in Job status counters.
  properties:
    failed:
      description: failed holds UIDs of failed Pods.
      items:
        type: string
      type: array
      x-kubernetes-list-type: set
    succeeded:
      description: succeeded holds UIDs of succeeded Pods.
      items:
        type: string
      type: array
      x-kubernetes-list-type: set
  type: object

```
