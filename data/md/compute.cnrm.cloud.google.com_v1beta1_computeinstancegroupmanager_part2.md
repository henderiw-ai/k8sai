# openAPI schema part2 for compute.cnrm.cloud.google.com.v1beta1.ComputeInstanceGroupManager

## schema

```yaml
properties:
  spec:
    properties:
      versions:
        properties:
          versions:
            description: Specifies the instance templates used by this managed instance
              group to create instances. Each version is defined by an `instanceTemplate`
              and a `name`. Every version can appear at most once per instance group.
              This field overrides the top-level `instanceTemplate` field. Read more
              about the [relationships between these fields](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#relationship_between_versions_and_instancetemplate_properties_for_a_managed_instance_group).
              Exactly one `version` must leave the `targetSize` field unset. That
              version will be applied to all remaining instances. For more information,
              read about [canary updates](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update).
            items:
              properties:
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
                        The URL of the instance template that is specified for this managed instance group. The group uses this template to create new instances in the managed instance group until the `targetSize` for this version is reached. The templates for existing instances in the group do not change unless you run `recreateInstances`, run `applyUpdatesToInstances`, or set the group's `updatePolicy.type` to `PROACTIVE`; in those cases, existing instances are updated until the `targetSize` for this version is reached.

                        Allowed value: The `selfLink` field of a `ComputeInstanceTemplate` resource.
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
                name:
                  description: Name of the version. Unique among all versions in the
                    scope of this managed instance group.
                  type: string
                targetSize:
                  description: 'Specifies the intended number of instances to be created
                    from the `instanceTemplate`. The final number of instances created
                    from the template will be equal to: - If expressed as a fixed
                    number, the minimum of either `targetSize.fixed` or `instanceGroupManager.targetSize`
                    is used. - if expressed as a `percent`, the `targetSize` would
                    be `(targetSize.percent/100 * InstanceGroupManager.targetSize)`
                    If there is a remainder, the number is rounded. If unset, this
                    version will update any remaining instances not updated by another
                    `version`. Read [Starting a canary update](/compute/docs/instance-groups/rolling-out-updates-to-managed-instance-groups#starting_a_canary_update)
                    for more information.'
                  properties:
                    calculated:
                      description: '[Output Only] Absolute value of VM instances calculated
                        based on the specific mode. - If the value is `fixed`, then
                        the `calculated` value is equal to the `fixed` value. - If
                        the value is a `percent`, then the `calculated` value is `percent`/100
                        * `targetSize`. For example, the `calculated` value of a 80%
                        of a managed instance group with 150 instances would be (80/100
                        * 150) = 120 VM instances. If there is a remainder, the number
                        is rounded.'
                      format: int64
                      type: integer
                    fixed:
                      description: Specifies a fixed number of VM instances. This
                        must be a positive integer.
                      format: int64
                      type: integer
                    percent:
                      description: Specifies a percentage of instances between 0 to
                        100%, inclusive. For example, specify `80` for 80%.
                      format: int64
                      type: integer
                  type: object
              type: object
            type: array
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
      creationTimestamp:
        description: The creation timestamp for this managed instance group in \[RFC3339\](https://www.ietf.org/rfc/rfc3339.txt)
          text format.
        type: string
      currentActions:
        description: '[Output Only] The list of instance actions and the number of
          instances in this managed instance group that are scheduled for each of
          those actions.'
        properties:
          abandoning:
            description: '[Output Only] The total number of instances in the managed
              instance group that are scheduled to be abandoned. Abandoning an instance
              removes it from the managed instance group without deleting it.'
            format: int64
            type: integer
          creating:
            description: '[Output Only] The number of instances in the managed instance
              group that are scheduled to be created or are currently being created.
              If the group fails to create any of these instances, it tries again
              until it creates the instance successfully. If you have disabled creation
              retries, this field will not be populated; instead, the `creatingWithoutRetries`
              field will be populated.'
            format: int64
            type: integer
          creatingWithoutRetries:
            description: '[Output Only] The number of instances that the managed instance
              group will attempt to create. The group attempts to create each instance
              only once. If the group fails to create any of these instances, it decreases
              the group''s `targetSize` value accordingly.'
            format: int64
            type: integer
          deleting:
            description: '[Output Only] The number of instances in the managed instance
              group that are scheduled to be deleted or are currently being deleted.'
            format: int64
            type: integer
          none:
            description: '[Output Only] The number of instances in the managed instance
              group that are running and have no scheduled actions.'
            format: int64
            type: integer
          recreating:
            description: '[Output Only] The number of instances in the managed instance
              group that are scheduled to be recreated or are currently being being
              recreated. Recreating an instance deletes the existing root persistent
              disk and creates a new disk from the image that is defined in the instance
              template.'
            format: int64
            type: integer
          refreshing:
            description: '[Output Only] The number of instances in the managed instance
              group that are being reconfigured with properties that do not require
              a restart or a recreate action. For example, setting or removing target
              pools for the instance.'
            format: int64
            type: integer
          restarting:
            description: '[Output Only] The number of instances in the managed instance
              group that are scheduled to be restarted or are currently being restarted.'
            format: int64
            type: integer
          verifying:
            description: '[Output Only] The number of instances in the managed instance
              group that are being verified. See the `managedInstances[].currentAction`
              property in the `listManagedInstances` method documentation.'
            format: int64
            type: integer
        type: object
      fingerprint:
        description: Fingerprint of this resource. This field may be used in optimistic
          locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date
          fingerprint must be provided in order to update the InstanceGroupManager,
          otherwise the request will fail with error `412 conditionNotMet`. To see
          the latest fingerprint, make a `get()` request to retrieve an InstanceGroupManager.
        type: string
      id:
        description: '[Output Only] A unique identifier for this resource type. The
          server generates this identifier.'
        format: int64
        type: integer
      instanceGroup:
        description: '[Output Only] The URL of the Instance Group resource.'
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      region:
        description: '[Output Only] The URL of the [region](/compute/docs/regions-zones/#available)
          where the managed instance group resides (for regional resources).'
        type: string
      selfLink:
        description: '[Output Only] The URL for this managed instance group. The server
          defines this URL.'
        type: string
      status:
        description: '[Output Only] The status of this managed instance group.'
        properties:
          autoscaler:
            description: '[Output Only] The URL of the [Autoscaler](/compute/docs/autoscaler/)
              that targets this instance group manager.'
            type: string
          isStable:
            description: '[Output Only] A bit indicating whether the managed instance
              group is in a stable state. A stable state means that: none of the instances
              in the managed instance group is currently undergoing any type of change
              (for example, creation, restart, or deletion); no future changes are
              scheduled for instances in the managed instance group; and the managed
              instance group itself is not being modified.'
            type: boolean
          stateful:
            description: '[Output Only] Stateful status of the given Instance Group
              Manager.'
            properties:
              hasStatefulConfig:
                description: '[Output Only] A bit indicating whether the managed instance
                  group has stateful configuration, that is, if you have configured
                  any items in a stateful policy or in per-instance configs. The group
                  might report that it has no stateful config even when there is still
                  some preserved state on a managed instance, for example, if you
                  have deleted all PICs but not yet applied those deletions.'
                type: boolean
              isStateful:
                description: '[Output Only] A bit indicating whether the managed instance
                  group has stateful configuration, that is, if you have configured
                  any items in a stateful policy or in per-instance configs. The group
                  might report that it has no stateful config even when there is still
                  some preserved state on a managed instance, for example, if you
                  have deleted all PICs but not yet applied those deletions. This
                  field is deprecated in favor of has_stateful_config.'
                type: boolean
              perInstanceConfigs:
                description: '[Output Only] Status of per-instance configs on the
                  instance.'
                properties:
                  allEffective:
                    description: A bit indicating if all of the group's per-instance
                      configs (listed in the output of a listPerInstanceConfigs API
                      call) have status `EFFECTIVE` or there are no per-instance-configs.
                    type: boolean
                type: object
            type: object
          versionTarget:
            description: '[Output Only] A status of consistency of Instances'' versions
              with their target version specified by `version` field on Instance Group
              Manager.'
            properties:
              isReached:
                description: '[Output Only] A bit indicating whether version target
                  has been reached in this managed instance group, i.e. all instances
                  are in their target version. Instances'' target version are specified
                  by `version` field on Instance Group Manager.'
                type: boolean
            type: object
        type: object
      updatePolicy:
        properties:
          maxSurge:
            properties:
              calculated:
                description: '[Output Only] Absolute value of VM instances calculated
                  based on the specific mode. - If the value is `fixed`, then the
                  `calculated` value is equal to the `fixed` value. - If the value
                  is a `percent`, then the `calculated` value is `percent`/100 * `targetSize`.
                  For example, the `calculated` value of a 80% of a managed instance
                  group with 150 instances would be (80/100 * 150) = 120 VM instances.
                  If there is a remainder, the number is rounded.'
                format: int64
                type: integer
            type: object
          maxUnavailable:
            properties:
              calculated:
                description: '[Output Only] Absolute value of VM instances calculated
                  based on the specific mode. - If the value is `fixed`, then the
                  `calculated` value is equal to the `fixed` value. - If the value
                  is a `percent`, then the `calculated` value is `percent`/100 * `targetSize`.
                  For example, the `calculated` value of a 80% of a managed instance
                  group with 150 instances would be (80/100 * 150) = 120 VM instances.
                  If there is a remainder, the number is rounded.'
                format: int64
                type: integer
            type: object
        type: object
      zone:
        description: '[Output Only] The URL of a [zone](/compute/docs/regions-zones/#available)
          where the managed instance group is located (for zonal resources).'
        type: string
    type: object
required:
- spec
type: object

```
