# openAPI schema part7 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocCluster

## schema

```yaml
properties:
  status:
    properties:
      status:
        properties:
          clusterUuid:
            description: Output only. A cluster UUID (Unique Universal Identifier).
              Dataproc generates this value when it creates the cluster.
            type: string
          conditions:
            description: Conditions represent the latest available observation of
              the resource's current state.
            items:
              properties:
                lastTransitionTime:
                  description: Last time the condition transitioned from one status
                    to another.
                  type: string
                message:
                  description: Human-readable message indicating details about last
                    transition.
                  type: string
                reason:
                  description: Unique, one-word, CamelCase reason for the condition's
                    last transition.
                  type: string
                status:
                  description: Status is the status of the condition. Can be True,
                    False, Unknown.
                  type: string
                type:
                  description: Type is the type of the condition.
                  type: string
              type: object
            type: array
          config:
            properties:
              endpointConfig:
                properties:
                  httpPorts:
                    additionalProperties:
                      type: string
                    description: Output only. The map of port descriptions to URLs.
                      Will only be populated if enable_http_port_access is true.
                    type: object
                type: object
              lifecycleConfig:
                properties:
                  idleStartTime:
                    description: Output only. The time when cluster became idle (most
                      recent job finished) and became eligible for deletion due to
                      idleness (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                    format: date-time
                    type: string
                type: object
              masterConfig:
                properties:
                  instanceNames:
                    description: Output only. The list of instance names. Dataproc
                      derives the names from `cluster_name`, `num_instances`, and
                      the instance group.
                    items:
                      type: string
                    type: array
                  instanceReferences:
                    description: Output only. List of references to Compute Engine
                      instances.
                    items:
                      properties:
                        instanceId:
                          description: The unique identifier of the Compute Engine
                            instance.
                          type: string
                        instanceName:
                          description: The user-friendly name of the Compute Engine
                            instance.
                          type: string
                        publicEciesKey:
                          description: The public ECIES key used for sharing data
                            with this instance.
                          type: string
                        publicKey:
                          description: The public RSA key used for sharing data with
                            this instance.
                          type: string
                      type: object
                    type: array
                  isPreemptible:
                    description: Output only. Specifies that this instance group contains
                      preemptible instances.
                    type: boolean
                  managedGroupConfig:
                    description: Output only. The config for Compute Engine Instance
                      Group Manager that manages this group. This is only used for
                      preemptible instance groups.
                    properties:
                      instanceGroupManagerName:
                        description: Output only. The name of the Instance Group Manager
                          for this group.
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
                      derives the names from `cluster_name`, `num_instances`, and
                      the instance group.
                    items:
                      type: string
                    type: array
                  instanceReferences:
                    description: Output only. List of references to Compute Engine
                      instances.
                    items:
                      properties:
                        instanceId:
                          description: The unique identifier of the Compute Engine
                            instance.
                          type: string
                        instanceName:
                          description: The user-friendly name of the Compute Engine
                            instance.
                          type: string
                        publicEciesKey:
                          description: The public ECIES key used for sharing data
                            with this instance.
                          type: string
                        publicKey:
                          description: The public RSA key used for sharing data with
                            this instance.
                          type: string
                      type: object
                    type: array
                  isPreemptible:
                    description: Output only. Specifies that this instance group contains
                      preemptible instances.
                    type: boolean
                  managedGroupConfig:
                    description: Output only. The config for Compute Engine Instance
                      Group Manager that manages this group. This is only used for
                      preemptible instance groups.
                    properties:
                      instanceGroupManagerName:
                        description: Output only. The name of the Instance Group Manager
                          for this group.
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
                      derives the names from `cluster_name`, `num_instances`, and
                      the instance group.
                    items:
                      type: string
                    type: array
                  instanceReferences:
                    description: Output only. List of references to Compute Engine
                      instances.
                    items:
                      properties:
                        instanceId:
                          description: The unique identifier of the Compute Engine
                            instance.
                          type: string
                        instanceName:
                          description: The user-friendly name of the Compute Engine
                            instance.
                          type: string
                        publicEciesKey:
                          description: The public ECIES key used for sharing data
                            with this instance.
                          type: string
                        publicKey:
                          description: The public RSA key used for sharing data with
                            this instance.
                          type: string
                      type: object
                    type: array
                  isPreemptible:
                    description: Output only. Specifies that this instance group contains
                      preemptible instances.
                    type: boolean
                  managedGroupConfig:
                    description: Output only. The config for Compute Engine Instance
                      Group Manager that manages this group. This is only used for
                      preemptible instance groups.
                    properties:
                      instanceGroupManagerName:
                        description: Output only. The name of the Instance Group Manager
                          for this group.
                        type: string
                      instanceTemplateName:
                        description: Output only. The name of the Instance Template
                          used for the Managed Instance Group.
                        type: string
                    type: object
                type: object
            type: object
          metrics:
            description: 'Output only. Contains cluster daemon metrics such as HDFS
              and YARN stats. **Beta Feature**: This report is available for testing
              purposes only. It may be changed before final release.'
            properties:
              hdfsMetrics:
                additionalProperties:
                  type: string
                description: The HDFS metrics.
                type: object
              yarnMetrics:
                additionalProperties:
                  type: string
                description: The YARN metrics.
                type: object
            type: object
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          status:
            description: Output only. Cluster status.
            properties:
              detail:
                description: Optional. Output only. Details of cluster's state.
                type: string
              state:
                description: 'Output only. The cluster''s state. Possible values:
                  UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING,
                  STOPPED, STARTING'
                type: string
              stateStartTime:
                description: Output only. Time when this state was entered (see JSON
                  representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                format: date-time
                type: string
              substate:
                description: 'Output only. Additional state information that includes
                  status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY,
                  STALE_STATUS'
                type: string
            type: object
          statusHistory:
            description: Output only. The previous cluster status.
            items:
              properties:
                detail:
                  description: Optional. Output only. Details of cluster's state.
                  type: string
                state:
                  description: 'Output only. The cluster''s state. Possible values:
                    UNKNOWN, CREATING, RUNNING, ERROR, DELETING, UPDATING, STOPPING,
                    STOPPED, STARTING'
                  type: string
                stateStartTime:
                  description: Output only. Time when this state was entered (see
                    JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                  format: date-time
                  type: string
                substate:
                  description: 'Output only. Additional state information that includes
                    status reported by the agent. Possible values: UNSPECIFIED, UNHEALTHY,
                    STALE_STATUS'
                  type: string
              type: object
            type: array
        type: object
required:
- spec
type: object

```
