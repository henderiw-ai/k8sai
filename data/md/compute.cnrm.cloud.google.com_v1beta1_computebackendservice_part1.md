# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeBackendService

## schema

```yaml
properties:
  spec:
    properties:
      affinityCookieTtlSec:
        description: |-
          Lifetime of cookies in seconds if session_affinity is
          GENERATED_COOKIE. If set to 0, the cookie is non-persistent and lasts
          only until the end of the browser session (or equivalent). The
          maximum allowed value for TTL is one day.

          When the load balancing scheme is INTERNAL, this field is not used.
        type: integer
      backend:
        description: The set of backends that serve this BackendService.
        items:
          properties:
            balancingMode:
              description: |-
                Specifies the balancing mode for this backend.

                For global HTTP(S) or TCP/SSL load balancing, the default is
                UTILIZATION. Valid values are UTILIZATION, RATE (for HTTP(S))
                and CONNECTION (for TCP/SSL).

                See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
                for an explanation of load balancing modes. Default value: "UTILIZATION" Possible values: ["UTILIZATION", "RATE", "CONNECTION"].
              type: string
            capacityScaler:
              description: |-
                A multiplier applied to the group's maximum servicing capacity
                (based on UTILIZATION, RATE or CONNECTION).

                Default value is 1, which means the group will serve up to 100%
                of its configured capacity (depending on balancingMode). A
                setting of 0 means the group is completely drained, offering
                0% of its available Capacity. Valid range is [0.0,1.0].
              type: number
            description:
              description: |-
                An optional description of this resource.
                Provide this property when you create the resource.
              type: string
            failover:
              description: |-
                This field designates whether this is a failover backend. More
                than one failover backend can be configured for a given RegionBackendService.
              type: boolean
            group:
              description: |-
                Reference to a ComputeInstanceGroup or ComputeNetworkEndpointGroup
                resource. In case of instance group this defines the list of
                instances that serve traffic. Member virtual machine instances from
                each instance group must live in the same zone as the instance
                group itself. No two backends in a backend service are allowed to
                use same Instance Group resource.

                For Network Endpoint Groups this defines list of endpoints. All
                endpoints of Network Endpoint Group must be hosted on instances
                located in the same zone as the Network Endpoint Group.

                Backend services cannot mix Instance Group and Network Endpoint
                Group backends.

                When the 'load_balancing_scheme' is INTERNAL, only instance groups
                are supported.
              oneOf:
              - required:
                - instanceGroupRef
              - required:
                - networkEndpointGroupRef
              properties:
                instanceGroupRef:
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
                      description: 'Allowed value: The `selfLink` field of a `ComputeInstanceGroup`
                        resource.'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
                networkEndpointGroupRef:
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
                      description: 'Allowed value: The `selfLink` field of a `ComputeNetworkEndpointGroup`
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
            maxConnections:
              description: |-
                The max number of simultaneous connections for the group. Can
                be used with either CONNECTION or UTILIZATION balancing modes.

                For CONNECTION mode, either maxConnections or one
                of maxConnectionsPerInstance or maxConnectionsPerEndpoint,
                as appropriate for group type, must be set.
              type: integer
            maxConnectionsPerEndpoint:
              description: |-
                The max number of simultaneous connections that a single backend
                network endpoint can handle. This is used to calculate the
                capacity of the group. Can be used in either CONNECTION or
                UTILIZATION balancing modes.

                For CONNECTION mode, either
                maxConnections or maxConnectionsPerEndpoint must be set.
              type: integer
            maxConnectionsPerInstance:
              description: |-
                The max number of simultaneous connections that a single
                backend instance can handle. This is used to calculate the
                capacity of the group. Can be used in either CONNECTION or
                UTILIZATION balancing modes.

                For CONNECTION mode, either maxConnections or
                maxConnectionsPerInstance must be set.
              type: integer
            maxRate:
              description: |-
                The max requests per second (RPS) of the group.

                Can be used with either RATE or UTILIZATION balancing modes,
                but required if RATE mode. For RATE mode, either maxRate or one
                of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
                group type, must be set.
              type: integer
            maxRatePerEndpoint:
              description: |-
                The max requests per second (RPS) that a single backend network
                endpoint can handle. This is used to calculate the capacity of
                the group. Can be used in either balancing mode. For RATE mode,
                either maxRate or maxRatePerEndpoint must be set.
              type: number
            maxRatePerInstance:
              description: |-
                The max requests per second (RPS) that a single backend
                instance can handle. This is used to calculate the capacity of
                the group. Can be used in either balancing mode. For RATE mode,
                either maxRate or maxRatePerInstance must be set.
              type: number
            maxUtilization:
              description: |-
                Used when balancingMode is UTILIZATION. This ratio defines the
                CPU utilization target for the group. Valid range is [0.0, 1.0].
              type: number
          required:
          - group
          type: object
        type: array
      cdnPolicy:
        description: Cloud CDN configuration for this BackendService.
        properties:
          bypassCacheOnRequestHeaders:
            description: |-
              Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified.
              The cache is bypassed for all cdnPolicy.cacheMode settings.
            items:
              properties:
                headerName:
                  description: The header field name to match on when bypassing cache.
                    Values are case-insensitive.
                  type: string
              required:
              - headerName
              type: object
            type: array
          cacheKeyPolicy:
            description: The CacheKeyPolicy for this CdnPolicy.
            properties:
              includeHost:
                description: If true requests to different hosts will be cached separately.
                type: boolean
              includeHttpHeaders:
                description: |-
                  Allows HTTP request headers (by name) to be used in the
                  cache key.
                items:
                  type: string
                type: array
              includeNamedCookies:
                description: Names of cookies to include in cache keys.
                items:
                  type: string
                type: array
              includeProtocol:
                description: If true, http and https requests will be cached separately.
                type: boolean
              includeQueryString:
                description: |-
                  If true, include query string parameters in the cache key
                  according to query_string_whitelist and
                  query_string_blacklist. If neither is set, the entire query
                  string will be included.

                  If false, the query string will be excluded from the cache
                  key entirely.
                type: boolean
              queryStringBlacklist:
                description: |-
                  Names of query string parameters to exclude in cache keys.

                  All other parameters will be included. Either specify
                  query_string_whitelist or query_string_blacklist, not both.
                  '&' and '=' will be percent encoded and not treated as
                  delimiters.
                items:
                  type: string
                type: array
              queryStringWhitelist:
                description: |-
                  Names of query string parameters to include in cache keys.

                  All other parameters will be excluded. Either specify
                  query_string_whitelist or query_string_blacklist, not both.
                  '&' and '=' will be percent encoded and not treated as
                  delimiters.
                items:
                  type: string
                type: array
            type: object
          cacheMode:
            description: |-
              Specifies the cache setting for all responses from this backend.
              The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC Possible values: ["USE_ORIGIN_HEADERS", "FORCE_CACHE_ALL", "CACHE_ALL_STATIC"].
            type: string
          clientTtl:
            description: Specifies the maximum allowed TTL for cached content served
              by this origin.
            type: integer
          defaultTtl:
            description: |-
              Specifies the default TTL for cached content served by this origin for responses
              that do not have an existing valid TTL (max-age or s-max-age).
            type: integer
          maxTtl:
            description: Specifies the maximum allowed TTL for cached content served
              by this origin.
            type: integer
          negativeCaching:
            description: Negative caching allows per-status code TTLs to be set, in
              order to apply fine-grained caching for common errors or redirects.
            type: boolean
          negativeCachingPolicy:
            description: |-
              Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
              Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
            items:
              properties:
                code:
                  description: |-
                    The HTTP status code to define a TTL against. Only HTTP status codes 300, 301, 308, 404, 405, 410, 421, 451 and 501
                    can be specified as values, and you cannot specify a status code more than once.
                  type: integer
                ttl:
                  description: |-
                    The TTL (in seconds) for which to cache responses with the corresponding status code. The maximum allowed value is 1800s
                    (30 minutes), noting that infrequently accessed objects may be evicted from the cache before the defined TTL.
                  type: integer
              type: object
            type: array
          serveWhileStale:
            description: Serve existing content from the cache (if available) when
              revalidating content with the origin, or when an error is encountered
              when refreshing the cache.
            type: integer
          signedUrlCacheMaxAgeSec:
            description: |-
              Maximum number of seconds the response to a signed URL request
              will be considered fresh, defaults to 1hr (3600s). After this
              time period, the response will be revalidated before
              being served.

              When serving responses to signed URL requests, Cloud CDN will
              internally behave as though all responses from this backend had a
              "Cache-Control: public, max-age=[TTL]" header, regardless of any
              existing Cache-Control header. The actual headers served in
              responses will not be altered.
            type: integer
        type: object
      circuitBreakers:
        description: |-
          Settings controlling the volume of connections to a backend service. This field
          is applicable only when the load_balancing_scheme is set to INTERNAL_SELF_MANAGED.
        properties:
          connectTimeout:
            description: The timeout for new network connections to hosts.
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
          maxConnections:
            description: |-
              The maximum number of connections to the backend cluster.
              Defaults to 1024.
            type: integer
          maxPendingRequests:
            description: |-
              The maximum number of pending requests to the backend cluster.
              Defaults to 1024.
            type: integer
          maxRequests:
            description: |-
              The maximum number of parallel requests to the backend cluster.
              Defaults to 1024.
            type: integer
          maxRequestsPerConnection:
            description: |-
              Maximum requests for a single backend connection. This parameter
              is respected by both the HTTP/1.1 and HTTP/2 implementations. If
              not specified, there is no limit. Setting this parameter to 1
              will effectively disable keep alive.
            type: integer
          maxRetries:
            description: |-
              The maximum number of parallel retries to the backend cluster.
              Defaults to 3.
            type: integer
        type: object
      compressionMode:
        description: 'Compress text responses using Brotli or gzip compression, based
          on the client''s Accept-Encoding header. Possible values: ["AUTOMATIC",
          "DISABLED"].'
        type: string
      connectionDrainingTimeoutSec:
        description: |-
          Time for which instance will be drained (not accept new
          connections, but still work to finish started).
        type: integer
      connectionTrackingPolicy:
        description: |-
          Connection Tracking configuration for this BackendService.
          This is available only for Layer 4 Internal Load Balancing and
          Network Load Balancing.
        properties:
          connectionPersistenceOnUnhealthyBackends:
            description: |-
              Specifies connection persistence when backends are unhealthy.

              If set to 'DEFAULT_FOR_PROTOCOL', the existing connections persist on
              unhealthy backends only for connection-oriented protocols (TCP and SCTP)
              and only if the Tracking Mode is PER_CONNECTION (default tracking mode)
              or the Session Affinity is configured for 5-tuple. They do not persist
              for UDP.

              If set to 'NEVER_PERSIST', after a backend becomes unhealthy, the existing
              connections on the unhealthy backend are never persisted on the unhealthy
              backend. They are always diverted to newly selected healthy backends
              (unless all backends are unhealthy).

              If set to 'ALWAYS_PERSIST', existing connections always persist on
              unhealthy backends regardless of protocol and session affinity. It is
              generally not recommended to use this mode overriding the default. Default value: "DEFAULT_FOR_PROTOCOL" Possible values: ["DEFAULT_FOR_PROTOCOL", "NEVER_PERSIST", "ALWAYS_PERSIST"].
            type: string
          idleTimeoutSec:
            description: |-
              Specifies how long to keep a Connection Tracking entry while there is
              no matching traffic (in seconds).

              For L4 ILB the minimum(default) is 10 minutes and maximum is 16 hours.

              For NLB the minimum(default) is 60 seconds and the maximum is 16 hours.
            type: integer
          trackingMode:
            description: |-
              Specifies the key used for connection tracking. There are two options:
              'PER_CONNECTION': The Connection Tracking is performed as per the
              Connection Key (default Hash Method) for the specific protocol.

              'PER_SESSION': The Connection Tracking is performed as per the
              configured Session Affinity. It matches the configured Session Affinity. Default value: "PER_CONNECTION" Possible values: ["PER_CONNECTION", "PER_SESSION"].
            type: string
        type: object
      spec:
        required:
        - location
        type: object
required:
- spec
type: object

```
