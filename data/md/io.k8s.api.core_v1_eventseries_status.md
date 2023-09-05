# openAPI schema status for io.k8s.api.core.v1.EventSeries

## description

EventSeries contain information on series of events, i.e. thing that was/is happening continuously for some time.

## schema

```yaml
|
  description: EventSeries contain information on series of events, i.e. thing that
    was/is happening continuously for some time.
  properties:
    count:
      description: Number of occurrences in this series up to the last heartbeat time
      format: int32
      type: integer
    lastObservedTime:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.MicroTime'
      description: Time of the last occurrence observed
  type: object

```