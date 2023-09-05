# openAPI schema part1 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      dagTimeout:
        description: Immutable. Optional. Timeout duration for the DAG of jobs, expressed
          in seconds (see [JSON representation of duration](https://developers.google.com/protocol-buffers/docs/proto3#json)).
          The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s").
          The timer begins when the first job is submitted. If the workflow is running
          at the end of the timeout period, any remaining jobs are cancelled, the
          workflow is ended, and if the workflow was running on a [managed cluster](/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster),
          the cluster is deleted.
        type: string
      spec:
        required:
        - jobs
        - location
        - placement
        type: object
required:
- spec
type: object

```
