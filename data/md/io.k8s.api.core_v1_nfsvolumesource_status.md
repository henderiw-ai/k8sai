# openAPI schema status for io.k8s.api.core.v1.NFSVolumeSource

## description

Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.

## schema

```yaml
|
  description: Represents an NFS mount that lasts the lifetime of a pod. NFS volumes
    do not support ownership management or SELinux relabeling.
  properties:
    path:
      description: 'path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
      type: string
    readOnly:
      description: 'readOnly here will force the NFS export to be mounted with read-only
        permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs'
      type: boolean
    server:
      description: 'server is the hostname or IP address of the NFS server. More info:
        https://kubernetes.io/docs/concepts/storage/volumes#nfs'
      type: string
  required:
  - server
  - path
  type: object

```