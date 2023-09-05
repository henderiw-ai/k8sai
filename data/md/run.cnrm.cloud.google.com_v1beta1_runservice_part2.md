# openAPI schema part2 for run.cnrm.cloud.google.com.v1beta1.RunService

## schema

```yaml
properties:
  spec:
    properties:
      template:
        properties:
          annotations:
            additionalProperties:
              type: string
            description: KRM-style annotations for the resource.
            type: object
          containerConcurrency:
            description: Sets the maximum number of requests that each serving instance
              can receive.
            format: int64
            type: integer
          containers:
            description: Holds the single container that defines the unit of execution
              for this Revision.
            items:
              properties:
                args:
                  description: 'Arguments to the entrypoint. The docker image''s CMD
                    is used if this is not provided. Variable references $(VAR_NAME)
                    are expanded using the container''s environment. If a variable
                    cannot be resolved, the reference in the input string will be
                    unchanged. The $(VAR_NAME) syntax can be escaped with a double
                    $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
                    regardless of whether the variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell'
                  items:
                    type: string
                  type: array
                command:
                  description: 'Entrypoint array. Not executed within a shell. The
                    docker image''s ENTRYPOINT is used if this is not provided. Variable
                    references $(VAR_NAME) are expanded using the container''s environment.
                    If a variable cannot be resolved, the reference in the input string
                    will be unchanged. The $(VAR_NAME) syntax can be escaped with
                    a double $$, ie: $$(VAR_NAME). Escaped references will never be
                    expanded, regardless of whether the variable exists or not. More
                    info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell'
                  items:
                    type: string
                  type: array
                env:
                  description: List of environment variables to set in the container.
                  items:
                    properties:
                      name:
                        description: Required. Name of the environment variable. Must
                          be a C_IDENTIFIER, and mnay not exceed 32768 characters.
                        type: string
                      value:
                        description: 'Variable references $(VAR_NAME) are expanded
                          using the previous defined environment variables in the
                          container and any route environment variables. If a variable
                          cannot be resolved, the reference in the input string will
                          be unchanged. The $(VAR_NAME) syntax can be escaped with
                          a double $$, ie: $$(VAR_NAME). Escaped references will never
                          be expanded, regardless of whether the variable exists or
                          not. Defaults to "", and the maximum length is 32768 bytes.'
                        type: string
                      valueSource:
                        description: Source for the environment variable's value.
                        properties:
                          secretKeyRef:
                            description: Selects a secret and a specific version from
                              Cloud Secret Manager.
                            properties:
                              secretRef:
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
                                      Required. The name of the secret in Cloud Secret Manager. Format: {secret_name} if the secret is in the same project. projects/{project}/secrets/{secret_name} if the secret is in a different project.

                                      Allowed value: The Google Cloud resource name of a `SecretManagerSecret` resource (format: `projects/{{project}}/secrets/{{name}}`).
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
                                    description: |-
                                      The Cloud Secret Manager secret version. Can be 'latest' for the latest value or an integer for a specific version.

                                      Allowed value: The Google Cloud resource name of a `SecretManagerSecretVersion` resource (format: `{{name}}`).
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
                            - secretRef
                            type: object
                        type: object
                    required:
                    - name
                    type: object
                  type: array
                image:
                  description: 'Required. URL of the Container image in Google Container
                    Registry or Docker More info: https://kubernetes.io/docs/concepts/containers/images'
                  type: string
                name:
                  description: Name of the container specified as a DNS_LABEL.
                  type: string
                ports:
                  description: List of ports to expose from the container. Only a
                    single port can be specified. The specified ports must be listening
                    on all interfaces (0.0.0.0) within the container to be accessible.
                    If omitted, a port number will be chosen and passed to the container
                    through the PORT environment variable for the container to listen
                    on.
                  items:
                    properties:
                      containerPort:
                        description: Port number the container listens on. This must
                          be a valid TCP port number, 0 < container_port < 65536.
                        format: int64
                        type: integer
                      name:
                        description: If specified, used to specify which protocol
                          to use. Allowed values are "http1" and "h2c".
                        type: string
                    type: object
                  type: array
                resources:
                  description: 'Compute Resource requirements by this container. More
                    info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources'
                  properties:
                    cpuIdle:
                      description: Determines whether CPU should be throttled or not
                        outside of requests.
                      type: boolean
                    limits:
                      additionalProperties:
                        type: string
                      description: 'Only memory and CPU are supported. Note: The only
                        supported values for CPU are ''1'', ''2'', and ''4''. Setting
                        4 CPU requires at least 2Gi of memory. The values of the map
                        is string form of the ''quantity'' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go'
                      type: object
                  type: object
                volumeMounts:
                  description: Volume to mount into the container's filesystem.
                  items:
                    properties:
                      mountPath:
                        description: Required. Path within the container at which
                          the volume should be mounted. Must not contain ':'. For
                          Cloud SQL volumes, it can be left empty, or must otherwise
                          be `/cloudsql`. All instances defined in the Volume will
                          be available as `/cloudsql/[instance]`. For more information
                          on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run
                        type: string
                      name:
                        description: Required. This must match the Name of a Volume.
                        type: string
                    required:
                    - mountPath
                    - name
                    type: object
                  type: array
              required:
              - image
              type: object
            type: array
          executionEnvironment:
            description: 'The sandbox environment to host this Revision. Possible
              values: EXECUTION_ENVIRONMENT_UNSPECIFIED, EXECUTION_ENVIRONMENT_GEN1,
              EXECUTION_ENVIRONMENT_GEN2'
            type: string
          labels:
            additionalProperties:
              type: string
            description: KRM-style labels for the resource.
            type: object
          revision:
            description: The unique name for the revision. If this field is omitted,
              it will be automatically generated based on the Service name.
            type: string
          scaling:
            description: Scaling settings for this Revision.
            properties:
              maxInstanceCount:
                description: Maximum number of serving instances that this resource
                  should have.
                format: int64
                type: integer
              minInstanceCount:
                description: Minimum number of serving instances that this resource
                  should have.
                format: int64
                type: integer
            type: object
          serviceAccountRef:
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
                  Email address of the IAM service account associated with the revision of the service. The service account represents the identity of the running revision, and determines what permissions the revision has. If not provided, the revision will use the project's default service account.

                  Allowed value: The `email` field of an `IAMServiceAccount` resource.
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          template:
            description: Required. The template used to create revisions for this
              Service.
            type: object
          timeout:
            description: Max allowed time for an instance to respond to a request.
            type: string
required:
- spec
type: object

```
