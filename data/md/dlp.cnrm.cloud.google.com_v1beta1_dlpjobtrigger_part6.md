# openAPI schema part6 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      projectRef:
        properties:
          projectRef:
            description: Immutable. The Project that this resource belongs to. Only
              one of [projectRef] may be specified.
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
                description: 'Allowed value: The Google Cloud resource name of a `Project`
                  resource (format: `projects/{{name}}`).'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
      status:
        description: 'Immutable. Required. A status for this trigger. Possible values:
          STATUS_UNSPECIFIED, HEALTHY, PAUSED, CANCELLED'
        type: string
      triggers:
        description: A list of triggers which will be OR'ed together. Only one in
          the list needs to trigger for a job to be started. The list may contain
          only a single Schedule trigger and must have at least one object.
        items:
          properties:
            manual:
              description: For use with hybrid jobs. Jobs must be manually created
                and finished.
              type: object
              x-kubernetes-preserve-unknown-fields: true
            schedule:
              description: Create a job on a repeating basis based on the elapse of
                time.
              properties:
                recurrencePeriodDuration:
                  description: 'With this option a job is started a regular periodic
                    basis. For example: every day (86400 seconds). A scheduled start
                    time will be skipped if the previous execution has not ended when
                    its scheduled time occurs. This value must be set to a time duration
                    greater than or equal to 1 day and can be no longer than 60 days.'
                  type: string
              type: object
          type: object
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
      createTime:
        description: Output only. The creation timestamp of a triggeredJob.
        format: date-time
        type: string
      errors:
        description: Output only. A stream of errors encountered when the trigger
          was activated. Repeated errors may result in the JobTrigger automatically
          being paused. Will return the last 100 errors. Whenever the JobTrigger is
          modified this list will be cleared.
        items:
          properties:
            details:
              description: Detailed error codes and messages.
              properties:
                code:
                  description: The status code, which should be an enum value of google.rpc.Code.
                  format: int64
                  type: integer
                details:
                  description: A list of messages that carry the error details. There
                    is a common set of message types for APIs to use.
                  items:
                    properties:
                      typeUrl:
                        description: 'A URL/resource name that uniquely identifies
                          the type of the serialized protocol buffer message. This
                          string must contain at least one "/" character. The last
                          segment of the URL''s path must represent the fully qualified
                          name of the type (as in `path/google.protobuf.Duration`).
                          The name should be in a canonical form (e.g., leading "."
                          is not accepted). In practice, teams usually precompile
                          into the binary all types that they expect it to use in
                          the context of Any. However, for URLs which use the scheme
                          `http`, `https`, or no scheme, one can optionally set up
                          a type server that maps type URLs to message definitions
                          as follows: * If no scheme is provided, `https` is assumed.
                          * An HTTP GET on the URL must yield a google.protobuf.Type
                          value in binary format, or produce an error. * Applications
                          are allowed to cache lookup results based on the URL, or
                          have them precompiled into a binary to avoid any lookup.
                          Therefore, binary compatibility needs to be preserved on
                          changes to types. (Use versioned type names to manage breaking
                          changes.) Note: this functionality is not currently available
                          in the official protobuf release, and it is not used for
                          type URLs beginning with type.googleapis.com. Schemes other
                          than `http`, `https` (or the empty scheme) might be used
                          with implementation specific semantics.'
                        type: string
                      value:
                        description: Must be a valid serialized protocol buffer of
                          the above specified type.
                        type: string
                    type: object
                  type: array
                message:
                  description: A developer-facing error message, which should be in
                    English. Any user-facing error message should be localized and
                    sent in the google.rpc.Status.details field, or localized by the
                    client.
                  type: string
              type: object
            timestamps:
              description: The times the error occurred.
              items:
                format: date-time
                type: string
              type: array
          type: object
        type: array
      lastRunTime:
        description: Output only. The timestamp of the last time this trigger executed.
        format: date-time
        type: string
      locationId:
        description: Output only. The geographic location where this resource is stored.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: Output only. The last update timestamp of a triggeredJob.
        format: date-time
        type: string
    type: object
required:
- spec
type: object

```
