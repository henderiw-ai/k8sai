# openAPI schema part2 for appengine.cnrm.cloud.google.com.v1alpha1.AppEngineFlexibleAppVersion

## schema

```yaml
properties:
  spec:
    properties:
      nobuildFilesRegex:
        properties:
          nobuildFilesRegex:
            description: Files that match this pattern will not be built into this
              version. Only applicable for Go runtimes.
            type: string
      noopOnDestroy:
        description: If set to 'true', the application version will not be deleted.
        type: boolean
      project:
        description: Immutable.
        type: string
      readinessCheck:
        description: Configures readiness health checking for instances. Unhealthy
          instances are not put into the backend traffic rotation.
        properties:
          appStartTimeout:
            description: |-
              A maximum time limit on application initialization, measured from moment the application successfully
              replies to a healthcheck until it is ready to serve traffic. Default: "300s".
            type: string
          checkInterval:
            description: 'Interval between health checks.  Default: "5s".'
            type: string
          failureThreshold:
            description: 'Number of consecutive failed checks required before removing
              traffic. Default: 2.'
            type: number
          host:
            description: 'Host header to send when performing a HTTP Readiness check.
              Example: "myapp.appspot.com".'
            type: string
          path:
            description: The request path.
            type: string
          successThreshold:
            description: 'Number of consecutive successful checks required before
              receiving traffic. Default: 2.'
            type: number
          timeout:
            description: 'Time before the check is considered failed. Default: "4s".'
            type: string
        required:
        - path
        type: object
      resourceID:
        description: Immutable. Optional. The versionId of the resource. Used for
          creation and acquisition. When unset, the value of `metadata.name` is used
          as the default.
        type: string
      resources:
        description: Machine resources for a version.
        properties:
          cpu:
            description: Number of CPU cores needed.
            type: integer
          diskGb:
            description: Disk size (GB) needed.
            type: integer
          memoryGb:
            description: Memory (GB) needed.
            type: number
          volumes:
            description: List of ports, or port pairs, to forward from the virtual
              machine to the application container.
            items:
              properties:
                name:
                  description: Unique name for the volume.
                  type: string
                sizeGb:
                  description: Volume size in gigabytes.
                  type: integer
                volumeType:
                  description: Underlying volume type, e.g. 'tmpfs'.
                  type: string
              required:
              - name
              - sizeGb
              - volumeType
              type: object
            type: array
        type: object
      runtime:
        description: Desired runtime. Example python27.
        type: string
      runtimeApiVersion:
        description: |-
          The version of the API in the given runtime environment.
          Please see the app.yaml reference for valid values at 'https://cloud.google.com/appengine/docs/standard/<language>/config/appref'\
          Substitute '<language>' with 'python', 'java', 'php', 'ruby', 'go' or 'nodejs'.
        type: string
      runtimeChannel:
        description: The channel of the runtime to use. Only available for some runtimes.
        type: string
      runtimeMainExecutablePath:
        description: The path or name of the app's main executable.
        type: string
      serviceAccount:
        description: |-
          The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as
          default if this field is neither provided in app.yaml file nor through CLI flag.
        type: string
      serviceRef:
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
            description: 'Allowed value: The `name` field of an `AppEngineService`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      servingStatus:
        description: 'Current serving status of this version. Only the versions with
          a SERVING status create instances and can be billed. Default value: "SERVING"
          Possible values: ["SERVING", "STOPPED"].'
        type: string
      vpcAccessConnector:
        description: Enables VPC connectivity for standard apps.
        properties:
          name:
            description: Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
            type: string
        required:
        - name
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
      name:
        description: Full path to the Version resource in the API. Example, "v1".
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
    type: object
required:
- spec
type: object

```
