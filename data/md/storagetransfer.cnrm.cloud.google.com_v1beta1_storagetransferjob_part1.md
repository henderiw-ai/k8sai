# openAPI schema part1 for storagetransfer.cnrm.cloud.google.com.v1beta1.StorageTransferJob

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: Unique description to identify the Transfer Job.
        type: string
      notificationConfig:
        description: Notification configuration.
        properties:
          eventTypes:
            description: Event types for which a notification is desired. If empty,
              send notifications for all event types. The valid types are "TRANSFER_OPERATION_SUCCESS",
              "TRANSFER_OPERATION_FAILED", "TRANSFER_OPERATION_ABORTED".
            items:
              type: string
            type: array
          payloadFormat:
            description: The desired format of the notification message payloads.
              One of "NONE" or "JSON".
            type: string
          topicRef:
            description: The PubSubTopic to which to publish notifications.
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
                description: 'Allowed value: string of the format `projects/{{project}}/topics/{{value}}`,
                  where {{value}} is the `name` field of a `PubSubTopic` resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
        required:
        - payloadFormat
        - topicRef
        type: object
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
      schedule:
        description: Schedule specification defining when the Transfer Job should
          be scheduled to start, end and what time to run.
        properties:
          repeatInterval:
            description: 'Interval between the start of each scheduled transfer. If
              unspecified, the default value is 24 hours. This value may not be less
              than 1 hour. A duration in seconds with up to nine fractional digits,
              terminated by ''s''. Example: "3.5s".'
            type: string
          scheduleEndDate:
            description: The last day the recurring transfer will be run. If schedule_end_date
              is the same as schedule_start_date, the transfer will be executed only
              once.
            properties:
              day:
                description: Day of month. Must be from 1 to 31 and valid for the
                  year and month.
                type: integer
              month:
                description: Month of year. Must be from 1 to 12.
                type: integer
              year:
                description: Year of date. Must be from 1 to 9999.
                type: integer
            required:
            - day
            - month
            - year
            type: object
          scheduleStartDate:
            description: The first day the recurring transfer is scheduled to run.
              If schedule_start_date is in the past, the transfer will run for the
              first time on the following day.
            properties:
              day:
                description: Day of month. Must be from 1 to 31 and valid for the
                  year and month.
                type: integer
              month:
                description: Month of year. Must be from 1 to 12.
                type: integer
              year:
                description: Year of date. Must be from 1 to 9999.
                type: integer
            required:
            - day
            - month
            - year
            type: object
          startTimeOfDay:
            description: The time in UTC at which the transfer will be scheduled to
              start in a day. Transfers may start later than this time. If not specified,
              recurring and one-time transfers that are scheduled to run today will
              run immediately; recurring transfers that are scheduled to run on a
              future date will start at approximately midnight UTC on that date. Note
              that when configuring a transfer with the Cloud Platform Console, the
              transfer's start time in a day is specified in your local timezone.
            properties:
              hours:
                description: Hours of day in 24 hour format. Should be from 0 to 23.
                type: integer
              minutes:
                description: Minutes of hour of day. Must be from 0 to 59.
                type: integer
              nanos:
                description: Fractions of seconds in nanoseconds. Must be from 0 to
                  999,999,999.
                type: integer
              seconds:
                description: Seconds of minutes of the time. Must normally be from
                  0 to 59.
                type: integer
            required:
            - hours
            - minutes
            - nanos
            - seconds
            type: object
        required:
        - scheduleStartDate
        type: object
      spec:
        required:
        - description
        - transferSpec
        type: object
      status:
        description: 'Status of the job. Default: ENABLED. NOTE: The effect of the
          new job status takes place during a subsequent job run. For example, if
          you change the job status from ENABLED to DISABLED, and an operation spawned
          by the transfer is running, the status change would not affect the current
          operation.'
        type: string
required:
- spec
type: object

```
