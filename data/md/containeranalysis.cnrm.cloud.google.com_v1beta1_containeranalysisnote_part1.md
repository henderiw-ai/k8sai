# openAPI schema part1 for containeranalysis.cnrm.cloud.google.com.v1beta1.ContainerAnalysisNote

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
            description: Output only. The time this note was created. This field can
              be used as a filter in list requests.
            format: date-time
            type: string
          image:
            properties:
              fingerprint:
                properties:
                  v2Name:
                    description: 'Output only. The name of the image''s v2 blobs computed
                      via: ) Only the name of the final blob is kept.'
                    type: string
                type: object
            type: object
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          updateTime:
            description: Output only. The time this note was last updated. This field
              can be used as a filter in list requests.
            format: date-time
            type: string
        type: object
type: object

```
