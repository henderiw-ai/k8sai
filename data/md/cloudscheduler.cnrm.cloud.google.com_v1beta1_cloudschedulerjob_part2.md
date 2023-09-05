# openAPI schema part2 for cloudscheduler.cnrm.cloud.google.com.v1beta1.CloudSchedulerJob

## schema

```yaml
properties:
  status:
    properties:
      status:
        properties:
          appEngineHttpTarget:
            properties:
              appEngineRouting:
                properties:
                  host:
                    description: 'Output only. The host that the job is sent to. For
                      more information about how App Engine requests are routed, see
                      [here](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed).
                      The host is constructed as: * `host = [application_domain_name]`
                      `| [service] + ''.'' + [application_domain_name]` `| [version]
                      + ''.'' + [application_domain_name]` `| [version_dot_service]+
                      ''.'' + [application_domain_name]` `| [instance] + ''.'' + [application_domain_name]`
                      `| [instance_dot_service] + ''.'' + [application_domain_name]`
                      `| [instance_dot_version] + ''.'' + [application_domain_name]`
                      `| [instance_dot_version_dot_service] + ''.'' + [application_domain_name]`
                      * `application_domain_name` = The domain name of the app, for
                      example .appspot.com, which is associated with the job''s project
                      ID. * `service =` service * `version =` version * `version_dot_service
                      =` version `+ ''.'' +` service * `instance =` instance * `instance_dot_service
                      =` instance `+ ''.'' +` service * `instance_dot_version =` instance
                      `+ ''.'' +` version * `instance_dot_version_dot_service =` instance
                      `+ ''.'' +` version `+ ''.'' +` service If service is empty,
                      then the job will be sent to the service which is the default
                      service when the job is attempted. If version is empty, then
                      the job will be sent to the version which is the default version
                      when the job is attempted. If instance is empty, then the job
                      will be sent to an instance which is available when the job
                      is attempted. If service, version, or instance is invalid, then
                      the job will be sent to the default version of the default service
                      when the job is attempted.'
                    type: string
                type: object
            type: object
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
          lastAttemptTime:
            description: Output only. The time the last job attempt started.
            format: date-time
            type: string
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          scheduleTime:
            description: Output only. The next time the job is scheduled. Note that
              this may be a retry of a previously failed attempt or the next execution
              time according to the schedule.
            format: date-time
            type: string
          state:
            description: 'Output only. State of the job. Possible values: STATE_UNSPECIFIED,
              ENABLED, PAUSED, DISABLED, UPDATE_FAILED'
            type: string
          status:
            description: Output only. The response from the target for the last attempted
              execution.
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
                      description: 'A URL/resource name that uniquely identifies the
                        type of the serialized protocol buffer message. This string
                        must contain at least one "/" character. The last segment
                        of the URL''s path must represent the fully qualified name
                        of the type (as in `path/google.protobuf.Duration`). The name
                        should be in a canonical form (e.g., leading "." is not accepted).
                        In practice, teams usually precompile into the binary all
                        types that they expect it to use in the context of Any. However,
                        for URLs which use the scheme `http`, `https`, or no scheme,
                        one can optionally set up a type server that maps type URLs
                        to message definitions as follows: * If no scheme is provided,
                        `https` is assumed. * An HTTP GET on the URL must yield a
                        google.protobuf.Type value in binary format, or produce an
                        error. * Applications are allowed to cache lookup results
                        based on the URL, or have them precompiled into a binary to
                        avoid any lookup. Therefore, binary compatibility needs to
                        be preserved on changes to types. (Use versioned type names
                        to manage breaking changes.) Note: this functionality is not
                        currently available in the official protobuf release, and
                        it is not used for type URLs beginning with type.googleapis.com.
                        Schemes other than `http`, `https` (or the empty scheme) might
                        be used with implementation specific semantics.'
                      type: string
                    value:
                      description: Must be a valid serialized protocol buffer of the
                        above specified type.
                      type: string
                  type: object
                type: array
              message:
                description: A developer-facing error message, which should be in
                  English. Any user-facing error message should be localized and sent
                  in the google.rpc.Status.details field, or localized by the client.
                type: string
            type: object
          userUpdateTime:
            description: Output only. The creation time of the job.
            format: date-time
            type: string
        type: object
required:
- spec
type: object

```
