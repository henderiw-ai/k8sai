# openAPI schema part2 for run.cnrm.cloud.google.com.v1beta1.RunJob

## schema

```yaml
properties:
  spec:
    properties:
      template:
        properties:
          annotations:
            additionalProperties:
              type: string
            description: |-
              Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.

              Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected.
              All system annotations in v1 now have a corresponding field in v2 ExecutionTemplate.

              This field follows Kubernetes annotations' namespacing, limits, and rules.
            type: object
          parallelism:
            description: Specifies the maximum desired number of tasks the execution
              should run at given time. Must be <= taskCount. When the job is run,
              if this field is 0 or unset, the maximum possible value will be used
              for that execution. The actual number of tasks running in steady state
              will be less than this number when there are fewer tasks waiting to
              be completed remaining, i.e. when the work left to do is less than max
              parallelism.
            type: integer
          taskCount:
            description: 'Specifies the desired number of tasks the execution should
              run. Setting to 1 means that parallelism is limited to 1 and the success
              of that task signals the success of the execution. More info: https://kubernetes.io/docs/concepts/workloads/controllers/jobs-run-to-completion/.'
            type: integer
          template:
            description: The template used to create executions for this Job.
            required:
            - template
            type: object
required:
- spec
type: object

```
