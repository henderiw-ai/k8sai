# openAPI schema part1 for redis.cnrm.cloud.google.com.v1beta1.RedisInstance

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
            description: |-
              The time the instance was created in RFC3339 UTC "Zulu" format,
              accurate to nanoseconds.
            type: string
          currentLocationId:
            description: |-
              The current zone where the Redis endpoint is placed.
              For Basic Tier instances, this will always be the same as the
              [locationId] provided by the user at creation time. For Standard Tier
              instances, this can be either [locationId] or [alternativeLocationId]
              and can change after a failover event.
            type: string
          host:
            description: |-
              Hostname or IP address of the exposed Redis endpoint used by clients
              to connect to the service.
            type: string
          maintenanceSchedule:
            description: Upcoming maintenance schedule.
            items:
              properties:
                endTime:
                  description: |-
                    Output only. The end time of any upcoming scheduled maintenance for this instance.
                    A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                    resolution and up to nine fractional digits.
                  type: string
                scheduleDeadlineTime:
                  description: |-
                    Output only. The deadline that the maintenance schedule start time
                    can not go beyond, including reschedule.
                    A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                    resolution and up to nine fractional digits.
                  type: string
                startTime:
                  description: |-
                    Output only. The start time of any upcoming scheduled maintenance for this instance.
                    A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
                    resolution and up to nine fractional digits.
                  type: string
              type: object
            type: array
          nodes:
            description: Output only. Info per node.
            items:
              properties:
                id:
                  description: Node identifying string. e.g. 'node-0', 'node-1'.
                  type: string
                zone:
                  description: Location of the node.
                  type: string
              type: object
            type: array
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          persistenceIamIdentity:
            description: |-
              Output only. Cloud IAM identity used by import / export operations
              to transfer data to/from Cloud Storage. Format is "serviceAccount:".
              The value may change over time for a given instance so should be
              checked before each import/export operation.
            type: string
          port:
            description: The port number of the exposed Redis endpoint.
            type: integer
          readEndpoint:
            description: |-
              Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.
              Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes
              will exhibit some lag behind the primary. Write requests must target 'host'.
            type: string
          readEndpointPort:
            description: |-
              Output only. The port number of the exposed readonly redis endpoint. Standard tier only.
              Write requests should target 'port'.
            type: integer
          serverCaCerts:
            description: List of server CA certificates for the instance.
            items:
              properties:
                cert:
                  description: The certificate data in PEM format.
                  type: string
                createTime:
                  description: The time when the certificate was created.
                  type: string
                expireTime:
                  description: The time when the certificate expires.
                  type: string
                serialNumber:
                  description: Serial number, as extracted from the certificate.
                  type: string
                sha1Fingerprint:
                  description: Sha1 Fingerprint of the certificate.
                  type: string
              type: object
            type: array
        type: object
required:
- spec
type: object

```
