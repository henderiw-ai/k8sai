# openAPI schema status for io.k8s.api.storage.v1.TokenRequest

## description

TokenRequest contains parameters of a service account token.

## schema

```yaml
|
  description: TokenRequest contains parameters of a service account token.
  properties:
    audience:
      description: audience is the intended audience of the token in "TokenRequestSpec".
        It will default to the audiences of kube apiserver.
      type: string
    expirationSeconds:
      description: expirationSeconds is the duration of validity of the token in "TokenRequestSpec".
        It has the same default value of "ExpirationSeconds" in "TokenRequestSpec".
      format: int64
      type: integer
  required:
  - audience
  type: object

```
