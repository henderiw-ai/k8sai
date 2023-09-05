# openAPI schema part4 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      placement:
        properties:
          managedCluster:
            properties:
              clusterName:
                description: Immutable. Required. The cluster name prefix. A unique
                  cluster name will be formed by appending a random suffix. The name
                  must contain only lower-case letters (a-z), numbers (0-9), and hyphens
                  (-). Must begin with a letter. Cannot begin or end with hyphen.
                  Must consist of between 2 and 35 characters.
                type: string
              managedCluster:
                description: Immutable. A cluster that is managed by the workflow.
                required:
                - clusterName
                - config
                type: object
required:
- spec
type: object

```
