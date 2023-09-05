# openAPI schema status for io.k8s.api.flowcontrol.v1beta2.GroupSubject

## description

GroupSubject holds detailed information for group-kind subject.

## schema

```yaml
|
  description: GroupSubject holds detailed information for group-kind subject.
  properties:
    name:
      description: name is the user group that matches, or "*" to match all user groups.
        See https://github.com/kubernetes/apiserver/blob/master/pkg/authentication/user/user.go
        for some well-known group names. Required.
      type: string
  required:
  - name
  type: object

```
