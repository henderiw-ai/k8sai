# openAPI schema part1 for logging.cnrm.cloud.google.com.v1beta1.LoggingLogMetric

## schema

```yaml
properties:
  status:
    properties:
      status:
        properties:
          conditions:
            description: Conditions represent the latest available observation of
              the resource's current state.
            items:
              properties:
                lastTransitionTime:
                  description: Last time the condition transitioned from one status
                    to another.
                  type: string
                message:
                  description: Human-readable message indicating details about last
                    transition.
                  type: string
                reason:
                  description: Unique, one-word, CamelCase reason for the condition's
                    last transition.
                  type: string
                status:
                  description: Status is the status of the condition. Can be True,
                    False, Unknown.
                  type: string
                type:
                  description: Type is the type of the condition.
                  type: string
              type: object
            type: array
          createTime:
            description: Output only. The creation timestamp of the metric. This field
              may not be present for older metrics.
            format: date-time
            type: string
          metricDescriptor:
            properties:
              description:
                description: A detailed description of the metric, which can be used
                  in documentation.
                type: string
              monitoredResourceTypes:
                description: Read-only. If present, then a time series, which is identified
                  partially by a metric type and a MonitoredResourceDescriptor, that
                  is associated with this metric type can only be associated with
                  one of the monitored resource types listed here.
                items:
                  type: string
                type: array
              name:
                description: The resource name of the metric descriptor.
                type: string
              type:
                description: 'The metric type, including its DNS name prefix. The
                  type is not URL-encoded. All user-defined metric types have the
                  DNS name `custom.googleapis.com` or `external.googleapis.com`. Metric
                  types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount"
                  "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies"'
                type: string
            type: object
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          updateTime:
            description: Output only. The last update timestamp of the metric. This
              field may not be present for older metrics.
            format: date-time
            type: string
        type: object
required:
- spec
type: object

```
