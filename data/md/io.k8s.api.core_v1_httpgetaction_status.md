# openAPI schema status for io.k8s.api.core.v1.HTTPGetAction

## description

HTTPGetAction describes an action based on HTTP Get requests.

## schema

```yaml
|
  description: HTTPGetAction describes an action based on HTTP Get requests.
  properties:
    host:
      description: Host name to connect to, defaults to the pod IP. You probably want
        to set "Host" in httpHeaders instead.
      type: string
    httpHeaders:
      description: Custom headers to set in the request. HTTP allows repeated headers.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.HTTPHeader'
      type: array
    path:
      description: Path to access on the HTTP server.
      type: string
    port:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.util.intstr.IntOrString'
      description: Name or number of the port to access on the container. Number must
        be in the range 1 to 65535. Name must be an IANA_SVC_NAME.
    scheme:
      description: Scheme to use for connecting to the host. Defaults to HTTP.
      type: string
  required:
  - port
  type: object

```