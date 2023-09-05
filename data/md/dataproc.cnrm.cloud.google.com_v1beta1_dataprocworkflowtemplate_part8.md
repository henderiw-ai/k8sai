# openAPI schema part8 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      projectRef:
        properties:
          projectRef:
            description: Immutable. The Project that this resource belongs to.
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
                  The project for the resource

                  Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
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
        description: Output only. The time template was created.
        format: date-time
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      placement:
        properties:
          managedCluster:
            properties:
              config:
                properties:
                  endpointConfig:
                    properties:
                      httpPorts:
                        additionalProperties:
                          type: string
                        description: Output only. The map of port descriptions to
                          URLs. Will only be populated if enable_http_port_access
                          is true.
                        type: object
                    type: object
                  lifecycleConfig:
                    properties:
                      idleStartTime:
                        description: Output only. The time when cluster became idle
                          (most recent job finished) and became eligible for deletion
                          due to idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                        format: date-time
                        type: string
                    type: object
                  masterConfig:
                    properties:
                      instanceNames:
                        description: Output only. The list of instance names. Dataproc
                          derives the names from `cluster_name`, `num_instances`,
                          and the instance group.
                        items:
                          type: string
                        type: array
                      isPreemptible:
                        description: Output only. Specifies that this instance group
                          contains preemptible instances.
                        type: boolean
                      managedGroupConfig:
                        description: Output only. The config for Compute Engine Instance
                          Group Manager that manages this group. This is only used
                          for preemptible instance groups.
                        properties:
                          instanceGroupManagerName:
                            description: Output only. The name of the Instance Group
                              Manager for this group.
                            type: string
                          instanceTemplateName:
                            description: Output only. The name of the Instance Template
                              used for the Managed Instance Group.
                            type: string
                        type: object
                    type: object
                  secondaryWorkerConfig:
                    properties:
                      instanceNames:
                        description: Output only. The list of instance names. Dataproc
                          derives the names from `cluster_name`, `num_instances`,
                          and the instance group.
                        items:
                          type: string
                        type: array
                      isPreemptible:
                        description: Output only. Specifies that this instance group
                          contains preemptible instances.
                        type: boolean
                      managedGroupConfig:
                        description: Output only. The config for Compute Engine Instance
                          Group Manager that manages this group. This is only used
                          for preemptible instance groups.
                        properties:
                          instanceGroupManagerName:
                            description: Output only. The name of the Instance Group
                              Manager for this group.
                            type: string
                          instanceTemplateName:
                            description: Output only. The name of the Instance Template
                              used for the Managed Instance Group.
                            type: string
                        type: object
                    type: object
                  workerConfig:
                    properties:
                      instanceNames:
                        description: Output only. The list of instance names. Dataproc
                          derives the names from `cluster_name`, `num_instances`,
                          and the instance group.
                        items:
                          type: string
                        type: array
                      isPreemptible:
                        description: Output only. Specifies that this instance group
                          contains preemptible instances.
                        type: boolean
                      managedGroupConfig:
                        description: Output only. The config for Compute Engine Instance
                          Group Manager that manages this group. This is only used
                          for preemptible instance groups.
                        properties:
                          instanceGroupManagerName:
                            description: Output only. The name of the Instance Group
                              Manager for this group.
                            type: string
                          instanceTemplateName:
                            description: Output only. The name of the Instance Template
                              used for the Managed Instance Group.
                            type: string
                        type: object
                    type: object
                type: object
            type: object
        type: object
      updateTime:
        description: Output only. The time template was last updated.
        format: date-time
        type: string
      version:
        description: Output only. The current version of this workflow template.
        format: int64
        type: integer
    type: object
required:
- spec
type: object

```
