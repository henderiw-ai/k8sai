# openAPI schema part1 for osconfig.cnrm.cloud.google.com.v1alpha1.OSConfigPatchDeployment

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: Immutable. Description of the patch deployment. Length of the
          description is limited to 1024 characters.
        type: string
      duration:
        description: |-
          Immutable. Duration of the patch. After the duration ends, the patch times out.
          A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
        type: string
      instanceFilter:
        description: Immutable. VM instances to patch.
        properties:
          all:
            description: Immutable. Target all VM instances in the project. If true,
              no other criteria is permitted.
            type: boolean
          groupLabels:
            description: Immutable. Targets VM instances matching ANY of these GroupLabels.
              This allows targeting of disparate groups of VM instances.
            items:
              properties:
                labels:
                  additionalProperties:
                    type: string
                  description: Immutable. Compute Engine instance labels that must
                    be present for a VM instance to be targeted by this filter.
                  type: object
              required:
              - labels
              type: object
            type: array
          instanceNamePrefixes:
            description: |-
              Immutable. Targets VMs whose name starts with one of these prefixes. Similar to labels, this is another way to group
              VMs when targeting configs, for example prefix="prod-".
            items:
              type: string
            type: array
          instances:
            description: |-
              Immutable. Targets any of the VM instances specified. Instances are specified by their URI in the 'form zones/{{zone}}/instances/{{instance_name}}',
              'projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}', or
              'https://www.googleapis.com/compute/v1/projects/{{project_id}}/zones/{{zone}}/instances/{{instance_name}}'.
            items:
              type: string
            type: array
          zones:
            description: Immutable. Targets VM instances in ANY of these zones. Leave
              empty to target VM instances in any zone.
            items:
              type: string
            type: array
        type: object
      oneTimeSchedule:
        description: Immutable. Schedule a one-time execution.
        properties:
          executeTime:
            description: |-
              Immutable. The desired patch job execution time. A timestamp in RFC3339 UTC "Zulu" format,
              accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
            type: string
        required:
        - executeTime
        type: object
      patchConfig:
        description: Immutable. Patch configuration that is applied.
        properties:
          apt:
            description: Immutable. Apt update settings. Use this setting to override
              the default apt patch rules.
            properties:
              excludes:
                description: Immutable. List of packages to exclude from update. These
                  packages will be excluded.
                items:
                  type: string
                type: array
              exclusivePackages:
                description: |-
                  Immutable. An exclusive list of packages to be updated. These are the only packages that will be updated.
                  If these packages are not installed, they will be ignored. This field cannot be specified with
                  any other patch configuration fields.
                items:
                  type: string
                type: array
              type:
                description: 'Immutable. By changing the type to DIST, the patching
                  is performed using apt-get dist-upgrade instead. Possible values:
                  ["DIST", "UPGRADE"].'
                type: string
            type: object
          goo:
            description: Immutable. goo update settings. Use this setting to override
              the default goo patch rules.
            properties:
              enabled:
                description: Immutable. goo update settings. Use this setting to override
                  the default goo patch rules.
                type: boolean
            required:
            - enabled
            type: object
          migInstancesAllowed:
            description: Immutable. Allows the patch job to run on Managed instance
              groups (MIGs).
            type: boolean
          postStep:
            description: Immutable. The ExecStep to run after the patch update.
            properties:
              linuxExecStepConfig:
                description: Immutable. The ExecStepConfig for all Linux VMs targeted
                  by the PatchJob.
                properties:
                  allowedSuccessCodes:
                    description: Immutable. Defaults to [0]. A list of possible return
                      values that the execution can return to indicate a success.
                    items:
                      type: integer
                    type: array
                  gcsObject:
                    description: Immutable. A Cloud Storage object containing the
                      executable.
                    properties:
                      bucket:
                        description: Immutable. Bucket of the Cloud Storage object.
                        type: string
                      generationNumber:
                        description: Immutable. Generation number of the Cloud Storage
                          object. This is used to ensure that the ExecStep specified
                          by this PatchJob does not change.
                        type: string
                      object:
                        description: Immutable. Name of the Cloud Storage object.
                        type: string
                    required:
                    - bucket
                    - generationNumber
                    - object
                    type: object
                  interpreter:
                    description: |-
                      Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
                      be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"].
                    type: string
                  localPath:
                    description: Immutable. An absolute path to the executable on
                      the VM.
                    type: string
                type: object
              windowsExecStepConfig:
                description: Immutable. The ExecStepConfig for all Windows VMs targeted
                  by the PatchJob.
                properties:
                  allowedSuccessCodes:
                    description: Immutable. Defaults to [0]. A list of possible return
                      values that the execution can return to indicate a success.
                    items:
                      type: integer
                    type: array
                  gcsObject:
                    description: Immutable. A Cloud Storage object containing the
                      executable.
                    properties:
                      bucket:
                        description: Immutable. Bucket of the Cloud Storage object.
                        type: string
                      generationNumber:
                        description: Immutable. Generation number of the Cloud Storage
                          object. This is used to ensure that the ExecStep specified
                          by this PatchJob does not change.
                        type: string
                      object:
                        description: Immutable. Name of the Cloud Storage object.
                        type: string
                    required:
                    - bucket
                    - generationNumber
                    - object
                    type: object
                  interpreter:
                    description: |-
                      Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
                      be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"].
                    type: string
                  localPath:
                    description: Immutable. An absolute path to the executable on
                      the VM.
                    type: string
                type: object
            type: object
          preStep:
            description: Immutable. The ExecStep to run before the patch update.
            properties:
              linuxExecStepConfig:
                description: Immutable. The ExecStepConfig for all Linux VMs targeted
                  by the PatchJob.
                properties:
                  allowedSuccessCodes:
                    description: Immutable. Defaults to [0]. A list of possible return
                      values that the execution can return to indicate a success.
                    items:
                      type: integer
                    type: array
                  gcsObject:
                    description: Immutable. A Cloud Storage object containing the
                      executable.
                    properties:
                      bucket:
                        description: Immutable. Bucket of the Cloud Storage object.
                        type: string
                      generationNumber:
                        description: Immutable. Generation number of the Cloud Storage
                          object. This is used to ensure that the ExecStep specified
                          by this PatchJob does not change.
                        type: string
                      object:
                        description: Immutable. Name of the Cloud Storage object.
                        type: string
                    required:
                    - bucket
                    - generationNumber
                    - object
                    type: object
                  interpreter:
                    description: |-
                      Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
                      be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"].
                    type: string
                  localPath:
                    description: Immutable. An absolute path to the executable on
                      the VM.
                    type: string
                type: object
              windowsExecStepConfig:
                description: Immutable. The ExecStepConfig for all Windows VMs targeted
                  by the PatchJob.
                properties:
                  allowedSuccessCodes:
                    description: Immutable. Defaults to [0]. A list of possible return
                      values that the execution can return to indicate a success.
                    items:
                      type: integer
                    type: array
                  gcsObject:
                    description: Immutable. A Cloud Storage object containing the
                      executable.
                    properties:
                      bucket:
                        description: Immutable. Bucket of the Cloud Storage object.
                        type: string
                      generationNumber:
                        description: Immutable. Generation number of the Cloud Storage
                          object. This is used to ensure that the ExecStep specified
                          by this PatchJob does not change.
                        type: string
                      object:
                        description: Immutable. Name of the Cloud Storage object.
                        type: string
                    required:
                    - bucket
                    - generationNumber
                    - object
                    type: object
                  interpreter:
                    description: |-
                      Immutable. The script interpreter to use to run the script. If no interpreter is specified the script will
                      be executed directly, which will likely only succeed for scripts with shebang lines. Possible values: ["SHELL", "POWERSHELL"].
                    type: string
                  localPath:
                    description: Immutable. An absolute path to the executable on
                      the VM.
                    type: string
                type: object
            type: object
          rebootConfig:
            description: 'Immutable. Post-patch reboot settings. Possible values:
              ["DEFAULT", "ALWAYS", "NEVER"].'
            type: string
          windowsUpdate:
            description: Immutable. Windows update settings. Use this setting to override
              the default Windows patch rules.
            properties:
              classifications:
                description: 'Immutable. Only apply updates of these windows update
                  classifications. If empty, all updates are applied. Possible values:
                  ["CRITICAL", "SECURITY", "DEFINITION", "DRIVER", "FEATURE_PACK",
                  "SERVICE_PACK", "TOOL", "UPDATE_ROLLUP", "UPDATE"].'
                items:
                  type: string
                type: array
              excludes:
                description: Immutable. List of KBs to exclude from update.
                items:
                  type: string
                type: array
              exclusivePatches:
                description: |-
                  Immutable. An exclusive list of kbs to be updated. These are the only patches that will be updated.
                  This field must not be used with other patch configurations.
                items:
                  type: string
                type: array
            type: object
          yum:
            description: Immutable. Yum update settings. Use this setting to override
              the default yum patch rules.
            properties:
              excludes:
                description: Immutable. List of packages to exclude from update. These
                  packages will be excluded.
                items:
                  type: string
                type: array
              exclusivePackages:
                description: |-
                  Immutable. An exclusive list of packages to be updated. These are the only packages that will be updated.
                  If these packages are not installed, they will be ignored. This field cannot be specified with
                  any other patch configuration fields.
                items:
                  type: string
                type: array
              minimal:
                description: Immutable. Will cause patch to run yum update-minimal
                  instead.
                type: boolean
              security:
                description: Immutable. Adds the --security flag to yum update. Not
                  supported on all platforms.
                type: boolean
            type: object
          zypper:
            description: Immutable. zypper update settings. Use this setting to override
              the default zypper patch rules.
            properties:
              categories:
                description: Immutable. Install only patches with these categories.
                  Common categories include security, recommended, and feature.
                items:
                  type: string
                type: array
              excludes:
                description: Immutable. List of packages to exclude from update.
                items:
                  type: string
                type: array
              exclusivePatches:
                description: |-
                  Immutable. An exclusive list of patches to be updated. These are the only patches that will be installed using 'zypper patch patch:' command.
                  This field must not be used with any other patch configuration fields.
                items:
                  type: string
                type: array
              severities:
                description: Immutable. Install only patches with these severities.
                  Common severities include critical, important, moderate, and low.
                items:
                  type: string
                type: array
              withOptional:
                description: Immutable. Adds the --with-optional flag to zypper patch.
                type: boolean
              withUpdate:
                description: Immutable. Adds the --with-update flag, to zypper patch.
                type: boolean
            type: object
        type: object
      patchDeploymentId:
        description: |-
          Immutable. A name for the patch deployment in the project. When creating a name the following rules apply:
          * Must contain only lowercase letters, numbers, and hyphens.
          * Must start with a letter.
          * Must be between 1-63 characters.
          * Must end with a number or a letter.
          * Must be unique within the project.
        type: string
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
      spec:
        required:
        - instanceFilter
        - patchDeploymentId
        - projectRef
        type: object
required:
- spec
type: object

```
