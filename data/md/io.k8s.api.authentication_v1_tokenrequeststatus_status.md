# openAPI schema status for io.k8s.api.authentication.v1.TokenRequestStatus

## description

TokenRequestStatus is the result of a token request.

## schema

```yaml
|
  description: TokenRequestStatus is the result of a token request.
  properties:
    expirationTimestamp:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: ExpirationTimestamp is the time of expiration of the returned token.
    token:
      description: Token is the opaque bearer token.
      type: string
  required:
  - token
  - expirationTimestamp
  type: object

```
