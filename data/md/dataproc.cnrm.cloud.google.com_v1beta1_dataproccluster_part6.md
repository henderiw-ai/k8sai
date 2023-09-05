# openAPI schema part6 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocCluster

## schema

```yaml
properties:
  spec:
    properties:
      virtualClusterConfig:
        properties:
          virtualClusterConfig:
            description: Immutable. Optional. The virtual cluster config is used when
              creating a Dataproc cluster that does not directly control the underlying
              compute resources, for example, when creating a [Dataproc-on-GKE cluster](https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke).
              Dataproc may set default values, and values may change when clusters
              are updated. Exactly one of config or virtual_cluster_config must be
              specified.
            properties:
              auxiliaryServicesConfig:
                description: Immutable. Optional. Configuration of auxiliary services
                  used by this cluster.
                properties:
                  metastoreConfig:
                    description: Immutable. Optional. The Hive Metastore configuration
                      for this workload.
                    properties:
                      dataprocMetastoreServiceRef:
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
                            description: 'Required. Resource name of an existing Dataproc
                              Metastore service. Example: * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`'
                            type: string
                          name:
                            description: |-
                              [WARNING] DataprocMetastoreService not yet supported in Config Connector, use 'external' field to reference existing resources.
                              Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                    required:
                    - dataprocMetastoreServiceRef
                    type: object
                  sparkHistoryServerConfig:
                    description: Immutable. Optional. The Spark History Server configuration
                      for the workload.
                    properties:
                      dataprocClusterRef:
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
                              Optional. Resource name of an existing Dataproc Cluster to act as a Spark History Server for the workload. Example: * `projects/[project_id]/regions/[region]/clusters/[cluster_name]`

                              Allowed value: The `selfLink` field of a `DataprocCluster` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                    type: object
                type: object
              kubernetesClusterConfig:
                description: Immutable. Required. The configuration for running the
                  Dataproc cluster on Kubernetes.
                properties:
                  gkeClusterConfig:
                    description: Immutable. Required. The configuration for running
                      the Dataproc cluster on GKE.
                    properties:
                      gkeClusterTargetRef:
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
                              Optional. A target GKE cluster to deploy to. It must be in the same project and region as the Dataproc cluster (the GKE cluster can be zonal or regional). Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'

                              Allowed value: The `selfLink` field of a `ContainerCluster` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                      nodePoolTarget:
                        description: Immutable. Optional. GKE node pools where workloads
                          will be scheduled. At least one node pool must be assigned
                          the `DEFAULT` GkeNodePoolTarget.Role. If a `GkeNodePoolTarget`
                          is not specified, Dataproc constructs a `DEFAULT` `GkeNodePoolTarget`.
                          Each role can be given to only one `GkeNodePoolTarget`.
                          All node pools must have the same location settings.
                        items:
                          properties:
                            nodePoolConfig:
                              description: Immutable. Input only. The configuration
                                for the GKE node pool. If specified, Dataproc attempts
                                to create a node pool with the specified shape. If
                                one with the same name already exists, it is verified
                                against all specified fields. If a field differs,
                                the virtual cluster creation will fail. If omitted,
                                any node pool with the specified name is used. If
                                a node pool with the specified name does not exist,
                                Dataproc create a node pool with default values. This
                                is an input only field. It will not be returned by
                                the API.
                              properties:
                                autoscaling:
                                  description: Immutable. Optional. The autoscaler
                                    configuration for this node pool. The autoscaler
                                    is enabled only when a valid configuration is
                                    present.
                                  properties:
                                    maxNodeCount:
                                      description: Immutable. The maximum number of
                                        nodes in the node pool. Must be >= min_node_count,
                                        and must be > 0. **Note:** Quota must be sufficient
                                        to scale up the cluster.
                                      format: int64
                                      type: integer
                                    minNodeCount:
                                      description: Immutable. The minimum number of
                                        nodes in the node pool. Must be >= 0 and <=
                                        max_node_count.
                                      format: int64
                                      type: integer
                                  type: object
                                config:
                                  description: Immutable. Optional. The node pool
                                    configuration.
                                  properties:
                                    accelerators:
                                      description: Immutable. Optional. A list of
                                        [hardware accelerators](https://cloud.google.com/compute/docs/gpus)
                                        to attach to each node.
                                      items:
                                        properties:
                                          acceleratorCount:
                                            description: Immutable. The number of
                                              accelerator cards exposed to an instance.
                                            format: int64
                                            type: integer
                                          acceleratorType:
                                            description: Immutable. The accelerator
                                              type resource namename (see GPUs on
                                              Compute Engine).
                                            type: string
                                          gpuPartitionSize:
                                            description: Immutable. Size of partitions
                                              to create on the GPU. Valid values are
                                              described in the NVIDIA [mig user guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
                                            type: string
                                        type: object
                                      type: array
                                    bootDiskKmsKey:
                                      description: 'Immutable. Optional. The [Customer
                                        Managed Encryption Key (CMEK)] (https://cloud.google.com/kubernetes-engine/docs/how-to/using-cmek)
                                        used to encrypt the boot disk attached to
                                        each node in the node pool. Specify the key
                                        using the following format: `projects/KEY_PROJECT_ID/locations/LOCATION/keyRings/RING_NAME/cryptoKeys/KEY_NAME`.'
                                      type: string
                                    ephemeralStorageConfig:
                                      description: Immutable. Optional. Parameters
                                        for the ephemeral storage filesystem. If unspecified,
                                        ephemeral storage is backed by the boot disk.
                                      properties:
                                        localSsdCount:
                                          description: Immutable. Number of local
                                            SSDs to use to back ephemeral storage.
                                            Uses NVMe interfaces. Each local SSD is
                                            375 GB in size. If zero, it means to disable
                                            using local SSDs as ephemeral storage.
                                          format: int64
                                          type: integer
                                      type: object
                                    localSsdCount:
                                      description: Immutable. Optional. The number
                                        of local SSD disks to attach to the node,
                                        which is limited by the maximum number of
                                        disks allowable per zone (see [Adding Local
                                        SSDs](https://cloud.google.com/compute/docs/disks/local-ssd)).
                                      format: int64
                                      type: integer
                                    machineType:
                                      description: Immutable. Optional. The name of
                                        a Compute Engine [machine type](https://cloud.google.com/compute/docs/machine-types).
                                      type: string
                                    minCpuPlatform:
                                      description: Immutable. Optional. [Minimum CPU
                                        platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
                                        to be used by this instance. The instance
                                        may be scheduled on the specified or a newer
                                        CPU platform. Specify the friendly names of
                                        CPU platforms, such as "Intel Haswell"` or
                                        Intel Sandy Bridge".
                                      type: string
                                    preemptible:
                                      description: Immutable. Optional. Whether the
                                        nodes are created as legacy [preemptible VM
                                        instances] (https://cloud.google.com/compute/docs/instances/preemptible).
                                        Also see Spot VMs, preemptible VM instances
                                        without a maximum lifetime. Legacy and Spot
                                        preemptible nodes cannot be used in a node
                                        pool with the `CONTROLLER` [role] (/dataproc/docs/reference/rest/v1/projects.regions.clusters#role)
                                        or in the DEFAULT node pool if the CONTROLLER
                                        role is not assigned (the DEFAULT node pool
                                        will assume the CONTROLLER role).
                                      type: boolean
                                    spot:
                                      description: Immutable. Optional. Whether the
                                        nodes are created as [Spot VM instances] (https://cloud.google.com/compute/docs/instances/spot).
                                        Spot VMs are the latest update to legacy preemptible
                                        VMs. Spot VMs do not have a maximum lifetime.
                                        Legacy and Spot preemptible nodes cannot be
                                        used in a node pool with the `CONTROLLER`
                                        [role](/dataproc/docs/reference/rest/v1/projects.regions.clusters#role)
                                        or in the DEFAULT node pool if the CONTROLLER
                                        role is not assigned (the DEFAULT node pool
                                        will assume the CONTROLLER role).
                                      type: boolean
                                  type: object
                                locations:
                                  description: Immutable. Optional. The list of Compute
                                    Engine [zones](https://cloud.google.com/compute/docs/zones#available)
                                    where node pool nodes associated with a Dataproc
                                    on GKE virtual cluster will be located. **Note:**
                                    All node pools associated with a virtual cluster
                                    must be located in the same region as the virtual
                                    cluster, and they must be located in the same
                                    zone within that region. If a location is not
                                    specified during node pool creation, Dataproc
                                    on GKE will choose the zone.
                                  items:
                                    type: string
                                  type: array
                              type: object
                            nodePoolRef:
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
                                    Required. The target GKE node pool. Format: 'projects/{project}/locations/{location}/clusters/{cluster}/nodePools/{node_pool}'

                                    Allowed value: The `selfLink` field of a `ContainerNodePool` resource.
                                  type: string
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                  type: string
                                namespace:
                                  description: 'Namespace of the referent. More info:
                                    https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                  type: string
                              type: object
                            roles:
                              description: Immutable. Required. The roles associated
                                with the GKE node pool.
                              items:
                                type: string
                              type: array
                          required:
                          - nodePoolRef
                          - roles
                          type: object
                        type: array
                    type: object
                  kubernetesNamespace:
                    description: Immutable. Optional. A namespace within the Kubernetes
                      cluster to deploy into. If this namespace does not exist, it
                      is created. If it exists, Dataproc verifies that another Dataproc
                      VirtualCluster is not installed into it. If not specified, the
                      name of the Dataproc Cluster is used.
                    type: string
                  kubernetesSoftwareConfig:
                    description: Immutable. Optional. The software configuration for
                      this Dataproc cluster running on Kubernetes.
                    properties:
                      componentVersion:
                        additionalProperties:
                          type: string
                        description: Immutable. The components that should be installed
                          in this Dataproc cluster. The key must be a string from
                          the KubernetesComponent enumeration. The value is the version
                          of the software to be installed. At least one entry must
                          be specified.
                        type: object
                      properties:
                        additionalProperties:
                          type: string
                        description: 'Immutable. The properties to set on daemon config
                          files. Property keys are specified in `prefix:property`
                          format, for example `spark:spark.kubernetes.container.image`.
                          The following are supported prefixes and their mappings:
                          * spark: `spark-defaults.conf` For more information, see
                          [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).'
                        type: object
                    type: object
                required:
                - gkeClusterConfig
                type: object
              stagingBucketRef:
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
                      Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging and temp buckets](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a `gs://...` URI to a Cloud Storage bucket.**

                      Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            required:
            - kubernetesClusterConfig
            type: object
required:
- spec
type: object

```
