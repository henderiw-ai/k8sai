# openAPI schema part0 for redis.cnrm.cloud.google.com.v1beta1.RedisInstance

## schema

```yaml
properties:
  apiVersion:
    description: 'apiVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
    type: string
  kind:
    description: 'kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    properties:
      alternativeLocationId:
        description: |-
          Immutable. Only applicable to STANDARD_HA tier which protects the instance
          against zonal failures by provisioning it across two zones.
          If provided, it must be a different zone from the one provided in
          [locationId].
        type: string
      authEnabled:
        description: |-
          Optional. Indicates whether OSS Redis AUTH is enabled for the
          instance. If set to "true" AUTH is enabled on the instance.
          Default value is "false" meaning AUTH is disabled.
        type: boolean
      authString:
        description: Output only. AUTH String set on the instance. This field will
          only be populated if auth_enabled is true.
        type: string
      authorizedNetworkRef:
        description: |-
          The network to which the instance is connected. If left
          unspecified, the default network will be used.
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
      connectMode:
        description: 'Immutable. The connection mode of the Redis instance. Default
          value: "DIRECT_PEERING" Possible values: ["DIRECT_PEERING", "PRIVATE_SERVICE_ACCESS"].'
        type: string
      customerManagedKeyRef:
        description: |-
          Immutable. Optional. The KMS key reference that you want to use to
          encrypt the data at rest for this Redis instance. If this is
          provided, CMEK is enabled.
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
      displayName:
        description: An arbitrary and optional user-provided name for the instance.
        type: string
      locationId:
        description: |-
          Immutable. The zone where the instance will be provisioned. If not provided,
          the service will choose a zone for the instance. For STANDARD_HA tier,
          instances will be created across two zones for protection against
          zonal failures. If [alternativeLocationId] is also provided, it must
          be different from [locationId].
        type: string
      maintenancePolicy:
        description: Maintenance policy for an instance.
        properties:
          createTime:
            description: |-
              Output only. The time when the policy was created.
              A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
              resolution and up to nine fractional digits.
            type: string
          description:
            description: |-
              Optional. Description of what this policy is for.
              Create/Update methods return INVALID_ARGUMENT if the
              length is greater than 512.
            type: string
          updateTime:
            description: |-
              Output only. The time when the policy was last updated.
              A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
              resolution and up to nine fractional digits.
            type: string
          weeklyMaintenanceWindow:
            description: |-
              Optional. Maintenance window that is applied to resources covered by this policy.
              Minimum 1. For the current version, the maximum number
              of weekly_window is expected to be one.
            items:
              properties:
                day:
                  description: |-
                    Required. The day of week that maintenance updates occur.

                    - DAY_OF_WEEK_UNSPECIFIED: The day of the week is unspecified.
                    - MONDAY: Monday
                    - TUESDAY: Tuesday
                    - WEDNESDAY: Wednesday
                    - THURSDAY: Thursday
                    - FRIDAY: Friday
                    - SATURDAY: Saturday
                    - SUNDAY: Sunday Possible values: ["DAY_OF_WEEK_UNSPECIFIED", "MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].
                  type: string
                duration:
                  description: |-
                    Output only. Duration of the maintenance window.
                    The current window is fixed at 1 hour.
                    A duration in seconds with up to nine fractional digits,
                    terminated by 's'. Example: "3.5s".
                  type: string
                startTime:
                  description: Required. Start time of the window in UTC time.
                  properties:
                    hours:
                      description: |-
                        Hours of day in 24 hour format. Should be from 0 to 23.
                        An API may choose to allow the value "24:00:00" for scenarios like business closing time.
                      type: integer
                    minutes:
                      description: Minutes of hour of day. Must be from 0 to 59.
                      type: integer
                    nanos:
                      description: Fractions of seconds in nanoseconds. Must be from
                        0 to 999,999,999.
                      type: integer
                    seconds:
                      description: |-
                        Seconds of minutes of the time. Must normally be from 0 to 59.
                        An API may allow the value 60 if it allows leap-seconds.
                      type: integer
                  type: object
              required:
              - day
              - startTime
              type: object
            type: array
        type: object
      maintenanceSchedule:
        description: Upcoming maintenance schedule.
        items:
          properties:
            endTime:
              description: |-
                Output only. The end time of any upcoming scheduled maintenance for this instance.
                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                resolution and up to nine fractional digits.
              type: string
            scheduleDeadlineTime:
              description: |-
                Output only. The deadline that the maintenance schedule start time
                can not go beyond, including reschedule.
                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                resolution and up to nine fractional digits.
              type: string
            startTime:
              description: |-
                Output only. The start time of any upcoming scheduled maintenance for this instance.
                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                resolution and up to nine fractional digits.
              type: string
          type: object
        type: array
      memorySizeGb:
        description: Redis memory size in GiB.
        type: integer
      persistenceConfig:
        description: Persistence configuration for an instance.
        properties:
          persistenceMode:
            description: "Optional. Controls whether Persistence features are enabled.
              If not provided, the existing value will be used.\n\n- DISABLED: \tPersistence
              is disabled for the instance, and any existing snapshots are deleted.\n-
              RDB: RDB based Persistence is enabled. Possible values: [\"DISABLED\",
              \"RDB\"]."
            type: string
          rdbNextSnapshotTime:
            description: |-
              Output only. The next time that a snapshot attempt is scheduled to occur.
              A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up
              to nine fractional digits.
              Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
            type: string
          rdbSnapshotPeriod:
            description: "Optional. Available snapshot periods for scheduling.\n\n-
              ONE_HOUR:\tSnapshot every 1 hour.\n- SIX_HOURS:\tSnapshot every 6 hours.\n-
              TWELVE_HOURS:\tSnapshot every 12 hours.\n- TWENTY_FOUR_HOURS:\tSnapshot
              every 24 hours. Possible values: [\"ONE_HOUR\", \"SIX_HOURS\", \"TWELVE_HOURS\",
              \"TWENTY_FOUR_HOURS\"]."
            type: string
          rdbSnapshotStartTime:
            description: |-
              Optional. Date and time that the first snapshot was/will be attempted,
              and to which future snapshots will be aligned. If not provided,
              the current time will be used.
              A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution
              and up to nine fractional digits.
              Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
            type: string
        type: object
      readReplicasMode:
        description: |-
          Optional. Read replica mode. Can only be specified when trying to create the instance.
          If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
          - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
          instance cannot scale up or down the number of replicas.
          - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
          can scale up and down the number of replicas. Possible values: ["READ_REPLICAS_DISABLED", "READ_REPLICAS_ENABLED"].
        type: string
      redisConfigs:
        additionalProperties:
          type: string
        description: |-
          Redis configuration parameters, according to http://redis.io/topics/config.
          Please check Memorystore documentation for the list of supported parameters:
          https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs.
        type: object
      redisVersion:
        description: |-
          The version of Redis software. If not provided, latest supported
          version will be used. Please check the API documentation linked
          at the top for the latest valid values.
        type: string
      region:
        description: Immutable. The name of the Redis region of the instance.
        type: string
      replicaCount:
        description: |-
          Optional. The number of replica nodes. The valid range for the Standard Tier with
          read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
          for a Standard Tier instance, the only valid value is 1 and the default is 1.
          The valid value for basic tier is 0 and the default is also 0.
        type: integer
      reservedIpRange:
        description: |-
          Immutable. The CIDR range of internal addresses that are reserved for this
          instance. If not provided, the service will choose an unused /29
          block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
          unique and non-overlapping with existing subnets in an authorized
          network.
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      secondaryIpRange:
        description: |-
          Optional. Additional IP range for node placement. Required when enabling read replicas on
          an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
          "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
          range associated with the private service access connection, or "auto".
        type: string
      tier:
        description: |-
          Immutable. The service tier of the instance. Must be one of these values:

          - BASIC: standalone instance
          - STANDARD_HA: highly available primary/replica instances Default value: "BASIC" Possible values: ["BASIC", "STANDARD_HA"].
        type: string
      transitEncryptionMode:
        description: |-
          Immutable. The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.

          - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication Default value: "DISABLED" Possible values: ["SERVER_AUTHENTICATION", "DISABLED"].
        type: string
    required:
    - memorySizeGb
    - region
    type: object
required:
- spec
type: object

```
