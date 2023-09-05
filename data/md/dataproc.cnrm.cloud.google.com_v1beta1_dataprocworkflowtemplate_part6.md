# openAPI schema part6 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      placement:
        properties:
          managedCluster:
            properties:
              config:
                properties:
                  masterConfig:
                    properties:
                      masterConfig:
                        description: Immutable. Optional. The Compute Engine config
                          settings for the master instance in a cluster.
                        properties:
                          accelerators:
                            description: Immutable. Optional. The Compute Engine accelerator
                              configuration for these instances.
                            items:
                              properties:
                                acceleratorCount:
                                  description: Immutable. The number of the accelerator
                                    cards of this type exposed to this instance.
                                  format: int64
                                  type: integer
                                acceleratorType:
                                  description: 'Immutable. Full URL, partial URI,
                                    or short name of the accelerator type resource
                                    to expose to this instance. See [Compute Engine
                                    AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes).
                                    Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                                    * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                                    * `nvidia-tesla-k80` **Auto Zone Exception**:
                                    If you are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                                    feature, you must use the short name of the accelerator
                                    type resource, for example, `nvidia-tesla-k80`.'
                                  type: string
                              type: object
                            type: array
                          diskConfig:
                            description: Immutable. Optional. Disk option config settings.
                            properties:
                              bootDiskSizeGb:
                                description: Immutable. Optional. Size in GB of the
                                  boot disk (default is 500GB).
                                format: int64
                                type: integer
                              bootDiskType:
                                description: 'Immutable. Optional. Type of the boot
                                  disk (default is "pd-standard"). Valid values: "pd-balanced"
                                  (Persistent Disk Balanced Solid State Drive), "pd-ssd"
                                  (Persistent Disk Solid State Drive), or "pd-standard"
                                  (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).'
                                type: string
                              numLocalSsds:
                                description: Immutable. Optional. Number of attached
                                  SSDs, from 0 to 4 (default is 0). If SSDs are not
                                  attached, the boot disk is used to store runtime
                                  logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html)
                                  data. If one or more SSDs are attached, this runtime
                                  bulk data is spread across them, and the boot disk
                                  contains only basic config and installed binaries.
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
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          machineType:
                            description: 'Immutable. Optional. The Compute Engine
                              machine type used for cluster instances. A full URL,
                              partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                              * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                              * `n1-standard-2` **Auto Zone Exception**: If you are
                              using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                              feature, you must use the short name of the machine
                              type resource, for example, `n1-standard-2`.'
                            type: string
                          minCpuPlatform:
                            description: Immutable. Optional. Specifies the minimum
                              cpu platform for the Instance Group. See [Dataproc ->
                              Minimum CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
                            type: string
                          numInstances:
                            description: Immutable. Optional. The number of VM instances
                              in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
                              [master_config](#FIELDS.master_config) groups, **must
                              be set to 3**. For standard cluster [master_config](#FIELDS.master_config)
                              groups, **must be set to 1**.
                            format: int64
                            type: integer
                          preemptibility:
                            description: 'Immutable. Optional. Specifies the preemptibility
                              of the instance group. The default value for master
                              and worker groups is `NON_PREEMPTIBLE`. This default
                              cannot be changed. The default value for secondary instances
                              is `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED,
                              NON_PREEMPTIBLE, PREEMPTIBLE'
                            type: string
                        type: object
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
                              description: Immutable. The number of the accelerator
                                cards of this type exposed to this instance.
                              format: int64
                              type: integer
                            acceleratorType:
                              description: 'Immutable. Full URL, partial URI, or short
                                name of the accelerator type resource to expose to
                                this instance. See [Compute Engine AcceleratorTypes](https://cloud.google.com/compute/docs/reference/beta/acceleratorTypes).
                                Examples: * `https://www.googleapis.com/compute/beta/projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                                * `projects/[project_id]/zones/us-east1-a/acceleratorTypes/nvidia-tesla-k80`
                                * `nvidia-tesla-k80` **Auto Zone Exception**: If you
                                are using the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                                feature, you must use the short name of the accelerator
                                type resource, for example, `nvidia-tesla-k80`.'
                              type: string
                          type: object
                        type: array
                      diskConfig:
                        description: Immutable. Optional. Disk option config settings.
                        properties:
                          bootDiskSizeGb:
                            description: Immutable. Optional. Size in GB of the boot
                              disk (default is 500GB).
                            format: int64
                            type: integer
                          bootDiskType:
                            description: 'Immutable. Optional. Type of the boot disk
                              (default is "pd-standard"). Valid values: "pd-balanced"
                              (Persistent Disk Balanced Solid State Drive), "pd-ssd"
                              (Persistent Disk Solid State Drive), or "pd-standard"
                              (Persistent Disk Hard Disk Drive). See [Disk types](https://cloud.google.com/compute/docs/disks#disk-types).'
                            type: string
                          numLocalSsds:
                            description: Immutable. Optional. Number of attached SSDs,
                              from 0 to 4 (default is 0). If SSDs are not attached,
                              the boot disk is used to store runtime logs and [HDFS](https://hadoop.apache.org/docs/r1.2.1/hdfs_user_guide.html)
                              data. If one or more SSDs are attached, this runtime
                              bulk data is spread across them, and the boot disk contains
                              only basic config and installed binaries.
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
                          type used for cluster instances. A full URL, partial URI,
                          or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                          * `projects/[project_id]/zones/us-east1-a/machineTypes/n1-standard-2`
                          * `n1-standard-2` **Auto Zone Exception**: If you are using
                          the Dataproc [Auto Zone Placement](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/auto-zone#using_auto_zone_placement)
                          feature, you must use the short name of the machine type
                          resource, for example, `n1-standard-2`.'
                        type: string
                      minCpuPlatform:
                        description: Immutable. Optional. Specifies the minimum cpu
                          platform for the Instance Group. See [Dataproc -> Minimum
                          CPU Platform](https://cloud.google.com/dataproc/docs/concepts/compute/dataproc-min-cpu).
                        type: string
                      numInstances:
                        description: Immutable. Optional. The number of VM instances
                          in the instance group. For [HA cluster](/dataproc/docs/concepts/configuring-clusters/high-availability)
                          [master_config](#FIELDS.master_config) groups, **must be
                          set to 3**. For standard cluster [master_config](#FIELDS.master_config)
                          groups, **must be set to 1**.
                        format: int64
                        type: integer
                      preemptibility:
                        description: 'Immutable. Optional. Specifies the preemptibility
                          of the instance group. The default value for master and
                          worker groups is `NON_PREEMPTIBLE`. This default cannot
                          be changed. The default value for secondary instances is
                          `PREEMPTIBLE`. Possible values: PREEMPTIBILITY_UNSPECIFIED,
                          NON_PREEMPTIBLE, PREEMPTIBLE'
                        type: string
                    type: object
required:
- spec
type: object

```
