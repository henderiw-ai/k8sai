# openAPI schema status for io.k8s.api.batch.v1.PodFailurePolicy

## description

PodFailurePolicy describes how failed pods influence the backoffLimit.

## schema

```yaml
|
  description: PodFailurePolicy describes how failed pods influence the backoffLimit.
  properties:
    rules:
      description: A list of pod failure policy rules. The rules are evaluated in order.
        Once a rule matches a Pod failure, the remaining of the rules are ignored. When
        no rule matches the Pod failure, the default handling applies - the counter
        of pod failures is incremented and it is checked against the backoffLimit. At
        most 20 elements are allowed.
      items:
        $ref: '#/definitions/io.k8s.api.batch.v1.PodFailurePolicyRule'
      type: array
      x-kubernetes-list-type: atomic
  required:
  - rules
  type: object

```