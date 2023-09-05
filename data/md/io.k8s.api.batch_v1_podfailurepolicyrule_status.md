# openAPI schema status for io.k8s.api.batch.v1.PodFailurePolicyRule

## description

PodFailurePolicyRule describes how a pod failure is handled when the requirements are met. One of onExitCodes and onPodConditions, but not both, can be used in each rule.

## schema

```yaml
|
  description: PodFailurePolicyRule describes how a pod failure is handled when the
    requirements are met. One of onExitCodes and onPodConditions, but not both, can
    be used in each rule.
  properties:
    action:
      description: |-
        Specifies the action taken on a pod failure when the requirements are satisfied. Possible values are:

        - FailJob: indicates that the pod's job is marked as Failed and all
          running pods are terminated.
        - FailIndex: indicates that the pod's index is marked as Failed and will
          not be restarted.
          This value is alpha-level. It can be used when the
          `JobBackoffLimitPerIndex` feature gate is enabled (disabled by default).
        - Ignore: indicates that the counter towards the .backoffLimit is not
          incremented and a replacement pod is created.
        - Count: indicates that the pod is handled in the default way - the
          counter towards the .backoffLimit is incremented.
        Additional values are considered to be added in the future. Clients should react to an unknown action by skipping the rule.
      type: string
    onExitCodes:
      $ref: '#/definitions/io.k8s.api.batch.v1.PodFailurePolicyOnExitCodesRequirement'
      description: Represents the requirement on the container exit codes.
    onPodConditions:
      description: Represents the requirement on the pod conditions. The requirement
        is represented as a list of pod condition patterns. The requirement is satisfied
        if at least one pattern matches an actual pod condition. At most 20 elements
        are allowed.
      items:
        $ref: '#/definitions/io.k8s.api.batch.v1.PodFailurePolicyOnPodConditionsPattern'
      type: array
      x-kubernetes-list-type: atomic
  required:
  - action
  type: object

```
