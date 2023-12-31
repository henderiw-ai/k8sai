# openAPI schema status for io.k8s.api.core.v1.PodCondition

## description

PodCondition contains details for the current condition of this pod.

## schema

```yaml
|
  description: PodCondition contains details for the current condition of this pod.
  properties:
    lastProbeTime:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: Last time we probed the condition.
    lastTransitionTime:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: Last time the condition transitioned from one status to another.
    message:
      description: Human-readable message indicating details about last transition.
      type: string
    reason:
      description: Unique, one-word, CamelCase reason for the condition's last transition.
      type: string
    status:
      description: 'Status is the status of the condition. Can be True, False, Unknown.
        More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions'
      type: string
    type:
      description: 'Type is the type of the condition. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#pod-conditions'
      type: string
  required:
  - type
  - status
  type: object

```
