# openAPI schema part2 for sql.cnrm.cloud.google.com.v1beta1.SQLInstance

## schema

```yaml
properties:
  spec:
    properties:
      settings:
        properties:
          settings:
            description: The settings to use for the database. The configuration is
              detailed below.
            properties:
              activationPolicy:
                description: This specifies when the instance should be active. Can
                  be either ALWAYS, NEVER or ON_DEMAND.
                type: string
              activeDirectoryConfig:
                properties:
                  domain:
                    description: Domain name of the Active Directory for SQL Server
                      (e.g., mydomain.com).
                    type: string
                required:
                - domain
                type: object
              advancedMachineFeatures:
                properties:
                  threadsPerCore:
                    description: The number of threads per physical core. Can be 1
                      or 2.
                    type: integer
                type: object
              authorizedGaeApplications:
                description: |-
                  DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
                  Specifying this field has no-ops; it's recommended to remove this field from your configuration.
                items:
                  type: string
                type: array
              availabilityType:
                description: |-
                  The availability type of the Cloud SQL instance, high availability
                  (REGIONAL) or single zone (ZONAL). For all instances, ensure that
                  settings.backup_configuration.enabled is set to true.
                  For MySQL instances, ensure that settings.backup_configuration.binary_log_enabled is set to true.
                  For Postgres instances, ensure that settings.backup_configuration.point_in_time_recovery_enabled
                  is set to true. Defaults to ZONAL.
                type: string
              backupConfiguration:
                properties:
                  backupRetentionSettings:
                    properties:
                      retainedBackups:
                        description: Number of backups to retain.
                        type: integer
                      retentionUnit:
                        description: The unit that 'retainedBackups' represents. Defaults
                          to COUNT.
                        type: string
                    required:
                    - retainedBackups
                    type: object
                  binaryLogEnabled:
                    description: True if binary logging is enabled. If settings.backup_configuration.enabled
                      is false, this must be as well. Can only be used with MySQL.
                    type: boolean
                  enabled:
                    description: True if backup configuration is enabled.
                    type: boolean
                  location:
                    description: Location of the backup configuration.
                    type: string
                  pointInTimeRecoveryEnabled:
                    description: True if Point-in-time recovery is enabled.
                    type: boolean
                  startTime:
                    description: HH:MM format time indicating when backup configuration
                      starts.
                    type: string
                  transactionLogRetentionDays:
                    description: The number of days of transaction logs we retain
                      for point in time restore, from 1-7.
                    type: integer
                type: object
              collation:
                description: Immutable. The name of server instance collation.
                type: string
              connectorEnforcement:
                description: Specifies if connections must use Cloud SQL connectors.
                type: string
              crashSafeReplication:
                description: |-
                  DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
                  Specifying this field has no-ops; it's recommended to remove this field from your configuration.
                type: boolean
              dataCacheConfig:
                description: Data cache configurations.
                properties:
                  dataCacheEnabled:
                    description: Whether data cache is enabled for the instance.
                    type: boolean
                type: object
              databaseFlags:
                items:
                  properties:
                    name:
                      description: Name of the flag.
                      type: string
                    value:
                      description: Value of the flag.
                      type: string
                  required:
                  - name
                  - value
                  type: object
                type: array
              deletionProtectionEnabled:
                description: Configuration to protect against accidental instance
                  deletion.
                type: boolean
              denyMaintenancePeriod:
                properties:
                  endDate:
                    description: End date before which maintenance will not take place.
                      The date is in format yyyy-mm-dd i.e., 2020-11-01, or mm-dd,
                      i.e., 11-01.
                    type: string
                  startDate:
                    description: Start date after which maintenance will not take
                      place. The date is in format yyyy-mm-dd i.e., 2020-11-01, or
                      mm-dd, i.e., 11-01.
                    type: string
                  time:
                    description: 'Time in UTC when the "deny maintenance period" starts
                      on start_date and ends on end_date. The time is in format: HH:mm:SS,
                      i.e., 00:00:00.'
                    type: string
                required:
                - endDate
                - startDate
                - time
                type: object
              diskAutoresize:
                description: Enables auto-resizing of the storage size. Defaults to
                  true.
                type: boolean
              diskAutoresizeLimit:
                description: The maximum size, in GB, to which storage capacity can
                  be automatically increased. The default value is 0, which specifies
                  that there is no limit.
                type: integer
              diskSize:
                description: The size of data disk, in GB. Size of a running instance
                  cannot be reduced but can be increased. The minimum value is 10GB.
                type: integer
              diskType:
                description: 'Immutable. The type of data disk: PD_SSD or PD_HDD.
                  Defaults to PD_SSD.'
                type: string
              edition:
                description: The edition of the instance, can be ENTERPRISE or ENTERPRISE_PLUS.
                type: string
              insightsConfig:
                description: Configuration of Query Insights.
                properties:
                  queryInsightsEnabled:
                    description: True if Query Insights feature is enabled.
                    type: boolean
                  queryPlansPerMinute:
                    description: Number of query execution plans captured by Insights
                      per minute for all queries combined. Between 0 and 20. Default
                      to 5.
                    type: integer
                  queryStringLength:
                    description: Maximum query length stored in bytes. Between 256
                      and 4500. Default to 1024.
                    type: integer
                  recordApplicationTags:
                    description: True if Query Insights will record application tags
                      from query when enabled.
                    type: boolean
                  recordClientAddress:
                    description: True if Query Insights will record client address
                      when enabled.
                    type: boolean
                type: object
              ipConfiguration:
                properties:
                  allocatedIpRange:
                    description: 'The name of the allocated ip range for the private
                      ip CloudSQL instance. For example: "google-managed-services-default".
                      If set, the instance ip will be created in the allocated range.
                      The range name must comply with RFC 1035. Specifically, the
                      name must be 1-63 characters long and match the regular expression
                      [a-z]([-a-z0-9]*[a-z0-9])?.'
                    type: string
                  authorizedNetworks:
                    items:
                      properties:
                        expirationTime:
                          type: string
                        name:
                          type: string
                        value:
                          type: string
                      required:
                      - value
                      type: object
                    type: array
                  enablePrivatePathForGoogleCloudServices:
                    description: Whether Google Cloud services such as BigQuery are
                      allowed to access data in this Cloud SQL instance over a private
                      IP connection. SQLSERVER database type is not supported.
                    type: boolean
                  ipv4Enabled:
                    description: Whether this Cloud SQL instance should be assigned
                      a public IPV4 address. At least ipv4_enabled must be enabled
                      or a private_network must be configured.
                    type: boolean
                  privateNetworkRef:
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
                  requireSsl:
                    type: boolean
                type: object
              locationPreference:
                properties:
                  followGaeApplication:
                    description: A Google App Engine application whose zone to remain
                      in. Must be in the same region as this instance.
                    type: string
                  secondaryZone:
                    description: The preferred Compute Engine zone for the secondary/failover.
                    type: string
                  zone:
                    description: The preferred compute engine zone.
                    type: string
                type: object
              maintenanceWindow:
                description: Declares a one-hour maintenance window when an Instance
                  can automatically restart to apply updates. The maintenance window
                  is specified in UTC time.
                properties:
                  day:
                    description: Day of week (1-7), starting on Monday.
                    type: integer
                  hour:
                    description: Hour of day (0-23), ignored if day not set.
                    type: integer
                  updateTrack:
                    description: Receive updates earlier (canary) or later (stable).
                    type: string
                type: object
              passwordValidationPolicy:
                properties:
                  complexity:
                    description: Password complexity.
                    type: string
                  disallowUsernameSubstring:
                    description: Disallow username as a part of the password.
                    type: boolean
                  enablePasswordPolicy:
                    description: Whether the password policy is enabled or not.
                    type: boolean
                  minLength:
                    description: Minimum number of characters allowed.
                    type: integer
                  passwordChangeInterval:
                    description: Minimum interval after which the password can be
                      changed. This flag is only supported for PostgresSQL.
                    type: string
                  reuseInterval:
                    description: Number of previous passwords that cannot be reused.
                    type: integer
                required:
                - enablePasswordPolicy
                type: object
              pricingPlan:
                description: Pricing plan for this instance, can only be PER_USE.
                type: string
              replicationType:
                description: |-
                  DEPRECATED. This property is only applicable to First Generation instances, and First Generation instances are now deprecated. see https://cloud.google.com/sql/docs/mysql/deprecation-notice for information on how to upgrade to Second Generation instances.
                  Specifying this field has no-ops; it's recommended to remove this field from your configuration.
                type: string
              sqlServerAuditConfig:
                properties:
                  bucketRef:
                    description: The name of the destination bucket (e.g., gs://mybucket).
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
                        description: 'Allowed value: The `url` field of a `StorageBucket`
                          resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  retentionInterval:
                    description: 'How long to keep generated audit files. A duration
                      in seconds with up to nine fractional digits, terminated by
                      ''s''. Example: "3.5s"..'
                    type: string
                  uploadInterval:
                    description: 'How often to upload generated audit files. A duration
                      in seconds with up to nine fractional digits, terminated by
                      ''s''. Example: "3.5s".'
                    type: string
                type: object
              tier:
                description: The machine type to use. See tiers for more details and
                  supported versions. Postgres supports only shared-core machine types,
                  and custom machine types such as db-custom-2-13312. See the Custom
                  Machine Type Documentation to learn about specifying custom machine
                  types.
                type: string
              timeZone:
                description: Immutable. The time_zone to be used by the database engine
                  (supported only for SQL Server), in SQL Server timezone format.
                type: string
            required:
            - tier
            type: object
  status:
    properties:
      availableMaintenanceVersions:
        description: Available Maintenance versions.
        items:
          type: string
        type: array
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
      connectionName:
        description: The connection name of the instance to be used in connection
          strings. For example, when connecting with Cloud SQL Proxy.
        type: string
      firstIpAddress:
        type: string
      instanceType:
        description: The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED',
          'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.
        type: string
      ipAddress:
        items:
          properties:
            ipAddress:
              type: string
            timeToRetire:
              type: string
            type:
              type: string
          type: object
        type: array
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      privateIpAddress:
        type: string
      publicIpAddress:
        type: string
      selfLink:
        description: The URI of the created resource.
        type: string
      serverCaCert:
        properties:
          cert:
            description: The CA Certificate used to connect to the SQL Instance via
              SSL.
            type: string
          commonName:
            description: The CN valid for the CA Cert.
            type: string
          createTime:
            description: Creation time of the CA Cert.
            type: string
          expirationTime:
            description: Expiration time of the CA Cert.
            type: string
          sha1Fingerprint:
            description: SHA Fingerprint of the CA Cert.
            type: string
        type: object
      serviceAccountEmailAddress:
        description: The service account email address assigned to the instance.
        type: string
    type: object
required:
- spec
type: object

```
