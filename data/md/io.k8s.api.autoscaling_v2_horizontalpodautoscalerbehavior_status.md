# openAPI schema status for io.k8s.api.autoscaling.v2.HorizontalPodAutoscalerBehavior

## description

HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).

## schema

```yaml
|
  description: HorizontalPodAutoscalerBehavior configures the scaling behavior of the
    target in both Up and Down directions (scaleUp and scaleDown fields respectively).
  properties:
    scaleDown:
      $ref: '#/definitions/io.k8s.api.autoscaling.v2.HPAScalingRules'
      description: scaleDown is scaling policy for scaling Down. If not set, the default
        value is to allow to scale down to minReplicas pods, with a 300 second stabilization
        window (i.e., the highest recommendation for the last 300sec is used).
    scaleUp:
      $ref: '#/definitions/io.k8s.api.autoscaling.v2.HPAScalingRules'
      description: |-
        scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of:
          * increase no more than 4 pods per 60 seconds
          * double the number of pods per 60 seconds
        No stabilization is used.
  type: object

```