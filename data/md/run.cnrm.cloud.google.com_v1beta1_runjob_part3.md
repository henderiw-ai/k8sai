# openAPI schema part3 for run.cnrm.cloud.google.com.v1beta1.RunJob

## schema

```yaml
properties:
  spec:
    properties:
      template:
        properties:
          template:
            properties:
              containers:
                description: Holds the single container that defines the unit of execution
                  for this task.
                items:
                  properties:
                    args:
                      description: 'Arguments to the entrypoint. The docker image''s
                        CMD is used if this is not provided. Variable references $(VAR_NAME)
                        are expanded using the container''s environment. If a variable
                        cannot be resolved, the reference in the input string will
                        be unchanged. The $(VAR_NAME) syntax can be escaped with a
                        double $$, ie: $$(VAR_NAME). Escaped references will never
                        be expanded, regardless of whether the variable exists or
                        not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell.'
                      items:
                        type: string
                      type: array
                    command:
                      description: 'Entrypoint array. Not executed within a shell.
                        The docker image''s ENTRYPOINT is used if this is not provided.
                        Variable references $(VAR_NAME) are expanded using the container''s
                        environment. If a variable cannot be resolved, the reference
                        in the input string will be unchanged. The $(VAR_NAME) syntax
                        can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
                        references will never be expanded, regardless of whether the
                        variable exists or not. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell.'
                      items:
                        type: string
                      type: array
                    env:
                      description: List of environment variables to set in the container.
                      items:
                        properties:
                          name:
                            description: Name of the environment variable. Must be
                              a C_IDENTIFIER, and mnay not exceed 32768 characters.
                            type: string
                          value:
                            description: 'Variable references $(VAR_NAME) are expanded
                              using the previous defined environment variables in
                              the container and any route environment variables. If
                              a variable cannot be resolved, the reference in the
                              input string will be unchanged. The $(VAR_NAME) syntax
                              can be escaped with a double $$, ie: $$(VAR_NAME). Escaped
                              references will never be expanded, regardless of whether
                              the variable exists or not. Defaults to "", and the
                              maximum length is 32768 bytes.'
                            type: string
                          valueSource:
                            description: Source for the environment variable's value.
                            properties:
                              secretKeyRef:
                                description: Selects a secret and a specific version
                                  from Cloud Secret Manager.
                                properties:
                                  secretRef:
                                    description: 'The name of the secret in Cloud
                                      Secret Manager. Format: {secretName} if the
                                      secret is in the same project. projects/{project}/secrets/{secretName}
                                      if the secret is in a different project.'
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
                                        description: 'Allowed value: The `name` field
                                          of a `SecretManagerSecret` resource.'
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
                                    description: The Cloud Secret Manager secret version.
                                      Can be 'latest' for the latest value or an integer
                                      for a specific version.
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
                                - secretRef
                                - versionRef
                                type: object
                            type: object
                        required:
                        - name
                        type: object
                      type: array
                    image:
                      description: 'URL of the Container image in Google Container
                        Registry or Google Artifact Registry. More info: https://kubernetes.io/docs/concepts/containers/images.'
                      type: string
                    livenessProbe:
                      description: |-
                        DEPRECATED. Cloud Run Job does not support liveness probe and `liveness_probe` field will be removed in a future major release. Periodic probe of container liveness. Container will be restarted if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
                        This field is not supported in Cloud Run Job currently.
                      properties:
                        failureThreshold:
                          description: Minimum consecutive failures for the probe
                            to be considered failed after having succeeded. Defaults
                            to 3. Minimum value is 1.
                          type: integer
                        httpGet:
                          description: HTTPGet specifies the http request to perform.
                            Exactly one of HTTPGet or TCPSocket must be specified.
                          properties:
                            httpHeaders:
                              description: Custom headers to set in the request. HTTP
                                allows repeated headers.
                              items:
                                properties:
                                  name:
                                    description: The header field name.
                                    type: string
                                  value:
                                    description: The header field value.
                                    type: string
                                required:
                                - name
                                type: object
                              type: array
                            path:
                              description: Path to access on the HTTP server. Defaults
                                to '/'.
                              type: string
                          type: object
                        initialDelaySeconds:
                          description: 'Number of seconds after the container has
                            started before the probe is initiated. Defaults to 0 seconds.
                            Minimum value is 0. Maximum value for liveness probe is
                            3600. Maximum value for startup probe is 240. More info:
                            https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes.'
                          type: integer
                        periodSeconds:
                          description: How often (in seconds) to perform the probe.
                            Default to 10 seconds. Minimum value is 1. Maximum value
                            for liveness probe is 3600. Maximum value for startup
                            probe is 240. Must be greater or equal than timeoutSeconds.
                          type: integer
                        tcpSocket:
                          description: TCPSocket specifies an action involving a TCP
                            port. Exactly one of HTTPGet or TCPSocket must be specified.
                          properties:
                            port:
                              description: Port number to access on the container.
                                Must be in the range 1 to 65535. If not specified,
                                defaults to 8080.
                              type: integer
                          type: object
                        timeoutSeconds:
                          description: 'Number of seconds after which the probe times
                            out. Defaults to 1 second. Minimum value is 1. Maximum
                            value is 3600. Must be smaller than periodSeconds. More
                            info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes.'
                          type: integer
                      type: object
                    name:
                      description: Name of the container specified as a DNS_LABEL.
                      type: string
                    ports:
                      description: |-
                        List of ports to expose from the container. Only a single port can be specified. The specified ports must be listening on all interfaces (0.0.0.0) within the container to be accessible.

                        If omitted, a port number will be chosen and passed to the container through the PORT environment variable for the container to listen on.
                      items:
                        properties:
                          containerPort:
                            description: Port number the container listens on. This
                              must be a valid TCP port number, 0 < containerPort <
                              65536.
                            type: integer
                          name:
                            description: If specified, used to specify which protocol
                              to use. Allowed values are "http1" and "h2c".
                            type: string
                        type: object
                      type: array
                    resources:
                      description: 'Compute Resource requirements by this container.
                        More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.'
                      properties:
                        limits:
                          additionalProperties:
                            type: string
                          description: 'Only memory and CPU are supported. Note: The
                            only supported values for CPU are ''1'', ''2'', ''4'',
                            and ''8''. Setting 4 CPU requires at least 2Gi of memory.
                            The values of the map is string form of the ''quantity''
                            k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go.'
                          type: object
                      type: object
                    startupProbe:
                      description: |-
                        DEPRECATED. Cloud Run Job does not support startup probe and `startup_probe` field will be removed in a future major release. Startup probe of application within the container. All other probes are disabled if a startup probe is provided, until it succeeds. Container will not be added to service endpoints if the probe fails. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
                        This field is not supported in Cloud Run Job currently.
                      properties:
                        failureThreshold:
                          description: Minimum consecutive failures for the probe
                            to be considered failed after having succeeded. Defaults
                            to 3. Minimum value is 1.
                          type: integer
                        httpGet:
                          description: HTTPGet specifies the http request to perform.
                            Exactly one of HTTPGet or TCPSocket must be specified.
                          properties:
                            httpHeaders:
                              description: Custom headers to set in the request. HTTP
                                allows repeated headers.
                              items:
                                properties:
                                  name:
                                    description: The header field name.
                                    type: string
                                  value:
                                    description: The header field value.
                                    type: string
                                required:
                                - name
                                type: object
                              type: array
                            path:
                              description: Path to access on the HTTP server. Defaults
                                to '/'.
                              type: string
                          type: object
                        initialDelaySeconds:
                          description: 'Number of seconds after the container has
                            started before the probe is initiated. Defaults to 0 seconds.
                            Minimum value is 0. Maximum value for liveness probe is
                            3600. Maximum value for startup probe is 240. More info:
                            https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes.'
                          type: integer
                        periodSeconds:
                          description: How often (in seconds) to perform the probe.
                            Default to 10 seconds. Minimum value is 1. Maximum value
                            for liveness probe is 3600. Maximum value for startup
                            probe is 240. Must be greater or equal than timeoutSeconds.
                          type: integer
                        tcpSocket:
                          description: TCPSocket specifies an action involving a TCP
                            port. Exactly one of HTTPGet or TCPSocket must be specified.
                          properties:
                            port:
                              description: Port number to access on the container.
                                Must be in the range 1 to 65535. If not specified,
                                defaults to 8080.
                              type: integer
                          type: object
                        timeoutSeconds:
                          description: 'Number of seconds after which the probe times
                            out. Defaults to 1 second. Minimum value is 1. Maximum
                            value is 3600. Must be smaller than periodSeconds. More
                            info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes.'
                          type: integer
                      type: object
                    volumeMounts:
                      description: Volume to mount into the container's filesystem.
                      items:
                        properties:
                          mountPath:
                            description: Path within the container at which the volume
                              should be mounted. Must not contain ':'. For Cloud SQL
                              volumes, it can be left empty, or must otherwise be
                              /cloudsql. All instances defined in the Volume will
                              be available as /cloudsql/[instance]. For more information
                              on Cloud SQL volumes, visit https://cloud.google.com/sql/docs/mysql/connect-run.
                            type: string
                          name:
                            description: This must match the Name of a Volume.
                            type: string
                        required:
                        - mountPath
                        - name
                        type: object
                      type: array
                    workingDir:
                      description: Container's working directory. If not specified,
                        the container runtime's default will be used, which might
                        be configured in the container image.
                      type: string
                  required:
                  - image
                  type: object
                type: array
              encryptionKeyRef:
                description: A reference to a customer managed encryption key (CMEK)
                  to use to encrypt this container image. For more information, go
                  to https://cloud.google.com/run/docs/securing/using-cmek
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
              executionEnvironment:
                description: 'The execution environment being used to host this Task.
                  Possible values: ["EXECUTION_ENVIRONMENT_GEN1", "EXECUTION_ENVIRONMENT_GEN2"].'
                type: string
              maxRetries:
                description: Number of retries allowed per Task, before marking this
                  Task failed.
                type: integer
              serviceAccountRef:
                description: Email address of the IAM service account associated with
                  the revision of the service. The service account represents the
                  identity of the running revision, and determines what permissions
                  the revision has. If not provided, the revision will use the project's
                  default service account.
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
              template:
                description: Describes the task(s) that will be created when executing
                  an execution.
                type: object
              timeout:
                description: |-
                  Max allowed time duration the Task may be active before the system will actively try to mark it failed and kill associated containers. This applies per attempt of a task, meaning each retry can run for the full timeout.

                  A duration in seconds with up to nine fractional digits, ending with 's'. Example: "3.5s".
                type: string
required:
- spec
type: object

```
