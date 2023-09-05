# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeForwardingRule

## schema

```yaml
properties:
  spec:
    properties:
      allPorts:
        description: |-
          Immutable. This field can only be used:
          * If 'IPProtocol' is one of TCP, UDP, or SCTP.
          * By internal TCP/UDP load balancers, backend service-based network load
          balancers, and internal and external protocol forwarding.


          Set this field to true to allow packets addressed to any port or packets
          lacking destination port information (for example, UDP fragments after the
          first fragment) to be forwarded to the backends configured with this
          forwarding rule.

          The 'ports', 'port_range', and
          'allPorts' fields are mutually exclusive.
        type: boolean
      allowGlobalAccess:
        description: |-
          This field is used along with the 'backend_service' field for
          internal load balancing or with the 'target' field for internal
          TargetInstance.

          If the field is set to 'TRUE', clients can access ILB from all
          regions.

          Otherwise only allows access from clients in the same region as the
          internal load balancer.
        type: boolean
      allowPscGlobalAccess:
        description: Immutable. This is used in PSC consumer ForwardingRule to control
          whether the PSC endpoint can be accessed from another region.
        type: boolean
      backendServiceRef:
        description: |-
          A ComputeBackendService to receive the matched traffic. This is
          used only for internal load balancing.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeBackendService`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      description:
        description: |-
          Immutable. An optional description of this resource. Provide this property when
          you create the resource.
        type: string
      ipAddress:
        description: |-
          The IP address that this forwarding rule is serving on behalf of.

          Addresses are restricted based on the forwarding rule's load
          balancing scheme (EXTERNAL or INTERNAL) and scope (global or
          regional).

          When the load balancing scheme is EXTERNAL, for global forwarding
          rules, the address must be a global IP, and for regional forwarding
          rules, the address must live in the same region as the forwarding
          rule. If this field is empty, an ephemeral IPv4 address from the
          same scope (global or regional) will be assigned. A regional
          forwarding rule supports IPv4 only. A global forwarding rule
          supports either IPv4 or IPv6.

          When the load balancing scheme is INTERNAL, this can only be an RFC
          1918 IP address belonging to the network/subnet configured for the
          forwarding rule. By default, if this field is empty, an ephemeral
          internal IP address will be automatically allocated from the IP
          range of the subnet or network configured for this forwarding rule.
        oneOf:
        - required:
          - addressRef
        - required:
          - ip
        properties:
          addressRef:
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
                description: 'Allowed value: The `address` field of a `ComputeAddress`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          ip:
            type: string
        type: object
      ipProtocol:
        description: |-
          Immutable. The IP protocol to which this rule applies.

          For protocol forwarding, valid
          options are 'TCP', 'UDP', 'ESP',
          'AH', 'SCTP', 'ICMP' and
          'L3_DEFAULT'.

          The valid IP protocols are different for different load balancing products
          as described in [Load balancing
          features](https://cloud.google.com/load-balancing/docs/features#protocols_from_the_load_balancer_to_the_backends). Possible values: ["TCP", "UDP", "ESP", "AH", "SCTP", "ICMP", "L3_DEFAULT"].
        type: string
      ipVersion:
        description: 'Immutable. The IP Version that will be used by this global forwarding
          rule. Possible values: ["IPV4", "IPV6"].'
        type: string
      isMirroringCollector:
        description: |-
          Immutable. Indicates whether or not this load balancer can be used as a collector for
          packet mirroring. To prevent mirroring loops, instances behind this
          load balancer will not have their traffic mirrored even if a
          'PacketMirroring' rule applies to them.

          This can only be set to true for load balancers that have their
          'loadBalancingScheme' set to 'INTERNAL'.
        type: boolean
      loadBalancingScheme:
        description: |-
          Immutable. Specifies the forwarding rule type.

          For more information about forwarding rules, refer to
          [Forwarding rule concepts](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "EXTERNAL_MANAGED", "INTERNAL", "INTERNAL_MANAGED"].
        type: string
      location:
        description: 'Location represents the geographical location of the ComputeForwardingRule.
          Specify a region name or "global" for global resources. Reference: GCP definition
          of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)'
        type: string
      metadataFilters:
        description: |-
          Immutable. Opaque filter criteria used by Loadbalancer to restrict routing
          configuration to a limited set xDS compliant clients. In their xDS
          requests to Loadbalancer, xDS clients present node metadata. If a
          match takes place, the relevant routing configuration is made available
          to those proxies.

          For each metadataFilter in this list, if its filterMatchCriteria is set
          to MATCH_ANY, at least one of the filterLabels must match the
          corresponding label provided in the metadata. If its filterMatchCriteria
          is set to MATCH_ALL, then all of its filterLabels must match with
          corresponding labels in the provided metadata.

          metadataFilters specified here can be overridden by those specified in
          the UrlMap that this ForwardingRule references.

          metadataFilters only applies to Loadbalancers that have their
          loadBalancingScheme set to INTERNAL_SELF_MANAGED.
        items:
          properties:
            filterLabels:
              description: |-
                Immutable. The list of label value pairs that must match labels in the
                provided metadata based on filterMatchCriteria

                This list must not be empty and can have at the most 64 entries.
              items:
                properties:
                  name:
                    description: |-
                      Immutable. Name of the metadata label. The length must be between
                      1 and 1024 characters, inclusive.
                    type: string
                  value:
                    description: |-
                      Immutable. The value that the label must match. The value has a maximum
                      length of 1024 characters.
                    type: string
                required:
                - name
                - value
                type: object
              type: array
            filterMatchCriteria:
              description: |-
                Immutable. Specifies how individual filterLabel matches within the list of
                filterLabels contribute towards the overall metadataFilter match.

                MATCH_ANY - At least one of the filterLabels must have a matching
                label in the provided metadata.
                MATCH_ALL - All filterLabels must have matching labels in the
                provided metadata. Possible values: ["MATCH_ANY", "MATCH_ALL"].
              type: string
          required:
          - filterLabels
          - filterMatchCriteria
          type: object
        type: array
      networkRef:
        description: |-
          This field is not used for external load balancing. For internal
          load balancing, this field identifies the network that the load
          balanced IP should belong to for this forwarding rule. If this
          field is not specified, the default network will be used.
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
      networkTier:
        description: |-
          Immutable. This signifies the networking tier used for configuring
          this load balancer and can only take the following values:
          'PREMIUM', 'STANDARD'.

          For regional ForwardingRule, the valid values are 'PREMIUM' and
          'STANDARD'. For GlobalForwardingRule, the valid value is
          'PREMIUM'.

          If this field is not specified, it is assumed to be 'PREMIUM'.
          If 'IPAddress' is specified, this value must be equal to the
          networkTier of the Address. Possible values: ["PREMIUM", "STANDARD"].
        type: string
      noAutomateDnsZone:
        description: Immutable. This is used in PSC consumer ForwardingRule to control
          whether it should try to auto-generate a DNS zone or not. Non-PSC forwarding
          rules do not use this field.
        type: boolean
      portRange:
        description: |-
          Immutable. This field can only be used:

          * If 'IPProtocol' is one of TCP, UDP, or SCTP.
          * By backend service-based network load balancers, target pool-based
          network load balancers, internal proxy load balancers, external proxy load
          balancers, Traffic Director, external protocol forwarding, and Classic VPN.
          Some products have restrictions on what ports can be used. See
          [port specifications](https://cloud.google.com/load-balancing/docs/forwarding-rule-concepts#port_specifications)
          for details.


          Only packets addressed to ports in the specified range will be forwarded to
          the backends configured with this forwarding rule.

          The 'ports' and 'port_range' fields are mutually exclusive.

          For external forwarding rules, two or more forwarding rules cannot use the
          same '[IPAddress, IPProtocol]' pair, and cannot have
          overlapping 'portRange's.

          For internal forwarding rules within the same VPC network, two or more
          forwarding rules cannot use the same '[IPAddress, IPProtocol]'
          pair, and cannot have overlapping 'portRange's.
        type: string
      ports:
        description: |-
          Immutable. This field can only be used:

          * If 'IPProtocol' is one of TCP, UDP, or SCTP.
          * By internal TCP/UDP load balancers, backend service-based network load
          balancers, and internal protocol forwarding.


          You can specify a list of up to five ports by number, separated by commas.
          The ports can be contiguous or discontiguous. Only packets addressed to
          these ports will be forwarded to the backends configured with this
          forwarding rule.

          For external forwarding rules, two or more forwarding rules cannot use the
          same '[IPAddress, IPProtocol]' pair, and cannot share any values
          defined in 'ports'.

          For internal forwarding rules within the same VPC network, two or more
          forwarding rules cannot use the same '[IPAddress, IPProtocol]'
          pair, and cannot share any values defined in 'ports'.

          The 'ports' and 'port_range' fields are mutually exclusive.
        items:
          type: string
        type: array
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      serviceDirectoryRegistrations:
        description: |-
          Immutable. Service Directory resources to register this forwarding rule with.

          Currently, only supports a single Service Directory resource.
        items:
          properties:
            namespace:
              description: Immutable. Service Directory namespace to register the
                forwarding rule under.
              type: string
            service:
              description: Immutable. Service Directory service to register the forwarding
                rule under.
              type: string
          type: object
        type: array
      serviceLabel:
        description: |-
          Immutable. An optional prefix to the service name for this Forwarding Rule.
          If specified, will be the first label of the fully qualified service
          name.

          The label must be 1-63 characters long, and comply with RFC1035.
          Specifically, the label must be 1-63 characters long and match the
          regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the first
          character must be a lowercase letter, and all following characters
          must be a dash, lowercase letter, or digit, except the last
          character, which cannot be a dash.

          This field is only used for INTERNAL load balancing.
        type: string
      sourceIpRanges:
        description: Immutable. If not empty, this Forwarding Rule will only forward
          the traffic when the source IP address matches one of the IP addresses or
          CIDR ranges set here. Note that a Forwarding Rule can only have up to 64
          source IP ranges, and this field can only be used with a regional Forwarding
          Rule whose scheme is EXTERNAL. Each sourceIpRange entry should be either
          an IP address (for example, 1.2.3.4) or a CIDR range (for example, 1.2.3.0/24).
        items:
          type: string
        type: array
      spec:
        required:
        - location
        type: object
      subnetworkRef:
        description: |-
          The subnetwork that the load balanced IP should belong to for this
          forwarding rule. This field is only used for internal load
          balancing.

          If the network specified is in auto subnet mode, this field is
          optional. However, if the network is in custom subnet mode, a
          subnetwork must be specified.
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
            description: 'Allowed value: The `name` field of a `ComputeSubnetwork`
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
- spec
type: object

```
