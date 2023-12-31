# openAPI schema status for io.k8s.api.core.v1.LoadBalancerStatus

## description

LoadBalancerStatus represents the status of a load-balancer.

## schema

```yaml
|
  description: LoadBalancerStatus represents the status of a load-balancer.
  properties:
    ingress:
      description: Ingress is a list containing ingress points for the load-balancer.
        Traffic intended for the service should be sent to these ingress points.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.LoadBalancerIngress'
      type: array
  type: object

```
