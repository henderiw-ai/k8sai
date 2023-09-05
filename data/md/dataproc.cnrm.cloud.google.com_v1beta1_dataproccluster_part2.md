# openAPI schema part2 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocCluster

## schema

```yaml
properties:
  spec:
    properties:
      config:
        properties:
          autoscalingConfig:
            description: Immutable. Optional. Autoscaling config for the policy associated
              with the cluster. Cluster does not autoscale if this field is unset.
            properties:
              policyRef:
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
                      Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` Note that the policy must be in the same project and Dataproc region.

                      Allowed value: The Google Cloud resource name of a `DataprocAutoscalingPolicy` resource (format: `projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}`).
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            type: object
          config:
            description: Immutable. The cluster config. Note that Dataproc may set
              default values, and values may change when clusters are updated.
            type: object
          dataprocMetricConfig:
            description: Immutable. Optional. The config for Dataproc metrics.
            properties:
              metrics:
                description: Immutable. Required. Metrics sources to enable.
                items:
                  properties:
                    metricOverrides:
                      description: 'Immutable. Optional. Specify one or more [available
                        OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics)
                        to collect for the metric course (for the `SPARK` metric source,
                        any [Spark metric] (https://spark.apache.org/docs/latest/monitoring.html#metrics)
                        can be specified). Provide metrics in the following format:
                        `METRIC_SOURCE:INSTANCE:GROUP:METRIC` Use camelcase as appropriate.
                        Examples: ``` yarn:ResourceManager:QueueMetrics:AppsCompleted
                        spark:driver:DAGScheduler:job.allJobs sparkHistoryServer:JVM:Memory:NonHeapMemoryUsage.committed
                        hiveserver2:JVM:Memory:NonHeapMemoryUsage.used ``` Notes:
                        * Only the specified overridden metrics will be collected
                        for the metric source. For example, if one or more `spark:executive`
                        metrics are listed as metric overrides, other `SPARK` metrics
                        will not be collected. The collection of the default metrics
                        for other OSS metric sources is unaffected. For example, if
                        both `SPARK` andd `YARN` metric sources are enabled, and overrides
                        are provided for Spark metrics only, all default YARN metrics
                        will be collected.'
                      items:
                        type: string
                      type: array
                    metricSource:
                      description: 'Immutable. Required. Default metrics are collected
                        unless `metricOverrides` are specified for the metric source
                        (see [Available OSS metrics] (https://cloud.google.com/dataproc/docs/guides/monitoring#available_oss_metrics)
                        for more information). Possible values: METRIC_SOURCE_UNSPECIFIED,
                        MONITORING_AGENT_DEFAULTS, HDFS, SPARK, YARN, SPARK_HISTORY_SERVER,
                        HIVESERVER2'
                      type: string
                  required:
                  - metricSource
                  type: object
                type: array
            required:
            - metrics
            type: object
          encryptionConfig:
            description: Immutable. Optional. Encryption settings for the cluster.
            properties:
              gcePdKmsKeyRef:
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
                      Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.

                      Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            type: object
          endpointConfig:
            description: Immutable. Optional. Port/endpoint configuration for this
              cluster
            properties:
              enableHttpPortAccess:
                description: Immutable. Optional. If true, enable http access to specific
                  ports on the cluster from external sources. Defaults to false.
                type: boolean
            type: object
          gceClusterConfig:
            description: Immutable. Optional. The shared Compute Engine config settings
              for all instances in a cluster.
            properties:
              confidentialInstanceConfig:
                description: Immutable. Optional. Confidential Instance Config for
                  clusters using [Confidential VMs](https://cloud.google.com/compute/confidential-vm/docs).
                properties:
                  enableConfidentialCompute:
                    description: Immutable. Optional. Defines whether the instance
                      should have confidential compute enabled.
                    type: boolean
                type: object
              internalIPOnly:
                description: Immutable. Optional. If true, all instances in the cluster
                  will only have internal IP addresses. By default, clusters are not
                  restricted to internal IP addresses, and will have ephemeral external
                  IP addresses assigned to each instance. This `internal_ip_only`
                  restriction can only be enabled for subnetwork enabled networks,
                  and all off-cluster dependencies must be configured to be accessible
                  without external IP addresses.
                type: boolean
              metadata:
                additionalProperties:
                  type: string
                description: Immutable. The Compute Engine metadata entries to add
                  to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
                type: object
              networkRef:
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
                      Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` * `default`

                      Allowed value: The `selfLink` field of a `ComputeNetwork` resource.
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              nodeGroupAffinity:
                description: Immutable. Optional. Node Group Affinity for sole-tenant
                  clusters.
                properties:
                  nodeGroupRef:
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
                          Required. The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `node-group-1`

                          Allowed value: The `selfLink` field of a `ComputeNodeGroup` resource.
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                required:
                - nodeGroupRef
                type: object
              privateIPv6GoogleAccess:
                description: 'Immutable. Optional. The type of IPv6 access for a cluster.
                  Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED, INHERIT_FROM_SUBNETWORK,
                  OUTBOUND, BIDIRECTIONAL'
                type: string
              reservationAffinity:
                description: Immutable. Optional. Reservation Affinity for consuming
                  Zonal reservation.
                properties:
                  consumeReservationType:
                    description: 'Immutable. Optional. Type of reservation to consume
                      Possible values: TYPE_UNSPECIFIED, NO_RESERVATION, ANY_RESERVATION,
                      SPECIFIC_RESERVATION'
                    type: string
                  key:
                    description: Immutable. Optional. Corresponds to the label key
                      of reservation resource.
                    type: string
                  values:
                    description: Immutable. Optional. Corresponds to the label values
                      of reservation resource.
                    items:
                      type: string
                    type: array
                type: object
              serviceAccountRef:
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
                      Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.

                      Allowed value: The `email` field of an `IAMServiceAccount` resource.
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              serviceAccountScopes:
                description: 'Immutable. Optional. The URIs of service account scopes
                  to be included in Compute Engine instances. The following base set
                  of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly
                  * https://www.googleapis.com/auth/devstorage.read_write * https://www.googleapis.com/auth/logging.write
                  If no scopes are specified, the following defaults are also provided:
                  * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table
                  * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control'
                items:
                  type: string
                type: array
              shieldedInstanceConfig:
                description: Immutable. Optional. Shielded Instance Config for clusters
                  using [Compute Engine Shielded VMs](https://cloud.google.com/security/shielded-cloud/shielded-vm).
                properties:
                  enableIntegrityMonitoring:
                    description: Immutable. Optional. Defines whether instances have
                      integrity monitoring enabled.
                    type: boolean
                  enableSecureBoot:
                    description: Immutable. Optional. Defines whether instances have
                      Secure Boot enabled.
                    type: boolean
                  enableVtpm:
                    description: Immutable. Optional. Defines whether instances have
                      the vTPM enabled.
                    type: boolean
                type: object
              subnetworkRef:
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
                      Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0` * `sub0`

                      Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              tags:
                description: Immutable. The Compute Engine tags to add to all instances
                  (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
                items:
                  type: string
                type: array
              zone:
                description: 'Immutable. Optional. The zone where the Compute Engine
                  cluster will be located. On a create request, it is required in
                  the "global" region. If omitted in a non-global Dataproc region,
                  the service will pick a zone in the corresponding Compute Engine
                  region. On a get request, zone will always be present. A full URL,
                  partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]`
                  * `projects/[project_id]/zones/[zone]` * `us-central1-f`'
                type: string
            type: object
required:
- spec
type: object

```
