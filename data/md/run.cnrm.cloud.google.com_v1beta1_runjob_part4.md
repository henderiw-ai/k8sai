# openAPI schema part4 for run.cnrm.cloud.google.com.v1beta1.RunJob

## schema

```yaml
properties:
  spec:
    properties:
      template:
        properties:
          template:
            properties:
              volumes:
                properties:
                  volumes:
                    description: A list of Volumes to make available to containers.
                    items:
                      properties:
                        emptyDir:
                          description: Ephemeral storage used as a shared volume.
                          properties:
                            medium:
                              description: 'The different types of medium supported
                                for EmptyDir. Default value: "MEMORY" Possible values:
                                ["MEMORY"].'
                              type: string
                            sizeLimit:
                              description: 'Limit on the storage usable by this EmptyDir
                                volume. The size limit is also applicable for memory
                                medium. The maximum usage on memory medium EmptyDir
                                would be the minimum value between the SizeLimit specified
                                here and the sum of memory limits of all containers
                                in a pod. This field''s values are of the ''Quantity''
                                k8s type: https://kubernetes.io/docs/reference/kubernetes-api/common-definitions/quantity/.
                                The default is nil which means that the limit is undefined.
                                More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir.'
                              type: string
                          type: object
                        name:
                          description: Volume's name.
                          type: string
                        secret:
                          description: 'Secret represents a secret that should populate
                            this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret.'
                          properties:
                            defaultMode:
                              description: Integer representation of mode bits to
                                use on created files by default. Must be a value between
                                0000 and 0777 (octal), defaulting to 0444. Directories
                                within the path are not affected by this setting.
                              type: integer
                            items:
                              description: If unspecified, the volume will expose
                                a file whose name is the secret, relative to VolumeMount.mount_path.
                                If specified, the key will be used as the version
                                to fetch from Cloud Secret Manager and the path will
                                be the name of the file exposed in the volume. When
                                items are defined, they must specify a path and a
                                version.
                              items:
                                properties:
                                  mode:
                                    description: Integer octal mode bits to use on
                                      this file, must be a value between 01 and 0777
                                      (octal). If 0 or not set, the Volume's default
                                      mode will be used.
                                    type: integer
                                  path:
                                    description: The relative path of the secret in
                                      the container.
                                    type: string
                                  versionRef:
                                    description: The Cloud Secret Manager secret version.
                                      Can be 'latest' for the latest value or an integer
                                      for a specific version
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
                                        description: 'Allowed value: The `version`
                                          field of a `SecretManagerSecretVersion`
                                          resource.'
                                        type: string
                                      name:
                                        description: 'Name of the referent. More info:
                                          https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                        type: string
                                      namespace:
                                        description: 'Namespace of the referent. More
                                          info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                        type: string
                                    type: object
                                required:
                                - mode
                                - path
                                - versionRef
                                type: object
                              type: array
                            secretRef:
                              description: 'Secret represents a secret that should
                                populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret'
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
                                  description: 'Allowed value: The `name` field of
                                    a `SecretManagerSecret` resource.'
                                  type: string
                                name:
                                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                  type: string
                                namespace:
                                  description: 'Namespace of the referent. More info:
                                    https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                  type: string
                              type: object
                          required:
                          - secretRef
                          type: object
                      required:
                      - name
                      type: object
                    type: array
              vpcAccess:
                description: VPC Access configuration to use for this Task. For more
                  information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
                properties:
                  egress:
                    description: 'Traffic VPC egress settings. Possible values: ["ALL_TRAFFIC",
                      "PRIVATE_RANGES_ONLY"].'
                    type: string
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
      etag:
        description: A system-generated fingerprint for this version of the resource.
          May be used to detect modification conflict during updates.
        type: string
      executionCount:
        description: Number of executions created for this job.
        type: integer
      latestCreatedExecution:
        description: Name of the last created execution.
        items:
          properties:
            completionTime:
              description: |-
                Completion timestamp of the execution.

                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
              type: string
            createTime:
              description: |-
                Creation timestamp of the execution.

                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
              type: string
            name:
              description: Name of the execution.
              type: string
          type: object
        type: array
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      reconciling:
        description: |-
          Returns true if the Job is currently being acted upon by the system to bring it into the desired state.

          When a new Job is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Job to the desired state. This process is called reconciliation. While reconciliation is in process, observedGeneration and latest_succeeded_execution, will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the state matches the Job, or there was an error, and reconciliation failed. This state can be found in terminalCondition.state.

          If reconciliation succeeded, the following fields will match: observedGeneration and generation, latest_succeeded_execution and latestCreatedExecution.

          If reconciliation failed, observedGeneration and latest_succeeded_execution will have the state of the last succeeded execution or empty for newly created Job. Additional information on the failure can be found in terminalCondition and conditions.
        type: boolean
      terminalCondition:
        description: The Condition of this Job, containing its readiness status, and
          detailed error information in case it did not reach the desired state.
        items:
          properties:
            executionReason:
              description: A reason for the execution condition.
              type: string
            lastTransitionTime:
              description: |-
                Last time the condition transitioned from one status to another.

                A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
              type: string
            message:
              description: Human readable message indicating details about the current
                status.
              type: string
            reason:
              description: A common (service-level) reason for this condition.
              type: string
            revisionReason:
              description: A reason for the revision condition.
              type: string
            severity:
              description: How to interpret failures of this condition, one of Error,
                Warning, Info.
              type: string
            state:
              description: State of the condition.
              type: string
            type:
              description: 'type is used to communicate the status of the reconciliation
                process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting
                Types common to all resources include: * "Ready": True when the Resource
                is ready.'
              type: string
          type: object
        type: array
      uid:
        description: Server assigned unique identifier for the Execution. The value
          is a UUID4 string and guaranteed to remain unchanged until the resource
          is deleted.
        type: string
    type: object
required:
- spec
type: object

```
