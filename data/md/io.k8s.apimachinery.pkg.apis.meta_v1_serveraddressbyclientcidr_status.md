# openAPI schema status for io.k8s.apimachinery.pkg.apis.meta.v1.ServerAddressByClientCIDR

## description

ServerAddressByClientCIDR helps the client to determine the server address that they should use, depending on the clientCIDR that they match.

## schema

```yaml
|
  description: ServerAddressByClientCIDR helps the client to determine the server address
    that they should use, depending on the clientCIDR that they match.
  properties:
    clientCIDR:
      description: The CIDR with which clients can match their IP to figure out the
        server address that they should use.
      type: string
    serverAddress:
      description: Address of this server, suitable for a client that matches the above
        CIDR. This can be a hostname, hostname:port, IP or IP:port.
      type: string
  required:
  - clientCIDR
  - serverAddress
  type: object

```
