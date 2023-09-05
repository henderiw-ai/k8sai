# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeInstanceGroupManager

## schema

```yaml
properties:
  spec:
    properties:
      autoHealingPolicies:
        description: The autohealing policy for this managed instance group. You can
          specify only one value.
        items:
          properties:
            healthCheckRef:
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
                    The URL for the health check that signals autohealing.

                    Allowed value: The `selfLink` field of a `ComputeHealthCheck` resource.
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
            initialDelaySec:
              description: The number of seconds that the managed instance group waits
                before it applies autohealing policies to new instances or recently
                recreated instances. This initial delay allows instances to initialize
                and run their startup scripts before the instance group determines
                that they are UNHEALTHY. This prevents the managed instance group
                from recreating its instances prematurely. This value must be from
                range [0, 3600].
              format: int64
              type: integer
          type: object
        type: array
      baseInstanceName:
        description: The base instance name to use for instances in this group. The
          value must be 1-58 characters long. Instances are named by appending a hyphen
          and a random four-character string to the base instance name. The base instance
          name must comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
        type: string
      description:
        description: Immutable. An optional description of this resource.
        type: string
      distributionPolicy:
        description: Policy specifying the intended distribution of managed instances
          across zones in a regional managed instance group.
        properties:
          targetShape:
            description: 'The distribution shape to which the group converges either
              proactively or on resize events (depending on the value set in `updatePolicy.instanceRedistributionType`).
              Possible values: TARGET_SHAPE_UNSPECIFIED, ANY, BALANCED, ANY_SINGLE_ZONE'
            type: string
          zones:
            description: Immutable. Zones where the regional managed instance group
              will create and manage its instances.
            items:
              properties:
                zone:
                  description: Immutable. The URL of the [zone](/compute/docs/regions-zones/#available).
                    The zone must exist in the region where the managed instance group
                    is located.
                  type: string
              type: object
            type: array
        type: object
      failoverAction:
        description: 'The action to perform in case of zone failure. Only one value
          is supported, `NO_FAILOVER`. The default is `NO_FAILOVER`. Possible values:
          UNKNOWN, NO_FAILOVER'
        type: string
      instanceTemplateRef:
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
              The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run `recreateInstances`, run `applyUpdatesToInstances`, or set the group's `updatePolicy.type` to `PROACTIVE`.

              Allowed value: The `selfLink` field of a `ComputeInstanceTemplate` resource.
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      location:
        description: Immutable. The location of this resource.
        type: string
      namedPorts:
        description: Immutable. Named ports configured for the Instance Groups complementary
          to this Instance Group Manager.
        items:
          properties:
            name:
              description: Immutable. The name for this named port. The name must
                be 1-63 characters long, and comply with [RFC1035](https://www.ietf.org/rfc/rfc1035.txt).
              type: string
            port:
              description: Immutable. The port number, which can be a value between
                1 and 65535.
              format: int64
              type: integer
          type: object
        type: array
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
            description: |-
              The service account to be used as credentials for all operations performed by the managed instance group on instances. The service accounts needs all permissions required to create and delete instances. By default, the service account: {projectNumber}@cloudservices.gserviceaccount.com is used.

              Allowed value: The `email` field of an `IAMServiceAccount` resource.
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      spec:
        required:
        - projectRef
        - targetSize
        type: object
      statefulPolicy:
        description: Stateful configuration for this Instanced Group Manager
        properties:
          preservedState:
            properties:
              disks:
                additionalProperties:
                  properties:
                    autoDelete:
                      description: 'These stateful disks will never be deleted during
                        autohealing, update or VM instance recreate operations. This
                        flag is used to configure if the disk should be deleted after
                        it is no longer used by the group, e.g. when the given instance
                        or the whole group is deleted. Note: disks attached in READ_ONLY
                        mode cannot be auto-deleted. Possible values: NEVER, ON_PERMANENT_INSTANCE_DELETION'
                      type: string
                  type: object
                description: Disks created on the instances that will be preserved
                  on instance delete, update, etc. This map is keyed with the device
                  names of the disks.
                type: object
              externalIps:
                additionalProperties:
                  properties:
                    autoDelete:
                      description: 'These stateful IPs will never be released during
                        autohealing, update or VM instance recreate operations. This
                        flag is used to configure if the IP reservation should be
                        deleted after it is no longer used by the group, e.g. when
                        the given instance or the whole group is deleted. Possible
                        values: NEVER, ON_PERMANENT_INSTANCE_DELETION'
                      type: string
                  type: object
                description: External network IPs assigned to the instances that will
                  be preserved on instance delete, update, etc. This map is keyed
                  with the network interface name.
                type: object
              internalIps:
                additionalProperties:
                  properties:
                    autoDelete:
                      description: 'These stateful IPs will never be released during
                        autohealing, update or VM instance recreate operations. This
                        flag is used to configure if the IP reservation should be
                        deleted after it is no longer used by the group, e.g. when
                        the given instance or the whole group is deleted. Possible
                        values: NEVER, ON_PERMANENT_INSTANCE_DELETION'
                      type: string
                  type: object
                description: Internal network IPs assigned to the instances that will
                  be preserved on instance delete, update, etc. This map is keyed
                  with the network interface name.
                type: object
            type: object
        type: object
      targetPools:
        items:
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
              description: 'Allowed value: The `selfLink` field of a `ComputeTargetPool`
                resource.'
              type: string
            name:
              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
              type: string
            namespace:
              description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
              type: string
          type: object
        type: array
      targetSize:
        description: The target number of running instances for this managed instance
          group. You can reduce this number by using the instanceGroupManager deleteInstances
          or abandonInstances methods. Resizing the group also changes this number.
        format: int64
        type: integer
      updatePolicy:
        description: The update policy for this managed instance group.
        properties:
          instanceRedistributionType:
            description: 'The [instance redistribution policy](/compute/docs/instance-groups/regional-migs#proactive_instance_redistribution)
              for regional managed instance groups. Valid values are: - `PROACTIVE`
              (default): The group attempts to maintain an even distribution of VM
              instances across zones in the region. - `NONE`: For non-autoscaled groups,
              proactive redistribution is disabled.'
            type: string
          maxSurge:
            description: The maximum number of instances that can be created above
              the specified `targetSize` during the update process. This value can
              be either a fixed number or, if the group has 10 or more instances,
              a percentage. If you set a percentage, the number of instances is rounded
              if necessary. The default value for `maxSurge` is a fixed value equal
              to the number of zones in which the managed instance group operates.
              At least one of either `maxSurge` or `maxUnavailable` must be greater
              than 0. Learn more about [`maxSurge`](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#max_surge).
            properties:
              fixed:
                description: Specifies a fixed number of VM instances. This must be
                  a positive integer.
                format: int64
                type: integer
              percent:
                description: Specifies a percentage of instances between 0 to 100%,
                  inclusive. For example, specify `80` for 80%.
                format: int64
                type: integer
            type: object
          maxUnavailable:
            description: 'The maximum number of instances that can be unavailable
              during the update process. An instance is considered available if all
              of the following conditions are satisfied: - The instance''s [status](/compute/docs/instances/checking-instance-status)
              is `RUNNING`. - If there is a [health check](/compute/docs/instance-groups/autohealing-instances-in-migs)
              on the instance group, the instance''s health check status must be `HEALTHY`
              at least once. If there is no health check on the group, then the instance
              only needs to have a status of `RUNNING` to be considered available.
              This value can be either a fixed number or, if the group has 10 or more
              instances, a percentage. If you set a percentage, the number of instances
              is rounded if necessary. The default value for `maxUnavailable` is a
              fixed value equal to the number of zones in which the managed instance
              group operates. At least one of either `maxSurge` or `maxUnavailable`
              must be greater than 0. Learn more about [`maxUnavailable`](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#max_unavailable).'
            properties:
              fixed:
                description: Specifies a fixed number of VM instances. This must be
                  a positive integer.
                format: int64
                type: integer
              percent:
                description: Specifies a percentage of instances between 0 to 100%,
                  inclusive. For example, specify `80` for 80%.
                format: int64
                type: integer
            type: object
          minReadySec:
            description: Minimum number of seconds to wait for after a newly created
              instance becomes available. This value must be from range [0, 3600].
            format: int64
            type: integer
          minimalAction:
            description: Minimal action to be taken on an instance. You can specify
              either `RESTART` to restart existing instances or `REPLACE` to delete
              and create new instances from the target template. If you specify a
              `RESTART`, the Updater will attempt to perform that action only. However,
              if the Updater determines that the minimal action you specify is not
              enough to perform the update, it might perform a more disruptive action.
            type: string
          mostDisruptiveAllowedAction:
            description: Most disruptive action that is allowed to be taken on an
              instance. You can specify either `NONE` to forbid any actions, `REFRESH`
              to allow actions that do not need instance restart, `RESTART` to allow
              actions that can be applied without instance replacing or `REPLACE`
              to allow all possible actions. If the Updater determines that the minimal
              update action needed is more disruptive than most disruptive allowed
              action you specify it will not perform the update at all.
            type: string
          replacementMethod:
            description: 'What action should be used to replace instances. See minimal_action.REPLACE
              Possible values: SUBSTITUTE, RECREATE'
            type: string
          type:
            description: The type of update process. You can specify either `PROACTIVE`
              so that the instance group manager proactively executes actions in order
              to bring instances to their target versions or `OPPORTUNISTIC` so that
              no action is proactively executed but the update will be performed as
              part of other actions (for example, resizes or `recreateInstances` calls).
            type: string
        type: object
required:
- spec
type: object

```
