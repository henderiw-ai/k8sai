# openAPI schema part3 for cloudbuild.cnrm.cloud.google.com.v1beta1.CloudBuildTrigger

## schema

```yaml
properties:
  spec:
    properties:
      build:
        properties:
          step:
            properties:
              step:
                description: The operations to be performed on the workspace.
                items:
                  properties:
                    allowExitCodes:
                      description: |-
                        Allow this build step to fail without failing the entire build if and
                        only if the exit code is one of the specified codes.

                        If 'allowFailure' is also specified, this field will take precedence.
                      items:
                        type: integer
                      type: array
                    allowFailure:
                      description: |-
                        Allow this build step to fail without failing the entire build.
                        If false, the entire build will fail if this step fails. Otherwise, the
                        build will succeed, but this step will still have a failure status.
                        Error information will be reported in the 'failureDetail' field.

                        'allowExitCodes' takes precedence over this field.
                      type: boolean
                    args:
                      description: |-
                        A list of arguments that will be presented to the step when it is started.

                        If the image used to run the step's container has an entrypoint, the args
                        are used as arguments to that entrypoint. If the image does not define an
                        entrypoint, the first element in args is used as the entrypoint, and the
                        remainder will be used as arguments.
                      items:
                        type: string
                      type: array
                    dir:
                      description: |-
                        Working directory to use when running this step's container.

                        If this value is a relative path, it is relative to the build's working
                        directory. If this value is absolute, it may be outside the build's working
                        directory, in which case the contents of the path may not be persisted
                        across build step executions, unless a 'volume' for that path is specified.

                        If the build specifies a 'RepoSource' with 'dir' and a step with a
                        'dir',
                        which specifies an absolute path, the 'RepoSource' 'dir' is ignored
                        for the step's execution.
                      type: string
                    entrypoint:
                      description: |-
                        Entrypoint to be used instead of the build step image's
                        default entrypoint.
                        If unset, the image's default entrypoint is used.
                      type: string
                    env:
                      description: |-
                        A list of environment variable definitions to be used when
                        running a step.

                        The elements are of the form "KEY=VALUE" for the environment variable
                        "KEY" being given the value "VALUE".
                      items:
                        type: string
                      type: array
                    id:
                      description: |-
                        Unique identifier for this build step, used in 'wait_for' to
                        reference this build step as a dependency.
                      type: string
                    name:
                      description: |-
                        The name of the container image that will run this particular build step.

                        If the image is available in the host's Docker daemon's cache, it will be
                        run directly. If not, the host will attempt to pull the image first, using
                        the builder service account's credentials if necessary.

                        The Docker daemon's cache will already have the latest versions of all of
                        the officially supported build steps (see https://github.com/GoogleCloudPlatform/cloud-builders
                        for images and examples).
                        The Docker daemon will also have cached many of the layers for some popular
                        images, like "ubuntu", "debian", but they will be refreshed at the time
                        you attempt to use them.

                        If you built an image in a previous build step, it will be stored in the
                        host's Docker daemon's cache and is available to use as the name for a
                        later build step.
                      type: string
                    script:
                      description: |-
                        A shell script to be executed in the step.
                        When script is provided, the user cannot specify the entrypoint or args.
                      type: string
                    secretEnv:
                      description: |-
                        A list of environment variables which are encrypted using
                        a Cloud Key
                        Management Service crypto key. These values must be specified in
                        the build's 'Secret'.
                      items:
                        type: string
                      type: array
                    timeout:
                      description: |-
                        Time limit for executing this build step. If not defined,
                        the step has no
                        time limit and will be allowed to continue to run until either it
                        completes or the build itself times out.
                      type: string
                    timing:
                      description: |-
                        Output only. Stores timing information for executing this
                        build step.
                      type: string
                    volumes:
                      description: |-
                        List of volumes to mount into the build step.

                        Each volume is created as an empty volume prior to execution of the
                        build step. Upon completion of the build, volumes and their contents
                        are discarded.

                        Using a named volume in only one step is not valid as it is
                        indicative of a build request with an incorrect configuration.
                      items:
                        properties:
                          name:
                            description: |-
                              Name of the volume to mount.

                              Volume names must be unique per build step and must be valid names for
                              Docker volumes. Each named volume must be used by at least two build steps.
                            type: string
                          path:
                            description: |-
                              Path at which to mount the volume.

                              Paths must be absolute and cannot conflict with other volume paths on
                              the same build step or with certain reserved volume paths.
                            type: string
                        required:
                        - name
                        - path
                        type: object
                      type: array
                    waitFor:
                      description: |-
                        The ID(s) of the step(s) that this build step depends on.

                        This build step will not start until all the build steps in 'wait_for'
                        have completed successfully. If 'wait_for' is empty, this build step
                        will start when all previous build steps in the 'Build.Steps' list
                        have completed successfully.
                      items:
                        type: string
                      type: array
                  required:
                  - name
                  type: object
                type: array
          substitutions:
            additionalProperties:
              type: string
            description: Substitutions data for Build resource.
            type: object
          tags:
            description: Tags for annotation of a Build. These are not docker tags.
            items:
              type: string
            type: array
          timeout:
            description: |-
              Amount of time that this build should be allowed to run, to second granularity.
              If this amount of time elapses, work on the build will cease and the build status will be TIMEOUT.
              This timeout must be equal to or greater than the sum of the timeouts for build steps within the build.
              The expected format is the number of seconds followed by s.
              Default time is ten minutes (600s).
            type: string
      description:
        description: Human-readable description of the trigger.
        type: string
      disabled:
        description: Whether the trigger is disabled or not. If true, the trigger
          will never result in a build.
        type: boolean
      filename:
        description: |-
          Path, from the source root, to a file whose contents is used for the template.
          Either a filename or build template must be provided. Set this only when using trigger_template or github.
          When using Pub/Sub, Webhook or Manual set the file name using git_file_source instead.
        type: string
      filter:
        description: A Common Expression Language string. Used only with Pub/Sub and
          Webhook.
        type: string
      gitFileSource:
        description: The file source describing the local or remote Build template.
        properties:
          githubEnterpriseConfigRef:
            description: |-
              Only `external` field is supported to configure the reference.

              The full resource name of the github enterprise config. Format:
              projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}.
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
                description: 'Allowed value: The `name` field of a `CloudBuildGithubEnterpriseConfig`
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
            description: The path of the file, with the repo root as the root of the
              path.
            type: string
          repoType:
            description: |-
              The type of the repo, since it may not be explicit from the repo field (e.g from a URL).
              Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
            type: string
          repositoryRef:
            description: |-
              Only `external` field is supported to configure the reference.

              The fully qualified resource name of the Repo API repository. The fully qualified resource name of the Repo API repository.
              If unspecified, the repo from which the trigger invocation originated is assumed to be the repo from which to read the specified path.
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
                description: 'Allowed value: The `name` field of a `CloudBuildV2Repository`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          revision:
            description: |-
              The branch, tag, arbitrary ref, or SHA version of the repo to use when resolving the
              filename (optional). This field respects the same syntax/resolution as described here: https://git-scm.com/docs/gitrevisions
              If unspecified, the revision from which the trigger invocation originated is assumed to be the revision from which to read the specified path.
            type: string
          uri:
            description: |-
              The URI of the repo (optional). If unspecified, the repo from which the trigger
              invocation originated is assumed to be the repo from which to read the specified path.
            type: string
        required:
        - path
        - repoType
        type: object
      github:
        description: |-
          Describes the configuration of a trigger that creates a build whenever a GitHub event is received.

          One of 'trigger_template', 'github', 'pubsub_config' or 'webhook_config' must be provided.
        properties:
          enterpriseConfigResourceNameRef:
            description: |-
              Only `external` field is supported to configure the reference.

              The full resource name of the github enterprise config. Format:
              projects/{project}/locations/{location}/githubEnterpriseConfigs/{id}.
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
                description: 'Allowed value: The `name` field of a `CloudBuildGithubEnterpriseConfig`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          name:
            description: |-
              Name of the repository. For example: The name for
              https://github.com/googlecloudplatform/cloud-builders is "cloud-builders".
            type: string
          owner:
            description: |-
              Owner of the repository. For example: The owner for
              https://github.com/googlecloudplatform/cloud-builders is "googlecloudplatform".
            type: string
          pullRequest:
            description: filter to match changes in pull requests. Specify only one
              of 'pull_request' or 'push'.
            properties:
              branch:
                description: Regex of branches to match.
                type: string
              commentControl:
                description: 'Whether to block builds on a "/gcbrun" comment from
                  a repository owner or collaborator. Possible values: ["COMMENTS_DISABLED",
                  "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].'
                type: string
              invertRegex:
                description: If true, branches that do NOT match the git_ref will
                  trigger a build.
                type: boolean
            required:
            - branch
            type: object
          push:
            description: filter to match changes in refs, like branches or tags. Specify
              only one of 'pull_request' or 'push'.
            properties:
              branch:
                description: Regex of branches to match.  Specify only one of branch
                  or tag.
                type: string
              invertRegex:
                description: When true, only trigger a build if the revision regex
                  does NOT match the git_ref regex.
                type: boolean
              tag:
                description: Regex of tags to match.  Specify only one of branch or
                  tag.
                type: string
            type: object
        type: object
      ignoredFiles:
        description: |-
          ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
          extended with support for '**'.

          If ignoredFiles and changed files are both empty, then they are not
          used to determine whether or not to trigger a build.

          If ignoredFiles is not empty, then we ignore any files that match any
          of the ignored_file globs. If the change has no files that are outside
          of the ignoredFiles globs, then we do not trigger a build.
        items:
          type: string
        type: array
      includeBuildLogs:
        description: |-
          Build logs will be sent back to GitHub as part of the checkrun
          result.  Values can be INCLUDE_BUILD_LOGS_UNSPECIFIED or
          INCLUDE_BUILD_LOGS_WITH_STATUS Possible values: ["INCLUDE_BUILD_LOGS_UNSPECIFIED", "INCLUDE_BUILD_LOGS_WITH_STATUS"].
        type: string
      includedFiles:
        description: |-
          ignoredFiles and includedFiles are file glob matches using https://golang.org/pkg/path/filepath/#Match
          extended with support for '**'.

          If any of the files altered in the commit pass the ignoredFiles filter
          and includedFiles is empty, then as far as this filter is concerned, we
          should trigger the build.

          If any of the files altered in the commit pass the ignoredFiles filter
          and includedFiles is not empty, then we make sure that at least one of
          those files matches a includedFiles glob. If not, then we do not trigger
          a build.
        items:
          type: string
        type: array
      location:
        description: |-
          Immutable. The [Cloud Build location](https://cloud.google.com/build/docs/locations) for the trigger.
          If not specified, "global" is used.
        type: string
type: object

```
