# openAPI schema status for io.k8s.api.apps.v1.DeploymentStatus

## description

DeploymentStatus is the most recently observed status of the Deployment.

## schema

```yaml
|
  description: DeploymentStatus is the most recently observed status of the Deployment.
  properties:
    availableReplicas:
      description: Total number of available pods (ready for at least minReadySeconds)
        targeted by this deployment.
      format: int32
      type: integer
    collisionCount:
      description: Count of hash collisions for the Deployment. The Deployment controller
        uses this field as a collision avoidance mechanism when it needs to create the
        name for the newest ReplicaSet.
      format: int32
      type: integer
    conditions:
      description: Represents the latest available observations of a deployment's current
        state.
      items:
        $ref: '#/definitions/io.k8s.api.apps.v1.DeploymentCondition'
      type: array
      x-kubernetes-patch-merge-key: type
      x-kubernetes-patch-strategy: merge
    observedGeneration:
      description: The generation observed by the deployment controller.
      format: int64
      type: integer
    readyReplicas:
      description: readyReplicas is the number of pods targeted by this Deployment with
        a Ready Condition.
      format: int32
      type: integer
    replicas:
      description: Total number of non-terminated pods targeted by this deployment (their
        labels match the selector).
      format: int32
      type: integer
    unavailableReplicas:
      description: Total number of unavailable pods targeted by this deployment. This
        is the total number of pods that are still required for the deployment to have
        100% available capacity. They may either be pods that are running but not yet
        available or pods that still have not been created.
      format: int32
      type: integer
    updatedReplicas:
      description: Total number of non-terminated pods targeted by this deployment that
        have the desired template spec.
      format: int32
      type: integer
  type: object

```