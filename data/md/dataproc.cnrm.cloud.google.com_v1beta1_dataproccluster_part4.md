# openAPI schema part4 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocCluster

## schema

```yaml
properties:
  spec:
    properties:
      config:
        properties:
          secondaryWorkerConfig:
            properties:
              secondaryWorkerConfig:
                description: Immutable. Optional. The Compute Engine config settings
                  for additional worker instances in a cluster.
                properties:
                  accelerators:
                    description: Immutable. Optional. The Compute Engine accelerator
                      configuration for these instances.
                    items:
                      properties:
                        acceleratorCount:
                          description: Immutable. The number of the accelerator cards
                            of this type exposed to this instance.
                          format: int64
                          type: integer
                        acceleratorType:
                          description: 'Immutable. Full URL, partial URI, or short
                            name of the accelerator type resource to expose to this
                            instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes).
                            Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                            * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                            * `nvidia-tesla-k80` **Auto Zone Exception**: If you are
                            using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                            feature, you must use the short name of the accelerator
                            type resource, for example, `nvidia-tesla-k80`.'
                          type: string
                      type: object
                    type: array
                  diskConfig:
                    description: Immutable. Optional. Disk option config settings.
                    properties:
                      bootDiskSizeGb:
                        description: Immutable. Optional. Size in GB of the boot disk
                          (default is 500GB).
                        format: int64
                        type: integer
                      bootDiskType:
                        description: 'Immutable. Optional. Type of the boot disk (default
                          is "pd-standard"). Valid values: "pd-balanced" (Persistent
                          Disk Balanced Solid State Drive), "pd-ssd" (Persistent Disk
                          Solid State Drive), or "pd-standard" (Persistent Disk Hard
                          Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).'
                        type: string
                      localSsdInterface:
                        description: 'Immutable. Optional. Interface type of local
                          SSDs (default is "scsi"). Valid values: "scsi" (Small Computer
                          System Interface), "nvme" (Non-Volatile Memory Express).
                          See [local SSD performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).'
                        type: string
                      numLocalSsds:
                        description: Immutable. Optional. Number of attached SSDs,
                          from 0 to 4 (default is 0). If SSDs are not attached, the
                          boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html)
                          data. If one or more SSDs are attached, this runtime bulk
                          data is spread across them, and the boot disk contains only
                          basic config and installed binaries.
                        format: int64
                        type: integer
                    type: object
                  imageRef:
                    description: Immutable.
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
                          Optional. The Compute Engine image resource used for cluster instances. The URI can represent an image or image family. Image examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/[image-id]` * `projects/[project_id]/global/images/[image-id]` * `image-id` Image family examples. Dataproc will use the most recent image from the family: * `https://www.googleapis.com/compute/beta/projects/[project_id]/global/images/family/[custom-image-family-name]` * `projects/[project_id]/global/images/family/[custom-image-family-name]` If the URI is unspecified, it will be inferred from `SoftwareConfig.image_version` or the system default.

                          Allowed value: The `selfLink` field of a `ComputeImage` resource.
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  machineType:
                    description: 'Immutable. Optional. The Compute Engine machine
                      type used for cluster instances. A full URL, partial URI, or
                      short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                      * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                      * `n1-standard-2` **Auto Zone Exception**: If you are using
                      the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                      feature, you must use the short name of the machine type resource,
                      for example, `n1-standard-2`.'
                    type: string
                  minCpuPlatform:
                    description: Immutable. Optional. Specifies the minimum cpu platform
                      for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
                    type: string
                  numInstances:
                    description: Immutable. Optional. The number of VM instances in
                      the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
                      [master_config](#FIELDS.master_config) groups, **must be set
                      to 3**. For standard cluster [master_config](#FIELDS.master_config)
                      groups, **must be set to 1**.
                    format: int64
                    type: integer
                  preemptibility:
                    description: 'Immutable. Optional. Specifies the preemptibility
                      of the instance group. The default value for master and worker
                      groups is `NON_PREEMPTIBLE`. This default cannot be changed.
                      The default value for secondary instances is `PREEMPTIBLE`.
                      Possible values: PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE,
                      PREEMPTIBLE'
                    type: string
                type: object
          securityConfig:
            description: Immutable. Optional. Security settings for the cluster.
            properties:
              identityConfig:
                description: Immutable. Optional. Identity related configuration,
                  including service account based secure multi-tenancy user mappings.
                properties:
                  userServiceAccountMapping:
                    additionalProperties:
                      type: string
                    description: Immutable. Required. Map of user to service account.
                    type: object
                required:
                - userServiceAccountMapping
                type: object
              kerberosConfig:
                description: Immutable. Optional. Kerberos related configuration.
                properties:
                  crossRealmTrustAdminServer:
                    description: Immutable. Optional. The admin server (IP or hostname)
                      for the remote trusted realm in a cross realm trust relationship.
                    type: string
                  crossRealmTrustKdc:
                    description: Immutable. Optional. The KDC (IP or hostname) for
                      the remote trusted realm in a cross realm trust relationship.
                    type: string
                  crossRealmTrustRealm:
                    description: Immutable. Optional. The remote realm the Dataproc
                      on-cluster KDC will trust, should the user enable cross realm
                      trust.
                    type: string
                  crossRealmTrustSharedPassword:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the shared password between the on-cluster
                      Kerberos realm and the remote trusted realm, in a cross realm
                      trust relationship.
                    type: string
                  enableKerberos:
                    description: 'Immutable. Optional. Flag to indicate whether to
                      Kerberize the cluster (default: false). Set this field to true
                      to enable Kerberos on a cluster.'
                    type: boolean
                  kdcDbKey:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the master key of the KDC database.
                    type: string
                  keyPassword:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the password to the user provided
                      key. For the self-signed certificate, this password is generated
                      by Dataproc.
                    type: string
                  keystore:
                    description: Immutable. Optional. The Cloud Storage URI of the
                      keystore file used for SSL encryption. If not provided, Dataproc
                      will provide a self-signed certificate.
                    type: string
                  keystorePassword:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the password to the user provided
                      keystore. For the self-signed certificate, this password is
                      generated by Dataproc.
                    type: string
                  kmsKeyRef:
                    description: Immutable.
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
                          Optional. The uri of the KMS key used to encrypt various sensitive files.

                          Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  realm:
                    description: Immutable. Optional. The name of the on-cluster Kerberos
                      realm. If not specified, the uppercased domain of hostnames
                      will be the realm.
                    type: string
                  rootPrincipalPassword:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the root principal password.
                    type: string
                  tgtLifetimeHours:
                    description: Immutable. Optional. The lifetime of the ticket granting
                      ticket, in hours. If not specified, or user specifies 0, then
                      default value 10 will be used.
                    format: int64
                    type: integer
                  truststore:
                    description: Immutable. Optional. The Cloud Storage URI of the
                      truststore file used for SSL encryption. If not provided, Dataproc
                      will provide a self-signed certificate.
                    type: string
                  truststorePassword:
                    description: Immutable. Optional. The Cloud Storage URI of a KMS
                      encrypted file containing the password to the user provided
                      truststore. For the self-signed certificate, this password is
                      generated by Dataproc.
                    type: string
                type: object
            type: object
          softwareConfig:
            description: Immutable. Optional. The config settings for software inside
              the cluster.
            properties:
              imageVersion:
                description: Immutable. Optional. The version of software inside the
                  cluster. It must be one of the supported [Dataproc Versions](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#supported_dataproc_versions),
                  such as "1.2" (including a subminor version, such as "1.2.29"),
                  or the ["preview" version](https://cloud.google.com/dataproc/docs/concepts/versioning/dataproc-versions#other_versions).
                  If unspecified, it defaults to the latest Debian version.
                type: string
              optionalComponents:
                description: Immutable. Optional. The set of components to activate
                  on the cluster.
                items:
                  type: string
                type: array
              properties:
                additionalProperties:
                  type: string
                description: 'Immutable. Optional. The properties to set on daemon
                  config files. Property keys are specified in `prefix:property` format,
                  for example `core:hadoop.tmp.dir`. The following are supported prefixes
                  and their mappings: * capacity-scheduler: `capacity-scheduler.xml`
                  * core: `core-site.xml` * distcp: `distcp-default.xml` * hdfs: `hdfs-site.xml`
                  * hive: `hive-site.xml` * mapred: `mapred-site.xml` * pig: `pig.properties`
                  * spark: `spark-defaults.conf` * yarn: `yarn-site.xml` For more
                  information, see [Cluster properties](https://cloud.google.com/dataproc/docs/concepts/cluster-properties).'
                type: object
            type: object
          stagingBucketRef:
            description: Immutable.
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
                  Optional. A Cloud Storage bucket used to stage job dependencies, config files, and job driver console output. If you do not specify a staging bucket, Cloud Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's staging bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket (see [Dataproc staging bucket](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/staging-bucket)). **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**

                  Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          tempBucketRef:
            description: Immutable.
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
                  Optional. A Cloud Storage bucket used to store ephemeral cluster and jobs data, such as Spark and MapReduce history files. If you do not specify a temp bucket, Dataproc will determine a Cloud Storage location (US, ASIA, or EU) for your cluster's temp bucket according to the Compute Engine zone where your cluster is deployed, and then create and manage this project-level, per-location bucket. The default bucket has a TTL of 90 days, but you can use any TTL (or none) if you specify a bucket. **This field requires a Cloud Storage bucket name, not a URI to a Cloud Storage bucket.**

                  Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
required:
- spec
type: object

```
