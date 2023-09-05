# openAPI schema status for io.k8s.api.authorization.v1.NonResourceAttributes

## description

NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface

## schema

```yaml
|
  description: NonResourceAttributes includes the authorization attributes available
    for non-resource requests to the Authorizer interface
  properties:
    path:
      description: Path is the URL path of the request
      type: string
    verb:
      description: Verb is the standard HTTP verb
      type: string
  type: object

```