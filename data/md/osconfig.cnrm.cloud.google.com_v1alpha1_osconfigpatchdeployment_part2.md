# openAPI schema part2 for osconfig.cnrm.cloud.google.com.v1alpha1.OSConfigPatchDeployment

## schema

```yaml
properties:
  spec:
    properties:
      recurringSchedule:
        properties:
          recurringSchedule:
            description: Immutable. Schedule recurring executions.
            properties:
              endTime:
                description: |-
                  Immutable. The end time at which a recurring patch deployment schedule is no longer active.
                  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
                type: string
              lastExecuteTime:
                description: |-
                  The time the last patch job ran successfully.
                  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
                type: string
              monthly:
                description: Immutable. Schedule with monthly executions.
                properties:
                  monthDay:
                    description: |-
                      Immutable. One day of the month. 1-31 indicates the 1st to the 31st day. -1 indicates the last day of the month.
                      Months without the target day will be skipped. For example, a schedule to run "every month on the 31st"
                      will not run in February, April, June, etc.
                    type: integer
                  weekDayOfMonth:
                    description: Immutable. Week day in a month.
                    properties:
                      dayOfWeek:
                        description: 'Immutable. A day of the week. Possible values:
                          ["MONDAY", "TUESDAY", "WEDNESDAY", "THURSDAY", "FRIDAY",
                          "SATURDAY", "SUNDAY"].'
                        type: string
                      weekOrdinal:
                        description: Immutable. Week number in a month. 1-4 indicates
                          the 1st to 4th week of the month. -1 indicates the last
                          week of the month.
                        type: integer
                    required:
                    - dayOfWeek
                    - weekOrdinal
                    type: object
                type: object
              nextExecuteTime:
                description: |-
                  The time the next patch job is scheduled to run.
                  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
                type: string
              startTime:
                description: |-
                  Immutable. The time that the recurring schedule becomes effective. Defaults to createTime of the patch deployment.
                  A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
                type: string
              timeOfDay:
                description: Immutable. Time of the day to run a recurring deployment.
                properties:
                  hours:
                    description: |-
                      Immutable. Hours of day in 24 hour format. Should be from 0 to 23.
                      An API may choose to allow the value "24:00:00" for scenarios like business closing time.
                    type: integer
                  minutes:
                    description: Immutable. Minutes of hour of day. Must be from 0
                      to 59.
                    type: integer
                  nanos:
                    description: Immutable. Fractions of seconds in nanoseconds. Must
                      be from 0 to 999,999,999.
                    type: integer
                  seconds:
                    description: Immutable. Seconds of minutes of the time. Must normally
                      be from 0 to 59. An API may allow the value 60 if it allows
                      leap-seconds.
                    type: integer
                type: object
              timeZone:
                description: |-
                  Immutable. Defines the time zone that timeOfDay is relative to. The rules for daylight saving time are
                  determined by the chosen time zone.
                properties:
                  id:
                    description: Immutable. IANA Time Zone Database time zone, e.g.
                      "America/New_York".
                    type: string
                  version:
                    description: Immutable. IANA Time Zone Database version number,
                      e.g. "2019a".
                    type: string
                required:
                - id
                type: object
              weekly:
                description: Immutable. Schedule with weekly executions.
                properties:
                  dayOfWeek:
                    description: 'Immutable. IANA Time Zone Database time zone, e.g.
                      "America/New_York". Possible values: ["MONDAY", "TUESDAY", "WEDNESDAY",
                      "THURSDAY", "FRIDAY", "SATURDAY", "SUNDAY"].'
                    type: string
                required:
                - dayOfWeek
                type: object
            required:
            - timeOfDay
            - timeZone
            type: object
      resourceID:
        description: Immutable. Optional. The service-generated name of the resource.
          Used for acquisition only. Leave unset to create a new resource.
        type: string
      rollout:
        description: Immutable. Rollout strategy of the patch job.
        properties:
          disruptionBudget:
            description: |-
              Immutable. The maximum number (or percentage) of VMs per zone to disrupt at any given moment. The number of VMs calculated from multiplying the percentage by the total number of VMs in a zone is rounded up.
              During patching, a VM is considered disrupted from the time the agent is notified to begin until patching has completed. This disruption time includes the time to complete reboot and any post-patch steps.
              A VM contributes to the disruption budget if its patching operation fails either when applying the patches, running pre or post patch steps, or if it fails to respond with a success notification before timing out. VMs that are not running or do not have an active agent do not count toward this disruption budget.
              For zone-by-zone rollouts, if the disruption budget in a zone is exceeded, the patch job stops, because continuing to the next zone requires completion of the patch process in the previous zone.
              For example, if the disruption budget has a fixed value of 10, and 8 VMs fail to patch in the current zone, the patch job continues to patch 2 VMs at a time until the zone is completed. When that zone is completed successfully, patching begins with 10 VMs at a time in the next zone. If 10 VMs in the next zone fail to patch, the patch job stops.
            properties:
              fixed:
                description: Immutable. Specifies a fixed value.
                type: integer
              percentage:
                description: Immutable. Specifies the relative value defined as a
                  percentage, which will be multiplied by a reference value.
                type: integer
            type: object
          mode:
            description: 'Immutable. Mode of the patch rollout. Possible values: ["ZONE_BY_ZONE",
              "CONCURRENT_ZONES"].'
            type: string
        required:
        - disruptionBudget
        - mode
        type: object
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
        description: |-
          Time the patch deployment was created. Timestamp is in RFC3339 text format.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        type: string
      lastExecuteTime:
        description: |-
          The last time a patch job was started by this deployment. Timestamp is in RFC3339 text format.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        type: string
      name:
        description: |-
          Unique name for the patch deployment resource in a project.
          The patch deployment name is in the form: projects/{project_id}/patchDeployments/{patchDeploymentId}.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: |-
          Time the patch deployment was last updated. Timestamp is in RFC3339 text format.
          A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        type: string
    type: object
required:
- spec
type: object

```
