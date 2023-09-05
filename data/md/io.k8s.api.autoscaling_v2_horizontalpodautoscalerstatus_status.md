# openAPI schema status for io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerStatus

## description

HorizontalPodAutoscalerStatus describes the current status of a horizontal pod autoscaler.

## schema

```yaml
|
  description: HorizontalPodAutoscalerStatus describes the current status of a horizontal
    pod autoscaler.
  properties:
    conditions:
      description: conditions is the set of conditions required for this autoscaler
        to scale its target, and indicates whether or not those conditions are met.
      items:
        $ref: '#/definitions/io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerCondition'
      type: array
      x-kubernetes-list-map-keys:
      - type
      x-kubernetes-list-type: map
      x-kubernetes-patch-merge-key: type
      x-kubernetes-patch-strategy: merge
    currentMetrics:
      description: currentMetrics is the last read state of the metrics used by this
        autoscaler.
      items:
        $ref: '#/definitions/io.k8s.api.autoscaling.v2.MetricStatus'
      type: array
      x-kubernetes-list-type: atomic
    currentReplicas:
      description: currentReplicas is current number of replicas of pods managed by
        this autoscaler, as last seen by the autoscaler.
      format: int32
      type: integer
    desiredReplicas:
      description: desiredReplicas is the desired number of replicas of pods managed
        by this autoscaler, as last calculated by the autoscaler.
      format: int32
      type: integer
    lastScaleTime:
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.Time'
      description: lastScaleTime is the last time the HorizontalPodAutoscaler scaled
        the number of pods, used by the autoscaler to control how often the number of
        pods is changed.
    observedGeneration:
      description: observedGeneration is the most recent generation observed by this
        autoscaler.
      format: int64
      type: integer
  required:
  - desiredReplicas
  type: object

```