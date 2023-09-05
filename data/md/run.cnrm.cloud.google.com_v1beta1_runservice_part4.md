# openAPI schema part4 for run.cnrm.cloud.google.com.v1beta1.RunService

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
            description: Output only. The creation time.
            format: date-time
            type: string
          creator:
            description: Output only. Email address of the authenticated creator.
            type: string
          deleteTime:
            description: Output only. The deletion time.
            format: date-time
            type: string
          etag:
            description: Output only. A system-generated fingerprint for this version
              of the resource. May be used to detect modification conflict during
              updates.
            type: string
          expireTime:
            description: Output only. For a deleted resource, the time after which
              it will be permamently deleted.
            format: date-time
            type: string
          labels:
            additionalProperties:
              type: string
            description: Map of string keys and values that can be used to organize
              and categorize objects. User-provided labels are shared with Google's
              billing system, so they can be used to filter, or break down billing
              charges by team, component, environment, state, etc. For more information,
              visit https://cloud.google.com/resource-manager/docs/creating-managing-labels
              or https://cloud.google.com/run/docs/configuring/labels Cloud Run will
              populate some labels with 'run.googleapis.com' or 'serving.knative.dev'
              namespaces. Those labels are read-only, and user changes will not be
              preserved.
            type: object
          lastModifier:
            description: Output only. Email address of the last authenticated modifier.
            type: string
          latestCreatedRevision:
            description: Output only. Name of the last created revision. See comments
              in `reconciling` for additional information on reconciliation process
              in Cloud Run.
            type: string
          latestReadyRevision:
            description: Output only. Name of the latest revision that is serving
              traffic. See comments in `reconciling` for additional information on
              reconciliation process in Cloud Run.
            type: string
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          reconciling:
            description: 'Output only. Returns true if the Service is currently being
              acted upon by the system to bring it into the desired state. When a
              new Service is created, or an existing one is updated, Cloud Run will
              asynchronously perform all necessary steps to bring the Service to the
              desired serving state. This process is called reconciliation. While
              reconciliation is in process, `observed_generation`, `latest_ready_revison`,
              `traffic_statuses`, and `uri` will have transient values that might
              mismatch the intended state: Once reconciliation is over (and this field
              is false), there are two possible outcomes: reconciliation succeeded
              and the serving state matches the Service, or there was an error, and
              reconciliation failed. This state can be found in `terminal_condition.state`.
              If reconciliation succeeded, the following fields will match: `traffic`
              and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision`
              and `latest_created_revision`. If reconciliation failed, `traffic_statuses`,
              `observed_generation`, and `latest_ready_revision` will have the state
              of the last serving revision, or empty for newly created Services. Additional
              information on the failure can be found in `terminal_condition` and
              `conditions`.'
            type: boolean
          resourceGeneration:
            description: Output only. A number that monotonically increases every
              time the user modifies the desired state.
            format: int64
            type: integer
          terminalCondition:
            description: Output only. The Condition of this Service, containing its
              readiness status, and detailed error information in case it did not
              reach a serving state. See comments in `reconciling` for additional
              information on reconciliation process in Cloud Run.
            properties:
              jobReason:
                description: 'A reason for the job condition. Possible values: JOB_REASON_UNDEFINED,
                  JOB_STATUS_SERVICE_POLLING_ERROR'
                type: string
              lastTransitionTime:
                description: Last time the condition transitioned from one status
                  to another.
                format: date-time
                type: string
              message:
                description: Human readable message indicating details about the current
                  status.
                type: string
              reason:
                description: 'A common (service-level) reason for this condition.
                  Possible values: COMMON_REASON_UNDEFINED, UNKNOWN, REVISION_FAILED,
                  PROGRESS_DEADLINE_EXCEEDED, CONTAINER_MISSING, CONTAINER_PERMISSION_DENIED,
                  CONTAINER_IMAGE_UNAUTHORIZED, CONTAINER_IMAGE_AUTHORIZATION_CHECK_FAILED,
                  ENCRYPTION_KEY_PERMISSION_DENIED, ENCRYPTION_KEY_CHECK_FAILED, SECRETS_ACCESS_CHECK_FAILED,
                  WAITING_FOR_OPERATION, IMMEDIATE_RETRY, POSTPONED_RETRY, INTERNAL'
                type: string
              revisionReason:
                description: 'A reason for the revision condition. Possible values:
                  REVISION_REASON_UNDEFINED, PENDING, RESERVE, RETIRED, RETIRING,
                  RECREATING, HEALTH_CHECK_CONTAINER_ERROR, CUSTOMIZED_PATH_RESPONSE_PENDING,
                  MIN_INSTANCES_NOT_PROVISIONED, ACTIVE_REVISION_LIMIT_REACHED, NO_DEPLOYMENT'
                type: string
              severity:
                description: 'How to interpret failures of this condition, one of
                  Error, Warning, Info Possible values: SEVERITY_UNSPECIFIED, ERROR,
                  WARNING, INFO'
                type: string
              state:
                description: 'State of the condition. Possible values: STATE_UNSPECIFIED,
                  CONDITION_PENDING, CONDITION_RECONCILING, CONDITION_FAILED, CONDITION_SUCCEEDED'
                type: string
              type:
                description: 'type is used to communicate the status of the reconciliation
                  process. See also: https://github.com/knative/serving/blob/main/docs/spec/errors.md#error-conditions-and-reporting
                  Types common to all resources include: * "Ready": True when the
                  Resource is ready.'
                type: string
            type: object
          trafficStatuses:
            description: Output only. Detailed status information for corresponding
              traffic targets. See comments in `reconciling` for additional information
              on reconciliation process in Cloud Run.
            items:
              properties:
                percent:
                  description: Specifies percent of the traffic to this Revision.
                  format: int64
                  type: integer
                revision:
                  description: Revision to which this traffic is sent.
                  type: string
                tag:
                  description: Indicates the string used in the URI to exclusively
                    reference this target.
                  type: string
                type:
                  description: 'The allocation type for this traffic target. Possible
                    values: TRAFFIC_TARGET_ALLOCATION_TYPE_UNSPECIFIED, TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST,
                    TRAFFIC_TARGET_ALLOCATION_TYPE_REVISION'
                  type: string
                uri:
                  description: Displays the target URI.
                  type: string
              type: object
            type: array
          uid:
            description: Output only. Server assigned unique identifier for the trigger.
              The value is a UUID4 string and guaranteed to remain unchanged until
              the resource is deleted.
            type: string
          updateTime:
            description: Output only. The last-modified time.
            format: date-time
            type: string
          uri:
            description: Output only. The main URI in which this Service is serving
              traffic.
            type: string
        type: object
required:
- spec
type: object

```
