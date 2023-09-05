# openAPI schema part2 for compute.cnrm.cloud.google.com.v1beta1.ComputeSecurityPolicy

## schema

```yaml
properties:
  spec:
    properties:
      rule:
        properties:
          rule:
            description: The set of rules that belong to this policy. There must always
              be a default rule (rule with priority 2147483647 and match "*"). If
              no rules are provided when creating a security policy, a default rule
              with action "allow" will be added.
            items:
              properties:
                action:
                  description: Action to take when match matches the request.
                  type: string
                description:
                  description: An optional description of this rule. Max size is 64.
                  type: string
                headerAction:
                  description: Additional actions that are performed on headers.
                  properties:
                    requestHeadersToAdds:
                      description: The list of request headers to add or overwrite
                        if they're already present.
                      items:
                        properties:
                          headerName:
                            description: The name of the header to set.
                            type: string
                          headerValue:
                            description: The value to set the named header to.
                            type: string
                        required:
                        - headerName
                        type: object
                      type: array
                  required:
                  - requestHeadersToAdds
                  type: object
                match:
                  description: A match condition that incoming traffic is evaluated
                    against. If it evaluates to true, the corresponding action is
                    enforced.
                  properties:
                    config:
                      description: The configuration options available when specifying
                        versioned_expr. This field must be specified if versioned_expr
                        is specified and cannot be specified if versioned_expr is
                        not specified.
                      properties:
                        srcIpRanges:
                          description: Set of IP addresses or ranges (IPV4 or IPV6)
                            in CIDR notation to match against inbound traffic. There
                            is a limit of 10 IP ranges per rule. A value of '*' matches
                            all IPs (can be used to override the default behavior).
                          items:
                            type: string
                          type: array
                      required:
                      - srcIpRanges
                      type: object
                    expr:
                      description: User defined CEVAL expression. A CEVAL expression
                        is used to specify match criteria such as origin.ip, source.region_code
                        and contents in the request header.
                      properties:
                        expression:
                          description: Textual representation of an expression in
                            Common Expression Language syntax. The application context
                            of the containing message determines which well-known
                            feature set of CEL is supported.
                          type: string
                      required:
                      - expression
                      type: object
                    versionedExpr:
                      description: 'Predefined rule expression. If this field is specified,
                        config must also be specified. Available options:   SRC_IPS_V1:
                        Must specify the corresponding src_ip_ranges field in config.'
                      type: string
                  type: object
                preconfiguredWafConfig:
                  description: Preconfigured WAF configuration to be applied for the
                    rule. If the rule does not evaluate preconfigured WAF rules, i.e.,
                    if evaluatePreconfiguredWaf() is not used, this field will have
                    no effect.
                  properties:
                    exclusion:
                      description: An exclusion to apply during preconfigured WAF
                        evaluation.
                      items:
                        properties:
                          requestCookie:
                            description: Request cookie whose value will be excluded
                              from inspection during preconfigured WAF evaluation.
                            items:
                              properties:
                                operator:
                                  description: 'You can specify an exact match or
                                    a partial match by using a field operator and
                                    a field value. Available options: EQUALS: The
                                    operator matches if the field value equals the
                                    specified value. STARTS_WITH: The operator matches
                                    if the field value starts with the specified value.
                                    ENDS_WITH: The operator matches if the field value
                                    ends with the specified value. CONTAINS: The operator
                                    matches if the field value contains the specified
                                    value. EQUALS_ANY: The operator matches if the
                                    field value is any value.'
                                  type: string
                                value:
                                  description: A request field matching the specified
                                    value will be excluded from inspection during
                                    preconfigured WAF evaluation. The field value
                                    must be given if the field operator is not EQUALS_ANY,
                                    and cannot be given if the field operator is EQUALS_ANY.
                                  type: string
                              required:
                              - operator
                              type: object
                            type: array
                          requestHeader:
                            description: Request header whose value will be excluded
                              from inspection during preconfigured WAF evaluation.
                            items:
                              properties:
                                operator:
                                  description: 'You can specify an exact match or
                                    a partial match by using a field operator and
                                    a field value. Available options: EQUALS: The
                                    operator matches if the field value equals the
                                    specified value. STARTS_WITH: The operator matches
                                    if the field value starts with the specified value.
                                    ENDS_WITH: The operator matches if the field value
                                    ends with the specified value. CONTAINS: The operator
                                    matches if the field value contains the specified
                                    value. EQUALS_ANY: The operator matches if the
                                    field value is any value.'
                                  type: string
                                value:
                                  description: A request field matching the specified
                                    value will be excluded from inspection during
                                    preconfigured WAF evaluation. The field value
                                    must be given if the field operator is not EQUALS_ANY,
                                    and cannot be given if the field operator is EQUALS_ANY.
                                  type: string
                              required:
                              - operator
                              type: object
                            type: array
                          requestQueryParam:
                            description: Request query parameter whose value will
                              be excluded from inspection during preconfigured WAF
                              evaluation.  Note that the parameter can be in the query
                              string or in the POST body.
                            items:
                              properties:
                                operator:
                                  description: 'You can specify an exact match or
                                    a partial match by using a field operator and
                                    a field value. Available options: EQUALS: The
                                    operator matches if the field value equals the
                                    specified value. STARTS_WITH: The operator matches
                                    if the field value starts with the specified value.
                                    ENDS_WITH: The operator matches if the field value
                                    ends with the specified value. CONTAINS: The operator
                                    matches if the field value contains the specified
                                    value. EQUALS_ANY: The operator matches if the
                                    field value is any value.'
                                  type: string
                                value:
                                  description: A request field matching the specified
                                    value will be excluded from inspection during
                                    preconfigured WAF evaluation. The field value
                                    must be given if the field operator is not EQUALS_ANY,
                                    and cannot be given if the field operator is EQUALS_ANY.
                                  type: string
                              required:
                              - operator
                              type: object
                            type: array
                          requestUri:
                            description: Request URI from the request line to be excluded
                              from inspection during preconfigured WAF evaluation.
                              When specifying this field, the query or fragment part
                              should be excluded.
                            items:
                              properties:
                                operator:
                                  description: 'You can specify an exact match or
                                    a partial match by using a field operator and
                                    a field value. Available options: EQUALS: The
                                    operator matches if the field value equals the
                                    specified value. STARTS_WITH: The operator matches
                                    if the field value starts with the specified value.
                                    ENDS_WITH: The operator matches if the field value
                                    ends with the specified value. CONTAINS: The operator
                                    matches if the field value contains the specified
                                    value. EQUALS_ANY: The operator matches if the
                                    field value is any value.'
                                  type: string
                                value:
                                  description: A request field matching the specified
                                    value will be excluded from inspection during
                                    preconfigured WAF evaluation. The field value
                                    must be given if the field operator is not EQUALS_ANY,
                                    and cannot be given if the field operator is EQUALS_ANY.
                                  type: string
                              required:
                              - operator
                              type: object
                            type: array
                          targetRuleIds:
                            description: A list of target rule IDs under the WAF rule
                              set to apply the preconfigured WAF exclusion. If omitted,
                              it refers to all the rule IDs under the WAF rule set.
                            items:
                              type: string
                            type: array
                          targetRuleSet:
                            description: Target WAF rule set to apply the preconfigured
                              WAF exclusion.
                            type: string
                        required:
                        - targetRuleSet
                        type: object
                      type: array
                  type: object
                preview:
                  description: When set to true, the action specified above is not
                    enforced. Stackdriver logs for requests that trigger a preview
                    action are annotated as such.
                  type: boolean
                priority:
                  description: An unique positive integer indicating the priority
                    of evaluation for a rule. Rules are evaluated from highest priority
                    (lowest numerically) to lowest priority (highest numerically)
                    in order.
                  type: integer
                rateLimitOptions:
                  description: Rate limit threshold for this security policy. Must
                    be specified if the action is "rate_based_ban" or "throttle".
                    Cannot be specified for any other actions.
                  properties:
                    banDurationSec:
                      description: Can only be specified if the action for the rule
                        is "rate_based_ban". If specified, determines the time (in
                        seconds) the traffic will continue to be banned by the rate
                        limit after the rate falls below the threshold.
                      type: integer
                    banThreshold:
                      description: Can only be specified if the action for the rule
                        is "rate_based_ban". If specified, the key will be banned
                        for the configured 'banDurationSec' when the number of requests
                        that exceed the 'rateLimitThreshold' also exceed this 'banThreshold'.
                      properties:
                        count:
                          description: Number of HTTP(S) requests for calculating
                            the threshold.
                          type: integer
                        intervalSec:
                          description: Interval over which the threshold is computed.
                          type: integer
                      required:
                      - count
                      - intervalSec
                      type: object
                    conformAction:
                      description: Action to take for requests that are under the
                        configured rate limit threshold. Valid option is "allow" only.
                      type: string
                    enforceOnKey:
                      description: Determines the key to enforce the rateLimitThreshold
                        on.
                      type: string
                    enforceOnKeyConfigs:
                      description: Immutable. Enforce On Key Config of this security
                        policy.
                      items:
                        properties:
                          enforceOnKeyName:
                            description: 'Rate limit key name applicable only for
                              the following key types: HTTP_HEADER -- Name of the
                              HTTP header whose value is taken as the key value. HTTP_COOKIE
                              -- Name of the HTTP cookie whose value is taken as the
                              key value.'
                            type: string
                          enforceOnKeyType:
                            description: Determines the key to enforce the rate_limit_threshold
                              on.
                            type: string
                        type: object
                      type: array
                    enforceOnKeyName:
                      description: 'Rate limit key name applicable only for the following
                        key types: HTTP_HEADER -- Name of the HTTP header whose value
                        is taken as the key value. HTTP_COOKIE -- Name of the HTTP
                        cookie whose value is taken as the key value.'
                      type: string
                    exceedAction:
                      description: Action to take for requests that are above the
                        configured rate limit threshold, to either deny with a specified
                        HTTP response code, or redirect to a different endpoint. Valid
                        options are "deny()" where valid values for status are 403,
                        404, 429, and 502, and "redirect" where the redirect parameters
                        come from exceedRedirectOptions below.
                      type: string
                    exceedRedirectOptions:
                      description: Parameters defining the redirect action that is
                        used as the exceed action. Cannot be specified if the exceed
                        action is not redirect.
                      properties:
                        target:
                          description: Target for the redirect action. This is required
                            if the type is EXTERNAL_302 and cannot be specified for
                            GOOGLE_RECAPTCHA.
                          type: string
                        type:
                          description: Type of the redirect action.
                          type: string
                      required:
                      - type
                      type: object
                    rateLimitThreshold:
                      description: Threshold at which to begin ratelimiting.
                      properties:
                        count:
                          description: Number of HTTP(S) requests for calculating
                            the threshold.
                          type: integer
                        intervalSec:
                          description: Interval over which the threshold is computed.
                          type: integer
                      required:
                      - count
                      - intervalSec
                      type: object
                  required:
                  - conformAction
                  - exceedAction
                  - rateLimitThreshold
                  type: object
                redirectOptions:
                  description: Parameters defining the redirect action. Cannot be
                    specified for any other actions.
                  properties:
                    target:
                      description: Target for the redirect action. This is required
                        if the type is EXTERNAL_302 and cannot be specified for GOOGLE_RECAPTCHA.
                      type: string
                    type:
                      description: 'Type of the redirect action. Available options:
                        EXTERNAL_302: Must specify the corresponding target field
                        in config. GOOGLE_RECAPTCHA: Cannot specify target field in
                        config.'
                      type: string
                  required:
                  - type
                  type: object
              required:
              - action
              - match
              - priority
              type: object
            type: array
      type:
        description: The type indicates the intended use of the security policy. CLOUD_ARMOR
          - Cloud Armor backend security policies can be configured to filter incoming
          HTTP requests targeting backend services. They filter requests before they
          hit the origin servers. CLOUD_ARMOR_EDGE - Cloud Armor edge security policies
          can be configured to filter incoming HTTP requests targeting backend services
          (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage).
          They filter requests before the request is served from Google's cache.
        type: string
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
      fingerprint:
        description: Fingerprint of this resource.
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
    type: object
type: object

```