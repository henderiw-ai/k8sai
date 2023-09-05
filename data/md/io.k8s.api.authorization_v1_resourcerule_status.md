# openAPI schema status for io.k8s.api.authorization.v1.ResourceRule

## description

ResourceRule is the list of actions the subject is allowed to perform on resources. The list ordering isn't significant, may contain duplicates, and possibly be incomplete.

## schema

```yaml
|
  description: ResourceRule is the list of actions the subject is allowed to perform
    on resources. The list ordering isn't significant, may contain duplicates, and possibly
    be incomplete.
  properties:
    apiGroups:
      description: APIGroups is the name of the APIGroup that contains the resources.  If
        multiple API groups are specified, any action requested against one of the enumerated
        resources in any API group will be allowed.  "*" means all.
      items:
        type: string
      type: array
    resourceNames:
      description: ResourceNames is an optional white list of names that the rule applies
        to.  An empty set means that everything is allowed.  "*" means all.
      items:
        type: string
      type: array
    resources:
      description: |-
        Resources is a list of resources this rule applies to.  "*" means all in the specified apiGroups.
         "*/foo" represents the subresource 'foo' for all resources in the specified apiGroups.
      items:
        type: string
      type: array
    verbs:
      description: 'Verb is a list of kubernetes resource API verbs, like: get, list,
        watch, create, update, delete, proxy.  "*" means all.'
      items:
        type: string
      type: array
  required:
  - verbs
  type: object

```