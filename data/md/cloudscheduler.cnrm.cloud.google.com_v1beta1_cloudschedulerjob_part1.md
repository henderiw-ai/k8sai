# openAPI schema part1 for cloudscheduler.cnrm.cloud.google.com.v1beta1.CloudSchedulerJob

## schema

```yaml
properties:
  spec:
    properties:
      spec:
        properties:
          appEngineHttpTarget:
            description: App Engine HTTP target.
            properties:
              appEngineRouting:
                description: App Engine Routing setting for the job.
                properties:
                  instance:
                    description: App instance. By default, the job is sent to an instance
                      which is available when the job is attempted. Requests can only
                      be sent to a specific instance if [manual scaling is used in
                      App Engine Standard](https://cloud.google.com/appengine/docs/python/an-overview-of-app-engine?hl=en_US#scaling_types_and_instance_classes).
                      App Engine Flex does not support instances. For more information,
                      see [App Engine Standard request routing](https://cloud.google.com/appengine/docs/standard/python/how-requests-are-routed)
                      and [App Engine Flex request routing](https://cloud.google.com/appengine/docs/flexible/python/how-requests-are-routed).
                    type: string
                  service:
                    description: App service. By default, the job is sent to the service
                      which is the default service when the job is attempted.
                    type: string
                  version:
                    description: App version. By default, the job is sent to the version
                      which is the default version when the job is attempted.
                    type: string
                type: object
              body:
                description: Body. HTTP request body. A request body is allowed only
                  if the HTTP method is POST or PUT. It will result in invalid argument
                  error to set a body on a job with an incompatible HttpMethod.
                type: string
              headers:
                additionalProperties:
                  type: string
                description: 'HTTP request headers. This map contains the header field
                  names and values. Headers can be set when the job is created. Cloud
                  Scheduler sets some headers to default values: * `User-Agent`: By
                  default, this header is `"App Engine-Google; (+http://code.google.com/appengine)"`.
                  This header can be modified, but Cloud Scheduler will append `"App
                  Engine-Google; (+http://code.google.com/appengine)"` to the modified
                  `User-Agent`. * `X-CloudScheduler`: This header will be set to true.
                  The headers below are output only. They cannot be set or overridden:
                  * `X-Google-*`: For Google internal use only. * `X-App Engine-*`:
                  For Google internal use only. In addition, some App Engine headers,
                  which contain job-specific information, are also be sent to the
                  job handler.'
                type: object
              httpMethod:
                description: 'The HTTP method to use for the request. PATCH and OPTIONS
                  are not permitted. Possible values: HTTP_METHOD_UNSPECIFIED, POST,
                  GET, HEAD, PUT, DELETE, PATCH, OPTIONS'
                type: string
              relativeUri:
                description: The relative URI. The relative URL must begin with "/"
                  and must be a valid HTTP relative URL. It can contain a path, query
                  string arguments, and `#` fragments. If the relative URL is empty,
                  then the root path "/" will be used. No spaces are allowed, and
                  the maximum length allowed is 2083 characters.
                type: string
            type: object
          attemptDeadline:
            description: 'The deadline for job attempts. If the request handler does
              not respond by this deadline then the request is cancelled and the attempt
              is marked as a `DEADLINE_EXCEEDED` failure. The failed attempt can be
              viewed in execution logs. Cloud Scheduler will retry the job according
              to the RetryConfig. The allowed duration for this deadline is: * For
              HTTP targets, between 15 seconds and 30 minutes. * For App Engine HTTP
              targets, between 15 seconds and 24 hours.'
            type: string
          description:
            description: Optionally caller-specified in CreateJob or UpdateJob. A
              human-readable description for the job. This string must not contain
              more than 500 characters.
            type: string
          httpTarget:
            description: HTTP target.
            properties:
              body:
                description: HTTP request body. A request body is allowed only if
                  the HTTP method is POST, PUT, or PATCH. It is an error to set body
                  on a job with an incompatible HttpMethod.
                type: string
              headers:
                additionalProperties:
                  type: string
                description: 'The user can specify HTTP request headers to send with
                  the job''s HTTP request. This map contains the header field names
                  and values. Repeated headers are not supported, but a header value
                  can contain commas. These headers represent a subset of the headers
                  that will accompany the job''s HTTP request. Some HTTP request headers
                  will be ignored or replaced. A partial list of headers that will
                  be ignored or replaced is below: - Host: This will be computed by
                  Cloud Scheduler and derived from uri. * `Content-Length`: This will
                  be computed by Cloud Scheduler. * `User-Agent`: This will be set
                  to `"Google-Cloud-Scheduler"`. * `X-Google-*`: Google internal use
                  only. * `X-appengine-*`: Google internal use only. The total size
                  of headers must be less than 80KB.'
                type: object
              httpMethod:
                description: 'Which HTTP method to use for the request. Possible values:
                  HTTP_METHOD_UNSPECIFIED, POST, GET, HEAD, PUT, DELETE, PATCH, OPTIONS'
                type: string
              oauthToken:
                description: If specified, an [OAuth token](https://developers.google.com/identity/protocols/OAuth2)
                  will be generated and attached as an `Authorization` header in the
                  HTTP request. This type of authorization should generally only be
                  used when calling Google APIs hosted on *.googleapis.com.
                properties:
                  scope:
                    description: OAuth scope to be used for generating OAuth access
                      token. If not specified, "https://www.googleapis.com/auth/cloud-platform"
                      will be used.
                    type: string
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
                        description: |-
                          [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OAuth token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.

                          Allowed value: The `email` field of an `IAMServiceAccount` resource.
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                type: object
              oidcToken:
                description: If specified, an [OIDC](https://developers.google.com/identity/protocols/OpenIDConnect)
                  token will be generated and attached as an `Authorization` header
                  in the HTTP request. This type of authorization can be used for
                  many scenarios, including calling Cloud Run, or endpoints where
                  you intend to validate the token yourself.
                properties:
                  audience:
                    description: Audience to be used when generating OIDC token. If
                      not specified, the URI specified in target will be used.
                    type: string
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
                        description: |-
                          [Service account email](https://cloud.google.com/iam/docs/service-accounts) to be used for generating OIDC token. The service account must be within the same project as the job. The caller must have iam.serviceAccounts.actAs permission for the service account.

                          Allowed value: The `email` field of an `IAMServiceAccount` resource.
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                type: object
              uri:
                description: 'Required. The full URI path that the request will be
                  sent to. This string must begin with either "http://" or "https://".
                  Some examples of valid values for uri are: `http://acme.com` and
                  `https://acme.com/sales:8080`. Cloud Scheduler will encode some
                  characters for safety and compatibility. The maximum allowed URL
                  length is 2083 characters after encoding.'
                type: string
            required:
            - uri
            type: object
          location:
            description: Immutable. The location for the resource
            type: string
          pubsubTarget:
            description: Pub/Sub target.
            properties:
              attributes:
                additionalProperties:
                  type: string
                description: Attributes for PubsubMessage. Pubsub message must contain
                  either non-empty data, or at least one attribute.
                type: object
              data:
                description: The message payload for PubsubMessage. Pubsub message
                  must contain either non-empty data, or at least one attribute.
                type: string
              topicRef:
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
                      Required. The name of the Cloud Pub/Sub topic to which messages will be published when a job is delivered. The topic name must be in the same format as required by Pub/Sub's [PublishRequest.name](https://cloud.google.com/pubsub/docs/reference/rpc/google.pubsub.v1#publishrequest), for example `projects/PROJECT_ID/topics/TOPIC_ID`. The topic must be in the same project as the Cloud Scheduler job.

                      Allowed value: The Google Cloud resource name of a `PubSubTopic` resource (format: `projects/{{project}}/topics/{{name}}`).
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            required:
            - topicRef
            type: object
          resourceID:
            description: Immutable. Optional. The name of the resource. Used for creation
              and acquisition. When unset, the value of `metadata.name` is used as
              the default.
            type: string
          retryConfig:
            description: Settings that determine the retry behavior.
            properties:
              maxBackoffDuration:
                description: The maximum amount of time to wait before retrying a
                  job after it fails. The default value of this field is 1 hour.
                type: string
              maxDoublings:
                description: The time between retries will double `max_doublings`
                  times. A job's retry interval starts at min_backoff_duration, then
                  doubles `max_doublings` times, then increases linearly, and finally
                  retries at intervals of max_backoff_duration up to retry_count times.
                  For example, if min_backoff_duration is 10s, max_backoff_duration
                  is 300s, and `max_doublings` is 3, then the a job will first be
                  retried in 10s. The retry interval will double three times, and
                  then increase linearly by 2^3 * 10s. Finally, the job will retry
                  at intervals of max_backoff_duration until the job has been attempted
                  retry_count times. Thus, the requests will retry at 10s, 20s, 40s,
                  80s, 160s, 240s, 300s, 300s, .... The default value of this field
                  is 5.
                format: int64
                type: integer
              maxRetryDuration:
                description: The time limit for retrying a failed job, measured from
                  time when an execution was first attempted. If specified with retry_count,
                  the job will be retried until both limits are reached. The default
                  value for max_retry_duration is zero, which means retry duration
                  is unlimited.
                type: string
              minBackoffDuration:
                description: The minimum amount of time to wait before retrying a
                  job after it fails. The default value of this field is 5 seconds.
                type: string
              retryCount:
                description: The number of attempts that the system will make to run
                  a job using the exponential backoff procedure described by max_doublings.
                  The default value of retry_count is zero. If retry_count is zero,
                  a job attempt will *not* be retried if it fails. Instead the Cloud
                  Scheduler system will wait for the next scheduled execution time.
                  If retry_count is set to a non-zero number then Cloud Scheduler
                  will retry failed attempts, using exponential backoff, retry_count
                  times, or until the next scheduled execution time, whichever comes
                  first. Values greater than 5 and negative values are not allowed.
                format: int64
                type: integer
            type: object
          schedule:
            description: 'Required, except when used with UpdateJob. Describes the
              schedule on which the job will be executed. The schedule can be either
              of the following types: * [Crontab](http://en.wikipedia.org/wiki/Cron#Overview)
              * English-like [schedule](https://cloud.google.com/scheduler/docs/configuring/cron-job-schedules)
              As a general rule, execution `n + 1` of a job will not begin until execution
              `n` has finished. Cloud Scheduler will never allow two simultaneously
              outstanding executions. For example, this implies that if the `n+1`th
              execution is scheduled to run at 16:00 but the `n`th execution takes
              until 16:15, the `n+1`th execution will not start until `16:15`. A scheduled
              start time will be delayed if the previous execution has not ended when
              its scheduled time occurs. If retry_count > 0 and a job attempt fails,
              the job will be tried a total of retry_count times, with exponential
              backoff, until the next scheduled start time.'
            type: string
          timeZone:
            description: Specifies the time zone to be used in interpreting schedule.
              The value of this field must be a time zone name from the [tz database](http://en.wikipedia.org/wiki/Tz_database).
              Note that some time zones include a provision for daylight savings time.
              The rules for daylight saving time are determined by the chosen tz.
              For UTC use the string "utc". If a time zone is not specified, the default
              will be in UTC (also known as GMT).
            type: string
        required:
        - location
        type: object
required:
- spec
type: object

```