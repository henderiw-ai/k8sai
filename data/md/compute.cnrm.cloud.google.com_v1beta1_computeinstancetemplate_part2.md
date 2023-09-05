# openAPI schema part2 for compute.cnrm.cloud.google.com.v1beta1.ComputeInstanceTemplate

## schema

```yaml
properties:
  spec:
    properties:
      networkInterface:
        properties:
          networkInterface:
            description: Immutable. Networks to attach to instances created from this
              template. This can be specified multiple times for multiple networks.
            items:
              properties:
                accessConfig:
                  items:
                    properties:
                      natIpRef:
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
                            description: 'Allowed value: The `address` field of a
                              `ComputeAddress` resource.'
                            type: string
                          name:
                            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                            type: string
                          namespace:
                            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                            type: string
                        type: object
                      networkTier:
                        description: 'Immutable. The networking tier used for configuring
                          this instance template. This field can take the following
                          values: PREMIUM, STANDARD, FIXED_STANDARD. If this field
                          is not specified, it is assumed to be PREMIUM.'
                        type: string
                      publicPtrDomainName:
                        description: The DNS domain name for the public PTR record.The
                          DNS domain name for the public PTR record.
                        type: string
                    type: object
                  type: array
                aliasIpRange:
                  description: Immutable. An array of alias IP ranges for this network
                    interface. Can only be specified for network interfaces on subnet-mode
                    networks.
                  items:
                    properties:
                      ipCidrRange:
                        description: Immutable. The IP CIDR range represented by this
                          alias IP range. This IP CIDR range must belong to the specified
                          subnetwork and cannot contain IP addresses reserved by system
                          or used by other network interfaces. At the time of writing
                          only a netmask (e.g. /24) may be supplied, with a CIDR format
                          resulting in an API error.
                        type: string
                      subnetworkRangeName:
                        description: Immutable. The subnetwork secondary range name
                          specifying the secondary range from which to allocate the
                          IP CIDR range for this alias IP range. If left unspecified,
                          the primary range of the subnetwork will be used.
                        type: string
                    required:
                    - ipCidrRange
                    type: object
                  type: array
                ipv6AccessConfig:
                  description: An array of IPv6 access configurations for this interface.
                    Currently, only one IPv6 access config, DIRECT_IPV6, is supported.
                    If there is no ipv6AccessConfig specified, then this instance
                    will have no external IPv6 Internet access.
                  items:
                    properties:
                      externalIpv6:
                        description: The first IPv6 address of the external IPv6 range
                          associated with this instance, prefix length is stored in
                          externalIpv6PrefixLength in ipv6AccessConfig. The field
                          is output only, an IPv6 address from a subnetwork associated
                          with the instance will be allocated dynamically.
                        type: string
                      externalIpv6PrefixLength:
                        description: The prefix length of the external IPv6 range.
                        type: string
                      networkTier:
                        description: The service-level to be provided for IPv6 traffic
                          when the subnet has an external subnet. Only PREMIUM tier
                          is valid for IPv6.
                        type: string
                      publicPtrDomainName:
                        description: The domain name to be used when creating DNSv6
                          records for the external IPv6 ranges.
                        type: string
                    required:
                    - networkTier
                    type: object
                  type: array
                ipv6AccessType:
                  description: One of EXTERNAL, INTERNAL to indicate whether the IP
                    can be accessed from the Internet. This field is always inherited
                    from its subnetwork.
                  type: string
                name:
                  description: The name of the network_interface.
                  type: string
                networkAttachment:
                  description: 'Immutable. The URL of the network attachment that
                    this interface should connect to in the following format: projects/{projectNumber}/regions/{region_name}/networkAttachments/{network_attachment_name}.'
                  type: string
                networkIp:
                  description: Immutable. The private IP address to assign to the
                    instance. If empty, the address will be automatically assigned.
                  type: string
                networkRef:
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
                      description: 'Allowed value: The `selfLink` field of a `ComputeNetwork`
                        resource.'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
                nicType:
                  description: Immutable. The type of vNIC to be used on this interface.
                    Possible values:GVNIC, VIRTIO_NET.
                  type: string
                queueCount:
                  description: Immutable. The networking queue count that's specified
                    by users for the network interface. Both Rx and Tx queues will
                    be set to this number. It will be empty if not specified.
                  type: integer
                stackType:
                  description: The stack type for this network interface to identify
                    whether the IPv6 feature is enabled or not. If not specified,
                    IPV4_ONLY will be used.
                  type: string
                subnetworkProject:
                  description: Immutable. The ID of the project in which the subnetwork
                    belongs. If it is not provided, the provider project is used.
                  type: string
                subnetworkRef:
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
                      description: 'Allowed value: The `selfLink` field of a `ComputeSubnetwork`
                        resource.'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
              type: object
            type: array
      networkPerformanceConfig:
        description: Immutable. Configures network performance settings for the instance.
          If not specified, the instance will be created with its default network
          performance configuration.
        properties:
          totalEgressBandwidthTier:
            description: Immutable. The egress bandwidth tier to enable. Possible
              values:TIER_1, DEFAULT.
            type: string
        required:
        - totalEgressBandwidthTier
        type: object
      region:
        description: Immutable. An instance template is a global resource that is
          not bound to a zone or a region. However, you can still specify some regional
          resources in an instance template, which restricts the template to the region
          where that resource resides. For example, a custom subnetwork resource is
          tied to a specific region. Defaults to the region of the Provider if no
          value is given.
        type: string
      reservationAffinity:
        description: Immutable. Specifies the reservations that this instance can
          consume from.
        properties:
          specificReservation:
            description: Immutable. Specifies the label selector for the reservation
              to use.
            properties:
              key:
                description: Immutable. Corresponds to the label key of a reservation
                  resource. To target a SPECIFIC_RESERVATION by name, specify compute.googleapis.com/reservation-name
                  as the key and specify the name of your reservation as the only
                  value.
                type: string
              values:
                description: Immutable. Corresponds to the label values of a reservation
                  resource.
                items:
                  type: string
                type: array
            required:
            - key
            - values
            type: object
          type:
            description: Immutable. The type of reservation from which this instance
              can consume resources.
            type: string
        required:
        - type
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      resourcePolicies:
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
              description: 'Allowed value: The `selfLink` field of a `ComputeResourcePolicy`
                resource.'
              type: string
            name:
              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
              type: string
            namespace:
              description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
              type: string
          type: object
        type: array
      scheduling:
        description: Immutable. The scheduling strategy to use.
        properties:
          automaticRestart:
            description: Immutable. Specifies whether the instance should be automatically
              restarted if it is terminated by Compute Engine (not terminated by a
              user). This defaults to true.
            type: boolean
          instanceTerminationAction:
            description: Immutable. Specifies the action GCE should take when SPOT
              VM is preempted.
            type: string
          maintenanceInterval:
            description: 'Specifies the frequency of planned maintenance events. The
              accepted values are: PERIODIC.'
            type: string
          maxRunDuration:
            description: Immutable. The timeout for new network connections to hosts.
            properties:
              nanos:
                description: |-
                  Immutable. Span of time that's a fraction of a second at nanosecond
                  resolution. Durations less than one second are represented
                  with a 0 seconds field and a positive nanos field. Must
                  be from 0 to 999,999,999 inclusive.
                type: integer
              seconds:
                description: |-
                  Immutable. Span of time at a resolution of a second.
                  Must be from 0 to 315,576,000,000 inclusive.
                type: integer
            required:
            - seconds
            type: object
          minNodeCpus:
            description: Minimum number of cpus for the instance.
            type: integer
          nodeAffinities:
            items:
              properties:
                value:
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
              type: object
            type: array
          onHostMaintenance:
            description: Immutable. Defines the maintenance behavior for this instance.
            type: string
          preemptible:
            description: Immutable. Allows instance to be preempted. This defaults
              to false.
            type: boolean
          provisioningModel:
            description: Immutable. Whether the instance is spot. If this is set as
              SPOT.
            type: string
        type: object
      serviceAccount:
        description: Immutable. Service account to attach to the instance.
        properties:
          scopes:
            description: Immutable. A list of service scopes. Both OAuth2 URLs and
              gcloud short names are supported. To allow full access to all Cloud
              APIs, use the cloud-platform scope.
            items:
              type: string
            type: array
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
        required:
        - scopes
        type: object
      shieldedInstanceConfig:
        description: 'Immutable. Enable Shielded VM on this instance. Shielded VM
          provides verifiable integrity to prevent against malware and rootkits. Defaults
          to disabled. Note: shielded_instance_config can only be used with boot images
          with shielded vm support.'
        properties:
          enableIntegrityMonitoring:
            description: Immutable. Compare the most recent boot measurements to the
              integrity policy baseline and return a pair of pass/fail results depending
              on whether they match or not. Defaults to true.
            type: boolean
          enableSecureBoot:
            description: Immutable. Verify the digital signature of all boot components,
              and halt the boot process if signature verification fails. Defaults
              to false.
            type: boolean
          enableVtpm:
            description: Immutable. Use a virtualized trusted platform module, which
              is a specialized computer chip you can use to encrypt objects like keys
              and certificates. Defaults to true.
            type: boolean
        type: object
      tags:
        description: Immutable. Tags to attach to the instance.
        items:
          type: string
        type: array
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
      metadataFingerprint:
        description: The unique fingerprint of the metadata.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      selfLink:
        description: The URI of the created resource.
        type: string
      selfLinkUnique:
        description: A special URI of the created resource that uniquely identifies
          this instance template.
        type: string
      tagsFingerprint:
        description: The unique fingerprint of the tags.
        type: string
    type: object
required:
- spec
type: object

```
