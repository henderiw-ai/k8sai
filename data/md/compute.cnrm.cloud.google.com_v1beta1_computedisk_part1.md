# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeDisk

## schema

```yaml
properties:
  spec:
    properties:
      asyncPrimaryDisk:
        description: Immutable. A nested object resource.
        properties:
          diskRef:
            description: Immutable. Primary disk for asynchronous disk replication.
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
        - diskRef
        type: object
      description:
        description: |-
          Immutable. An optional description of this resource. Provide this property when
          you create the resource.
        type: string
      diskEncryptionKey:
        description: |-
          Immutable. Encrypts the disk using a customer-supplied encryption key.

          After you encrypt a disk with a customer-supplied key, you must
          provide the same key if you use the disk later (e.g. to create a disk
          snapshot or an image, or to attach the disk to a virtual machine).

          Customer-supplied encryption keys do not protect access to metadata of
          the disk.

          If you do not provide an encryption key when creating the disk, then
          the disk will be encrypted using an automatically generated key and
          you do not need to provide a key to use the disk later.
        properties:
          kmsKeyRef:
            description: |-
              The encryption key used to encrypt the disk. Your project's Compute
              Engine System service account
              ('service-{{PROJECT_NUMBER}}@compute-system.iam.gserviceaccount.com')
              must have 'roles/cloudkms.cryptoKeyEncrypterDecrypter' to use this
              feature. See
              https://cloud.google.com/compute/docs/disks/customer-managed-encryption#encrypt_a_new_persistent_disk_with_your_own_keys
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
              The service account used for the encryption request for the given KMS key.
              If absent, the Compute Engine Service Agent service account is used.
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
          rawKey:
            description: |-
              Immutable. Specifies a 256-bit customer-supplied encryption key, encoded in
              RFC 4648 base64 to either encrypt or decrypt this resource.
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
                    - key
                    - name
                    type: object
                type: object
            type: object
          rsaEncryptedKey:
            description: |-
              Immutable. Specifies an RFC 4648 base64 encoded, RSA-wrapped 2048-bit
              customer-supplied encryption key to either encrypt or decrypt
              this resource. You can provide either the rawKey or the rsaEncryptedKey.
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
          sha256:
            description: |-
              The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
              encryption key that protects this resource.
            type: string
        type: object
      enableConfidentialCompute:
        description: |-
          Immutable. Whether this disk is using confidential compute mode.
          Note: Only supported on hyperdisk skus, disk_encryption_key is required when setting to true.
        type: boolean
      guestOsFeatures:
        description: |-
          Immutable. A list of features to enable on the guest operating system.
          Applicable only for bootable disks.
        items:
          properties:
            type:
              description: 'Immutable. The type of supported feature. Read [Enabling
                guest operating system features](https://cloud.google.com/compute/docs/images/create-delete-deprecate-private-images#guest-os-features)
                to see a list of available options. Possible values: ["MULTI_IP_SUBNET",
                "SECURE_BOOT", "SEV_CAPABLE", "UEFI_COMPATIBLE", "VIRTIO_SCSI_MULTIQUEUE",
                "WINDOWS", "GVNIC", "SEV_LIVE_MIGRATABLE", "SEV_SNP_CAPABLE", "SUSPEND_RESUME_COMPATIBLE",
                "TDX_CAPABLE"].'
              type: string
          required:
          - type
          type: object
        type: array
      imageRef:
        description: The image from which to initialize this disk.
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
      interface:
        description: DEPRECATED. This field is no longer in use, disk interfaces will
          be automatically determined on attachment. To resolve this issue, remove
          this field from your config. Immutable. Specifies the disk interface to
          use for attaching this disk, which is either SCSI or NVME. The default is
          SCSI.
        type: string
      licenses:
        description: Immutable. Any applicable license URI.
        items:
          type: string
        type: array
      location:
        description: 'Location represents the geographical location of the ComputeDisk.
          Specify a region name or a zone name. Reference: GCP definition of regions/zones
          (https://cloud.google.com/compute/docs/regions-zones/)'
        type: string
      multiWriter:
        description: Immutable. Indicates whether or not the disk can be read/write
          attached to more than one instance.
        type: boolean
      physicalBlockSizeBytes:
        description: |-
          Immutable. Physical block size of the persistent disk, in bytes. If not present
          in a request, a default value is used. Currently supported sizes
          are 4096 and 16384, other sizes may be added in the future.
          If an unsupported value is requested, the error message will list
          the supported values for the caller's project.
        type: integer
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
      provisionedIops:
        description: |-
          Indicates how many IOPS must be provisioned for the disk.
          Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
          allows for an update of IOPS every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it.
        type: integer
      provisionedThroughput:
        description: |-
          Indicates how much Throughput must be provisioned for the disk.
          Note: Updating currently is only supported by hyperdisk skus without the need to delete and recreate the disk, hyperdisk
          allows for an update of Throughput every 4 hours. To update your hyperdisk more frequently, you'll need to manually delete and recreate it.
        type: integer
      replicaZones:
        description: Immutable. URLs of the zones where the disk should be replicated
          to.
        items:
          type: string
        type: array
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      resourcePolicies:
        items:
          description: Resource policies applied to this disk for automatic snapshot
            creations.
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
      size:
        description: |-
          Size of the persistent disk, specified in GB. You can specify this
          field when creating a persistent disk using the 'image' or
          'snapshot' parameter, or specify it alone to create an empty
          persistent disk.

          If you specify this field along with 'image' or 'snapshot',
          the value must not be less than the size of the image
          or the size of the snapshot.

          Upsizing the disk is mutable, but downsizing the disk
          requires re-creating the resource.
        type: integer
      snapshotRef:
        description: The source snapshot used to create this disk.
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
      sourceDiskRef:
        description: The source disk used to create this disk.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeDisk` resource.'
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
        - location
        type: object
required:
- spec
type: object

```
