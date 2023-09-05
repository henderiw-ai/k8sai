# openAPI schema part2 for compute.cnrm.cloud.google.com.v1beta1.ComputeBackendService

## schema

```yaml
properties:
  spec:
    properties:
      consistentHash:
        properties:
          consistentHash:
            description: |-
              Consistent Hash-based load balancing can be used to provide soft session
              affinity based on HTTP headers, cookies or other properties. This load balancing
              policy is applicable only for HTTP connections. The affinity to a particular
              destination host will be lost when one or more hosts are added/removed from the
              destination service. This field specifies parameters that control consistent
              hashing. This field only applies if the load_balancing_scheme is set to
              INTERNAL_SELF_MANAGED. This field is only applicable when locality_lb_policy is
              set to MAGLEV or RING_HASH.
            properties:
              httpCookie:
                description: |-
                  Hash is based on HTTP Cookie. This field describes a HTTP cookie
                  that will be used as the hash key for the consistent hash load
                  balancer. If the cookie is not present, it will be generated.
                  This field is applicable if the sessionAffinity is set to HTTP_COOKIE.
                properties:
                  name:
                    description: Name of the cookie.
                    type: string
                  path:
                    description: Path to set for the cookie.
                    type: string
                  ttl:
                    description: Lifetime of the cookie.
                    properties:
                      nanos:
                        description: |-
                          Span of time that's a fraction of a second at nanosecond
                          resolution. Durations less than one second are represented
                          with a 0 seconds field and a positive nanos field. Must
                          be from 0 to 999,999,999 inclusive.
                        type: integer
                      seconds:
                        description: |-
                          Span of time at a resolution of a second.
                          Must be from 0 to 315,576,000,000 inclusive.
                        type: integer
                    required:
                    - seconds
                    type: object
                type: object
              httpHeaderName:
                description: |-
                  The hash based on the value of the specified header field.
                  This field is applicable if the sessionAffinity is set to HEADER_FIELD.
                type: string
              minimumRingSize:
                description: |-
                  The minimum number of virtual nodes to use for the hash ring.
                  Larger ring sizes result in more granular load
                  distributions. If the number of hosts in the load balancing pool
                  is larger than the ring size, each host will be assigned a single
                  virtual node.
                  Defaults to 1024.
                type: integer
            type: object
      customRequestHeaders:
        description: |-
          Headers that the HTTP/S load balancer should add to proxied
          requests.
        items:
          type: string
        type: array
      customResponseHeaders:
        description: |-
          Headers that the HTTP/S load balancer should add to proxied
          responses.
        items:
          type: string
        type: array
      description:
        description: An optional description of this resource.
        type: string
      edgeSecurityPolicyRef:
        description: |-
          The resource URL for the edge security policy associated with this
          backend service.
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
            description: 'Allowed value: The `selfLink` field of a `ComputeSecurityPolicy`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      enableCdn:
        description: If true, enable Cloud CDN for this BackendService.
        type: boolean
      failoverPolicy:
        description: Policy for failovers.
        properties:
          disableConnectionDrainOnFailover:
            description: |-
              On failover or failback, this field indicates whether connection drain
              will be honored. Setting this to true has the following effect: connections
              to the old active pool are not drained. Connections to the new active pool
              use the timeout of 10 min (currently fixed). Setting to false has the
              following effect: both old and new connections will have a drain timeout
              of 10 min.
              This can be set to true only if the protocol is TCP.
              The default is false.
            type: boolean
          dropTrafficIfUnhealthy:
            description: |-
              This option is used only when no healthy VMs are detected in the primary
              and backup instance groups. When set to true, traffic is dropped. When
              set to false, new connections are sent across all VMs in the primary group.
              The default is false.
            type: boolean
          failoverRatio:
            description: |-
              The value of the field must be in [0, 1]. If the ratio of the healthy
              VMs in the primary backend is at or below this number, traffic arriving
              at the load-balanced IP will be directed to the failover backend.
              In case where 'failoverRatio' is not set or all the VMs in the backup
              backend are unhealthy, the traffic will be directed back to the primary
              backend in the "force" mode, where traffic will be spread to the healthy
              VMs with the best effort, or to all VMs when no VM is healthy.
              This field is only used with l4 load balancing.
            type: number
        type: object
      healthChecks:
        items:
          description: |-
            The health check resources for health checking this
            ComputeBackendService. Currently at most one health check can be
            specified, and a health check is required.
          oneOf:
          - required:
            - healthCheckRef
          - required:
            - httpHealthCheckRef
          properties:
            healthCheckRef:
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
                  description: 'Allowed value: The `selfLink` field of a `ComputeHealthCheck`
                    resource.'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
            httpHealthCheckRef:
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
                  description: 'Allowed value: The `selfLink` field of a `ComputeHTTPHealthCheck`
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
      iap:
        description: Settings for enabling Cloud Identity Aware Proxy.
        oneOf:
        - required:
          - oauth2ClientId
        - required:
          - oauth2ClientIdRef
        properties:
          oauth2ClientId:
            description: DEPRECATED. Although this field is still available, there
              is limited support. We recommend that you use `spec.iap.oauth2ClientIdRef`
              instead.
            type: string
          oauth2ClientIdRef:
            description: OAuth2 Client ID for IAP.
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
                description: 'Allowed value: The `name` field of an `IAPIdentityAwareProxyClient`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          oauth2ClientSecret:
            description: OAuth2 Client Secret for IAP.
            oneOf:
            - not:
                required:
                - valueFrom
              required:
              - value
            - not:
                required:
                - value
              required:
              - valueFrom
            properties:
              value:
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - key
                    - name
                    type: object
                type: object
            type: object
          oauth2ClientSecretSha256:
            description: OAuth2 Client Secret SHA-256 for IAP.
            type: string
        type: object
      loadBalancingScheme:
        description: |-
          Immutable. Indicates whether the backend service will be used with internal or
          external load balancing. A backend service created for one type of
          load balancing cannot be used with the other. For more information, refer to
          [Choosing a load balancer](https://cloud.google.com/load-balancing/docs/backend-service). Default value: "EXTERNAL" Possible values: ["EXTERNAL", "INTERNAL_SELF_MANAGED", "EXTERNAL_MANAGED"].
        type: string
      localityLbPolicies:
        description: |-
          A list of locality load balancing policies to be used in order of
          preference. Either the policy or the customPolicy field should be set.
          Overrides any value set in the localityLbPolicy field.

          localityLbPolicies is only supported when the BackendService is referenced
          by a URL Map that is referenced by a target gRPC proxy that has the
          validateForProxyless field set to true.
        items:
          properties:
            customPolicy:
              description: |-
                The configuration for a custom policy implemented by the user and
                deployed with the client.
              properties:
                data:
                  description: |-
                    An optional, arbitrary JSON object with configuration data, understood
                    by a locally installed custom policy implementation.
                  type: string
                name:
                  description: |-
                    Identifies the custom policy.

                    The value should match the type the custom implementation is registered
                    with on the gRPC clients. It should follow protocol buffer
                    message naming conventions and include the full path (e.g.
                    myorg.CustomLbPolicy). The maximum length is 256 characters.

                    Note that specifying the same custom policy more than once for a
                    backend is not a valid configuration and will be rejected.
                  type: string
              required:
              - name
              type: object
            policy:
              description: The configuration for a built-in load balancing policy.
              properties:
                name:
                  description: |-
                    The name of a locality load balancer policy to be used. The value
                    should be one of the predefined ones as supported by localityLbPolicy,
                    although at the moment only ROUND_ROBIN is supported.

                    This field should only be populated when the customPolicy field is not
                    used.

                    Note that specifying the same policy more than once for a backend is
                    not a valid configuration and will be rejected.

                    The possible values are:

                    * 'ROUND_ROBIN': This is a simple policy in which each healthy backend
                                    is selected in round robin order.

                    * 'LEAST_REQUEST': An O(1) algorithm which selects two random healthy
                                      hosts and picks the host which has fewer active requests.

                    * 'RING_HASH': The ring/modulo hash load balancer implements consistent
                                  hashing to backends. The algorithm has the property that the
                                  addition/removal of a host from a set of N hosts only affects
                                  1/N of the requests.

                    * 'RANDOM': The load balancer selects a random healthy host.

                    * 'ORIGINAL_DESTINATION': Backend host is selected based on the client
                                              connection metadata, i.e., connections are opened
                                              to the same address as the destination address of
                                              the incoming connection before the connection
                                              was redirected to the load balancer.

                    * 'MAGLEV': used as a drop in replacement for the ring hash load balancer.
                                Maglev is not as stable as ring hash but has faster table lookup
                                build times and host selection times. For more information about
                                Maglev, refer to https://ai.google/research/pubs/pub44824 Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV"].
                  type: string
              required:
              - name
              type: object
          type: object
        type: array
required:
- spec
type: object

```
