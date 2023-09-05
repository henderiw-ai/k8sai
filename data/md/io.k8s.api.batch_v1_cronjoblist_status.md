# openAPI schema status for io.k8s.api.batch.v1.CronJobList

## description

CronJobList is a collection of cron jobs.

## schema

```yaml
|
  description: CronJobList is a collection of cron jobs.
  properties:
    apiVersion:
      description: 'APIVersion defines the versioned schema of this representation of
        an object. Servers should convert recognized schemas to the latest internal
        value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
      type: string
    items:
      description: items is the list of CronJobs.
      items:
        $ref: '#/definitions/io.k8s.api.batch.v1.CronJob'
      type: array
    kind:
      description: 'Kind is a string value representing the REST resource this object
        represents. Servers may infer this from the endpoint the client submits requests
        to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
      type: string
    metadata:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ListMeta'
      description: 'Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
  required:
  - items
  type: object
  x-kubernetes-group-version-kind:
  - group: batch
    kind: CronJobList
    version: v1

```