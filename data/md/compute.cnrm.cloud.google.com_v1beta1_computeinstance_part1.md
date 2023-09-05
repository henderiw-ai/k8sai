# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeInstance

## schema

```yaml
properties:
  spec:
    properties:
      advancedMachineFeatures:
        description: Controls for advanced machine-related behavior features.
        properties:
          enableNestedVirtualization:
            description: Whether to enable nested virtualization or not.
            type: boolean
          threadsPerCore:
            description: The number of threads per physical core. To disable simultaneous
              multithreading (SMT) set this to 1. If unset, the maximum number of
              threads supported per core by the underlying processor is assumed.
            type: integer
          visibleCoreCount:
            description: The number of physical cores to expose to an instance. Multiply
              by the number of threads per core to compute the total number of virtual
              CPUs to expose to the instance. If unset, the number of cores is inferred
              from the instance\'s nominal CPU count and the underlying platform\'s
              SMT width.
            type: integer
        type: object
      attachedDisk:
        description: List of disks attached to the instance.
        items:
          properties:
            deviceName:
              description: Name with which the attached disk is accessible under /dev/disk/by-id/.
              type: string
            diskEncryptionKeyRaw:
              description: A 256-bit customer-supplied encryption key, encoded in
                RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link
                and disk_encryption_key_raw may be set.
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
                      description: Reference to a value with the given key in the
                        given Secret in the resource's namespace.
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
            diskEncryptionKeySha256:
              description: The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
                encryption key that protects this resource.
              type: string
            kmsKeyRef:
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
            mode:
              description: Read/write mode for the disk. One of "READ_ONLY" or "READ_WRITE".
              type: string
            sourceDiskRef:
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
                  description: 'Allowed value: The `selfLink` field of a `ComputeDisk`
                    resource.'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
          required:
          - sourceDiskRef
          type: object
        type: array
      bootDisk:
        description: Immutable. The boot disk for the instance.
        properties:
          autoDelete:
            description: Immutable. Whether the disk will be auto-deleted when the
              instance is deleted.
            type: boolean
          deviceName:
            description: Immutable. Name with which attached disk will be accessible
              under /dev/disk/by-id/.
            type: string
          diskEncryptionKeyRaw:
            description: Immutable. A 256-bit customer-supplied encryption key, encoded
              in RFC 4648 base64 to encrypt this disk. Only one of kms_key_self_link
              and disk_encryption_key_raw may be set.
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
          diskEncryptionKeySha256:
            description: The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
              encryption key that protects this resource.
            type: string
          initializeParams:
            description: Immutable. Parameters with which a disk was created alongside
              the instance.
            properties:
              labels:
                description: Immutable. A set of key/value label pairs assigned to
                  the disk.
                type: object
                x-kubernetes-preserve-unknown-fields: true
              resourceManagerTags:
                description: Immutable. A map of resource manager tags. Resource manager
                  tag keys and values have the same definition as resource manager
                  tags. Keys must be in the format tagKeys/{tag_key_id}, and values
                  are in the format tagValues/456. The field is ignored (both PUT
                  & PATCH) when empty.
                type: object
                x-kubernetes-preserve-unknown-fields: true
              size:
                description: Immutable. The size of the image in gigabytes.
                type: integer
              sourceImageRef:
                description: Immutable. The image from which to initialize this disk.
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
                    description: 'Allowed value: The `selfLink` field of a `ComputeImage`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              type:
                description: Immutable. The Google Compute Engine disk type. Such
                  as pd-standard, pd-ssd or pd-balanced.
                type: string
            type: object
          kmsKeyRef:
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
          mode:
            description: Immutable. Read/write mode for the disk. One of "READ_ONLY"
              or "READ_WRITE".
            type: string
          sourceDiskRef:
            description: Immutable. The source disk used to create this disk.
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
                description: 'Allowed value: The `selfLink` field of a `ComputeDisk`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
        type: object
      canIpForward:
        description: Whether sending and receiving of packets with non-matching source
          or destination IPs is allowed.
        type: boolean
      confidentialInstanceConfig:
        description: Immutable. The Confidential VM config being used by the instance.  on_host_maintenance
          has to be set to TERMINATE or this will fail to create.
        properties:
          enableConfidentialCompute:
            description: Defines whether the instance should have confidential compute
              enabled.
            type: boolean
        required:
        - enableConfidentialCompute
        type: object
      deletionProtection:
        description: Whether deletion protection is enabled on this instance.
        type: boolean
      description:
        description: Immutable. A brief description of the resource.
        type: string
      desiredStatus:
        description: Desired status of the instance. Either "RUNNING" or "TERMINATED".
        type: string
      enableDisplay:
        description: Whether the instance has virtual displays enabled.
        type: boolean
      guestAccelerator:
        description: Immutable. List of the type and count of accelerator cards attached
          to the instance.
        items:
          properties:
            count:
              description: Immutable. The number of the guest accelerator cards exposed
                to this instance.
              type: integer
            type:
              description: Immutable. The accelerator type resource exposed to this
                instance. E.g. nvidia-tesla-k80.
              type: string
          required:
          - count
          - type
          type: object
        type: array
      hostname:
        description: Immutable. A custom hostname for the instance. Must be a fully
          qualified DNS name and RFC-1035-valid. Valid format is a series of labels
          1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]),
          concatenated with periods. The entire hostname must not exceed 253 characters.
          Changing this forces a new resource to be created.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeInstanceTemplate`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      machineType:
        description: The machine type to create.
        type: string
      metadata:
        items:
          properties:
            key:
              type: string
            value:
              type: string
          required:
          - key
          - value
          type: object
        type: array
      metadataStartupScript:
        description: Immutable. Metadata startup scripts made available within the
          instance.
        type: string
      minCpuPlatform:
        description: The minimum CPU platform specified for the VM instance.
        type: string
      spec:
        anyOf:
        - required:
          - bootDisk
          - machineType
          - networkInterface
          - zone
        - required:
          - instanceTemplateRef
          - zone
        type: object
required:
- spec
type: object

```
