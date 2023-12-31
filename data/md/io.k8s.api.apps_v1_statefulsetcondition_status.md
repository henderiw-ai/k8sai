# openAPI schema status for io.k8s.api.apps.v1.StatefulSetCondition

## description

StatefulSetCondition describes the state of a statefulset at a certain point.

## schema

```yaml
|
  description: StatefulSetCondition describes the state of a statefulset at a certain
    point.
  properties:
    lastTransitionTime:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: Last time the condition transitioned from one status to another.
    message:
      description: A human readable message indicating details about the transition.
      type: string
    reason:
      description: The reason for the condition's last transition.
      type: string
    status:
      description: Status of the condition, one of True, False, Unknown.
      type: string
    type:
      description: Type of statefulset condition.
      type: string
  required:
  - type
  - status
  type: object

```
