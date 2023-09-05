# openAPI schema part1 for container.cnrm.cloud.google.com.v1beta1.ContainerCluster

## schema

```yaml
properties:
  spec:
    properties:
      addonsConfig:
        description: The configuration for addons supported by GKE.
        properties:
          cloudrunConfig:
            description: The status of the CloudRun addon. It is disabled by default.
              Set disabled = false to enable.
            properties:
              disabled:
                type: boolean
              loadBalancerType:
                type: string
            required:
            - disabled
            type: object
          configConnectorConfig:
            description: The of the Config Connector addon.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          dnsCacheConfig:
            description: The status of the NodeLocal DNSCache addon. It is disabled
              by default. Set enabled = true to enable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          gcePersistentDiskCsiDriverConfig:
            description: Whether this cluster should enable the Google Compute Engine
              Persistent Disk Container Storage Interface (CSI) Driver. Defaults to
              enabled; set disabled = true to disable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          gcpFilestoreCsiDriverConfig:
            description: The status of the Filestore CSI driver addon, which allows
              the usage of filestore instance as volumes. Defaults to disabled; set
              enabled = true to enable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          gcsFuseCsiDriverConfig:
            description: The status of the GCS Fuse CSI driver addon, which allows
              the usage of gcs bucket as volumes. Defaults to disabled; set enabled
              = true to enable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          gkeBackupAgentConfig:
            description: The status of the Backup for GKE Agent addon. It is disabled
              by default. Set enabled = true to enable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          horizontalPodAutoscaling:
            description: The status of the Horizontal Pod Autoscaling addon, which
              increases or decreases the number of replica pods a replication controller
              has based on the resource usage of the existing pods. It ensures that
              a Heapster pod is running in the cluster, which is also used by the
              Cloud Monitoring service. It is enabled by default; set disabled = true
              to disable.
            properties:
              disabled:
                type: boolean
            required:
            - disabled
            type: object
          httpLoadBalancing:
            description: The status of the HTTP (L7) load balancing controller addon,
              which makes it easy to set up HTTP load balancers for services in a
              cluster. It is enabled by default; set disabled = true to disable.
            properties:
              disabled:
                type: boolean
            required:
            - disabled
            type: object
          istioConfig:
            description: The status of the Istio addon.
            properties:
              auth:
                description: The authentication type between services in Istio. Available
                  options include AUTH_MUTUAL_TLS.
                type: string
              disabled:
                description: The status of the Istio addon, which makes it easy to
                  set up Istio for services in a cluster. It is disabled by default.
                  Set disabled = false to enable.
                type: boolean
            required:
            - disabled
            type: object
          kalmConfig:
            description: Configuration for the KALM addon, which manages the lifecycle
              of k8s. It is disabled by default; Set enabled = true to enable.
            properties:
              enabled:
                type: boolean
            required:
            - enabled
            type: object
          networkPolicyConfig:
            description: Whether we should enable the network policy addon for the
              master. This must be enabled in order to enable network policy for the
              nodes. To enable this, you must also define a network_policy block,
              otherwise nothing will happen. It can only be disabled if the nodes
              already do not have network policies enabled. Defaults to disabled;
              set disabled = false to enable.
            properties:
              disabled:
                type: boolean
            required:
            - disabled
            type: object
        type: object
      authenticatorGroupsConfig:
        description: Configuration for the Google Groups for GKE feature.
        properties:
          securityGroup:
            description: The name of the RBAC security group for use with Google security
              groups in Kubernetes RBAC. Group name must be in format gke-security-groups@yourdomain.com.
            type: string
        required:
        - securityGroup
        type: object
      binaryAuthorization:
        description: Configuration options for the Binary Authorization feature.
        properties:
          enabled:
            description: DEPRECATED. Deprecated in favor of evaluation_mode. Enable
              Binary Authorization for this cluster.
            type: boolean
          evaluationMode:
            description: Mode of operation for Binary Authorization policy evaluation.
            type: string
        type: object
      clusterAutoscaling:
        description: Per-cluster configuration of Node Auto-Provisioning with Cluster
          Autoscaler to automatically adjust the size of the cluster and create/delete
          node pools based on the current needs of the cluster's workload. See the
          guide to using Node Auto-Provisioning for more details.
        properties:
          autoProvisioningDefaults:
            description: Contains defaults for a node pool created by NAP.
            properties:
              bootDiskKMSKeyRef:
                description: |-
                  Immutable. The Customer Managed Encryption Key used to encrypt the
                  boot disk attached to each node in the node pool.
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
              diskSize:
                description: Size of the disk attached to each node, specified in
                  GB. The smallest allowed disk size is 10GB.
                type: integer
              imageType:
                description: The default image type used by NAP once a new node pool
                  is being created.
                type: string
              management:
                description: NodeManagement configuration for this NodePool.
                properties:
                  autoRepair:
                    description: Specifies whether the node auto-repair is enabled
                      for the node pool. If enabled, the nodes in this node pool will
                      be monitored and, if they fail health checks too many times,
                      an automatic repair action will be triggered.
                    type: boolean
                  autoUpgrade:
                    description: Specifies whether node auto-upgrade is enabled for
                      the node pool. If enabled, node auto-upgrade helps keep the
                      nodes in your node pool up to date with the latest release version
                      of Kubernetes.
                    type: boolean
                  upgradeOptions:
                    description: Specifies the Auto Upgrade knobs for the node pool.
                    items:
                      properties:
                        autoUpgradeStartTime:
                          description: This field is set when upgrades are about to
                            commence with the approximate start time for the upgrades,
                            in RFC3339 text format.
                          type: string
                        description:
                          description: This field is set when upgrades are about to
                            commence with the description of the upgrade.
                          type: string
                      type: object
                    type: array
                type: object
              minCpuPlatform:
                description: Minimum CPU platform to be used by this instance. The
                  instance may be scheduled on the specified or newer CPU platform.
                  Applicable values are the friendly names of CPU platforms, such
                  as Intel Haswell.
                type: string
              oauthScopes:
                description: Scopes that are used by NAP when creating node pools.
                items:
                  type: string
                type: array
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
                description: Shielded Instance options.
                properties:
                  enableIntegrityMonitoring:
                    description: Defines whether the instance has integrity monitoring
                      enabled.
                    type: boolean
                  enableSecureBoot:
                    description: Defines whether the instance has Secure Boot enabled.
                    type: boolean
                type: object
              upgradeSettings:
                description: Specifies the upgrade settings for NAP created node pools.
                properties:
                  blueGreenSettings:
                    description: Settings for blue-green upgrade strategy.
                    properties:
                      nodePoolSoakDuration:
                        description: "Time needed after draining entire blue pool.
                          After this period, blue pool will be cleaned up.\n\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tA
                          duration in seconds with up to nine fractional digits, ending
                          with 's'. Example: \"3.5s\"."
                        type: string
                      standardRolloutPolicy:
                        description: Standard policy for the blue-green upgrade.
                        properties:
                          batchNodeCount:
                            description: Number of blue nodes to drain in a batch.
                            type: integer
                          batchPercentage:
                            description: Percentage of the bool pool nodes to drain
                              in a batch. The range of this field should be (0.0,
                              1.0].
                            type: number
                          batchSoakDuration:
                            description: "Soak time after each batch gets drained.\n\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\tA
                              duration in seconds with up to nine fractional digits,
                              ending with 's'. Example: \"3.5s\"."
                            type: string
                        type: object
                    type: object
                  maxSurge:
                    description: The maximum number of nodes that can be created beyond
                      the current size of the node pool during the upgrade process.
                    type: integer
                  maxUnavailable:
                    description: The maximum number of nodes that can be simultaneously
                      unavailable during the upgrade process.
                    type: integer
                  strategy:
                    description: Update strategy of the node pool.
                    type: string
                type: object
            type: object
          autoscalingProfile:
            description: Configuration options for the Autoscaling profile feature,
              which lets you choose whether the cluster autoscaler should optimize
              for resource utilization or resource availability when deciding to remove
              nodes from a cluster. Can be BALANCED or OPTIMIZE_UTILIZATION. Defaults
              to BALANCED.
            type: string
          enabled:
            description: Whether node auto-provisioning is enabled. Resource limits
              for cpu and memory must be defined to enable node auto-provisioning.
            type: boolean
          resourceLimits:
            description: Global constraints for machine resources in the cluster.
              Configuring the cpu and memory types is required if node auto-provisioning
              is enabled. These limits will apply to node pool autoscaling in addition
              to node auto-provisioning.
            items:
              properties:
                maximum:
                  description: Maximum amount of the resource in the cluster.
                  type: integer
                minimum:
                  description: Minimum amount of the resource in the cluster.
                  type: integer
                resourceType:
                  description: The type of the resource. For example, cpu and memory.
                    See the guide to using Node Auto-Provisioning for a list of types.
                  type: string
              required:
              - resourceType
              type: object
            type: array
        type: object
      clusterIpv4Cidr:
        description: Immutable. The IP address range of the Kubernetes pods in this
          cluster in CIDR notation (e.g. 10.96.0.0/14). Leave blank to have one automatically
          chosen or specify a /14 block in 10.0.0.0/8. This field will only work for
          routes-based clusters, where ip_allocation_policy is not defined.
        type: string
      clusterTelemetry:
        description: Telemetry integration for the cluster.
        properties:
          type:
            description: Type of the integration.
            type: string
        required:
        - type
        type: object
      confidentialNodes:
        description: 'Immutable. Configuration for the confidential nodes feature,
          which makes nodes run on confidential VMs. Warning: This configuration can''t
          be changed (or added/removed) after cluster creation without deleting and
          recreating the entire cluster.'
        properties:
          enabled:
            description: Immutable. Whether Confidential Nodes feature is enabled
              for all nodes in this cluster.
            type: boolean
        required:
        - enabled
        type: object
      costManagementConfig:
        description: Cost management configuration for the cluster.
        properties:
          enabled:
            description: Whether to enable GKE cost allocation. When you enable GKE
              cost allocation, the cluster name and namespace of your GKE workloads
              appear in the labels field of the billing export to BigQuery. Defaults
              to false.
            type: boolean
        required:
        - enabled
        type: object
      databaseEncryption:
        description: 'Application-layer Secrets Encryption settings. The object format
          is {state = string, key_name = string}. Valid values of state are: "ENCRYPTED";
          "DECRYPTED". key_name is the name of a CloudKMS key.'
        properties:
          keyName:
            description: The key to use to encrypt/decrypt secrets.
            type: string
          state:
            description: ENCRYPTED or DECRYPTED.
            type: string
        required:
        - state
        type: object
      datapathProvider:
        description: Immutable. The desired datapath provider for this cluster. By
          default, uses the IPTables-based kube-proxy implementation.
        type: string
      defaultMaxPodsPerNode:
        description: Immutable. The default maximum number of pods per node in this
          cluster. This doesn't work on "routes-based" clusters, clusters that don't
          have IP Aliasing enabled.
        type: integer
      defaultSnatStatus:
        description: Whether the cluster disables default in-node sNAT rules. In-node
          sNAT rules will be disabled when defaultSnatStatus is disabled.
        properties:
          disabled:
            description: When disabled is set to false, default IP masquerade rules
              will be applied to the nodes to prevent sNAT on cluster internal traffic.
            type: boolean
        required:
        - disabled
        type: object
      description:
        description: Immutable.  Description of the cluster.
        type: string
      dnsConfig:
        description: Immutable. Configuration for Cloud DNS for Kubernetes Engine.
        properties:
          clusterDns:
            description: Which in-cluster DNS provider should be used.
            type: string
          clusterDnsDomain:
            description: The suffix used for all cluster service records.
            type: string
          clusterDnsScope:
            description: The scope of access to cluster DNS records.
            type: string
        type: object
      enableAutopilot:
        description: Immutable. Enable Autopilot for this cluster.
        type: boolean
      enableBinaryAuthorization:
        description: DEPRECATED. Deprecated in favor of binary_authorization. Enable
          Binary Authorization for this cluster. If enabled, all container images
          will be validated by Google Binary Authorization.
        type: boolean
      enableIntranodeVisibility:
        description: Whether Intra-node visibility is enabled for this cluster. This
          makes same node pod to pod traffic visible for VPC network.
        type: boolean
      enableKubernetesAlpha:
        description: Immutable. Whether to enable Kubernetes Alpha features for this
          cluster. Note that when this option is enabled, the cluster cannot be upgraded
          and will be automatically deleted after 30 days.
        type: boolean
      enableL4IlbSubsetting:
        description: Whether L4ILB Subsetting is enabled for this cluster.
        type: boolean
      spec:
        required:
        - location
        type: object
required:
- spec
type: object

```
