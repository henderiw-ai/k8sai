# openAPI schema part5 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

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
                  autoscalingConfig:
                    description: Immutable. Optional. Autoscaling config for the policy
                      associated with the cluster. Cluster does not autoscale if this
                      field is unset.
                    properties:
                      policyRef:
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
                              Optional. The autoscaling policy used by the cluster. Only resource names including projectid and location (region) are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` * `projects/[project_id]/locations/[dataproc_region]/autoscalingPolicies/[policy_id]` Note that the policy must be in the same project and Dataproc region.

                              Allowed value: The Google Cloud resource name of a `DataprocAutoscalingPolicy` resource (format: `projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{name}}`).
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                    type: object
                  config:
                    description: Immutable. Required. The cluster configuration.
                    type: object
                  encryptionConfig:
                    description: Immutable. Optional. Encryption settings for the
                      cluster.
                    properties:
                      gcePdKmsKeyRef:
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
                              Optional. The Cloud KMS key name to use for PD disk encryption for all instances in the cluster.

                              Allowed value: The `selfLink` field of a `KMSCryptoKey` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                    type: object
                  endpointConfig:
                    description: Immutable. Optional. Port/endpoint configuration
                      for this cluster
                    properties:
                      enableHttpPortAccess:
                        description: Immutable. Optional. If true, enable http access
                          to specific ports on the cluster from external sources.
                          Defaults to false.
                        type: boolean
                    type: object
                  gceClusterConfig:
                    description: Immutable. Optional. The shared Compute Engine config
                      settings for all instances in a cluster.
                    properties:
                      internalIPOnly:
                        description: Immutable. Optional. If true, all instances in
                          the cluster will only have internal IP addresses. By default,
                          clusters are not restricted to internal IP addresses, and
                          will have ephemeral external IP addresses assigned to each
                          instance. This `internal_ip_only` restriction can only be
                          enabled for subnetwork enabled networks, and all off-cluster
                          dependencies must be configured to be accessible without
                          external IP addresses.
                        type: boolean
                      metadata:
                        additionalProperties:
                          type: string
                        description: Immutable. The Compute Engine metadata entries
                          to add to all instances (see [Project and instance metadata](https://cloud.google.com/compute/docs/storing-retrieving-metadata#project_and_instance_metadata)).
                        type: object
                      networkRef:
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
                              Optional. The Compute Engine network to be used for machine communications. Cannot be specified with subnetwork_uri. If neither `network_uri` nor `subnetwork_uri` is specified, the "default" network of the project is used, if it exists. Cannot be a "Custom Subnet Network" (see [Using Subnetworks](https://cloud.google.com/compute/docs/subnetworks) for more information). A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/global/default` * `projects/[project_id]/regions/global/default` * `default`

                              Allowed value: The `selfLink` field of a `ComputeNetwork` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                      nodeGroupAffinity:
                        description: Immutable. Optional. Node Group Affinity for
                          sole-tenant clusters.
                        properties:
                          nodeGroupRef:
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
                                  Required. The URI of a sole-tenant [node group resource](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups) that the cluster will be created on. A full URL, partial URI, or node group name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `projects/[project_id]/zones/us-central1-a/nodeGroups/node-group-1` * `node-group-1`

                                  Allowed value: The `selfLink` field of a `ComputeNodeGroup` resource.
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
                        - nodeGroupRef
                        type: object
                      privateIPv6GoogleAccess:
                        description: 'Immutable. Optional. The type of IPv6 access
                          for a cluster. Possible values: PRIVATE_IPV6_GOOGLE_ACCESS_UNSPECIFIED,
                          INHERIT_FROM_SUBNETWORK, OUTBOUND, BIDIRECTIONAL'
                        type: string
                      reservationAffinity:
                        description: Immutable. Optional. Reservation Affinity for
                          consuming Zonal reservation.
                        properties:
                          consumeReservationType:
                            description: 'Immutable. Optional. Type of reservation
                              to consume Possible values: TYPE_UNSPECIFIED, NO_RESERVATION,
                              ANY_RESERVATION, SPECIFIC_RESERVATION'
                            type: string
                          key:
                            description: Immutable. Optional. Corresponds to the label
                              key of reservation resource.
                            type: string
                          values:
                            description: Immutable. Optional. Corresponds to the label
                              values of reservation resource.
                            items:
                              type: string
                            type: array
                        type: object
                      serviceAccountRef:
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
                              Optional. The [Dataproc service account](https://cloud.google.com/dataproc/docs/concepts/configuring-clusters/service-accounts#service_accounts_in_dataproc) (also see [VM Data Plane identity](https://cloud.google.com/dataproc/docs/concepts/iam/dataproc-principals#vm_service_account_data_plane_identity)) used by Dataproc cluster VM instances to access Google Cloud Platform services. If not specified, the [Compute Engine default service account](https://cloud.google.com/compute/docs/access/service-accounts#default_service_account) is used.

                              Allowed value: The `email` field of an `IAMServiceAccount` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                      serviceAccountScopes:
                        description: 'Immutable. Optional. The URIs of service account
                          scopes to be included in Compute Engine instances. The following
                          base set of scopes is always included: * https://www.googleapis.com/auth/cloud.useraccounts.readonly
                          * https://www.googleapis.com/auth/devstorage.read_write
                          * https://www.googleapis.com/auth/logging.write If no scopes
                          are specified, the following defaults are also provided:
                          * https://www.googleapis.com/auth/bigquery * https://www.googleapis.com/auth/bigtable.admin.table
                          * https://www.googleapis.com/auth/bigtable.data * https://www.googleapis.com/auth/devstorage.full_control'
                        items:
                          type: string
                        type: array
                      shieldedInstanceConfig:
                        description: Immutable. Optional. Shielded Instance Config
                          for clusters using Compute Engine Shielded VMs.
                        properties:
                          enableIntegrityMonitoring:
                            description: Immutable. Optional. Defines whether instances
                              have integrity monitoring enabled. Integrity monitoring
                              compares the most recent boot measurements to the integrity
                              policy baseline and returns a pair of pass/fail results
                              depending on whether they match or not.
                            type: boolean
                          enableSecureBoot:
                            description: Immutable. Optional. Defines whether the
                              instances have Secure Boot enabled. Secure Boot helps
                              ensure that the system only runs authentic software
                              by verifying the digital signature of all boot components,
                              and halting the boot process if signature verification
                              fails.
                            type: boolean
                          enableVtpm:
                            description: Immutable. Optional. Defines whether the
                              instance have the vTPM enabled. Virtual Trusted Platform
                              Module protects objects like keys, certificates and
                              enables Measured Boot by performing the measurements
                              needed to create a known good boot baseline, called
                              the integrity policy baseline.
                            type: boolean
                        type: object
                      subnetworkRef:
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
                              Optional. The Compute Engine subnetwork to be used for machine communications. Cannot be specified with network_uri. A full URL, partial URI, or short name are valid. Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/regions/us-east1/subnetworks/sub0` * `projects/[project_id]/regions/us-east1/subnetworks/sub0` * `sub0`

                              Allowed value: The `selfLink` field of a `ComputeSubnetwork` resource.
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                      tags:
                        description: Immutable. The Compute Engine tags to add to
                          all instances (see [Tagging instances](https://cloud.google.com/compute/docs/label-or-tag-resources#tags)).
                        items:
                          type: string
                        type: array
                      zone:
                        description: 'Immutable. Optional. The zone where the Compute
                          Engine cluster will be located. On a create request, it
                          is required in the "global" region. If omitted in a non-global
                          Dataproc region, the service will pick a zone in the corresponding
                          Compute Engine region. On a get request, zone will always
                          be present. A full URL, partial URI, or short name are valid.
                          Examples: * `https://www.googleapis.com/compute/v1/projects/[project_id]/zones/[zone]`
                          * `projects/[project_id]/zones/[zone]` * `us-central1-f`'
                        type: string
                    type: object
                  initializationActions:
                    description: 'Immutable. Optional. Commands to execute on each
                      node after config is completed. By default, executables are
                      run on master and all worker nodes. You can test a node''s `role`
                      metadata to run an executable on a master or worker node, as
                      shown below using `curl` (you can also use `wget`): ROLE=$(curl
                      -H Metadata-Flavor:Google http://metadata/computeMetadata/v1/instance/attributes/dataproc-role)
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
                            Cluster creation fails with an explanatory error message
                            (the name of the executable that caused the error and
                            the exceeded timeout period) if the executable is not
                            completed at end of the timeout period.
                          type: string
                      type: object
                    type: array
                  lifecycleConfig:
                    description: Immutable. Optional. Lifecycle setting for the cluster.
                    properties:
                      autoDeleteTime:
                        description: Immutable. Optional. The time when cluster will
                          be auto-deleted (see JSON representation of [Timestamp](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                        format: date-time
                        type: string
                      autoDeleteTtl:
                        description: Immutable. Optional. The lifetime duration of
                          cluster. The cluster will be auto-deleted at the end of
                          this period. Minimum value is 10 minutes; maximum value
                          is 14 days (see JSON representation of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                        type: string
                      idleDeleteTtl:
                        description: Immutable. Optional. The duration to keep the
                          cluster alive while idling (when no jobs are running). Passing
                          this threshold will cause the cluster to be deleted. Minimum
                          value is 5 minutes; maximum value is 14 days (see JSON representation
                          of [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
                        type: string
                    type: object
required:
- spec
type: object

```
