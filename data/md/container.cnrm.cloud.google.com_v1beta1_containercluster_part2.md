# openAPI schema part2 for container.cnrm.cloud.google.com.v1beta1.ContainerCluster

## schema

```yaml
properties:
  spec:
    properties:
      enableLegacyAbac:
        properties:
          enableLegacyAbac:
            description: Whether the ABAC authorizer is enabled for this cluster.
              When enabled, identities in the system, including service accounts,
              nodes, and controllers, will have statically granted permissions beyond
              those provided by the RBAC configuration or IAM. Defaults to false.
            type: boolean
      enableShieldedNodes:
        description: Enable Shielded Nodes features on all nodes in this cluster.
          Defaults to true.
        type: boolean
      enableTpu:
        description: Immutable. Whether to enable Cloud TPU resources in this cluster.
        type: boolean
      gatewayApiConfig:
        description: Configuration for GKE Gateway API controller.
        properties:
          channel:
            description: The Gateway API release channel to use for Gateway API.
            type: string
        required:
        - channel
        type: object
      identityServiceConfig:
        description: Configuration for Identity Service which allows customers to
          use external identity providers with the K8S API.
        properties:
          enabled:
            description: Whether to enable the Identity Service component.
            type: boolean
        type: object
      initialNodeCount:
        description: Immutable. The number of nodes to create in this cluster's default
          node pool. In regional or multi-zonal clusters, this is the number of nodes
          per zone. Must be set if node_pool is not set. If you're using google_container_node_pool
          objects with no default node pool, you'll need to set this to a value of
          at least 1, alongside setting remove_default_node_pool to true.
        type: integer
      ipAllocationPolicy:
        description: Immutable. Configuration of cluster IP allocation for VPC-native
          clusters. Adding this block enables IP aliasing, making the cluster VPC-native
          instead of routes-based.
        properties:
          clusterIpv4CidrBlock:
            description: Immutable. The IP address range for the cluster pod IPs.
              Set to blank to have a range chosen with the default size. Set to /netmask
              (e.g. /14) to have a range chosen with a specific netmask. Set to a
              CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
              (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific
              range to use.
            type: string
          clusterSecondaryRangeName:
            description: Immutable. The name of the existing secondary range in the
              cluster's subnetwork to use for pod IP addresses. Alternatively, cluster_ipv4_cidr_block
              can be used to automatically create a GKE-managed one.
            type: string
          podCidrOverprovisionConfig:
            description: Immutable. Configuration for cluster level pod cidr overprovision.
              Default is disabled=false.
            properties:
              disabled:
                type: boolean
            required:
            - disabled
            type: object
          servicesIpv4CidrBlock:
            description: Immutable. The IP address range of the services IPs in this
              cluster. Set to blank to have a range chosen with the default size.
              Set to /netmask (e.g. /14) to have a range chosen with a specific netmask.
              Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private
              networks (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a
              specific range to use.
            type: string
          servicesSecondaryRangeName:
            description: Immutable. The name of the existing secondary range in the
              cluster's subnetwork to use for service ClusterIPs. Alternatively, services_ipv4_cidr_block
              can be used to automatically create a GKE-managed one.
            type: string
          stackType:
            description: Immutable. The IP Stack type of the cluster. Choose between
              IPV4 and IPV4_IPV6. Default type is IPV4 Only if not set.
            type: string
        type: object
      location:
        description: Immutable. The location (region or zone) in which the cluster
          master will be created, as well as the default node location. If you specify
          a zone (such as us-central1-a), the cluster will be a zonal cluster with
          a single cluster master. If you specify a region (such as us-west1), the
          cluster will be a regional cluster with multiple masters spread across zones
          in the region, and with default node locations in those zones as well.
        type: string
      loggingConfig:
        description: Logging configuration for the cluster.
        properties:
          enableComponents:
            description: GKE components exposing logs. Valid values include SYSTEM_COMPONENTS,
              APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
            items:
              type: string
            type: array
        required:
        - enableComponents
        type: object
      loggingService:
        description: The logging service that the cluster should write logs to. Available
          options include logging.googleapis.com(Legacy Stackdriver), logging.googleapis.com/kubernetes(Stackdriver
          Kubernetes Engine Logging), and none. Defaults to logging.googleapis.com/kubernetes.
        type: string
      maintenancePolicy:
        description: The maintenance policy to use for the cluster.
        properties:
          dailyMaintenanceWindow:
            description: 'Time window specified for daily maintenance operations.
              Specify start_time in RFC3339 format "HH:MM”, where HH : [00-23] and
              MM : [00-59] GMT.'
            properties:
              duration:
                type: string
              startTime:
                type: string
            required:
            - startTime
            type: object
          maintenanceExclusion:
            description: Exceptions to maintenance window. Non-emergency maintenance
              should not occur in these windows.
            items:
              properties:
                endTime:
                  type: string
                exclusionName:
                  type: string
                exclusionOptions:
                  description: Maintenance exclusion related options.
                  properties:
                    scope:
                      description: The scope of automatic upgrades to restrict in
                        the exclusion window.
                      type: string
                  required:
                  - scope
                  type: object
                startTime:
                  type: string
              required:
              - endTime
              - exclusionName
              - startTime
              type: object
            type: array
          recurringWindow:
            description: Time window for recurring maintenance operations.
            properties:
              endTime:
                type: string
              recurrence:
                type: string
              startTime:
                type: string
            required:
            - endTime
            - recurrence
            - startTime
            type: object
        type: object
      masterAuth:
        description: DEPRECATED. Basic authentication was removed for GKE cluster
          versions >= 1.19. The authentication information for accessing the Kubernetes
          master. Some values in this block are only returned by the API if your service
          account has permission to get credentials for your GKE cluster. If you see
          an unexpected diff unsetting your client cert, ensure you have the container.clusters.getCredentials
          permission.
        properties:
          clientCertificate:
            description: Base64 encoded public certificate used by clients to authenticate
              to the cluster endpoint.
            type: string
          clientCertificateConfig:
            description: Immutable. Whether client certificate authorization is enabled
              for this cluster.
            properties:
              issueClientCertificate:
                description: Immutable. Whether client certificate authorization is
                  enabled for this cluster.
                type: boolean
            required:
            - issueClientCertificate
            type: object
          clientKey:
            description: Base64 encoded private key used by clients to authenticate
              to the cluster endpoint.
            type: string
          clusterCaCertificate:
            description: Base64 encoded public certificate that is the root of trust
              for the cluster.
            type: string
          password:
            description: The password to use for HTTP basic authentication when accessing
              the Kubernetes master endpoint.
            oneOf:
            - not:
                required:
                - valueFrom
              required:
              - value
            - not:
                required:
                - value
              required:
              - valueFrom
            properties:
              value:
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - name
                    - key
                    type: object
                type: object
            type: object
          username:
            description: The username to use for HTTP basic authentication when accessing
              the Kubernetes master endpoint. If not present basic auth will be disabled.
            type: string
        type: object
      masterAuthorizedNetworksConfig:
        description: The desired configuration options for master authorized networks.
          Omit the nested cidr_blocks attribute to disallow external access (except
          the cluster node IPs, which GKE automatically whitelists).
        properties:
          cidrBlocks:
            description: External networks that can access the Kubernetes cluster
              master through HTTPS.
            items:
              properties:
                cidrBlock:
                  description: External network that can access Kubernetes master
                    through HTTPS. Must be specified in CIDR notation.
                  type: string
                displayName:
                  description: Field for users to identify CIDR blocks.
                  type: string
              required:
              - cidrBlock
              type: object
            type: array
          gcpPublicCidrsAccessEnabled:
            description: Whether master is accessbile via Google Compute Engine Public
              IP addresses.
            type: boolean
        type: object
      meshCertificates:
        description: If set, and enable_certificates=true, the GKE Workload Identity
          Certificates controller and node agent will be deployed in the cluster.
        properties:
          enableCertificates:
            description: When enabled the GKE Workload Identity Certificates controller
              and node agent will be deployed in the cluster.
            type: boolean
        required:
        - enableCertificates
        type: object
      minMasterVersion:
        description: The minimum version of the master. GKE will auto-update the master
          to new versions, so this does not guarantee the current master version--use
          the read-only master_version field to obtain that. If unset, the cluster's
          version will be set by GKE to the version of the most recent official release
          (which is not necessarily the latest version).
        type: string
      monitoringConfig:
        description: Monitoring configuration for the cluster.
        properties:
          enableComponents:
            description: GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS,
              APISERVER, CONTROLLER_MANAGER, SCHEDULER, and WORKLOADS.
            items:
              type: string
            type: array
          managedPrometheus:
            description: Configuration for Google Cloud Managed Services for Prometheus.
            properties:
              enabled:
                description: Whether or not the managed collection is enabled.
                type: boolean
            required:
            - enabled
            type: object
        type: object
      monitoringService:
        description: The monitoring service that the cluster should write metrics
          to. Automatically send metrics from pods in the cluster to the Google Cloud
          Monitoring API. VM metrics will be collected by Google Compute Engine regardless
          of this setting Available options include monitoring.googleapis.com(Legacy
          Stackdriver), monitoring.googleapis.com/kubernetes(Stackdriver Kubernetes
          Engine Monitoring), and none. Defaults to monitoring.googleapis.com/kubernetes.
        type: string
      networkPolicy:
        description: Configuration options for the NetworkPolicy feature.
        properties:
          enabled:
            description: Whether network policy is enabled on the cluster.
            type: boolean
          provider:
            description: The selected network policy provider. Defaults to PROVIDER_UNSPECIFIED.
            type: string
        required:
        - enabled
        type: object
      networkRef:
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
            description: 'Allowed value: The `selfLink` field of a `ComputeNetwork`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      networkingMode:
        description: Immutable. Determines whether alias IPs or routes will be used
          for pod IPs in the cluster.
        type: string
required:
- spec
type: object

```
