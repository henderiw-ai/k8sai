# openAPI schema part2 for container.cnrm.cloud.google.com.v1beta1.ContainerNodePool

## schema

```yaml
properties:
  spec:
    properties:
      nodeConfig:
        properties:
          nodeConfig:
            description: Immutable. The configuration of the nodepool.
            properties:
              advancedMachineFeatures:
                description: Immutable. Specifies options for controlling advanced
                  machine features.
                properties:
                  threadsPerCore:
                    description: Immutable. The number of threads per physical core.
                      To disable simultaneous multithreading (SMT) set this to 1.
                      If unset, the maximum number of threads supported per core by
                      the underlying processor is assumed.
                    type: integer
                required:
                - threadsPerCore
                type: object
              bootDiskKMSCryptoKeyRef:
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
                    description: 'Allowed value: The `selfLink` field of a `KMSCryptoKey`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              diskSizeGb:
                description: Immutable. Size of the disk attached to each node, specified
                  in GB. The smallest allowed disk size is 10GB.
                type: integer
              diskType:
                description: Immutable. Type of the disk attached to each node. Such
                  as pd-standard, pd-balanced or pd-ssd.
                type: string
              ephemeralStorageConfig:
                description: Immutable. Parameters for the ephemeral storage filesystem.
                  If unspecified, ephemeral storage is backed by the boot disk.
                properties:
                  localSsdCount:
                    description: Immutable. Number of local SSDs to use to back ephemeral
                      storage. Uses NVMe interfaces. Each local SSD must be 375 or
                      3000 GB in size, and all local SSDs must share the same size.
                    type: integer
                required:
                - localSsdCount
                type: object
              ephemeralStorageLocalSsdConfig:
                description: Immutable. Parameters for the ephemeral storage filesystem.
                  If unspecified, ephemeral storage is backed by the boot disk.
                properties:
                  localSsdCount:
                    description: Immutable. Number of local SSDs to use to back ephemeral
                      storage. Uses NVMe interfaces. Each local SSD must be 375 or
                      3000 GB in size, and all local SSDs must share the same size.
                    type: integer
                required:
                - localSsdCount
                type: object
              gcfsConfig:
                description: Immutable. GCFS configuration for this node.
                properties:
                  enabled:
                    description: Immutable. Whether or not GCFS is enabled.
                    type: boolean
                required:
                - enabled
                type: object
              guestAccelerator:
                description: Immutable. List of the type and count of accelerator
                  cards attached to the instance.
                items:
                  properties:
                    count:
                      description: Immutable. The number of the accelerator cards
                        exposed to an instance.
                      type: integer
                    gpuDriverInstallationConfig:
                      description: Immutable. Configuration for auto installation
                        of GPU driver.
                      properties:
                        gpuDriverVersion:
                          description: Immutable. Mode for how the GPU driver is installed.
                          type: string
                      required:
                      - gpuDriverVersion
                      type: object
                    gpuPartitionSize:
                      description: Immutable. Size of partitions to create on the
                        GPU. Valid values are described in the NVIDIA mig user guide
                        (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
                      type: string
                    gpuSharingConfig:
                      description: Immutable. Configuration for GPU sharing.
                      properties:
                        gpuSharingStrategy:
                          description: Immutable. The type of GPU sharing strategy
                            to enable on the GPU node. Possible values are described
                            in the API package (https://pkg.go.dev/google.golang.org/api/container/v1#GPUSharingConfig).
                          type: string
                        maxSharedClientsPerGpu:
                          description: Immutable. The maximum number of containers
                            that can share a GPU.
                          type: integer
                      required:
                      - gpuSharingStrategy
                      - maxSharedClientsPerGpu
                      type: object
                    type:
                      description: Immutable. The accelerator type resource name.
                      type: string
                  required:
                  - count
                  - type
                  type: object
                type: array
              gvnic:
                description: Immutable. Enable or disable gvnic in the node pool.
                properties:
                  enabled:
                    description: Immutable. Whether or not gvnic is enabled.
                    type: boolean
                required:
                - enabled
                type: object
              imageType:
                description: The image type to use for this node. Note that for a
                  given image type, the latest version of it will be used.
                type: string
              kubeletConfig:
                description: Node kubelet configs.
                properties:
                  cpuCfsQuota:
                    description: Enable CPU CFS quota enforcement for containers that
                      specify CPU limits.
                    type: boolean
                  cpuCfsQuotaPeriod:
                    description: Set the CPU CFS quota period value 'cpu.cfs_period_us'.
                    type: string
                  cpuManagerPolicy:
                    description: Control the CPU management policy on the node.
                    type: string
                  podPidsLimit:
                    description: Controls the maximum number of processes allowed
                      to run in a pod.
                    type: integer
                required:
                - cpuManagerPolicy
                type: object
              labels:
                additionalProperties:
                  type: string
                description: The map of Kubernetes labels (key/value pairs) to be
                  applied to each node. These will added in addition to any default
                  label(s) that Kubernetes may apply to the node.
                type: object
              linuxNodeConfig:
                description: Parameters that can be configured on Linux nodes.
                properties:
                  sysctls:
                    additionalProperties:
                      type: string
                    description: The Linux kernel parameters to be applied to the
                      nodes and all pods running on the nodes.
                    type: object
                required:
                - sysctls
                type: object
              localNvmeSsdBlockConfig:
                description: Immutable. Parameters for raw-block local NVMe SSDs.
                properties:
                  localSsdCount:
                    description: Immutable. Number of raw-block local NVMe SSD disks
                      to be attached to the node. Each local SSD is 375 GB in size.
                    type: integer
                required:
                - localSsdCount
                type: object
              localSsdCount:
                description: Immutable. The number of local SSD disks to be attached
                  to the node.
                type: integer
              loggingVariant:
                description: Type of logging agent that is used as the default value
                  for node pools in the cluster. Valid values include DEFAULT and
                  MAX_THROUGHPUT.
                type: string
              machineType:
                description: Immutable. The name of a Google Compute Engine machine
                  type.
                type: string
              metadata:
                additionalProperties:
                  type: string
                description: Immutable. The metadata key/value pairs assigned to instances
                  in the cluster.
                type: object
              minCpuPlatform:
                description: Immutable. Minimum CPU platform to be used by this instance.
                  The instance may be scheduled on the specified or newer CPU platform.
                type: string
              nodeGroupRef:
                description: |-
                  Immutable. Setting this field will assign instances
                  of this pool to run on the specified node group. This is useful
                  for running workloads on sole tenant nodes.
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
                    description: 'Allowed value: The `name` field of a `ComputeNodeGroup`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              oauthScopes:
                description: Immutable. The set of Google API scopes to be made available
                  on all of the node VMs.
                items:
                  type: string
                type: array
              preemptible:
                description: Immutable. Whether the nodes are created as preemptible
                  VM instances.
                type: boolean
              reservationAffinity:
                description: Immutable. The reservation affinity configuration for
                  the node pool.
                properties:
                  consumeReservationType:
                    description: Immutable. Corresponds to the type of reservation
                      consumption.
                    type: string
                  key:
                    description: Immutable. The label key of a reservation resource.
                    type: string
                  values:
                    description: Immutable. The label values of the reservation resource.
                    items:
                      type: string
                    type: array
                required:
                - consumeReservationType
                type: object
              resourceLabels:
                additionalProperties:
                  type: string
                description: The GCE resource labels (a map of key/value pairs) to
                  be applied to the node pool.
                type: object
              sandboxConfig:
                description: Immutable. Sandbox configuration for this node.
                properties:
                  sandboxType:
                    description: Type of the sandbox to use for the node (e.g. 'gvisor').
                    type: string
                required:
                - sandboxType
                type: object
              serviceAccountRef:
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
                    description: 'Allowed value: The `email` field of an `IAMServiceAccount`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              shieldedInstanceConfig:
                description: Immutable. Shielded Instance options.
                properties:
                  enableIntegrityMonitoring:
                    description: Immutable. Defines whether the instance has integrity
                      monitoring enabled.
                    type: boolean
                  enableSecureBoot:
                    description: Immutable. Defines whether the instance has Secure
                      Boot enabled.
                    type: boolean
                type: object
              soleTenantConfig:
                description: Immutable. Node affinity options for sole tenant node
                  pools.
                properties:
                  nodeAffinity:
                    description: Immutable. .
                    items:
                      properties:
                        key:
                          description: Immutable. .
                          type: string
                        operator:
                          description: Immutable. .
                          type: string
                        values:
                          description: Immutable. .
                          items:
                            type: string
                          type: array
                      required:
                      - key
                      - operator
                      - values
                      type: object
                    type: array
                required:
                - nodeAffinity
                type: object
              spot:
                description: Immutable. Whether the nodes are created as spot VM instances.
                type: boolean
              tags:
                description: The list of instance tags applied to all nodes.
                items:
                  type: string
                type: array
              taint:
                description: Immutable. List of Kubernetes taints to be applied to
                  each node.
                items:
                  properties:
                    effect:
                      description: Immutable. Effect for taint.
                      type: string
                    key:
                      description: Immutable. Key for taint.
                      type: string
                    value:
                      description: Immutable. Value for taint.
                      type: string
                  required:
                  - effect
                  - key
                  - value
                  type: object
                type: array
              workloadMetadataConfig:
                description: The workload metadata configuration for this node.
                properties:
                  mode:
                    description: Mode is the configuration for how to expose metadata
                      to workloads running on the node.
                    type: string
                  nodeMetadata:
                    description: DEPRECATED. Deprecated in favor of mode. NodeMetadata
                      is the configuration for how to expose metadata to the workloads
                      running on the node.
                    type: string
                type: object
            type: object
      nodeCount:
        description: The number of nodes per instance group. This field can be used
          to update the number of nodes per instance group but should not be used
          alongside autoscaling.
        type: integer
      nodeLocations:
        description: The list of zones in which the node pool's nodes should be located.
          Nodes must be in the region of their regional cluster or in the same region
          as their cluster's zone for zonal clusters. If unspecified, the cluster-level
          node_locations will be used.
        items:
          type: string
        type: array
      placementPolicy:
        description: Immutable. Specifies the node placement policy.
        properties:
          tpuTopology:
            description: TPU placement topology for pod slice node pool. https://cloud.google.com/tpu/docs/types-topologies#tpu_topologies.
            type: string
          type:
            description: Type defines the type of placement policy.
            type: string
        required:
        - type
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      upgradeSettings:
        description: Specify node upgrade settings to change how many nodes GKE attempts
          to upgrade at once. The number of nodes upgraded simultaneously is the sum
          of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously
          is limited to 20.
        properties:
          blueGreenSettings:
            description: Settings for BlueGreen node pool upgrade.
            properties:
              nodePoolSoakDuration:
                description: Time needed after draining entire blue pool. After this
                  period, blue pool will be cleaned up.
                type: string
              standardRolloutPolicy:
                description: Standard rollout policy is the default policy for blue-green.
                properties:
                  batchNodeCount:
                    description: Number of blue nodes to drain in a batch.
                    type: integer
                  batchPercentage:
                    description: Percentage of the blue pool nodes to drain in a batch.
                    type: number
                  batchSoakDuration:
                    description: Soak time after each batch gets drained.
                    type: string
                type: object
            required:
            - standardRolloutPolicy
            type: object
          maxSurge:
            description: The number of additional nodes that can be added to the node
              pool during an upgrade. Increasing max_surge raises the number of nodes
              that can be upgraded simultaneously. Can be set to 0 or greater.
            type: integer
          maxUnavailable:
            description: The number of nodes that can be simultaneously unavailable
              during an upgrade. Increasing max_unavailable raises the number of nodes
              that can be upgraded in parallel. Can be set to 0 or greater.
            type: integer
          strategy:
            description: Update strategy for the given nodepool.
            type: string
        type: object
      version:
        type: string
required:
- spec
type: object

```
