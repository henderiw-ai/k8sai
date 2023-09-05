# openAPI schema part3 for compute.cnrm.cloud.google.com.v1beta1.ComputeURLMap

## schema

```yaml
properties:
  spec:
    properties:
      defaultService:
        properties:
          defaultService:
            description: |-
              The defaultService resource to which traffic is directed if none of
              the hostRules match.
              For the Global URL Map, it should be a reference to the backend
              service or backend bucket.
              For the Regional URL Map, it should be a reference to the backend
              service.
              If defaultRouteAction is additionally specified, advanced routing
              actions like URL Rewrites, etc. take effect prior to sending the
              request to the backend. However, if defaultService is specified,
              defaultRouteAction cannot contain any weightedBackendServices.
              Conversely, if routeAction specifies any weightedBackendServices,
              service must not be specified. Only one of defaultService,
              defaultUrlRedirect or defaultRouteAction.weightedBackendService
              must be set.
            oneOf:
            - required:
              - backendBucketRef
            - required:
              - backendServiceRef
            properties:
              backendBucketRef:
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
                    description: 'Allowed value: The `selfLink` field of a `ComputeBackendBucket`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
              backendServiceRef:
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
                    description: 'Allowed value: The `selfLink` field of a `ComputeBackendService`
                      resource.'
                    type: string
                  name:
                    description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            type: object
      defaultUrlRedirect:
        description: |-
          When none of the specified hostRules match, the request is redirected to a URL specified
          by defaultUrlRedirect. If defaultUrlRedirect is specified, defaultService or
          defaultRouteAction must not be set.
        properties:
          hostRedirect:
            description: |-
              The host that will be used in the redirect response instead of the one that was
              supplied in the request. The value must be between 1 and 255 characters.
            type: string
          httpsRedirect:
            description: |-
              If set to true, the URL scheme in the redirected request is set to https. If set to
              false, the URL scheme of the redirected request will remain the same as that of the
              request. This must only be set for UrlMaps used in TargetHttpProxys. Setting this
              true for TargetHttpsProxy is not permitted. The default is set to false.
            type: boolean
          pathRedirect:
            description: |-
              The path that will be used in the redirect response instead of the one that was
              supplied in the request. pathRedirect cannot be supplied together with
              prefixRedirect. Supply one alone or neither. If neither is supplied, the path of the
              original request will be used for the redirect. The value must be between 1 and 1024
              characters.
            type: string
          prefixRedirect:
            description: |-
              The prefix that replaces the prefixMatch specified in the HttpRouteRuleMatch,
              retaining the remaining portion of the URL before redirecting the request.
              prefixRedirect cannot be supplied together with pathRedirect. Supply one alone or
              neither. If neither is supplied, the path of the original request will be used for
              the redirect. The value must be between 1 and 1024 characters.
            type: string
          redirectResponseCode:
            description: |-
              The HTTP Status code to use for this RedirectAction. Supported values are:

              * MOVED_PERMANENTLY_DEFAULT, which is the default value and corresponds to 301.

              * FOUND, which corresponds to 302.

              * SEE_OTHER which corresponds to 303.

              * TEMPORARY_REDIRECT, which corresponds to 307. In this case, the request method
              will be retained.

              * PERMANENT_REDIRECT, which corresponds to 308. In this case,
              the request method will be retained. Possible values: ["FOUND", "MOVED_PERMANENTLY_DEFAULT", "PERMANENT_REDIRECT", "SEE_OTHER", "TEMPORARY_REDIRECT"].
            type: string
          stripQuery:
            description: |-
              If set to true, any accompanying query portion of the original URL is removed prior
              to redirecting the request. If set to false, the query portion of the original URL is
              retained.
               This field is required to ensure an empty block is not set. The normal default value is false.
            type: boolean
        required:
        - stripQuery
        type: object
      description:
        description: |-
          An optional description of this resource. Provide this property when
          you create the resource.
        type: string
      headerAction:
        description: |-
          Specifies changes to request and response headers that need to take effect for
          the selected backendService. The headerAction specified here take effect after
          headerAction specified under pathMatcher.
        properties:
          requestHeadersToAdd:
            description: |-
              Headers to add to a matching request prior to forwarding the request to the
              backendService.
            items:
              properties:
                headerName:
                  description: The name of the header.
                  type: string
                headerValue:
                  description: The value of the header to add.
                  type: string
                replace:
                  description: |-
                    If false, headerValue is appended to any values that already exist for the
                    header. If true, headerValue is set for the header, discarding any values that
                    were set for that header.
                  type: boolean
              required:
              - headerName
              - headerValue
              - replace
              type: object
            type: array
          requestHeadersToRemove:
            description: |-
              A list of header names for headers that need to be removed from the request
              prior to forwarding the request to the backendService.
            items:
              type: string
            type: array
          responseHeadersToAdd:
            description: Headers to add the response prior to sending the response
              back to the client.
            items:
              properties:
                headerName:
                  description: The name of the header.
                  type: string
                headerValue:
                  description: The value of the header to add.
                  type: string
                replace:
                  description: |-
                    If false, headerValue is appended to any values that already exist for the
                    header. If true, headerValue is set for the header, discarding any values that
                    were set for that header.
                  type: boolean
              required:
              - headerName
              - headerValue
              - replace
              type: object
            type: array
          responseHeadersToRemove:
            description: |-
              A list of header names for headers that need to be removed from the response
              prior to sending the response back to the client.
            items:
              type: string
            type: array
        type: object
      hostRule:
        description: The list of HostRules to use against the URL.
        items:
          properties:
            description:
              description: |-
                An optional description of this HostRule. Provide this property
                when you create the resource.
              type: string
            hosts:
              description: |-
                The list of host patterns to match. They must be valid
                hostnames, except * will match any string of ([a-z0-9-.]*). In
                that case, * must be the first character and must be followed in
                the pattern by either - or ..
              items:
                type: string
              type: array
            pathMatcher:
              description: |-
                The name of the PathMatcher to use to match the path portion of
                the URL if the hostRule matches the URL's host portion.
              type: string
          required:
          - hosts
          - pathMatcher
          type: object
        type: array
      location:
        description: 'Location represents the geographical location of the ComputeURLMap.
          Specify a region name or "global" for global resources. Reference: GCP definition
          of regions/zones (https://cloud.google.com/compute/docs/regions-zones/)'
        type: string
required:
- spec
type: object

```
