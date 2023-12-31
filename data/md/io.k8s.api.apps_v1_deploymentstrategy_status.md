# openAPI schema status for io.k8s.api.apps.v1.DeploymentStrategy

## description

DeploymentStrategy describes how to replace existing pods with new ones.

## schema

```yaml
|
  description: DeploymentStrategy describes how to replace existing pods with new ones.
  properties:
    rollingUpdate:
      $ref: '#/definitions/io.k8s.api.apps.v1.RollingUpdateDeployment'
      description: Rolling update config params. Present only if DeploymentStrategyType
        = RollingUpdate.
    type:
      description: Type of deployment. Can be "Recreate" or "RollingUpdate". Default
        is RollingUpdate.
      type: string
  type: object

```
