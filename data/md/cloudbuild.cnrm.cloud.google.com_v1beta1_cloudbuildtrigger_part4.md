# openAPI schema part4 for cloudbuild.cnrm.cloud.google.com.v1beta1.CloudBuildTrigger

## schema

```yaml
properties:
  spec:
    properties:
      pubsubConfig:
        properties:
          pubsubConfig:
            description: |-
              PubsubConfig describes the configuration of a trigger that creates
              a build whenever a Pub/Sub message is published.

              One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
            properties:
              serviceAccountRef:
                description: Service account that will make the push request.
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
              state:
                description: |-
                  Potential issues with the underlying Pub/Sub subscription configuration.
                  Only populated on get requests.
                type: string
              subscription:
                description: Output only. Name of the subscription.
                type: string
              topicRef:
                description: |-
                  The name of the topic from which this subscription
                  is receiving messages.
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
                    description: 'Allowed value: string of the format `projects/{{project}}/topics/{{value}}`,
                      where {{value}} is the `name` field of a `PubSubTopic` resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            required:
            - topicRef
            type: object
      repositoryEventConfig:
        description: The configuration of a trigger that creates a build whenever
          an event from Repo API is received.
        properties:
          pullRequest:
            description: Contains filter properties for matching Pull Requests.
            properties:
              branch:
                description: |-
                  Regex of branches to match.

                  The syntax of the regular expressions accepted is the syntax accepted by
                  RE2 and described at https://github.com/google/re2/wiki/Syntax.
                type: string
              commentControl:
                description: 'Configure builds to run whether a repository owner or
                  collaborator need to comment ''/gcbrun''. Possible values: ["COMMENTS_DISABLED",
                  "COMMENTS_ENABLED", "COMMENTS_ENABLED_FOR_EXTERNAL_CONTRIBUTORS_ONLY"].'
                type: string
              invertRegex:
                description: If true, branches that do NOT match the git_ref will
                  trigger a build.
                type: boolean
            type: object
          push:
            description: Contains filter properties for matching git pushes.
            properties:
              branch:
                description: |-
                  Regex of branches to match.

                  The syntax of the regular expressions accepted is the syntax accepted by
                  RE2 and described at https://github.com/google/re2/wiki/Syntax.
                type: string
              invertRegex:
                description: If true, only trigger a build if the revision regex does
                  NOT match the git_ref regex.
                type: boolean
              tag:
                description: |-
                  Regex of tags to match.

                  The syntax of the regular expressions accepted is the syntax accepted by
                  RE2 and described at https://github.com/google/re2/wiki/Syntax.
                type: string
            type: object
          repository:
            description: The resource name of the Repo API resource.
            type: string
        type: object
      serviceAccountRef:
        description: |-
          The service account used for all user-controlled operations including
          triggers.patch, triggers.run, builds.create, and builds.cancel.

          If no service account is set, then the standard Cloud Build service account
          ([PROJECT_NUM]@system.gserviceaccount.com) will be used instead.

          When populating via the external field, the following format is supported:
          projects/{PROJECT_ID}/serviceAccounts/{SERVICE_ACCOUNT_EMAIL}
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
            description: 'Allowed value: string of the format `projects/{{project}}/serviceAccounts/{{value}}`,
              where {{value}} is the `email` field of an `IAMServiceAccount` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      sourceToBuild:
        description: |-
          The repo and ref of the repository from which to build.
          This field is used only for those triggers that do not respond to SCM events.
          Triggers that respond to such events build source at whatever commit caused the event.
          This field is currently only used by Webhook, Pub/Sub, Manual, and Cron triggers.

          One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
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
          ref:
            description: The branch or tag to use. Must start with "refs/" (required).
            type: string
          repoType:
            description: |-
              The type of the repo, since it may not be explicit from the repo field (e.g from a URL).
              Values can be UNKNOWN, CLOUD_SOURCE_REPOSITORIES, GITHUB, BITBUCKET_SERVER Possible values: ["UNKNOWN", "CLOUD_SOURCE_REPOSITORIES", "GITHUB", "BITBUCKET_SERVER"].
            type: string
          repositoryRef:
            description: |-
              Only `external` field is supported to configure the reference.

              The qualified resource name of the Repo API repository.
              Either uri or repository can be specified and is required.
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
          uri:
            description: The URI of the repo.
            type: string
        required:
        - ref
        - repoType
        type: object
      substitutions:
        additionalProperties:
          type: string
        description: Substitutions data for Build resource.
        type: object
      tags:
        description: Tags for annotation of a BuildTrigger.
        items:
          type: string
        type: array
      triggerTemplate:
        description: |-
          Template describing the types of source changes to trigger a build.

          Branch and tag names in trigger templates are interpreted as regular
          expressions. Any branch or tag change that matches that regular
          expression will trigger a build.

          One of 'trigger_template', 'github', 'pubsub_config', 'webhook_config' or 'source_to_build' must be provided.
        properties:
          branchName:
            description: |-
              Name of the branch to build. Exactly one a of branch name, tag, or commit SHA must be provided.
              This field is a regular expression.
            type: string
          commitSha:
            description: Explicit commit SHA to build. Exactly one of a branch name,
              tag, or commit SHA must be provided.
            type: string
          dir:
            description: |-
              Directory, relative to the source root, in which to run the build.

              This must be a relative path. If a step's dir is specified and
              is an absolute path, this value is ignored for that step's
              execution.
            type: string
          invertRegex:
            description: Only trigger a build if the revision regex does NOT match
              the revision regex.
            type: boolean
          repoRef:
            description: |-
              The Cloud Source Repository to build. If omitted, the repo with
              name "default" is assumed.
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
          tagName:
            description: |-
              Name of the tag to build. Exactly one of a branch name, tag, or commit SHA must be provided.
              This field is a regular expression.
            type: string
        type: object
      webhookConfig:
        description: |-
          WebhookConfig describes the configuration of a trigger that creates
          a build whenever a webhook is sent to a trigger's webhook URL.

          One of 'trigger_template', 'github', 'pubsub_config' 'webhook_config' or 'source_to_build' must be provided.
        properties:
          secretRef:
            description: The secret required
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
                description: 'Allowed value: The `name` field of a `SecretManagerSecret`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          state:
            description: |-
              Potential issues with the underlying Pub/Sub subscription configuration.
              Only populated on get requests.
            type: string
        required:
        - secretRef
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
      createTime:
        description: Time when the trigger was created.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      triggerId:
        description: The unique identifier for the trigger.
        type: string
    type: object
type: object

```
