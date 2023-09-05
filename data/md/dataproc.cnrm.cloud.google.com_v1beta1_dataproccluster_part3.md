# openAPI schema part3 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocCluster

## schema

```yaml
properties:
  spec:
    properties:
      config:
        properties:
          initializationActions:
            properties:
              initializationActions:
                description: 'Immutable. Optional. Commands to execute on each node
                  after config is completed. By default, executables are run on master
                  and all worker nodes. You can test a node''s `role` metadata to
                  run an executable on a master or worker node, as shown below using
                  `curl` (you can also use `wget`): ROLE=$(curl -H Metadata-Flavor:Google
                  http://metadata/computeMetadata/v1/instance/attributes/dataproc-role)
                  if [[ "${ROLE}" == ''Master'' ]]; then ... master specific actions
                  ... else ... worker specific actions ... fi'
                items:
                  properties:
                    executableFile:
                      description: Immutable. Required. Cloud Storage URI of executable
                        file.
                      type: string
                    executionTimeout:
                      description: Immutable. Optional. Amount of time executable
                        has to complete. Default is 10 minutes (see JSON representation
                        of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                        Cluster creation fails with an explanatory error message (the
                        name of the executable that caused the error and the exceeded
                        timeout period) if the executable is not completed at end
                        of the timeout period.
                      type: string
                  required:
                  - executableFile
                  type: object
                type: array
          lifecycleConfig:
            description: Immutable. Optional. Lifecycle setting for the cluster.
            properties:
              autoDeleteTime:
                description: Immutable. Optional. The time when cluster will be auto-deleted
                  (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                format: date-time
                type: string
              autoDeleteTtl:
                description: Immutable. Optional. The lifetime duration of cluster.
                  The cluster will be auto-deleted at the end of this period. Minimum
                  value is 10 minutes; maximum value is 14 days (see JSON representation
                  of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                type: string
              idleDeleteTtl:
                description: Immutable. Optional. The duration to keep the cluster
                  alive while idling (when no jobs are running). Passing this threshold
                  will cause the cluster to be deleted. Minimum value is 5 minutes;
                  maximum value is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                type: string
            type: object
          masterConfig:
            description: Immutable. Optional. The Compute Engine config settings for
              the master instance in a cluster.
            properties:
              accelerators:
                description: Immutable. Optional. The Compute Engine accelerator configuration
                  for these instances.
                items:
                  properties:
                    acceleratorCount:
                      description: Immutable. The number of the accelerator cards
                        of this type exposed to this instance.
                      format: int64
                      type: integer
                    acceleratorType:
                      description: 'Immutable. Full URL, partial URI, or short name
                        of the accelerator type resource to expose to this instance.
                        See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes).
                        Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                        * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                        * `nvidia-tesla-k80` **Auto Zone Exception**: If you are using
                        the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                        feature, you must use the short name of the accelerator type
                        resource, for example, `nvidia-tesla-k80`.'
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
                      is "pd-standard"). Valid values: "pd-balanced" (Persistent Disk
                      Balanced Solid State Drive), "pd-ssd" (Persistent Disk Solid
                      State Drive), or "pd-standard" (Persistent Disk Hard Disk Drive).
                      See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).'
                    type: string
                  localSsdInterface:
                    description: 'Immutable. Optional. Interface type of local SSDs
                      (default is "scsi"). Valid values: "scsi" (Small Computer System
                      Interface), "nvme" (Non-Volatile Memory Express). See [local
                      SSD performance](https://cloud.google.com/compute/docs/disks/local-ssd#performance).'
                    type: string
                  numLocalSsds:
                    description: Immutable. Optional. Number of attached SSDs, from
                      0 to 4 (default is 0). If SSDs are not attached, the boot disk
                      is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html)
                      data. If one or more SSDs are attached, this runtime bulk data
                      is spread across them, and the boot disk contains only basic
                      config and installed binaries.
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
                description: 'Immutable. Optional. The Compute Engine machine type
                  used for cluster instances. A full URL, partial URI, or short name
                  are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                  * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                  * `n1-standard-2` **Auto Zone Exception**: If you are using the
                  Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                  feature, you must use the short name of the machine type resource,
                  for example, `n1-standard-2`.'
                type: string
              minCpuPlatform:
                description: Immutable. Optional. Specifies the minimum cpu platform
                  for the Instance Group. See [Dataproc -> Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
                type: string
              numInstances:
                description: Immutable. Optional. The number of VM instances in the
                  instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
                  [master_config](#FIELDS.master_config) groups, **must be set to
                  3**. For standard cluster [master_config](#FIELDS.master_config)
                  groups, **must be set to 1**.
                format: int64
                type: integer
              preemptibility:
                description: 'Immutable. Optional. Specifies the preemptibility of
                  the instance group. The default value for master and worker groups
                  is `NON_PREEMPTIBLE`. This default cannot be changed. The default
                  value for secondary instances is `PREEMPTIBLE`. Possible values:
                  PREEMPTIBILITY_UNSPECIFIED, NON_PREEMPTIBLE, PREEMPTIBLE'
                type: string
            type: object
          metastoreConfig:
            description: Immutable. Optional. Metastore configuration.
            properties:
              dataprocMetastoreServiceRef:
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
                    description: 'Required. Resource name of an existing Dataproc
                      Metastore service. Example: * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`'
                    type: string
                  name:
                    description: |-
                      [WARNING] DataprocMetastoreService not yet supported in Config Connector, use 'external' field to reference existing resources.
                      Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            required:
            - dataprocMetastoreServiceRef
            type: object
required:
- spec
type: object

```
