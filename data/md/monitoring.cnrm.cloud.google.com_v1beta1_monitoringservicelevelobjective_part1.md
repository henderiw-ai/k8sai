# openAPI schema part1 for monitoring.cnrm.cloud.google.com.v1beta1.MonitoringServiceLevelObjective

## schema

```yaml
properties:
  spec:
    properties:
      calendarPeriod:
        description: 'A calendar period, semantically "since the start of the current
          ``". At this time, only `DAY`, `WEEK`, `FORTNIGHT`, and `MONTH` are supported.
          Possible values: CALENDAR_PERIOD_UNSPECIFIED, DAY, WEEK, FORTNIGHT, MONTH,
          QUARTER, HALF, YEAR'
        type: string
      displayName:
        description: Name used for UI elements listing this SLO.
        type: string
      goal:
        description: The fraction of service that must be good in order for this objective
          to be met. `0 < goal <= 0.999`.
        format: double
        type: number
      projectRef:
        description: Immutable. The Project that this resource belongs to.
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
              The project for the resource

              Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      rollingPeriod:
        description: A rolling time period, semantically "in the past ``". Must be
          an integer multiple of 1 day no larger than 30 days.
        type: string
      serviceLevelIndicator:
        description: The definition of good service, used to measure and calculate
          the quality of the `Service`'s performance with respect to a single aspect
          of service quality.
        properties:
          basicSli:
            description: Basic SLI on a well-known service type.
            properties:
              availability:
                description: Good service is defined to be the count of requests made
                  to this service that return successfully.
                type: object
                x-kubernetes-preserve-unknown-fields: true
              latency:
                description: Good service is defined to be the count of requests made
                  to this service that are fast enough with respect to `latency.threshold`.
                properties:
                  experience:
                    description: 'A description of the experience associated with
                      failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,
                      DELIGHTING, SATISFYING, ANNOYING'
                    type: string
                  threshold:
                    description: Good service is defined to be the count of requests
                      made to this service that return in no more than `threshold`.
                    type: string
                type: object
              location:
                description: 'OPTIONAL: The set of locations to which this SLI is
                  relevant. Telemetry from other locations will not be used to calculate
                  performance for this SLI. If omitted, this SLI applies to all locations
                  in which the Service has activity. For service types that don''t
                  support breaking down by location, setting this field will result
                  in an error.'
                items:
                  type: string
                type: array
              method:
                description: 'OPTIONAL: The set of RPCs to which this SLI is relevant.
                  Telemetry from other methods will not be used to calculate performance
                  for this SLI. If omitted, this SLI applies to all the Service''s
                  methods. For service types that don''t support breaking down by
                  method, setting this field will result in an error.'
                items:
                  type: string
                type: array
              operationAvailability:
                description: Good service is defined to be the count of operations
                  performed by this service that return successfully
                type: object
                x-kubernetes-preserve-unknown-fields: true
              operationLatency:
                description: Good service is defined to be the count of operations
                  performed by this service that are fast enough with respect to `operation_latency.threshold`.
                properties:
                  experience:
                    description: 'A description of the experience associated with
                      failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,
                      DELIGHTING, SATISFYING, ANNOYING'
                    type: string
                  threshold:
                    description: Good service is defined to be the count of operations
                      that are completed in no more than `threshold`.
                    type: string
                type: object
              version:
                description: 'OPTIONAL: The set of API versions to which this SLI
                  is relevant. Telemetry from other API versions will not be used
                  to calculate performance for this SLI. If omitted, this SLI applies
                  to all API versions. For service types that don''t support breaking
                  down by version, setting this field will result in an error.'
                items:
                  type: string
                type: array
            type: object
          requestBased:
            description: Request-based SLIs
            properties:
              distributionCut:
                description: '`distribution_cut` is used when `good_service` is a
                  count of values aggregated in a `Distribution` that fall into a
                  good range. The `total_service` is the total count of all values
                  aggregated in the `Distribution`.'
                properties:
                  distributionFilter:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying a `TimeSeries` aggregating values. Must have `ValueType
                      = DISTRIBUTION` and `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.
                    type: string
                  range:
                    description: Range of values considered "good." For a one-sided
                      range, set one bound to an infinite value.
                    properties:
                      max:
                        description: Range maximum.
                        format: double
                        type: number
                      min:
                        description: Range minimum.
                        format: double
                        type: number
                    type: object
                type: object
              goodTotalRatio:
                description: '`good_total_ratio` is used when the ratio of `good_service`
                  to `total_service` is computed from two `TimeSeries`.'
                properties:
                  badServiceFilter:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying a `TimeSeries` quantifying bad service, either demanded
                      service that was not provided or demanded service that was of
                      inadequate quality. Must have `ValueType = DOUBLE` or `ValueType
                      = INT64` and must have `MetricKind = DELTA` or `MetricKind =
                      CUMULATIVE`.
                    type: string
                  goodServiceFilter:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying a `TimeSeries` quantifying good service provided.
                      Must have `ValueType = DOUBLE` or `ValueType = INT64` and must
                      have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.
                    type: string
                  totalServiceFilter:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying a `TimeSeries` quantifying total demanded service.
                      Must have `ValueType = DOUBLE` or `ValueType = INT64` and must
                      have `MetricKind = DELTA` or `MetricKind = CUMULATIVE`.
                    type: string
                type: object
            type: object
          windowsBased:
            description: Windows-based SLIs
            properties:
              goodBadMetricFilter:
                description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                  specifying a `TimeSeries` with `ValueType = BOOL`. The window is
                  good if any `true` values appear in the window.
                type: string
              goodTotalRatioThreshold:
                description: A window is good if its `performance` is high enough.
                properties:
                  basicSliPerformance:
                    description: '`BasicSli` to evaluate to judge window quality.'
                    properties:
                      availability:
                        description: Good service is defined to be the count of requests
                          made to this service that return successfully.
                        type: object
                        x-kubernetes-preserve-unknown-fields: true
                      latency:
                        description: Good service is defined to be the count of requests
                          made to this service that are fast enough with respect to
                          `latency.threshold`.
                        properties:
                          experience:
                            description: 'A description of the experience associated
                              with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,
                              DELIGHTING, SATISFYING, ANNOYING'
                            type: string
                          threshold:
                            description: Good service is defined to be the count of
                              requests made to this service that return in no more
                              than `threshold`.
                            type: string
                        type: object
                      location:
                        description: 'OPTIONAL: The set of locations to which this
                          SLI is relevant. Telemetry from other locations will not
                          be used to calculate performance for this SLI. If omitted,
                          this SLI applies to all locations in which the Service has
                          activity. For service types that don''t support breaking
                          down by location, setting this field will result in an error.'
                        items:
                          type: string
                        type: array
                      method:
                        description: 'OPTIONAL: The set of RPCs to which this SLI
                          is relevant. Telemetry from other methods will not be used
                          to calculate performance for this SLI. If omitted, this
                          SLI applies to all the Service''s methods. For service types
                          that don''t support breaking down by method, setting this
                          field will result in an error.'
                        items:
                          type: string
                        type: array
                      operationAvailability:
                        description: Good service is defined to be the count of operations
                          performed by this service that return successfully
                        type: object
                        x-kubernetes-preserve-unknown-fields: true
                      operationLatency:
                        description: Good service is defined to be the count of operations
                          performed by this service that are fast enough with respect
                          to `operation_latency.threshold`.
                        properties:
                          experience:
                            description: 'A description of the experience associated
                              with failing requests. Possible values: LATENCY_EXPERIENCE_UNSPECIFIED,
                              DELIGHTING, SATISFYING, ANNOYING'
                            type: string
                          threshold:
                            description: Good service is defined to be the count of
                              operations that are completed in no more than `threshold`.
                            type: string
                        type: object
                      version:
                        description: 'OPTIONAL: The set of API versions to which this
                          SLI is relevant. Telemetry from other API versions will
                          not be used to calculate performance for this SLI. If omitted,
                          this SLI applies to all API versions. For service types
                          that don''t support breaking down by version, setting this
                          field will result in an error.'
                        items:
                          type: string
                        type: array
                    type: object
                  performance:
                    description: '`RequestBasedSli` to evaluate to judge window quality.'
                    properties:
                      distributionCut:
                        description: '`distribution_cut` is used when `good_service`
                          is a count of values aggregated in a `Distribution` that
                          fall into a good range. The `total_service` is the total
                          count of all values aggregated in the `Distribution`.'
                        properties:
                          distributionFilter:
                            description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                              specifying a `TimeSeries` aggregating values. Must have
                              `ValueType = DISTRIBUTION` and `MetricKind = DELTA`
                              or `MetricKind = CUMULATIVE`.
                            type: string
                          range:
                            description: Range of values considered "good." For a
                              one-sided range, set one bound to an infinite value.
                            properties:
                              max:
                                description: Range maximum.
                                format: double
                                type: number
                              min:
                                description: Range minimum.
                                format: double
                                type: number
                            type: object
                        type: object
                      goodTotalRatio:
                        description: '`good_total_ratio` is used when the ratio of
                          `good_service` to `total_service` is computed from two `TimeSeries`.'
                        properties:
                          badServiceFilter:
                            description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                              specifying a `TimeSeries` quantifying bad service, either
                              demanded service that was not provided or demanded service
                              that was of inadequate quality. Must have `ValueType
                              = DOUBLE` or `ValueType = INT64` and must have `MetricKind
                              = DELTA` or `MetricKind = CUMULATIVE`.
                            type: string
                          goodServiceFilter:
                            description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                              specifying a `TimeSeries` quantifying good service provided.
                              Must have `ValueType = DOUBLE` or `ValueType = INT64`
                              and must have `MetricKind = DELTA` or `MetricKind =
                              CUMULATIVE`.
                            type: string
                          totalServiceFilter:
                            description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                              specifying a `TimeSeries` quantifying total demanded
                              service. Must have `ValueType = DOUBLE` or `ValueType
                              = INT64` and must have `MetricKind = DELTA` or `MetricKind
                              = CUMULATIVE`.
                            type: string
                        type: object
                    type: object
                  threshold:
                    description: If window `performance >= threshold`, the window
                      is counted as good.
                    format: double
                    type: number
                type: object
              metricMeanInRange:
                description: A window is good if the metric's value is in a good range,
                  averaged across returned streams.
                properties:
                  range:
                    description: Range of values considered "good." For a one-sided
                      range, set one bound to an infinite value.
                    properties:
                      max:
                        description: Range maximum.
                        format: double
                        type: number
                      min:
                        description: Range minimum.
                        format: double
                        type: number
                    type: object
                  timeSeries:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying the `TimeSeries` to use for evaluating window quality.
                    type: string
                type: object
              metricSumInRange:
                description: A window is good if the metric's value is in a good range,
                  summed across returned streams.
                properties:
                  range:
                    description: Range of values considered "good." For a one-sided
                      range, set one bound to an infinite value.
                    properties:
                      max:
                        description: Range maximum.
                        format: double
                        type: number
                      min:
                        description: Range minimum.
                        format: double
                        type: number
                    type: object
                  timeSeries:
                    description: A [monitoring filter](https://cloud.google.com/monitoring/api/v3/filters)
                      specifying the `TimeSeries` to use for evaluating window quality.
                    type: string
                type: object
              windowPeriod:
                description: Duration over which window quality is evaluated. Must
                  be an integer fraction of a day and at least `60s`.
                type: string
            type: object
        type: object
      spec:
        required:
        - goal
        - projectRef
        - serviceRef
        type: object
required:
- spec
type: object

```
