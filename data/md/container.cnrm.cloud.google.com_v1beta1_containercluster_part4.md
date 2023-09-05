# openAPI schema part4 for container.cnrm.cloud.google.com.v1beta1.ContainerCluster

## schema

```yaml
properties:
  spec:
    properties:
      notificationConfig:
        properties:
          notificationConfig:
            description: The notification config for sending cluster upgrade notifications.
            properties:
              pubsub:
                description: Notification config for Cloud Pub/Sub.
                properties:
                  enabled:
                    description: Whether or not the notification config is enabled.
                    type: boolean
                  filter:
                    description: Allows filtering to one or more specific event types.
                      If event types are present, those and only those event types
                      will be transmitted to the cluster. Other types will be skipped.
                      If no filter is specified, or no event types are present, all
                      event types will be sent.
                    properties:
                      eventType:
                        description: Can be used to filter what notifications are
                          sent. Valid values include include UPGRADE_AVAILABLE_EVENT,
                          UPGRADE_EVENT and SECURITY_BULLETIN_EVENT.
                        items:
                          type: string
                        type: array
                    required:
                    - eventType
                    type: object
                  topicRef:
                    description: The PubSubTopic to send the notification to.
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
                        description: 'Allowed value: string of the format `projects/{{project}}/topics/{{value}}`,
                          where {{value}} is the `name` field of a `PubSubTopic` resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                required:
                - enabled
                type: object
            required:
            - pubsub
            type: object
      podSecurityPolicyConfig:
        description: Configuration for the PodSecurityPolicy feature.
        properties:
          enabled:
            description: Enable the PodSecurityPolicy controller for this cluster.
              If enabled, pods must be valid under a PodSecurityPolicy to be created.
            type: boolean
        required:
        - enabled
        type: object
      privateClusterConfig:
        description: Configuration for private clusters, clusters with private nodes.
        properties:
          enablePrivateEndpoint:
            description: When true, the cluster's private endpoint is used as the
              cluster endpoint and access through the public endpoint is disabled.
              When false, either endpoint can be used. This field only applies to
              private clusters, when enable_private_nodes is true.
            type: boolean
          enablePrivateNodes:
            description: Immutable. Enables the private cluster feature, creating
              a private endpoint on the cluster. In a private cluster, nodes only
              have RFC 1918 private addresses and communicate with the master's private
              endpoint via private networking.
            type: boolean
          masterGlobalAccessConfig:
            description: Controls cluster master global access settings.
            properties:
              enabled:
                description: Whether the cluster master is accessible globally or
                  not.
                type: boolean
            required:
            - enabled
            type: object
          masterIpv4CidrBlock:
            description: Immutable. The IP range in CIDR notation to use for the hosted
              master network. This range will be used for assigning private IP addresses
              to the cluster master(s) and the ILB VIP. This range must not overlap
              with any other ranges in use within the cluster's network, and it must
              be a /28 subnet. See Private Cluster Limitations for more details. This
              field only applies to private clusters, when enable_private_nodes is
              true.
            type: string
          peeringName:
            description: The name of the peering between this cluster and the Google
              owned VPC.
            type: string
          privateEndpoint:
            description: The internal IP address of this cluster's master endpoint.
            type: string
          privateEndpointSubnetworkRef:
            description: |-
              Immutable. Subnetwork in cluster's network where master's endpoint
              will be provisioned.
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
                description: 'Allowed value: The `selfLink` field of a `ComputeSubnetwork`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          publicEndpoint:
            description: The external IP address of this cluster's master endpoint.
            type: string
        type: object
      privateIpv6GoogleAccess:
        description: The desired state of IPv6 connectivity to Google Services. By
          default, no private IPv6 access to or from Google Services (all access will
          be via IPv4).
        type: string
      protectConfig:
        description: Enable/Disable Protect API features for the cluster.
        properties:
          workloadConfig:
            description: WorkloadConfig defines which actions are enabled for a cluster's
              workload configurations.
            properties:
              auditMode:
                description: Sets which mode of auditing should be used for the cluster's
                  workloads. Accepted values are DISABLED, BASIC.
                type: string
            required:
            - auditMode
            type: object
          workloadVulnerabilityMode:
            description: Sets which mode to use for Protect workload vulnerability
              scanning feature. Accepted values are DISABLED, BASIC.
            type: string
        type: object
      releaseChannel:
        description: Configuration options for the Release channel feature, which
          provide more control over automatic upgrades of your GKE clusters. Note
          that removing this field from your config will not unenroll it. Instead,
          use the "UNSPECIFIED" channel.
        properties:
          channel:
            description: |-
              The selected release channel. Accepted values are:
              * UNSPECIFIED: Not set.
              * RAPID: Weekly upgrade cadence; Early testers and developers who requires new features.
              * REGULAR: Multiple per month upgrade cadence; Production users who need features not yet offered in the Stable channel.
              * STABLE: Every few months upgrade cadence; Production users who need stability above all else, and for whom frequent upgrades are too risky.
            type: string
        required:
        - channel
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      resourceUsageExportConfig:
        description: Configuration for the ResourceUsageExportConfig feature.
        properties:
          bigqueryDestination:
            description: Parameters for using BigQuery as the destination of resource
              usage export.
            properties:
              datasetId:
                description: The ID of a BigQuery Dataset.
                type: string
            required:
            - datasetId
            type: object
          enableNetworkEgressMetering:
            description: Whether to enable network egress metering for this cluster.
              If enabled, a daemonset will be created in the cluster to meter network
              egress traffic.
            type: boolean
          enableResourceConsumptionMetering:
            description: Whether to enable resource consumption metering on this cluster.
              When enabled, a table will be created in the resource export BigQuery
              dataset to store resource consumption data. The resulting table can
              be joined with the resource usage table or with BigQuery billing export.
              Defaults to true.
            type: boolean
        required:
        - bigqueryDestination
        type: object
      securityPostureConfig:
        description: Defines the config needed to enable/disable features for the
          Security Posture API.
        properties:
          mode:
            description: Sets the mode of the Kubernetes security posture API's off-cluster
              features. Available options include DISABLED and BASIC.
            type: string
          vulnerabilityMode:
            description: Sets the mode of the Kubernetes security posture API's workload
              vulnerability scanning. Available options include VULNERABILITY_DISABLED
              and VULNERABILITY_BASIC.
            type: string
        type: object
      serviceExternalIpsConfig:
        description: If set, and enabled=true, services with external ips field will
          not be blocked.
        properties:
          enabled:
            description: When enabled, services with exterenal ips specified will
              be allowed.
            type: boolean
        required:
        - enabled
        type: object
      subnetworkRef:
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
            description: 'Allowed value: The `selfLink` field of a `ComputeSubnetwork`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      verticalPodAutoscaling:
        description: Vertical Pod Autoscaling automatically adjusts the resources
          of pods controlled by it.
        properties:
          enabled:
            description: Enables vertical pod autoscaling.
            type: boolean
        required:
        - enabled
        type: object
      workloadIdentityConfig:
        description: Configuration for the use of Kubernetes Service Accounts in GCP
          IAM policies.
        properties:
          identityNamespace:
            description: |-
              DEPRECATED. This field will be removed in a future major release as it has been deprecated in the API. Use `workloadPool` instead; `workloadPool` field will supersede this field.
              Enables workload identity.
            type: string
          workloadPool:
            description: The workload pool to attach all Kubernetes service accounts
              to.
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
      endpoint:
        description: The IP address of this cluster's Kubernetes master.
        type: string
      labelFingerprint:
        description: The fingerprint of the set of labels for this cluster.
        type: string
      masterVersion:
        description: The current version of the master in the cluster. This may be
          different than the min_master_version set in the config if the master has
          been updated by GKE.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      operation:
        type: string
      selfLink:
        description: Server-defined URL for the resource.
        type: string
      servicesIpv4Cidr:
        description: The IP address range of the Kubernetes services in this cluster,
          in CIDR notation (e.g. 1.2.3.4/29). Service addresses are typically put
          in the last /16 from the container CIDR.
        type: string
      tpuIpv4CidrBlock:
        description: The IP address range of the Cloud TPUs in this cluster, in CIDR
          notation (e.g. 1.2.3.4/29).
        type: string
    type: object
required:
- spec
type: object

```
