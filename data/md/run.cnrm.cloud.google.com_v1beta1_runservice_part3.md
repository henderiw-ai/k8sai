# openAPI schema part3 for run.cnrm.cloud.google.com.v1beta1.RunService

## schema

```yaml
properties:
  spec:
    properties:
      template:
        properties:
          volumes:
            properties:
              volumes:
                description: A list of Volumes to make available to containers.
                items:
                  properties:
                    cloudSqlInstance:
                      description: For Cloud SQL volumes, contains the specific instances
                        that should be mounted. Visit https://cloud.google.com/sql/docs/mysql/connect-run
                        for more information on how to connect Cloud SQL and Cloud
                        Run.
                      properties:
                        instances:
                          items:
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
                                description: 'Allowed value: The `instanceName` field
                                  of a `SQLInstance` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          type: array
                      type: object
                    name:
                      description: Required. Volume's name.
                      type: string
                    secret:
                      description: 'Secret represents a secret that should populate
                        this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret'
                      properties:
                        defaultMode:
                          description: 'Integer representation of mode bits to use
                            on created files by default. Must be a value between 0000
                            and 0777 (octal), defaulting to 0644. Directories within
                            the path are not affected by this setting. Notes * Internally,
                            a umask of 0222 will be applied to any non-zero value.
                            * This is an integer representation of the mode bits.
                            So, the octal integer value should look exactly as the
                            chmod numeric notation with a leading zero. Some examples:
                            for chmod 777 (a=rwx), set to 0777 (octal) or 511 (base-10).
                            For chmod 640 (u=rw,g=r), set to 0640 (octal) or 416 (base-10).
                            For chmod 755 (u=rwx,g=rx,o=rx), set to 0755 (octal) or
                            493 (base-10). * This might be in conflict with other
                            options that affect the file mode, like fsGroup, and the
                            result can be other mode bits set. This might be in conflict
                            with other options that affect the file mode, like fsGroup,
                            and as a result, other mode bits could be set.'
                          format: int64
                          type: integer
                        items:
                          description: If unspecified, the volume will expose a file
                            whose name is the secret, relative to VolumeMount.mount_path.
                            If specified, the key will be used as the version to fetch
                            from Cloud Secret Manager and the path will be the name
                            of the file exposed in the volume. When items are defined,
                            they must specify a path and a version.
                          items:
                            properties:
                              mode:
                                description: 'Integer octal mode bits to use on this
                                  file, must be a value between 01 and 0777 (octal).
                                  If 0 or not set, the Volume''s default mode will
                                  be used. Notes * Internally, a umask of 0222 will
                                  be applied to any non-zero value. * This is an integer
                                  representation of the mode bits. So, the octal integer
                                  value should look exactly as the chmod numeric notation
                                  with a leading zero. Some examples: for chmod 777
                                  (a=rwx), set to 0777 (octal) or 511 (base-10). For
                                  chmod 640 (u=rw,g=r), set to 0640 (octal) or 416
                                  (base-10). For chmod 755 (u=rwx,g=rx,o=rx), set
                                  to 0755 (octal) or 493 (base-10). * This might be
                                  in conflict with other options that affect the file
                                  mode, like fsGroup, and the result can be other
                                  mode bits set.'
                                format: int64
                                type: integer
                              path:
                                description: Required. The relative path of the secret
                                  in the container.
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
                            - path
                            type: object
                          type: array
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
                                Required. The name of the secret in Cloud Secret Manager. Format: {secret} if the secret is in the same project. projects/{project}/secrets/{secret} if the secret is in a different project.

                                Allowed value: The Google Cloud resource name of a `SecretManagerSecret` resource (format: `projects/{{project}}/secrets/{{name}}`).
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
            description: VPC Access configuration to use for this Revision. For more
              information, visit https://cloud.google.com/run/docs/configuring/connecting-vpc.
            properties:
              connectorRef:
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
                      VPC Access connector name. Format: projects/{project}/locations/{location}/connectors/{connector}

                      Allowed value: The Google Cloud resource name of a `VPCAccessConnector` resource (format: `projects/{{project}}/locations/{{location}}/connectors/{{name}}`).
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              egress:
                description: 'Traffic VPC egress settings. Possible values: VPC_EGRESS_UNSPECIFIED,
                  ALL_TRAFFIC, PRIVATE_RANGES_ONLY'
                type: string
            type: object
      traffic:
        description: Specifies how to distribute traffic over a collection of Revisions
          belonging to the Service. If traffic is empty or not provided, defaults
          to 100% traffic to the latest `Ready` Revision.
        items:
          properties:
            percent:
              description: Specifies percent of the traffic to this Revision. This
                defaults to zero if unspecified. Cloud Run currently requires 100
                percent for a single TrafficTarget entry.
              format: int64
              type: integer
            revision:
              description: Revision to which to send this portion of traffic, if traffic
                allocation is by revision.
              type: string
            tag:
              description: Indicates a string to be part of the URI to exclusively
                reference this target.
              type: string
            type:
              description: 'The allocation type for this traffic target. Possible
                values: TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED, TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST,
                TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION'
              type: string
          type: object
        type: array
required:
- spec
type: object

```
