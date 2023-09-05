# openAPI schema part2 for storagetransfer.cnrm.cloud.google.com.v1beta1.StorageTransferJob

## schema

```yaml
properties:
  spec:
    properties:
      transferSpec:
        properties:
          transferSpec:
            description: Transfer specification.
            properties:
              awsS3DataSource:
                description: An AWS S3 data source.
                properties:
                  awsAccessKey:
                    description: AWS credentials block.
                    properties:
                      accessKeyId:
                        description: AWS Key ID.
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
                            description: Value of the field. Cannot be used if 'valueFrom'
                              is specified.
                            type: string
                          valueFrom:
                            description: Source for the field's value. Cannot be used
                              if 'value' is specified.
                            properties:
                              secretKeyRef:
                                description: Reference to a value with the given key
                                  in the given Secret in the resource's namespace.
                                properties:
                                  key:
                                    description: Key that identifies the value to
                                      be extracted.
                                    type: string
                                  name:
                                    description: Name of the Secret to extract a value
                                      from.
                                    type: string
                                required:
                                - name
                                - key
                                type: object
                            type: object
                        type: object
                      secretAccessKey:
                        description: AWS Secret Access Key.
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
                            description: Value of the field. Cannot be used if 'valueFrom'
                              is specified.
                            type: string
                          valueFrom:
                            description: Source for the field's value. Cannot be used
                              if 'value' is specified.
                            properties:
                              secretKeyRef:
                                description: Reference to a value with the given key
                                  in the given Secret in the resource's namespace.
                                properties:
                                  key:
                                    description: Key that identifies the value to
                                      be extracted.
                                    type: string
                                  name:
                                    description: Name of the Secret to extract a value
                                      from.
                                    type: string
                                required:
                                - name
                                - key
                                type: object
                            type: object
                        type: object
                    required:
                    - accessKeyId
                    - secretAccessKey
                    type: object
                  bucketName:
                    description: S3 Bucket name.
                    type: string
                  path:
                    description: S3 Bucket path in bucket to transfer.
                    type: string
                  roleArn:
                    description: The Amazon Resource Name (ARN) of the role to support
                      temporary credentials via 'AssumeRoleWithWebIdentity'. For more
                      information about ARNs, see [IAM ARNs](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns).
                      When a role ARN is provided, Transfer Service fetches temporary
                      credentials for the session using a 'AssumeRoleWithWebIdentity'
                      call for the provided role using the [GoogleServiceAccount][]
                      for this project.
                    type: string
                required:
                - bucketName
                type: object
              azureBlobStorageDataSource:
                description: An Azure Blob Storage data source.
                properties:
                  azureCredentials:
                    description: ' Credentials used to authenticate API requests to
                      Azure.'
                    properties:
                      sasToken:
                        description: Azure shared access signature.
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
                            description: Value of the field. Cannot be used if 'valueFrom'
                              is specified.
                            type: string
                          valueFrom:
                            description: Source for the field's value. Cannot be used
                              if 'value' is specified.
                            properties:
                              secretKeyRef:
                                description: Reference to a value with the given key
                                  in the given Secret in the resource's namespace.
                                properties:
                                  key:
                                    description: Key that identifies the value to
                                      be extracted.
                                    type: string
                                  name:
                                    description: Name of the Secret to extract a value
                                      from.
                                    type: string
                                required:
                                - name
                                - key
                                type: object
                            type: object
                        type: object
                    required:
                    - sasToken
                    type: object
                  container:
                    description: The container to transfer from the Azure Storage
                      account.
                    type: string
                  path:
                    description: Root path to transfer objects. Must be an empty string
                      or full path name that ends with a '/'. This field is treated
                      as an object prefix. As such, it should generally not begin
                      with a '/'.
                    type: string
                  storageAccount:
                    description: The name of the Azure Storage account.
                    type: string
                required:
                - azureCredentials
                - container
                - storageAccount
                type: object
              gcsDataSink:
                description: A Google Cloud Storage data sink.
                properties:
                  bucketRef:
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
                        description: 'Allowed value: The `name` field of a `StorageBucket`
                          resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  path:
                    description: Google Cloud Storage path in bucket to transfer.
                    type: string
                required:
                - bucketRef
                type: object
              gcsDataSource:
                description: A Google Cloud Storage data source.
                properties:
                  bucketRef:
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
                        description: 'Allowed value: The `name` field of a `StorageBucket`
                          resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  path:
                    description: Google Cloud Storage path in bucket to transfer.
                    type: string
                required:
                - bucketRef
                type: object
              httpDataSource:
                description: A HTTP URL data source.
                properties:
                  listUrl:
                    description: The URL that points to the file that stores the object
                      list entries. This file must allow public access. Currently,
                      only URLs with HTTP and HTTPS schemes are supported.
                    type: string
                required:
                - listUrl
                type: object
              objectConditions:
                description: Only objects that satisfy these object conditions are
                  included in the set of data source and data sink objects. Object
                  conditions based on objects' last_modification_time do not exclude
                  objects in a data sink.
                properties:
                  excludePrefixes:
                    description: exclude_prefixes must follow the requirements described
                      for include_prefixes.
                    items:
                      type: string
                    type: array
                  includePrefixes:
                    description: If include_refixes is specified, objects that satisfy
                      the object conditions must have names that start with one of
                      the include_prefixes and that do not start with any of the exclude_prefixes.
                      If include_prefixes is not specified, all objects except those
                      that have names starting with one of the exclude_prefixes must
                      satisfy the object conditions.
                    items:
                      type: string
                    type: array
                  lastModifiedBefore:
                    description: 'If specified, only objects with a "last modification
                      time" before this timestamp and objects that don''t have a "last
                      modification time" are transferred. A timestamp in RFC3339 UTC
                      "Zulu" format, with nanosecond resolution and up to nine fractional
                      digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".'
                    type: string
                  lastModifiedSince:
                    description: 'If specified, only objects with a "last modification
                      time" on or after this timestamp and objects that don''t have
                      a "last modification time" are transferred. A timestamp in RFC3339
                      UTC "Zulu" format, with nanosecond resolution and up to nine
                      fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".'
                    type: string
                  maxTimeElapsedSinceLastModification:
                    description: 'A duration in seconds with up to nine fractional
                      digits, terminated by ''s''. Example: "3.5s".'
                    type: string
                  minTimeElapsedSinceLastModification:
                    description: 'A duration in seconds with up to nine fractional
                      digits, terminated by ''s''. Example: "3.5s".'
                    type: string
                type: object
              posixDataSink:
                description: A POSIX filesystem data sink.
                properties:
                  rootDirectory:
                    description: Root directory path to the filesystem.
                    type: string
                required:
                - rootDirectory
                type: object
              posixDataSource:
                description: A POSIX filesystem data source.
                properties:
                  rootDirectory:
                    description: Root directory path to the filesystem.
                    type: string
                required:
                - rootDirectory
                type: object
              sinkAgentPoolName:
                description: Immutable. Specifies the agent pool name associated with
                  the posix data source. When unspecified, the default name is used.
                type: string
              sourceAgentPoolName:
                description: Immutable. Specifies the agent pool name associated with
                  the posix data source. When unspecified, the default name is used.
                type: string
              transferOptions:
                description: Characteristics of how to treat files from datasource
                  and sink during job. If the option delete_objects_unique_in_sink
                  is true, object conditions based on objects' last_modification_time
                  are ignored and do not exclude objects in a data source or a data
                  sink.
                properties:
                  deleteObjectsFromSourceAfterTransfer:
                    description: Whether objects should be deleted from the source
                      after they are transferred to the sink. Note that this option
                      and delete_objects_unique_in_sink are mutually exclusive.
                    type: boolean
                  deleteObjectsUniqueInSink:
                    description: Whether objects that exist only in the sink should
                      be deleted. Note that this option and delete_objects_from_source_after_transfer
                      are mutually exclusive.
                    type: boolean
                  overwriteObjectsAlreadyExistingInSink:
                    description: Whether overwriting objects that already exist in
                      the sink is allowed.
                    type: boolean
                  overwriteWhen:
                    description: When to overwrite objects that already exist in the
                      sink. If not set, overwrite behavior is determined by overwriteObjectsAlreadyExistingInSink.
                    type: string
                type: object
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
      creationTime:
        description: When the Transfer Job was created.
        type: string
      deletionTime:
        description: When the Transfer Job was deleted.
        type: string
      lastModificationTime:
        description: When the Transfer Job was last modified.
        type: string
      name:
        description: The name of the Transfer Job.
        type: string
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
