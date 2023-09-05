# openAPI schema part4 for accesscontextmanager.cnrm.cloud.google.com.v1beta1.AccessContextManagerServicePerimeter

## schema

```yaml
properties:
  spec:
    properties:
      status:
        properties:
          accessLevels:
            items:
              description: |-
                (Optional) A list of AccessLevel resource names that allow resources within
                the ServicePerimeter to be accessed from the internet. AccessLevels listed
                must be in the same policy as this ServicePerimeter.
                Referencing a nonexistent AccessLevel is a syntax error. If no
                AccessLevel names are listed, resources within the perimeter can
                only be accessed via GCP calls with request origins within the
                perimeter. For Service Perimeter Bridge, must be empty.
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
                  description: 'Allowed value: string of the format `{{parent}}/accessLevels/{{value}}`,
                    where {{value}} is the `name` field of an `AccessContextManagerAccessLevel`
                    resource.'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
            type: array
          egressPolicies:
            description: |-
              List of EgressPolicies to apply to the perimeter. A perimeter may
              have multiple EgressPolicies, each of which is evaluated separately.
              Access is granted if any EgressPolicy grants it. Must be empty for
              a perimeter bridge.
            items:
              properties:
                egressFrom:
                  description: Defines conditions on the source of a request causing
                    this 'EgressPolicy' to apply.
                  properties:
                    identities:
                      items:
                        description: |-
                          (Optional) A list of identities that are allowed access through this
                          EgressPolicy. Should be in the format of email address. The email
                          address should represent individual user or service account only.
                        oneOf:
                        - required:
                          - serviceAccountRef
                        - required:
                          - user
                        properties:
                          serviceAccountRef:
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
                                description: 'Allowed value: string of the format
                                  `serviceAccount:{{value}}`, where {{value}} is the
                                  `email` field of an `IAMServiceAccount` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          user:
                            type: string
                        type: object
                      type: array
                    identityType:
                      description: |-
                        Specifies the type of identities that are allowed access to outside the
                        perimeter. If left unspecified, then members of 'identities' field will
                        be allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"].
                      type: string
                  type: object
                egressTo:
                  description: |-
                    Defines the conditions on the 'ApiOperation' and destination resources that
                    cause this 'EgressPolicy' to apply.
                  properties:
                    externalResources:
                      description: |-
                        A list of external resources that are allowed to be accessed. A request
                        matches if it contains an external resource in this list (Example:
                        s3://bucket/path). Currently '*' is not allowed.
                      items:
                        type: string
                      type: array
                    operations:
                      description: |-
                        A list of 'ApiOperations' that this egress rule applies to. A request matches
                        if it contains an operation/service in this list.
                      items:
                        properties:
                          methodSelectors:
                            description: |-
                              API methods or permissions to allow. Method or permission must belong
                              to the service specified by 'serviceName' field. A single MethodSelector
                              entry with '*' specified for the 'method' field will allow all methods
                              AND permissions for the service specified in 'serviceName'.
                            items:
                              properties:
                                method:
                                  description: |-
                                    Value for 'method' should be a valid method name for the corresponding
                                    'serviceName' in 'ApiOperation'. If '*' used as value for method,
                                    then ALL methods and permissions are allowed.
                                  type: string
                                permission:
                                  description: |-
                                    Value for permission should be a valid Cloud IAM permission for the
                                    corresponding 'serviceName' in 'ApiOperation'.
                                  type: string
                              type: object
                            type: array
                          serviceName:
                            description: |-
                              The name of the API whose methods or permissions the 'IngressPolicy' or
                              'EgressPolicy' want to allow. A single 'ApiOperation' with serviceName
                              field set to '*' will allow all methods AND permissions for all services.
                            type: string
                        type: object
                      type: array
                    resources:
                      items:
                        description: |-
                          (Optional) A list of resources, currently only projects in the form
                          "projects/{project_number}". A request
                          matches if it contains a resource in this list.
                        properties:
                          projectRef:
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
                                description: 'Allowed value: string of the format
                                  `projects/{{value}}`, where {{value}} is the `number`
                                  field of a `Project` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                        type: object
                      type: array
                  type: object
              type: object
            type: array
          ingressPolicies:
            description: |-
              List of 'IngressPolicies' to apply to the perimeter. A perimeter may
              have multiple 'IngressPolicies', each of which is evaluated
              separately. Access is granted if any 'Ingress Policy' grants it.
              Must be empty for a perimeter bridge.
            items:
              properties:
                ingressFrom:
                  description: |-
                    Defines the conditions on the source of a request causing this 'IngressPolicy'
                    to apply.
                  properties:
                    identities:
                      items:
                        description: |-
                          (Optional) A list of identities that are allowed access through this
                          EgressPolicy. Should be in the format of email address. The email
                          address should represent individual user or service account only.
                        oneOf:
                        - required:
                          - serviceAccountRef
                        - required:
                          - user
                        properties:
                          serviceAccountRef:
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
                                description: 'Allowed value: string of the format
                                  `serviceAccount:{{value}}`, where {{value}} is the
                                  `email` field of an `IAMServiceAccount` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          user:
                            type: string
                        type: object
                      type: array
                    identityType:
                      description: |-
                        Specifies the type of identities that are allowed access from outside the
                        perimeter. If left unspecified, then members of 'identities' field will be
                        allowed access. Possible values: ["IDENTITY_TYPE_UNSPECIFIED", "ANY_IDENTITY", "ANY_USER_ACCOUNT", "ANY_SERVICE_ACCOUNT"].
                      type: string
                    sources:
                      description: Sources that this 'IngressPolicy' authorizes access
                        from.
                      items:
                        properties:
                          accessLevelRef:
                            description: |-
                              An AccessLevel resource name that allow resources within the
                              ServicePerimeters to be accessed from the internet. AccessLevels
                              listed must be in the same policy as this ServicePerimeter.
                              Referencing a nonexistent AccessLevel will cause an error. If no
                              AccessLevel names are listed, resources within the perimeter can
                              only be accessed via Google Cloud calls with request origins within
                              the perimeter.
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
                                description: 'Allowed value: string of the format
                                  `{{parent}}/accessLevels/{{value}}`, where {{value}}
                                  is the `name` field of an `AccessContextManagerAccessLevel`
                                  resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          projectRef:
                            description: |-
                              (Optional) A Google Cloud resource that is allowed to ingress the
                              perimeter. Requests from these resources will be allowed to access
                              perimeter data. Currently only projects are allowed. Format
                              "projects/{project_number}" The project may be in any Google Cloud
                              organization, not just the organization that the perimeter is defined in.
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
                                description: 'Allowed value: string of the format
                                  `projects/{{value}}`, where {{value}} is the `number`
                                  field of a `Project` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                        type: object
                      type: array
                  type: object
                ingressTo:
                  description: |-
                    Defines the conditions on the 'ApiOperation' and request destination that cause
                    this 'IngressPolicy' to apply.
                  properties:
                    operations:
                      description: |-
                        A list of 'ApiOperations' the sources specified in corresponding 'IngressFrom'
                        are allowed to perform in this 'ServicePerimeter'.
                      items:
                        properties:
                          methodSelectors:
                            description: |-
                              API methods or permissions to allow. Method or permission must belong to
                              the service specified by serviceName field. A single 'MethodSelector' entry
                              with '*' specified for the method field will allow all methods AND
                              permissions for the service specified in 'serviceName'.
                            items:
                              properties:
                                method:
                                  description: |-
                                    Value for method should be a valid method name for the corresponding
                                    serviceName in 'ApiOperation'. If '*' used as value for 'method', then
                                    ALL methods and permissions are allowed.
                                  type: string
                                permission:
                                  description: |-
                                    Value for permission should be a valid Cloud IAM permission for the
                                    corresponding 'serviceName' in 'ApiOperation'.
                                  type: string
                              type: object
                            type: array
                          serviceName:
                            description: |-
                              The name of the API whose methods or permissions the 'IngressPolicy' or
                              'EgressPolicy' want to allow. A single 'ApiOperation' with 'serviceName'
                              field set to '*' will allow all methods AND permissions for all services.
                            type: string
                        type: object
                      type: array
                    resources:
                      items:
                        description: |-
                          A list of resources, currently only projects in the form
                          "projects/{project_number}", protected by this ServicePerimeter
                          that are allowed to be accessed by sources defined in the
                          corresponding IngressFrom. A request matches if it contains a
                          resource in this list.
                        properties:
                          projectRef:
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
                                description: 'Allowed value: string of the format
                                  `projects/{{value}}`, where {{value}} is the `number`
                                  field of a `Project` resource.'
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                        type: object
                      type: array
                  type: object
              type: object
            type: array
          resources:
            items:
              description: |-
                (Optional) A list of GCP resources that are inside of the service perimeter.
                Currently only projects are allowed.
              properties:
                projectRef:
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
                      description: 'Allowed value: string of the format `projects/{{value}}`,
                        where {{value}} is the `number` field of a `Project` resource.'
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
              type: object
            type: array
          restrictedServices:
            description: |-
              GCP services that are subject to the Service Perimeter
              restrictions. Must contain a list of services. For example, if
              'storage.googleapis.com' is specified, access to the storage
              buckets inside the perimeter must meet the perimeter's access
              restrictions.
            items:
              type: string
            type: array
          status:
            description: |-
              ServicePerimeter configuration. Specifies sets of resources,
              restricted services and access levels that determine
              perimeter content and boundaries.
            type: object
required:
- spec
type: object

```
