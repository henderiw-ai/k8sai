# openAPI schema part2 for monitoring.cnrm.cloud.google.com.v1beta1.MonitoringServiceLevelObjective

## schema

```yaml
properties:
  spec:
    properties:
      serviceRef:
        properties:
          serviceRef:
            description: Immutable.
            oneOf:
            - not:
                required:
                - external
              required:
              - name
            - not:
                anyOf:
                - required:
                  - name
                - required:
                  - namespace
              required:
              - external
            properties:
              external:
                description: |-
                  The service for the resource

                  Allowed value: The Google Cloud resource name of a `MonitoringService` resource (format: `projects/{{project}}/services/{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
  status:
    properties:
      conditions:
        description: Conditions represent the latest available observation of the
          resource's current state.
        items:
          properties:
            lastTransitionTime:
              description: Last time the condition transitioned from one status to
                another.
              type: string
            message:
              description: Human-readable message indicating details about last transition.
              type: string
            reason:
              description: Unique, one-word, CamelCase reason for the condition's
                last transition.
              type: string
            status:
              description: Status is the status of the condition. Can be True, False,
                Unknown.
              type: string
            type:
              description: Type is the type of the condition.
              type: string
          type: object
        type: array
      createTime:
        description: Time stamp of the `Create` or most recent `Update` command on
          this `Slo`.
        format: date-time
        type: string
      deleteTime:
        description: Time stamp of the `Update` or `Delete` command that made this
          no longer a current `Slo`. This field is not populated in `ServiceLevelObjective`s
          returned from calls to `GetServiceLevelObjective` and `ListServiceLevelObjectives`,
          because it is always empty in the current version. It is populated in `ServiceLevelObjective`s
          representing previous versions in the output of `ListServiceLevelObjectiveVersions`.
          Because all old configuration versions are stored, `Update` operations mark
          the obsoleted version as deleted.
        format: date-time
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      serviceManagementOwned:
        description: Output only. If set, this SLO is managed at the [Service Management](https://cloud.google.com/service-management/overview)
          level. Therefore the service yaml file is the source of truth for this SLO,
          and API `Update` and `Delete` operations are forbidden.
        type: boolean
    type: object
required:
- spec
type: object

```
