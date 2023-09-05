# openAPI schema part0 for compute.cnrm.cloud.google.com.v1alpha1.ComputeRegionPerInstanceConfig

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
      minimalAction:
        description: |-
          The minimal action to perform on the instance during an update.
          Default is 'NONE'. Possible values are:
          * REPLACE
          * RESTART
          * REFRESH
          * NONE.
        type: string
      mostDisruptiveAllowedAction:
        description: |-
          The most disruptive action to perform on the instance during an update.
          Default is 'REPLACE'. Possible values are:
          * REPLACE
          * RESTART
          * REFRESH
          * NONE.
        type: string
      preservedState:
        description: The preserved state for this instance.
        properties:
          disk:
            description: Stateful disks for the instance.
            items:
              properties:
                deleteRule:
                  description: |-
                    A value that prescribes what should happen to the stateful disk when the VM instance is deleted.
                    The available options are 'NEVER' and 'ON_PERMANENT_INSTANCE_DELETION'.
                    'NEVER' - detach the disk when the VM is deleted, but do not delete the disk.
                    'ON_PERMANENT_INSTANCE_DELETION' will delete the stateful disk when the VM is permanently
                    deleted from the instance group. Default value: "NEVER" Possible values: ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"].
                  type: string
                deviceName:
                  description: A unique device name that is reflected into the /dev/
                    tree of a Linux operating system running within the instance.
                  type: string
                mode:
                  description: 'The mode of the disk. Default value: "READ_WRITE"
                    Possible values: ["READ_ONLY", "READ_WRITE"].'
                  type: string
                source:
                  description: |-
                    The URI of an existing persistent disk to attach under the specified device-name in the format
                    'projects/project-id/zones/zone/disks/disk-name'.
                  type: string
              required:
              - deviceName
              - source
              type: object
            type: array
          externalIp:
            description: Preserved external IPs defined for this instance. This map
              is keyed with the name of the network interface.
            items:
              properties:
                autoDelete:
                  description: 'These stateful IPs will never be released during autohealing,
                    update or VM instance recreate operations. This flag is used to
                    configure if the IP reservation should be deleted after it is
                    no longer used by the group, e.g. when the given instance or the
                    whole group is deleted. Default value: "NEVER" Possible values:
                    ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"].'
                  type: string
                interfaceName:
                  type: string
                ipAddress:
                  description: Ip address representation.
                  properties:
                    address:
                      description: The URL of the reservation for this IP address.
                      type: string
                  type: object
              required:
              - interfaceName
              type: object
            type: array
          internalIp:
            description: Preserved internal IPs defined for this instance. This map
              is keyed with the name of the network interface.
            items:
              properties:
                autoDelete:
                  description: 'These stateful IPs will never be released during autohealing,
                    update or VM instance recreate operations. This flag is used to
                    configure if the IP reservation should be deleted after it is
                    no longer used by the group, e.g. when the given instance or the
                    whole group is deleted. Default value: "NEVER" Possible values:
                    ["NEVER", "ON_PERMANENT_INSTANCE_DELETION"].'
                  type: string
                interfaceName:
                  type: string
                ipAddress:
                  description: Ip address representation.
                  properties:
                    address:
                      description: The URL of the reservation for this IP address.
                      type: string
                  type: object
              required:
              - interfaceName
              type: object
            type: array
          metadata:
            additionalProperties:
              type: string
            description: Preserved metadata defined for this instance. This is a list
              of key->value pairs.
            type: object
        type: object
      projectRef:
        description: The project that this resource belongs to.
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
            description: 'Allowed value: The `name` field of a `Project` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      region:
        description: Immutable. Region where the containing instance group manager
          is located.
        type: string
      regionInstanceGroupManagerRef:
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
            description: 'Allowed value: The `name` field of a `ComputeRegionInstanceGroupManager`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      removeInstanceStateOnDestroy:
        description: |-
          When true, deleting this config will immediately remove any specified state from the underlying instance.
          When false, deleting this config will *not* immediately remove any state from the underlying instance.
          State will be removed on the next instance recreation or update.
        type: boolean
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
    required:
    - projectRef
    - region
    - regionInstanceGroupManagerRef
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
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
    type: object
required:
- spec
type: object

```