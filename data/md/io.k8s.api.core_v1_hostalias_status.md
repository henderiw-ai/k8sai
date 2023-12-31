# openAPI schema status for io.k8s.api.core.v1.HostAlias

## description

HostAlias holds the mapping between IP and hostnames that will be injected as an entry in the pod's hosts file.

## schema

```yaml
|
  description: HostAlias holds the mapping between IP and hostnames that will be injected
    as an entry in the pod's hosts file.
  properties:
    hostnames:
      description: Hostnames for the above IP address.
      items:
        type: string
      type: array
    ip:
      description: IP address of the host file entry.
      type: string
  type: object

```
