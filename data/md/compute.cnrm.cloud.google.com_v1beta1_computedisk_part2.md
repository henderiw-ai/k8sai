# openAPI schema part2 for compute.cnrm.cloud.google.com.v1beta1.ComputeDisk

## schema

```yaml
properties:
  spec:
    properties:
      sourceImageEncryptionKey:
        properties:
          sourceImageEncryptionKey:
            description: |-
              Immutable. The customer-supplied encryption key of the source image. Required if
              the source image is protected by a customer-supplied encryption key.
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
                type: string
              sha256:
                description: |-
                  The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
                  encryption key that protects this resource.
                type: string
            type: object
      sourceSnapshotEncryptionKey:
        description: |-
          Immutable. The customer-supplied encryption key of the source snapshot. Required
          if the source snapshot is protected by a customer-supplied encryption
          key.
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
            type: string
          sha256:
            description: |-
              The RFC 4648 base64 encoded SHA-256 hash of the customer-supplied
              encryption key that protects this resource.
            type: string
        type: object
      type:
        description: |-
          Immutable. URL of the disk type resource describing which disk type to use to
          create the disk. Provide this when creating the disk.
        type: string
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
        description: Creation timestamp in RFC3339 text format.
        type: string
      labelFingerprint:
        description: |-
          The fingerprint used for optimistic locking of this resource.  Used
          internally during updates.
        type: string
      lastAttachTimestamp:
        description: Last attach timestamp in RFC3339 text format.
        type: string
      lastDetachTimestamp:
        description: Last detach timestamp in RFC3339 text format.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      selfLink:
        type: string
      sourceDiskId:
        description: |-
          The ID value of the disk used to create this image. This value may
          be used to determine whether the image was taken from the current
          or a previous instance of a given disk name.
        type: string
      sourceImageId:
        description: |-
          The ID value of the image used to create this disk. This value
          identifies the exact image that was used to create this persistent
          disk. For example, if you created the persistent disk from an image
          that was later deleted and recreated under the same name, the source
          image ID would identify the exact version of the image that was used.
        type: string
      sourceSnapshotId:
        description: |-
          The unique ID of the snapshot used to create this disk. This value
          identifies the exact snapshot that was used to create this persistent
          disk. For example, if you created the persistent disk from a snapshot
          that was later deleted and recreated under the same name, the source
          snapshot ID would identify the exact version of the snapshot that was
          used.
        type: string
      users:
        description: |-
          Links to the users of the disk (attached instances) in form:
          project/zones/zone/instances/instance.
        items:
          type: string
        type: array
    type: object
required:
- spec
type: object

```
