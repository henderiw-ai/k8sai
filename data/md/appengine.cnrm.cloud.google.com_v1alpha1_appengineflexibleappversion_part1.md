# openAPI schema part1 for appengine.cnrm.cloud.google.com.v1alpha1.AppEngineFlexibleAppVersion

## schema

```yaml
properties:
  spec:
    properties:
      apiConfig:
        description: Serving configuration for Google Cloud Endpoints.
        properties:
          authFailAction:
            description: 'Action to take when users access resources that require
              authentication. Default value: "AUTH_FAIL_ACTION_REDIRECT" Possible
              values: ["AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"].'
            type: string
          login:
            description: 'Level of login required to access this resource. Default
              value: "LOGIN_OPTIONAL" Possible values: ["LOGIN_OPTIONAL", "LOGIN_ADMIN",
              "LOGIN_REQUIRED"].'
            type: string
          script:
            description: Path to the script from the application root directory.
            type: string
          securityLevel:
            description: 'Security (HTTPS) enforcement for this URL. Possible values:
              ["SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"].'
            type: string
          url:
            description: URL to serve the endpoint at.
            type: string
        required:
        - script
        type: object
      automaticScaling:
        description: Automatic scaling is based on request rate, response latencies,
          and other application metrics.
        properties:
          coolDownPeriod:
            description: |-
              The time period that the Autoscaler should wait before it starts collecting information from a new instance.
              This prevents the autoscaler from collecting information when the instance is initializing,
              during which the collected usage would not be reliable. Default: 120s.
            type: string
          cpuUtilization:
            description: Target scaling by CPU usage.
            properties:
              aggregationWindowLength:
                description: Period of time over which CPU utilization is calculated.
                type: string
              targetUtilization:
                description: Target CPU utilization ratio to maintain when scaling.
                  Must be between 0 and 1.
                type: number
            required:
            - targetUtilization
            type: object
          diskUtilization:
            description: Target scaling by disk usage.
            properties:
              targetReadBytesPerSecond:
                description: Target bytes read per second.
                type: integer
              targetReadOpsPerSecond:
                description: Target ops read per seconds.
                type: integer
              targetWriteBytesPerSecond:
                description: Target bytes written per second.
                type: integer
              targetWriteOpsPerSecond:
                description: Target ops written per second.
                type: integer
            type: object
          maxConcurrentRequests:
            description: |-
              Number of concurrent requests an automatic scaling instance can accept before the scheduler spawns a new instance.

              Defaults to a runtime-specific value.
            type: integer
          maxIdleInstances:
            description: Maximum number of idle instances that should be maintained
              for this version.
            type: integer
          maxPendingLatency:
            description: Maximum amount of time that a request should wait in the
              pending queue before starting a new instance to handle it.
            type: string
          maxTotalInstances:
            description: 'Maximum number of instances that should be started to handle
              requests for this version. Default: 20.'
            type: integer
          minIdleInstances:
            description: Minimum number of idle instances that should be maintained
              for this version. Only applicable for the default version of a service.
            type: integer
          minPendingLatency:
            description: Minimum amount of time a request should wait in the pending
              queue before starting a new instance to handle it.
            type: string
          minTotalInstances:
            description: 'Minimum number of running instances that should be maintained
              for this version. Default: 2.'
            type: integer
          networkUtilization:
            description: Target scaling by network usage.
            properties:
              targetReceivedBytesPerSecond:
                description: Target bytes received per second.
                type: integer
              targetReceivedPacketsPerSecond:
                description: Target packets received per second.
                type: integer
              targetSentBytesPerSecond:
                description: Target bytes sent per second.
                type: integer
              targetSentPacketsPerSecond:
                description: Target packets sent per second.
                type: integer
            type: object
          requestUtilization:
            description: Target scaling by request utilization.
            properties:
              targetConcurrentRequests:
                description: Target number of concurrent requests.
                type: number
              targetRequestCountPerSecond:
                description: Target requests per second.
                type: string
            type: object
        required:
        - cpuUtilization
        type: object
      betaSettings:
        additionalProperties:
          type: string
        description: Metadata settings that are supplied to this version to enable
          beta runtime features.
        type: object
      defaultExpiration:
        description: |-
          Duration that static files should be cached by web proxies and browsers.
          Only applicable if the corresponding StaticFilesHandler does not specify its own expiration time.
        type: string
      deleteServiceOnDestroy:
        description: If set to 'true', the service will be deleted if it is the last
          version.
        type: boolean
      deployment:
        description: Code and application artifacts that make up this version.
        properties:
          cloudBuildOptions:
            description: Options for the build operations performed as a part of the
              version deployment. Only applicable when creating a version using source
              code directly.
            properties:
              appYamlPath:
                description: Path to the yaml file used in deployment, used to determine
                  runtime configuration details.
                type: string
              cloudBuildTimeout:
                description: |-
                  The Cloud Build timeout used as part of any dependent builds performed by version creation. Defaults to 10 minutes.

                  A duration in seconds with up to nine fractional digits, terminated by 's'. Example: "3.5s".
                type: string
            required:
            - appYamlPath
            type: object
          container:
            description: The Docker image for the container that runs the version.
            properties:
              image:
                description: |-
                  URI to the hosted container image in Google Container Registry. The URI must be fully qualified and include a tag or digest.
                  Examples: "gcr.io/my-project/image:tag" or "gcr.io/my-project/image@digest".
                type: string
            required:
            - image
            type: object
          files:
            description: |-
              Manifest of the files stored in Google Cloud Storage that are included as part of this version.
              All files must be readable using the credentials supplied with this call.
            items:
              properties:
                name:
                  type: string
                sha1Sum:
                  description: SHA1 checksum of the file.
                  type: string
                sourceUrl:
                  description: Source URL.
                  type: string
              required:
              - name
              - sourceUrl
              type: object
            type: array
          zip:
            description: Zip File.
            properties:
              filesCount:
                description: files count.
                type: integer
              sourceUrl:
                description: Source URL.
                type: string
            required:
            - sourceUrl
            type: object
        type: object
      endpointsApiService:
        description: Code and application artifacts that make up this version.
        properties:
          configId:
            description: |-
              Endpoints service configuration ID as specified by the Service Management API. For example "2016-09-19r1".

              By default, the rollout strategy for Endpoints is "FIXED". This means that Endpoints starts up with a particular configuration ID.
              When a new configuration is rolled out, Endpoints must be given the new configuration ID. The configId field is used to give the configuration ID
              and is required in this case.

              Endpoints also has a rollout strategy called "MANAGED". When using this, Endpoints fetches the latest configuration and does not need
              the configuration ID. In this case, configId must be omitted.
            type: string
          disableTraceSampling:
            description: Enable or disable trace sampling. By default, this is set
              to false for enabled.
            type: boolean
          name:
            description: |-
              Endpoints service name which is the name of the "service" resource in the Service Management API.
              For example "myapi.endpoints.myproject.cloud.goog".
            type: string
          rolloutStrategy:
            description: 'Endpoints rollout strategy. If FIXED, configId must be specified.
              If MANAGED, configId must be omitted. Default value: "FIXED" Possible
              values: ["FIXED", "MANAGED"].'
            type: string
        required:
        - name
        type: object
      entrypoint:
        description: The entrypoint for the application.
        properties:
          shell:
            description: The format should be a shell command that can be fed to bash
              -c.
            type: string
        required:
        - shell
        type: object
      envVariables:
        additionalProperties:
          type: string
        type: object
      handlers:
        description: |-
          An ordered list of URL-matching patterns that should be applied to incoming requests.
          The first matching URL handles the request and other request handlers are not attempted.
        items:
          properties:
            authFailAction:
              description: 'Actions to take when the user is not logged in. Possible
                values: ["AUTH_FAIL_ACTION_REDIRECT", "AUTH_FAIL_ACTION_UNAUTHORIZED"].'
              type: string
            login:
              description: 'Methods to restrict access to a URL based on login status.
                Possible values: ["LOGIN_OPTIONAL", "LOGIN_ADMIN", "LOGIN_REQUIRED"].'
              type: string
            redirectHttpResponseCode:
              description: '30x code to use when performing redirects for the secure
                field. Possible values: ["REDIRECT_HTTP_RESPONSE_CODE_301", "REDIRECT_HTTP_RESPONSE_CODE_302",
                "REDIRECT_HTTP_RESPONSE_CODE_303", "REDIRECT_HTTP_RESPONSE_CODE_307"].'
              type: string
            script:
              description: |-
                Executes a script to handle the requests that match this URL pattern.
                Only the auto value is supported for Node.js in the App Engine standard environment, for example "script:" "auto".
              properties:
                scriptPath:
                  description: Path to the script from the application root directory.
                  type: string
              required:
              - scriptPath
              type: object
            securityLevel:
              description: 'Security (HTTPS) enforcement for this URL. Possible values:
                ["SECURE_DEFAULT", "SECURE_NEVER", "SECURE_OPTIONAL", "SECURE_ALWAYS"].'
              type: string
            staticFiles:
              description: |-
                Files served directly to the user for a given URL, such as images, CSS stylesheets, or JavaScript source files.
                Static file handlers describe which files in the application directory are static files, and which URLs serve them.
              properties:
                applicationReadable:
                  description: |-
                    Whether files should also be uploaded as code data. By default, files declared in static file handlers are
                    uploaded as static data and are only served to end users; they cannot be read by the application. If enabled,
                    uploads are charged against both your code and static data storage resource quotas.
                  type: boolean
                expiration:
                  description: |-
                    Time a static file served by this handler should be cached by web proxies and browsers.
                    A duration in seconds with up to nine fractional digits, terminated by 's'. Example "3.5s".
                    Default is '0s'.
                  type: string
                httpHeaders:
                  additionalProperties:
                    type: string
                  description: |-
                    HTTP headers to use for all responses from these URLs.
                    An object containing a list of "key:value" value pairs.".
                  type: object
                mimeType:
                  description: |-
                    MIME type used to serve all files served by this handler.
                    Defaults to file-specific MIME types, which are derived from each file's filename extension.
                  type: string
                path:
                  description: |-
                    Path to the static files matched by the URL pattern, from the application root directory.
                    The path can refer to text matched in groupings in the URL pattern.
                  type: string
                requireMatchingFile:
                  description: Whether this handler should match the request if the
                    file referenced by the handler does not exist.
                  type: boolean
                uploadPathRegex:
                  description: Regular expression that matches the file paths for
                    all files that should be referenced by this handler.
                  type: string
              type: object
            urlRegex:
              description: |-
                URL prefix. Uses regular expression syntax, which means regexp special characters must be escaped, but should not contain groupings.
                All URLs that begin with this prefix are handled by this handler, using the portion of the URL after the prefix as part of the file path.
              type: string
          type: object
        type: array
      inboundServices:
        description: 'A list of the types of messages that this application is able
          to receive. Possible values: ["INBOUND_SERVICE_MAIL", "INBOUND_SERVICE_MAIL_BOUNCE",
          "INBOUND_SERVICE_XMPP_ERROR", "INBOUND_SERVICE_XMPP_MESSAGE", "INBOUND_SERVICE_XMPP_SUBSCRIBE",
          "INBOUND_SERVICE_XMPP_PRESENCE", "INBOUND_SERVICE_CHANNEL_PRESENCE", "INBOUND_SERVICE_WARMUP"].'
        items:
          type: string
        type: array
      instanceClass:
        description: |-
          Instance class that is used to run this version. Valid values are
          AutomaticScaling: F1, F2, F4, F4_1G
          ManualScaling: B1, B2, B4, B8, B4_1G
          Defaults to F1 for AutomaticScaling and B1 for ManualScaling.
        type: string
      livenessCheck:
        description: Health checking configuration for VM instances. Unhealthy instances
          are killed and replaced with new instances.
        properties:
          checkInterval:
            description: Interval between health checks.
            type: string
          failureThreshold:
            description: 'Number of consecutive failed checks required before considering
              the VM unhealthy. Default: 4.'
            type: number
          host:
            description: 'Host header to send when performing a HTTP Readiness check.
              Example: "myapp.appspot.com".'
            type: string
          initialDelay:
            description: 'The initial delay before starting to execute the checks.
              Default: "300s".'
            type: string
          path:
            description: The request path.
            type: string
          successThreshold:
            description: 'Number of consecutive successful checks required before
              considering the VM healthy. Default: 2.'
            type: number
          timeout:
            description: 'Time before the check is considered failed. Default: "4s".'
            type: string
        required:
        - path
        type: object
      manualScaling:
        description: A service with manual scaling runs continuously, allowing you
          to perform complex initialization and rely on the state of its memory over
          time.
        properties:
          instances:
            description: |-
              Number of instances to assign to the service at the start.

              **Note:** When managing the number of instances at runtime through the App Engine Admin API or the (now deprecated) Python 2
              Modules API set_num_instances() you must use 'lifecycle.ignore_changes = ["manual_scaling"[0].instances]' to prevent drift detection.
            type: integer
        required:
        - instances
        type: object
      network:
        description: Extra network settings.
        properties:
          forwardedPorts:
            description: List of ports, or port pairs, to forward from the virtual
              machine to the application container.
            items:
              type: string
            type: array
          instanceTag:
            description: Tag to apply to the instance during creation.
            type: string
          name:
            description: Google Compute Engine network where the virtual machines
              are created. Specify the short name, not the resource path.
            type: string
          sessionAffinity:
            description: Enable session affinity.
            type: boolean
          subnetwork:
            description: |-
              Google Cloud Platform sub-network where the virtual machines are created. Specify the short name, not the resource path.

              If the network that the instance is being created in is a Legacy network, then the IP address is allocated from the IPv4Range.
              If the network that the instance is being created in is an auto Subnet Mode Network, then only network name should be specified (not the subnetworkName) and the IP address is created from the IPCidrRange of the subnetwork that exists in that zone for that network.
              If the network that the instance is being created in is a custom Subnet Mode Network, then the subnetworkName must be specified and the IP address is created from the IPCidrRange of the subnetwork.
              If specified, the subnetwork must exist in the same region as the App Engine flexible environment application.
            type: string
        required:
        - name
        type: object
      spec:
        required:
        - livenessCheck
        - readinessCheck
        - runtime
        - serviceRef
        type: object
required:
- spec
type: object

```