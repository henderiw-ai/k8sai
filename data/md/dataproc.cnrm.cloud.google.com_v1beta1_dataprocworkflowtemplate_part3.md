# openAPI schema part3 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      placement:
        properties:
          clusterSelector:
            description: Immutable. Optional. A selector that chooses target cluster
              for jobs based on metadata. The selector is evaluated at the time each
              job is submitted.
            properties:
              clusterLabels:
                additionalProperties:
                  type: string
                description: Immutable. Required. The cluster labels. Cluster must
                  have all labels to match.
                type: object
              zone:
                description: Immutable. Optional. The zone where workflow process
                  executes. This parameter does not affect the selection of the cluster.
                  If unspecified, the zone of the first cluster matching the selector
                  is used.
                type: string
            required:
            - clusterLabels
            type: object
          placement:
            description: Immutable. Required. WorkflowTemplate scheduling information.
            type: object
required:
- spec
type: object

```
