# openAPI schema part2 for osconfig.cnrm.cloud.google.com.v1beta1.OSConfigGuestPolicy

## schema

```yaml
properties:
  spec:
    properties:
      recipes:
        properties:
          recipes:
            description: Optional. A list of Recipes to install on the VM.
            items:
              properties:
                artifacts:
                  description: Resources available to be used in the steps in the
                    recipe.
                  items:
                    properties:
                      allowInsecure:
                        description: 'Defaults to false. When false, recipes are subject
                          to validations based on the artifact type: Remote: A checksum
                          must be specified, and only protocols with transport-layer
                          security are permitted. GCS: An object generation number
                          must be specified.'
                        type: boolean
                      gcs:
                        description: A Google Cloud Storage artifact.
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
                                description: |-
                                  Bucket of the Google Cloud Storage object. Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567` this value would be `my-bucket`.

                                  Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          generation:
                            description: Must be provided if allow_insecure is false.
                              Generation number of the Google Cloud Storage object.
                              `https://storage.googleapis.com/my-bucket/foo/bar#1234567`
                              this value would be `1234567`.
                            format: int64
                            type: integer
                          object:
                            description: 'Name of the Google Cloud Storage object.
                              As specified [here] (https://cloud.google.com/storage/docs/naming#objectnames)
                              Given an example URL: `https://storage.googleapis.com/my-bucket/foo/bar#1234567`
                              this value would be `foo/bar`.'
                            type: string
                        type: object
                      id:
                        description: Required. Id of the artifact, which the installation
                          and update steps of this recipe can reference. Artifacts
                          in a recipe cannot have the same id.
                        type: string
                      remote:
                        description: A generic remote artifact.
                        properties:
                          checksum:
                            description: Must be provided if `allow_insecure` is `false`.
                              SHA256 checksum in hex format, to compare to the checksum
                              of the artifact. If the checksum is not empty and it
                              doesn't match the artifact then the recipe installation
                              fails before running any of the steps.
                            type: string
                          uri:
                            description: 'URI from which to fetch the object. It should
                              contain both the protocol and path following the format:
                              {protocol}://{location}.'
                            type: string
                        type: object
                    type: object
                  type: array
                desiredState:
                  description: 'Default is INSTALLED. The desired state the agent
                    should maintain for this recipe. INSTALLED: The software recipe
                    is installed on the instance but won''t be updated to new versions.
                    UPDATED: The software recipe is installed on the instance. The
                    recipe is updated to a higher version, if a higher version of
                    the recipe is assigned to this instance. REMOVE: Remove is unsupported
                    for software recipes and attempts to create or update a recipe
                    to the REMOVE state is rejected. Possible values: DESIRED_STATE_UNSPECIFIED,
                    INSTALLED, REMOVED'
                  type: string
                installSteps:
                  description: Actions to be taken for installing this recipe. On
                    failure it stops executing steps and does not attempt another
                    installation. Any steps taken (including partially completed steps)
                    are not rolled back.
                  items:
                    properties:
                      archiveExtraction:
                        description: Extracts an archive into the specified directory.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          destination:
                            description: Directory to extract archive to. Defaults
                              to `/` on Linux or `C:` on Windows.
                            type: string
                          type:
                            description: 'Required. The type of the archive to extract.
                              Possible values: TYPE_UNSPECIFIED, VALIDATION, DESIRED_STATE_CHECK,
                              DESIRED_STATE_ENFORCEMENT, DESIRED_STATE_CHECK_POST_ENFORCEMENT'
                            type: string
                        type: object
                      dpkgInstallation:
                        description: Installs a deb file via dpkg.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                        type: object
                      fileCopy:
                        description: Copies a file onto the instance.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          destination:
                            description: Required. The absolute path on the instance
                              to put the file.
                            type: string
                          overwrite:
                            description: Whether to allow this step to overwrite existing
                              files. If this is false and the file already exists
                              the file is not overwritten and the step is considered
                              a success. Defaults to false.
                            type: boolean
                          permissions:
                            description: 'Consists of three octal digits which represent,
                              in order, the permissions of the owner, group, and other
                              users for the file (similarly to the numeric mode used
                              in the linux chmod utility). Each digit represents a
                              three bit number with the 4 bit corresponding to the
                              read permissions, the 2 bit corresponds to the write
                              bit, and the one bit corresponds to the execute permission.
                              Default behavior is 755. Below are some examples of
                              permissions and their associated values: read, write,
                              and execute: 7 read and execute: 5 read and write: 6
                              read only: 4'
                            type: string
                        type: object
                      fileExec:
                        description: Executes an artifact or local file.
                        properties:
                          allowedExitCodes:
                            description: Defaults to [0]. A list of possible return
                              values that the program can return to indicate a success.
                            items:
                              format: int64
                              type: integer
                            type: array
                          args:
                            description: Arguments to be passed to the provided executable.
                            items:
                              type: string
                            type: array
                          artifactId:
                            description: The id of the relevant artifact in the recipe.
                            type: string
                          localPath:
                            description: The absolute path of the file on the local
                              filesystem.
                            type: string
                        type: object
                      msiInstallation:
                        description: Installs an MSI file.
                        properties:
                          allowedExitCodes:
                            description: Return codes that indicate that the software
                              installed or updated successfully. Behaviour defaults
                              to [0]
                            items:
                              format: int64
                              type: integer
                            type: array
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          flags:
                            description: The flags to use when installing the MSI
                              defaults to ["/i"] (i.e. the install flag).
                            items:
                              type: string
                            type: array
                        type: object
                      rpmInstallation:
                        description: Installs an rpm file via the rpm utility.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                        type: object
                      scriptRun:
                        description: Runs commands in a shell.
                        properties:
                          allowedExitCodes:
                            description: Return codes that indicate that the software
                              installed or updated successfully. Behaviour defaults
                              to [0]
                            items:
                              format: int64
                              type: integer
                            type: array
                          interpreter:
                            description: 'The script interpreter to use to run the
                              script. If no interpreter is specified the script is
                              executed directly, which likely only succeed for scripts
                              with [shebang lines](https://en.wikipedia.org/wiki/Shebang_(Unix)).
                              Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL,
                              POWERSHELL'
                            type: string
                          script:
                            description: Required. The shell script to be executed.
                            type: string
                        type: object
                    type: object
                  type: array
                name:
                  description: Required. Unique identifier for the recipe. Only one
                    recipe with a given name is installed on an instance. Names are
                    also used to identify resources which helps to determine whether
                    guest policies have conflicts. This means that requests to create
                    multiple recipes with the same name and version are rejected since
                    they could potentially have conflicting assignments.
                  type: string
                updateSteps:
                  description: Actions to be taken for updating this recipe. On failure
                    it stops executing steps and does not attempt another update for
                    this recipe. Any steps taken (including partially completed steps)
                    are not rolled back.
                  items:
                    properties:
                      archiveExtraction:
                        description: Extracts an archive into the specified directory.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          destination:
                            description: Directory to extract archive to. Defaults
                              to `/` on Linux or `C:` on Windows.
                            type: string
                          type:
                            description: 'Required. The type of the archive to extract.
                              Possible values: TYPE_UNSPECIFIED, VALIDATION, DESIRED_STATE_CHECK,
                              DESIRED_STATE_ENFORCEMENT, DESIRED_STATE_CHECK_POST_ENFORCEMENT'
                            type: string
                        type: object
                      dpkgInstallation:
                        description: Installs a deb file via dpkg.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                        type: object
                      fileCopy:
                        description: Copies a file onto the instance.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          destination:
                            description: Required. The absolute path on the instance
                              to put the file.
                            type: string
                          overwrite:
                            description: Whether to allow this step to overwrite existing
                              files. If this is false and the file already exists
                              the file is not overwritten and the step is considered
                              a success. Defaults to false.
                            type: boolean
                          permissions:
                            description: 'Consists of three octal digits which represent,
                              in order, the permissions of the owner, group, and other
                              users for the file (similarly to the numeric mode used
                              in the linux chmod utility). Each digit represents a
                              three bit number with the 4 bit corresponding to the
                              read permissions, the 2 bit corresponds to the write
                              bit, and the one bit corresponds to the execute permission.
                              Default behavior is 755. Below are some examples of
                              permissions and their associated values: read, write,
                              and execute: 7 read and execute: 5 read and write: 6
                              read only: 4'
                            type: string
                        type: object
                      fileExec:
                        description: Executes an artifact or local file.
                        properties:
                          allowedExitCodes:
                            description: Defaults to [0]. A list of possible return
                              values that the program can return to indicate a success.
                            items:
                              format: int64
                              type: integer
                            type: array
                          args:
                            description: Arguments to be passed to the provided executable.
                            items:
                              type: string
                            type: array
                          artifactId:
                            description: The id of the relevant artifact in the recipe.
                            type: string
                          localPath:
                            description: The absolute path of the file on the local
                              filesystem.
                            type: string
                        type: object
                      msiInstallation:
                        description: Installs an MSI file.
                        properties:
                          allowedExitCodes:
                            description: Return codes that indicate that the software
                              installed or updated successfully. Behaviour defaults
                              to [0]
                            items:
                              format: int64
                              type: integer
                            type: array
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                          flags:
                            description: The flags to use when installing the MSI
                              defaults to ["/i"] (i.e. the install flag).
                            items:
                              type: string
                            type: array
                        type: object
                      rpmInstallation:
                        description: Installs an rpm file via the rpm utility.
                        properties:
                          artifactId:
                            description: Required. The id of the relevant artifact
                              in the recipe.
                            type: string
                        type: object
                      scriptRun:
                        description: Runs commands in a shell.
                        properties:
                          allowedExitCodes:
                            description: Return codes that indicate that the software
                              installed or updated successfully. Behaviour defaults
                              to [0]
                            items:
                              format: int64
                              type: integer
                            type: array
                          interpreter:
                            description: 'The script interpreter to use to run the
                              script. If no interpreter is specified the script is
                              executed directly, which likely only succeed for scripts
                              with [shebang lines](https://en.wikipedia.org/wiki/Shebang_(Unix)).
                              Possible values: INTERPRETER_UNSPECIFIED, NONE, SHELL,
                              POWERSHELL'
                            type: string
                          script:
                            description: Required. The shell script to be executed.
                            type: string
                        type: object
                    type: object
                  type: array
                version:
                  description: The version of this software recipe. Version can be
                    up to 4 period separated numbers (e.g. 12.34.56.78).
                  type: string
              type: object
            type: array
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
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
      createTime:
        description: Output only. Time this GuestPolicy was created.
        format: date-time
        type: string
      etag:
        description: The etag for this GuestPolicy. If this is provided on update,
          it must match the server's etag.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: Output only. Last time this GuestPolicy was updated.
        format: date-time
        type: string
    type: object
type: object

```
