# openAPI schema part1 for run.cnrm.cloud.google.com.v1beta1.RunJob

## schema

```yaml
properties:
  spec:
    properties:
      annotations:
        additionalProperties:
          type: string
        description: |-
          Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects.

          Cloud Run API v2 does not support annotations with 'run.googleapis.com', 'cloud.googleapis.com', 'serving.knative.dev', or 'autoscaling.knative.dev' namespaces, and they will be rejected on new resources.
          All system annotations in v1 now have a corresponding field in v2 Job.

          This field follows Kubernetes annotations' namespacing, limits, and rules.
        type: object
      binaryAuthorization:
        description: Settings for the Binary Authorization feature.
        properties:
          breakglassJustification:
            description: If present, indicates to use Breakglass using this justification.
              If useDefault is False, then it must be empty. For more information
              on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass.
            type: string
          useDefault:
            description: If True, indicates to use the default project's binary authorization
              policy. If False, binary authorization will be disabled.
            type: boolean
        type: object
      client:
        description: Arbitrary identifier for the API client.
        type: string
      clientVersion:
        description: Arbitrary version identifier for the API client.
        type: string
      launchStage:
        description: |-
          The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/products#product-launch-stages). Cloud Run supports ALPHA, BETA, and GA.
          If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features.

          For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output. Possible values: ["UNIMPLEMENTED", "PRELAUNCH", "EARLY_ACCESS", "ALPHA", "BETA", "GA", "DEPRECATED"].
        type: string
      location:
        description: Immutable. The location of the cloud run job.
        type: string
      projectRef:
        description: The project that this resource belongs to.
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
            description: 'Allowed value: The `name` field of a `Project` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      spec:
        required:
        - location
        - projectRef
        - template
        type: object
required:
- spec
type: object

```
