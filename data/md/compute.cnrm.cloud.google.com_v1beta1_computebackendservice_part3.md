# openAPI schema part3 for compute.cnrm.cloud.google.com.v1beta1.ComputeBackendService

## schema

```yaml
properties:
  spec:
    properties:
      localityLbPolicy:
        properties:
          localityLbPolicy:
            description: |-
              The load balancing algorithm used within the scope of the locality.
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
                          Maglev, refer to https://ai.google/research/pubs/pub44824

              * 'WEIGHTED_MAGLEV': Per-instance weighted Load Balancing via health check
                                   reported weights. If set, the Backend Service must
                                   configure a non legacy HTTP-based Health Check, and
                                   health check replies are expected to contain
                                   non-standard HTTP response header field
                                   X-Load-Balancing-Endpoint-Weight to specify the
                                   per-instance weights. If set, Load Balancing is weight
                                   based on the per-instance weights reported in the last
                                   processed health check replies, as long as every
                                   instance either reported a valid weight or had
                                   UNAVAILABLE_WEIGHT. Otherwise, Load Balancing remains
                                   equal-weight.


              This field is applicable to either:

              * A regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2,
                and loadBalancingScheme set to INTERNAL_MANAGED.
              * A global backend service with the load_balancing_scheme set to INTERNAL_SELF_MANAGED.
              * A regional backend service with loadBalancingScheme set to EXTERNAL (External Network
                Load Balancing). Only MAGLEV and WEIGHTED_MAGLEV values are possible for External
                Network Load Balancing. The default is MAGLEV.


              If session_affinity is not NONE, and this field is not set to MAGLEV, WEIGHTED_MAGLEV,
              or RING_HASH, session affinity settings will not take effect.

              Only ROUND_ROBIN and RING_HASH are supported when the backend service is referenced
              by a URL map that is bound to target gRPC proxy that has validate_for_proxyless
              field set to true. Possible values: ["ROUND_ROBIN", "LEAST_REQUEST", "RING_HASH", "RANDOM", "ORIGINAL_DESTINATION", "MAGLEV", "WEIGHTED_MAGLEV"].
            type: string
      location:
        description: 'Location represents the geographical location of the ComputeBackendService.
          Specify a region name or "global" for global resources. Reference: GCP definition
          of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)'
        type: string
      logConfig:
        description: |-
          This field denotes the logging options for the load balancer traffic served by this backend service.
          If logging is enabled, logs will be exported to Stackdriver.
        properties:
          enable:
            description: Whether to enable logging for the load balancer traffic served
              by this backend service.
            type: boolean
          sampleRate:
            description: |-
              This field can only be specified if logging is enabled for this backend service. The value of
              the field must be in [0, 1]. This configures the sampling rate of requests to the load balancer
              where 1.0 means all logged requests are reported and 0.0 means no logged requests are reported.
              The default value is 1.0.
            type: number
        type: object
      networkRef:
        description: |-
          The network to which this backend service belongs.  This field can
          only be specified when the load balancing scheme is set to
          INTERNAL.
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
      outlierDetection:
        description: |-
          Settings controlling eviction of unhealthy hosts from the load balancing pool.
          This field is applicable only when the load_balancing_scheme is set
          to INTERNAL_SELF_MANAGED.
        properties:
          baseEjectionTime:
            description: |-
              The base time that a host is ejected for. The real time is equal to the base
              time multiplied by the number of times the host has been ejected. Defaults to
              30000ms or 30s.
            properties:
              nanos:
                description: |-
                  Span of time that's a fraction of a second at nanosecond resolution. Durations
                  less than one second are represented with a 0 'seconds' field and a positive
                  'nanos' field. Must be from 0 to 999,999,999 inclusive.
                type: integer
              seconds:
                description: |-
                  Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
                  inclusive.
                type: integer
            required:
            - seconds
            type: object
          consecutiveErrors:
            description: |-
              Number of errors before a host is ejected from the connection pool. When the
              backend host is accessed over HTTP, a 5xx return code qualifies as an error.
              Defaults to 5.
            type: integer
          consecutiveGatewayFailure:
            description: |-
              The number of consecutive gateway failures (502, 503, 504 status or connection
              errors that are mapped to one of those status codes) before a consecutive
              gateway failure ejection occurs. Defaults to 5.
            type: integer
          enforcingConsecutiveErrors:
            description: |-
              The percentage chance that a host will be actually ejected when an outlier
              status is detected through consecutive 5xx. This setting can be used to disable
              ejection or to ramp it up slowly. Defaults to 100.
            type: integer
          enforcingConsecutiveGatewayFailure:
            description: |-
              The percentage chance that a host will be actually ejected when an outlier
              status is detected through consecutive gateway failures. This setting can be
              used to disable ejection or to ramp it up slowly. Defaults to 0.
            type: integer
          enforcingSuccessRate:
            description: |-
              The percentage chance that a host will be actually ejected when an outlier
              status is detected through success rate statistics. This setting can be used to
              disable ejection or to ramp it up slowly. Defaults to 100.
            type: integer
          interval:
            description: |-
              Time interval between ejection sweep analysis. This can result in both new
              ejections as well as hosts being returned to service. Defaults to 10 seconds.
            properties:
              nanos:
                description: |-
                  Span of time that's a fraction of a second at nanosecond resolution. Durations
                  less than one second are represented with a 0 'seconds' field and a positive
                  'nanos' field. Must be from 0 to 999,999,999 inclusive.
                type: integer
              seconds:
                description: |-
                  Span of time at a resolution of a second. Must be from 0 to 315,576,000,000
                  inclusive.
                type: integer
            required:
            - seconds
            type: object
          maxEjectionPercent:
            description: |-
              Maximum percentage of hosts in the load balancing pool for the backend service
              that can be ejected. Defaults to 10%.
            type: integer
          successRateMinimumHosts:
            description: |-
              The number of hosts in a cluster that must have enough request volume to detect
              success rate outliers. If the number of hosts is less than this setting, outlier
              detection via success rate statistics is not performed for any host in the
              cluster. Defaults to 5.
            type: integer
          successRateRequestVolume:
            description: |-
              The minimum number of total requests that must be collected in one interval (as
              defined by the interval duration above) to include this host in success rate
              based outlier detection. If the volume is lower than this setting, outlier
              detection via success rate statistics is not performed for that host. Defaults
              to 100.
            type: integer
          successRateStdevFactor:
            description: |-
              This factor is used to determine the ejection threshold for success rate outlier
              ejection. The ejection threshold is the difference between the mean success
              rate, and the product of this factor and the standard deviation of the mean
              success rate: mean - (stdev * success_rate_stdev_factor). This factor is divided
              by a thousand to get a double. That is, if the desired factor is 1.9, the
              runtime value should be 1900. Defaults to 1900.
            type: integer
        type: object
      portName:
        description: |-
          Name of backend port. The same name should appear in the instance
          groups referenced by this service. Required when the load balancing
          scheme is EXTERNAL.
        type: string
      protocol:
        description: |-
          The protocol this BackendService uses to communicate with backends.
          The default is HTTP. **NOTE**: HTTP2 is only valid for beta HTTP/2 load balancer
          types and may result in errors if used with the GA API. Possible values: ["HTTP", "HTTPS", "HTTP2", "TCP", "SSL", "GRPC"].
        type: string
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      securityPolicyRef:
        description: The security policy associated with this backend service.
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
      securitySettings:
        description: |-
          The security settings that apply to this backend service. This field is applicable to either
          a regional backend service with the service_protocol set to HTTP, HTTPS, or HTTP2, and
          load_balancing_scheme set to INTERNAL_MANAGED; or a global backend service with the
          load_balancing_scheme set to INTERNAL_SELF_MANAGED.
        properties:
          clientTLSPolicyRef:
            description: |-
              ClientTlsPolicy is a resource that specifies how a client should
              authenticate connections to backends of a service. This resource itself
              does not affect configuration unless it is attached to a backend
              service resource.
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
                description: 'Allowed value: The `name` field of a `NetworkSecurityClientTLSPolicy`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          subjectAltNames:
            description: |-
              A list of alternate names to verify the subject identity in the certificate.
              If specified, the client will verify that the server certificate's subject
              alt name matches one of the specified values.
            items:
              type: string
            type: array
        required:
        - clientTLSPolicyRef
        - subjectAltNames
        type: object
      sessionAffinity:
        description: |-
          Type of session affinity to use. The default is NONE. Session affinity is
          not applicable if the protocol is UDP. Possible values: ["NONE", "CLIENT_IP", "CLIENT_IP_PORT_PROTO", "CLIENT_IP_PROTO", "GENERATED_COOKIE", "HEADER_FIELD", "HTTP_COOKIE"].
        type: string
      subsetting:
        description: Subsetting configuration for this BackendService. Currently this
          is applicable only for Internal TCP/UDP load balancing and Internal HTTP(S)
          load balancing.
        properties:
          policy:
            description: 'The algorithm used for subsetting. Possible values: ["CONSISTENT_HASH_SUBSETTING"].'
            type: string
        required:
        - policy
        type: object
      timeoutSec:
        description: |-
          How many seconds to wait for the backend before considering it a
          failed request. Default is 30 seconds. Valid range is [1, 86400].
        type: integer
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
      creationTimestamp:
        description: Creation timestamp in RFC3339 text format.
        type: string
      fingerprint:
        description: |-
          Fingerprint of this resource. A hash of the contents stored in this
          object. This field is used in optimistic locking.
        type: string
      generatedId:
        description: The unique identifier for the resource. This identifier is defined
          by the server.
        type: integer
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      selfLink:
        type: string
    type: object
required:
- spec
type: object

```
