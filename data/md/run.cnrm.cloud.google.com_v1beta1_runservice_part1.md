# openAPI schema part1 for run.cnrm.cloud.google.com.v1beta1.RunService

## schema

```yaml
properties:
  spec:
    properties:
      annotations:
        additionalProperties:
          type: string
        description: 'Unstructured key value map that may be set by external tools
          to store and arbitrary metadata. They are not queryable and should be preserved
          when modifying objects. Cloud Run will populate some annotations using ''run.googleapis.com''
          or ''serving.knative.dev'' namespaces. This field follows Kubernetes annotations''
          namespacing, limits, and rules. More info: http://kubernetes.io/docs/user-guide/annotations'
        type: object
      binaryAuthorization:
        description: Settings for the Binary Authorization feature.
        properties:
          breakglassJustification:
            description: If present, indicates to use Breakglass using this justification.
              For more information on breakglass, see https://cloud.google.com/binary-authorization/docs/using-breakglass
            type: string
          useDefault:
            description: If True, indicates to use the default project's binary authorization
              policy. If False, binary authorization will be disabled
            type: boolean
        type: object
      client:
        description: Arbitrary identifier for the API client.
        type: string
      clientVersion:
        description: Arbitrary version identifier for the API client.
        type: string
      description:
        description: User-provided description of the Service.
        type: string
      ingress:
        description: Provides the ingress settings for this Service. On output, returns
          the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED
          if no revision is active.
        type: string
      launchStage:
        description: 'The launch stage as defined by [Google Cloud Platform Launch
          Stages](http://cloud.google.com/terms/launch-stages). Cloud Run supports
          `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Possible
          values: LAUNCH_STAGE_UNSPECIFIED, UNIMPLEMENTED, PRELAUNCH, EARLY_ACCESS,
          ALPHA, BETA, GA, DEPRECATED'
        type: string
      location:
        description: Immutable. The location for the resource
        type: string
      projectRef:
        description: Immutable. The Project that this resource belongs to.
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
            description: |-
              The project for the resource

              Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
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
