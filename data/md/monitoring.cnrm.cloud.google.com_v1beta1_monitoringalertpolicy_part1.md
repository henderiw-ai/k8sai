# openAPI schema part1 for monitoring.cnrm.cloud.google.com.v1beta1.MonitoringAlertPolicy

## schema

```yaml
properties:
  spec:
    properties:
      alertStrategy:
        description: Control over how this alert policy's notification channels are
          notified.
        properties:
          autoClose:
            description: If an alert policy that was active has no data for this long,
              any open incidents will close.
            type: string
          notificationChannelStrategy:
            description: |-
              Control over how the notification channels in 'notification_channels'
              are notified when this alert fires, on a per-channel basis.
            items:
              properties:
                notificationChannelNames:
                  description: |-
                    The notification channels that these settings apply to. Each of these
                    correspond to the name field in one of the NotificationChannel objects
                    referenced in the notification_channels field of this AlertPolicy. The format is
                    'projects/[PROJECT_ID_OR_NUMBER]/notificationChannels/[CHANNEL_ID]'.
                  items:
                    type: string
                  type: array
                renotifyInterval:
                  description: The frequency at which to send reminder notifications
                    for open incidents.
                  type: string
              type: object
            type: array
          notificationRateLimit:
            description: |-
              Required for alert policies with a LogMatch condition.
              This limit is not implemented for alert policies that are not log-based.
            properties:
              period:
                description: Not more than one notification per period.
                type: string
            type: object
        type: object
      combiner:
        description: |-
          How to combine the results of multiple conditions to
          determine if an incident should be opened. Possible values: ["AND", "OR", "AND_WITH_MATCHING_RESOURCE"].
        type: string
      spec:
        required:
        - combiner
        - conditions
        - displayName
        type: object
required:
- spec
type: object

```
