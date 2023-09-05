# openAPI schema part2 for cloudbuild.cnrm.cloud.google.com.v1beta1.CloudBuildTrigger

## schema

```yaml
properties:
  spec:
    properties:
      build:
        properties:
          artifacts:
            description: Artifacts produced by the build that should be uploaded upon
              successful completion of all build steps.
            properties:
              images:
                description: |-
                  A list of images to be pushed upon the successful completion of all build steps.

                  The images will be pushed using the builder service account's credentials.

                  The digests of the pushed images will be stored in the Build resource's results field.

                  If any of the images fail to be pushed, the build is marked FAILURE.
                items:
                  type: string
                type: array
              objects:
                description: |-
                  A list of objects to be uploaded to Cloud Storage upon successful completion of all build steps.

                  Files in the workspace matching specified paths globs will be uploaded to the
                  Cloud Storage location using the builder service account's credentials.

                  The location and generation of the uploaded objects will be stored in the Build resource's results field.

                  If any objects fail to be pushed, the build is marked FAILURE.
                properties:
                  location:
                    description: |-
                      Cloud Storage bucket and optional object path, in the form "gs://bucket/path/to/somewhere/".

                      Files in the workspace matching any path pattern will be uploaded to Cloud Storage with
                      this location as a prefix.
                    type: string
                  paths:
                    description: Path globs used to match files in the build's workspace.
                    items:
                      type: string
                    type: array
                  timing:
                    description: Output only. Stores timing information for pushing
                      all artifact objects.
                    items:
                      properties:
                        endTime:
                          description: |-
                            End of time span.

                            A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
                            nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
                          type: string
                        startTime:
                          description: |-
                            Start of time span.

                            A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to
                            nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
                          type: string
                      type: object
                    type: array
                type: object
            type: object
          availableSecrets:
            description: Secrets and secret environment variables.
            properties:
              secretManager:
                description: Pairs a secret environment variable with a SecretVersion
                  in Secret Manager.
                items:
                  properties:
                    env:
                      description: |-
                        Environment variable name to associate with the secret. Secret environment
                        variables must be unique across all of a build's secrets, and must be used
                        by at least one build step.
                      type: string
                    versionRef:
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
                          description: 'Allowed value: The `name` field of a `SecretManagerSecretVersion`
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
                  - env
                  - versionRef
                  type: object
                type: array
            required:
            - secretManager
            type: object
          build:
            description: Contents of the build template. Either a filename or build
              template must be provided.
            required:
            - step
            type: object
          images:
            description: |-
              A list of images to be pushed upon the successful completion of all build steps.
              The images are pushed using the builder service account's credentials.
              The digests of the pushed images will be stored in the Build resource's results field.
              If any of the images fail to be pushed, the build status is marked FAILURE.
            items:
              type: string
            type: array
          logsBucketRef:
            description: |-
              Google Cloud Storage bucket where logs should be written. Logs file
              names will be of the format ${logsBucket}/log-${build_id}.txt.
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
                description: 'Allowed value: The `url` field of a `StorageBucket`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          options:
            description: Special options for this build.
            properties:
              diskSizeGb:
                description: |-
                  Requested disk size for the VM that runs the build. Note that this is NOT "disk free";
                  some of the space will be used by the operating system and build utilities.
                  Also note that this is the minimum disk size that will be allocated for the build --
                  the build may run with a larger disk than requested. At present, the maximum disk size
                  is 1000GB; builds that request more than the maximum are rejected with an error.
                type: integer
              dynamicSubstitutions:
                description: |-
                  Option to specify whether or not to apply bash style string operations to the substitutions.

                  NOTE this is always enabled for triggered builds and cannot be overridden in the build configuration file.
                type: boolean
              env:
                description: |-
                  A list of global environment variable definitions that will exist for all build steps
                  in this build. If a variable is defined in both globally and in a build step,
                  the variable will use the build step value.

                  The elements are of the form "KEY=VALUE" for the environment variable "KEY" being given the value "VALUE".
                items:
                  type: string
                type: array
              logStreamingOption:
                description: 'Option to define build log streaming behavior to Google
                  Cloud Storage. Possible values: ["STREAM_DEFAULT", "STREAM_ON",
                  "STREAM_OFF"].'
                type: string
              logging:
                description: 'Option to specify the logging mode, which determines
                  if and where build logs are stored. Possible values: ["LOGGING_UNSPECIFIED",
                  "LEGACY", "GCS_ONLY", "STACKDRIVER_ONLY", "CLOUD_LOGGING_ONLY",
                  "NONE"].'
                type: string
              machineType:
                description: 'Compute Engine machine type on which to run the build.
                  Possible values: ["UNSPECIFIED", "N1_HIGHCPU_8", "N1_HIGHCPU_32",
                  "E2_HIGHCPU_8", "E2_HIGHCPU_32"].'
                type: string
              requestedVerifyOption:
                description: 'Requested verifiability options. Possible values: ["NOT_VERIFIED",
                  "VERIFIED"].'
                type: string
              secretEnv:
                description: |-
                  A list of global environment variables, which are encrypted using a Cloud Key Management
                  Service crypto key. These values must be specified in the build's Secret. These variables
                  will be available to all build steps in this build.
                items:
                  type: string
                type: array
              sourceProvenanceHash:
                description: 'Requested hash for SourceProvenance. Possible values:
                  ["NONE", "SHA256", "MD5"].'
                items:
                  type: string
                type: array
              substitutionOption:
                description: |-
                  Option to specify behavior when there is an error in the substitution checks.

                  NOTE this is always set to ALLOW_LOOSE for triggered builds and cannot be overridden
                  in the build configuration file. Possible values: ["MUST_MATCH", "ALLOW_LOOSE"].
                type: string
              volumes:
                description: |-
                  Global list of volumes to mount for ALL build steps

                  Each volume is created as an empty volume prior to starting the build process.
                  Upon completion of the build, volumes and their contents are discarded. Global
                  volume names and paths cannot conflict with the volumes defined a build step.

                  Using a global volume in a build with only one step is not valid as it is indicative
                  of a build request with an incorrect configuration.
                items:
                  properties:
                    name:
                      description: |-
                        Name of the volume to mount.

                        Volume names must be unique per build step and must be valid names for Docker volumes.
                        Each named volume must be used by at least two build steps.
                      type: string
                    path:
                      description: |-
                        Path at which to mount the volume.

                        Paths must be absolute and cannot conflict with other volume paths on the same
                        build step or with certain reserved volume paths.
                      type: string
                  type: object
                type: array
              workerPool:
                description: |-
                  Option to specify a WorkerPool for the build. Format projects/{project}/workerPools/{workerPool}

                  This field is experimental.
                type: string
            type: object
          queueTtl:
            description: |-
              TTL in queue for this build. If provided and the build is enqueued longer than this value,
              the build will expire and the build status will be EXPIRED.
              The TTL starts ticking from createTime.
              A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
            type: string
          secret:
            description: Secrets to decrypt using Cloud Key Management Service.
            items:
              properties:
                kmsKeyRef:
                  description: KMS crypto key to use to decrypt these envs.
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
                secretEnv:
                  additionalProperties:
                    type: string
                  description: |-
                    Map of environment variable name to its encrypted value.
                    Secret environment variables must be unique across all of a build's secrets,
                    and must be used by at least one build step. Values can be at most 64 KB in size.
                    There can be at most 100 secret values across all of a build's secrets.
                  type: object
              required:
              - kmsKeyRef
              type: object
            type: array
          source:
            description: |-
              The location of the source files to build.

              One of 'storageSource' or 'repoSource' must be provided.
            properties:
              repoSource:
                description: Location of the source in a Google Cloud Source Repository.
                properties:
                  branchName:
                    description: |-
                      Regex matching branches to build. Exactly one a of branch name, tag, or commit SHA must be provided.
                      The syntax of the regular expressions accepted is the syntax accepted by RE2 and
                      described at https://github.com/google/re2/wiki/Syntax.
                    type: string
                  commitSha:
                    description: Explicit commit SHA to build. Exactly one a of branch
                      name, tag, or commit SHA must be provided.
                    type: string
                  dir:
                    description: |-
                      Directory, relative to the source root, in which to run the build.
                      This must be a relative path. If a step's dir is specified and is an absolute path,
                      this value is ignored for that step's execution.
                    type: string
                  invertRegex:
                    description: Only trigger a build if the revision regex does NOT
                      match the revision regex.
                    type: boolean
                  projectId:
                    description: |-
                      ID of the project that owns the Cloud Source Repository.
                      If omitted, the project ID requesting the build is assumed.
                    type: string
                  repoRef:
                    description: |-
                      The desired Cloud Source Repository. If omitted, "default" is
                      assumed.
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
                        description: 'Allowed value: The `name` field of a `SourceRepoRepository`
                          resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  substitutions:
                    additionalProperties:
                      type: string
                    description: Substitutions to use in a triggered build. Should
                      only be used with triggers.run.
                    type: object
                  tagName:
                    description: |-
                      Regex matching tags to build. Exactly one a of branch name, tag, or commit SHA must be provided.
                      The syntax of the regular expressions accepted is the syntax accepted by RE2 and
                      described at https://github.com/google/re2/wiki/Syntax.
                    type: string
                required:
                - repoRef
                type: object
              storageSource:
                description: Location of the source in an archive file in Google Cloud
                  Storage.
                properties:
                  bucketRef:
                    description: Google Cloud Storage bucket containing the source.
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
                  generation:
                    description: |-
                      Google Cloud Storage generation for the object.
                      If the generation is omitted, the latest generation will be used.
                    type: string
                  object:
                    description: |-
                      Google Cloud Storage object containing the source.
                      This object must be a gzipped archive file (.tar.gz) containing source to build.
                    type: string
                required:
                - bucketRef
                - object
                type: object
            type: object
type: object

```
