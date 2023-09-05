# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeInstanceTemplate

## schema

```yaml
properties:
  spec:
    properties:
      advancedMachineFeatures:
        description: Immutable. Controls for advanced machine-related behavior features.
        properties:
          enableNestedVirtualization:
            description: Immutable. Whether to enable nested virtualization or not.
            type: boolean
          threadsPerCore:
            description: Immutable. The number of threads per physical core. To disable
              simultaneous multithreading (SMT) set this to 1. If unset, the maximum
              number of threads supported per core by the underlying processor is
              assumed.
            type: integer
          visibleCoreCount:
            description: Immutable. The number of physical cores to expose to an instance.
              Multiply by the number of threads per core to compute the total number
              of virtual CPUs to expose to the instance. If unset, the number of cores
              is inferred from the instance\'s nominal CPU count and the underlying
              platform\'s SMT width.
            type: integer
        type: object
      canIpForward:
        description: Immutable. Whether to allow sending and receiving of packets
          with non-matching source or destination IPs. This defaults to false.
        type: boolean
      confidentialInstanceConfig:
        description: Immutable. The Confidential VM config being used by the instance.
          on_host_maintenance has to be set to TERMINATE or this will fail to create.
        properties:
          enableConfidentialCompute:
            description: Immutable. Defines whether the instance should have confidential
              compute enabled.
            type: boolean
        required:
        - enableConfidentialCompute
        type: object
      description:
        description: Immutable. A brief description of this resource.
        type: string
      disk:
        description: Immutable. Disks to attach to instances created from this template.
          This can be specified multiple times for multiple disks.
        items:
          properties:
            autoDelete:
              description: Immutable. Whether or not the disk should be auto-deleted.
                This defaults to true.
              type: boolean
            boot:
              description: Immutable. Indicates that this is a boot disk.
              type: boolean
            deviceName:
              description: Immutable. A unique device name that is reflected into
                the /dev/ tree of a Linux operating system running within the instance.
                If not specified, the server chooses a default device name to apply
                to this disk.
              type: string
            diskEncryptionKey:
              description: Immutable. Encrypts or decrypts a disk using a customer-supplied
                encryption key.
              properties:
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
              required:
              - kmsKeyRef
              type: object
            diskName:
              description: Immutable. Name of the disk. When not provided, this defaults
                to the name of the instance.
              type: string
            diskSizeGb:
              description: Immutable. The size of the image in gigabytes. If not specified,
                it will inherit the size of its base image. For SCRATCH disks, the
                size must be one of 375 or 3000 GB, with a default of 375 GB.
              type: integer
            diskType:
              description: Immutable. The Google Compute Engine disk type. Such as
                "pd-ssd", "local-ssd", "pd-balanced" or "pd-standard".
              type: string
            interface:
              description: Immutable. Specifies the disk interface to use for attaching
                this disk.
              type: string
            labels:
              additionalProperties:
                type: string
              description: Immutable. A set of key/value label pairs to assign to
                disks,.
              type: object
            mode:
              description: Immutable. The mode in which to attach this disk, either
                READ_WRITE or READ_ONLY. If you are attaching or creating a boot disk,
                this must read-write mode.
              type: string
            resourcePolicies:
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
                    description: 'Allowed value: The `selfLink` field of a `ComputeResourcePolicy`
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
                  description: 'Allowed value: The `name` field of a `ComputeDisk`
                    resource.'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
            sourceImageEncryptionKey:
              description: |-
                Immutable. The customer-supplied encryption key of the source
                image. Required if the source image is protected by a
                customer-supplied encryption key.

                Instance templates do not store customer-supplied
                encryption keys, so you cannot create disks for
                instances in a managed instance group if the source
                images are encrypted with your own keys.
              properties:
                kmsKeySelfLinkRef:
                  description: |-
                    The self link of the encryption key that is stored in Google Cloud
                    KMS.
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
                kmsKeyServiceAccountRef:
                  description: |-
                    The service account being used for the encryption request for the
                    given KMS key. If absent, the Compute Engine default service account
                    is used.
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
              required:
              - kmsKeySelfLinkRef
              type: object
            sourceImageRef:
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
            sourceSnapshotEncryptionKey:
              description: Immutable. The customer-supplied encryption key of the
                source snapshot.
              properties:
                kmsKeySelfLinkRef:
                  description: |-
                    The self link of the encryption key that is stored in Google Cloud
                    KMS.
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
                kmsKeyServiceAccountRef:
                  description: |-
                    The service account being used for the encryption request for the
                    given KMS key. If absent, the Compute Engine default service account
                    is used.
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
              required:
              - kmsKeySelfLinkRef
              type: object
            sourceSnapshotRef:
              description: |-
                The source snapshot to create this disk. When creating a new
                instance, one of initializeParams.sourceSnapshot,
                initializeParams.sourceImage, or disks.source is required except for
                local SSD.
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
                  description: 'Allowed value: The `selfLink` field of a `ComputeSnapshot`
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
              description: Immutable. The type of Google Compute Engine disk, can
                be either "SCRATCH" or "PERSISTENT".
              type: string
          type: object
        type: array
      enableDisplay:
        description: 'Immutable. Enable Virtual Displays on this instance. Note: allow_stopping_for_update
          must be set to true in order to update this field.'
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
              description: Immutable. The accelerator type resource to expose to this
                instance. E.g. nvidia-tesla-k80.
              type: string
          required:
          - count
          - type
          type: object
        type: array
      instanceDescription:
        description: Immutable. A description of the instance.
        type: string
      machineType:
        description: Immutable. The machine type to create. To create a machine with
          a custom type (such as extended memory), format the value like custom-VCPUS-MEM_IN_MB
          like custom-6-20480 for 6 vCPU and 20GB of RAM.
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
        description: Immutable. An alternative to using the startup-script metadata
          key, mostly to match the compute_instance resource. This replaces the startup-script
          metadata key on the created instance and thus the two mechanisms are not
          allowed to be used simultaneously.
        type: string
      minCpuPlatform:
        description: Immutable. Specifies a minimum CPU platform. Applicable values
          are the friendly names of CPU platforms, such as Intel Haswell or Intel
          Skylake.
        type: string
      namePrefix:
        description: Immutable. Creates a unique name beginning with the specified
          prefix. Conflicts with name.
        type: string
      spec:
        required:
        - disk
        - machineType
        type: object
required:
- spec
type: object

```
