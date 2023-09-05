# openAPI schema status for io.k8s.api.admissionregistration.v1.RuleWithOperations

## description

RuleWithOperations is a tuple of Operations and Resources. It is recommended to make sure that all the tuple expansions are valid.

## schema

```yaml
|
  description: RuleWithOperations is a tuple of Operations and Resources. It is recommended
    to make sure that all the tuple expansions are valid.
  properties:
    apiGroups:
      description: APIGroups is the API groups the resources belong to. '*' is all groups.
        If '*' is present, the length of the slice must be one. Required.
      items:
        type: string
      type: array
      x-kubernetes-list-type: atomic
    apiVersions:
      description: APIVersions is the API versions the resources belong to. '*' is all
        versions. If '*' is present, the length of the slice must be one. Required.
      items:
        type: string
      type: array
      x-kubernetes-list-type: atomic
    operations:
      description: Operations is the operations the admission hook cares about - CREATE,
        UPDATE, DELETE, CONNECT or * for all of those operations and any future admission
        operations that are added. If '*' is present, the length of the slice must be
        one. Required.
      items:
        type: string
      type: array
      x-kubernetes-list-type: atomic
    resources:
      description: |-
        Resources is a list of resources this rule applies to.

        For example: 'pods' means pods. 'pods/log' means the log subresource of pods. '*' means all resources, but not subresources. 'pods/*' means all subresources of pods. '*/scale' means all scale subresources. '*/*' means all resources and their subresources.

        If wildcard is present, the validation rule will ensure resources do not overlap with each other.

        Depending on the enclosing object, subresources might not be allowed. Required.
      items:
        type: string
      type: array
      x-kubernetes-list-type: atomic
    scope:
      description: scope specifies the scope of this rule. Valid values are "Cluster",
        "Namespaced", and "*" "Cluster" means that only cluster-scoped resources will
        match this rule. Namespace API objects are cluster-scoped. "Namespaced" means
        that only namespaced resources will match this rule. "*" means that there are
        no scope restrictions. Subresources match the scope of their parent resource.
        Default is "*".
      type: string
  type: object

```