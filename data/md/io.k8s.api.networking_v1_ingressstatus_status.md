# openAPI schema status for io.k8s.api.networking.v1.IngressStatus

## description

IngressStatus describe the current state of the Ingress.

## schema

```yaml
|
  description: IngressStatus describe the current state of the Ingress.
  properties:
    loadBalancer:
      $ref: '#/definitions/io.k8s.api.networking.v1.IngressLoadBalancerStatus'
      description: loadBalancer contains the current status of the load-balancer.
  type: object

```
