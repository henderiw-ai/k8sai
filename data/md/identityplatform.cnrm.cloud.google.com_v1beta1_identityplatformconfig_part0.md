# openAPI schema part0 for identityplatform.cnrm.cloud.google.com.v1beta1.IdentityPlatformConfig

## schema

```yaml
properties:
  apiVersion:
    description: 'apiVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources'
    type: string
  kind:
    description: 'kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    properties:
      authorizedDomains:
        description: List of domains authorized for OAuth redirects
        items:
          type: string
        type: array
      blockingFunctions:
        description: Configuration related to blocking functions.
        properties:
          triggers:
            additionalProperties:
              properties:
                functionUriRef:
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
                        HTTP URI trigger for the Cloud Function.

                        Allowed value: The `httpsTrigger.url` field of a `CloudFunctionsFunction` resource.
                      type: string
                    name:
                      description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                      type: string
                    namespace:
                      description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                      type: string
                  type: object
                updateTime:
                  description: When the trigger was changed.
                  format: date-time
                  type: string
              type: object
            description: 'Map of Trigger to event type. Key should be one of the supported
              event types: "beforeCreate", "beforeSignIn"'
            type: object
        type: object
      client:
        description: Options related to how clients making requests on behalf of a
          project should be configured.
        properties:
          permissions:
            description: Configuration related to restricting a user's ability to
              affect their account.
            properties:
              disabledUserDeletion:
                description: When true, end users cannot delete their account on the
                  associated project through any of our API methods
                type: boolean
              disabledUserSignup:
                description: When true, end users cannot sign up for a new account
                  on the associated project through any of our API methods
                type: boolean
            type: object
        type: object
      mfa:
        description: Configuration for this project's multi-factor authentication,
          including whether it is active and what factors can be used for the second
          factor
        properties:
          state:
            description: 'Whether MultiFactor Authentication has been enabled for
              this project. Possible values: STATE_UNSPECIFIED, DISABLED, ENABLED,
              MANDATORY'
            type: string
        type: object
      monitoring:
        description: Configuration related to monitoring project activity.
        properties:
          requestLogging:
            description: Configuration for logging requests made to this project to
              Stackdriver Logging
            properties:
              enabled:
                description: Whether logging is enabled for this project or not.
                type: boolean
            type: object
        type: object
      multiTenant:
        description: Configuration related to multi-tenant functionality.
        properties:
          allowTenants:
            description: Whether this project can have tenants or not.
            type: boolean
          defaultTenantLocationRef:
            oneOf:
            - not:
                required:
                - external
              required:
              - name
              - kind
            - not:
                anyOf:
                - required:
                  - name
                - required:
                  - namespace
                - required:
                  - kind
              required:
              - external
            properties:
              external:
                description: |-
                  The default cloud parent org or folder that the tenant project should be created under. The parent resource name should be in the format of "<type>/<number>", such as "folders/123" or "organizations/456". If the value is not set, the tenant will be created under the same organization or folder as the agent project.

                  Allowed values:
                  * The Google Cloud resource name of a `Folder` resource (format: `folders/{{name}}`).
                  * The Google Cloud resource name of a Google Cloud Organization (format: `organizations/{{name}}`).
                type: string
              kind:
                description: 'Kind of the referent. Allowed values: Folder'
                type: string
              name:
                description: |-
                  [WARNING] Organization not yet supported in Config Connector, use 'external' field to reference existing resources.
                  Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
        type: object
      notification:
        description: Configuration related to sending notifications to users.
        properties:
          defaultLocale:
            description: Default locale used for email and SMS in IETF BCP 47 format.
            type: string
          sendEmail:
            description: Options for email sending.
            properties:
              callbackUri:
                description: action url in email template.
                type: string
              changeEmailTemplate:
                description: Email template for change email
                properties:
                  body:
                    description: Immutable. Email body
                    type: string
                  bodyFormat:
                    description: 'Email body format Possible values: BODY_FORMAT_UNSPECIFIED,
                      PLAIN_TEXT, HTML'
                    type: string
                  replyTo:
                    description: Reply-to address
                    type: string
                  senderDisplayName:
                    description: Sender display name
                    type: string
                  senderLocalPart:
                    description: Local part of From address
                    type: string
                  subject:
                    description: Subject of the email
                    type: string
                type: object
              dnsInfo:
                description: Information of custom domain DNS verification.
                properties:
                  useCustomDomain:
                    description: Whether to use custom domain.
                    type: boolean
                type: object
              method:
                description: 'The method used for sending an email. Possible values:
                  METHOD_UNSPECIFIED, DEFAULT, CUSTOM_SMTP'
                type: string
              resetPasswordTemplate:
                description: Email template for reset password
                properties:
                  body:
                    description: Email body
                    type: string
                  bodyFormat:
                    description: 'Email body format Possible values: BODY_FORMAT_UNSPECIFIED,
                      PLAIN_TEXT, HTML'
                    type: string
                  replyTo:
                    description: Reply-to address
                    type: string
                  senderDisplayName:
                    description: Sender display name
                    type: string
                  senderLocalPart:
                    description: Local part of From address
                    type: string
                  subject:
                    description: Subject of the email
                    type: string
                type: object
              revertSecondFactorAdditionTemplate:
                description: Email template for reverting second factor addition emails
                properties:
                  body:
                    description: Immutable. Email body
                    type: string
                  bodyFormat:
                    description: 'Email body format Possible values: BODY_FORMAT_UNSPECIFIED,
                      PLAIN_TEXT, HTML'
                    type: string
                  replyTo:
                    description: Reply-to address
                    type: string
                  senderDisplayName:
                    description: Sender display name
                    type: string
                  senderLocalPart:
                    description: Local part of From address
                    type: string
                  subject:
                    description: Subject of the email
                    type: string
                type: object
              smtp:
                description: Use a custom SMTP relay
                properties:
                  host:
                    description: SMTP relay host
                    type: string
                  password:
                    description: SMTP relay password
                    oneOf:
                    - not:
                        required:
                        - valueFrom
                      required:
                      - value
                    - not:
                        required:
                        - value
                      required:
                      - valueFrom
                    properties:
                      value:
                        description: Value of the field. Cannot be used if 'valueFrom'
                          is specified.
                        type: string
                      valueFrom:
                        description: Source for the field's value. Cannot be used
                          if 'value' is specified.
                        properties:
                          secretKeyRef:
                            description: Reference to a value with the given key in
                              the given Secret in the resource's namespace.
                            properties:
                              key:
                                description: Key that identifies the value to be extracted.
                                type: string
                              name:
                                description: Name of the Secret to extract a value
                                  from.
                                type: string
                            required:
                            - name
                            - key
                            type: object
                        type: object
                    type: object
                  port:
                    description: SMTP relay port
                    format: int64
                    type: integer
                  securityMode:
                    description: 'SMTP security mode. Possible values: SECURITY_MODE_UNSPECIFIED,
                      SSL, START_TLS'
                    type: string
                  senderEmail:
                    description: Sender email for the SMTP relay
                    type: string
                  username:
                    description: SMTP relay username
                    type: string
                type: object
              verifyEmailTemplate:
                description: Email template for verify email
                properties:
                  body:
                    description: Immutable. Email body
                    type: string
                  bodyFormat:
                    description: 'Email body format Possible values: BODY_FORMAT_UNSPECIFIED,
                      PLAIN_TEXT, HTML'
                    type: string
                  replyTo:
                    description: Reply-to address
                    type: string
                  senderDisplayName:
                    description: Sender display name
                    type: string
                  senderLocalPart:
                    description: Local part of From address
                    type: string
                  subject:
                    description: Subject of the email
                    type: string
                type: object
            type: object
          sendSms:
            description: Options for SMS sending.
            properties:
              useDeviceLocale:
                description: Whether to use the accept_language header for SMS.
                type: boolean
            type: object
        type: object
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
              The project of the resource

              Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      quota:
        description: Configuration related to quotas.
        properties:
          signUpQuotaConfig:
            description: Quota for the Signup endpoint, if overwritten. Signup quota
              is measured in sign ups per project per hour per IP.
            properties:
              quota:
                description: Corresponds to the 'refill_token_count' field in QuotaServer
                  config
                format: int64
                type: integer
              quotaDuration:
                description: How long this quota will be active for
                type: string
              startTime:
                description: When this quota will take affect
                format: date-time
                type: string
            type: object
        type: object
      signIn:
        description: Configuration related to local sign in methods.
        properties:
          allowDuplicateEmails:
            description: Whether to allow more than one account to have the same email.
            type: boolean
          anonymous:
            description: Configuration options related to authenticating an anonymous
              user.
            properties:
              enabled:
                description: Whether anonymous user auth is enabled for the project
                  or not.
                type: boolean
            type: object
          email:
            description: Configuration options related to authenticating a user by
              their email address.
            properties:
              enabled:
                description: Whether email auth is enabled for the project or not.
                type: boolean
              passwordRequired:
                description: Whether a password is required for email auth or not.
                  If true, both an email and password must be provided to sign in.
                  If false, a user may sign in via either email/password or email
                  link.
                type: boolean
            type: object
          phoneNumber:
            description: Configuration options related to authenticated a user by
              their phone number.
            properties:
              enabled:
                description: Whether phone number auth is enabled for the project
                  or not.
                type: boolean
              testPhoneNumbers:
                additionalProperties:
                  type: string
                description: A map of that can be used for phone auth testing.
                type: object
            type: object
        type: object
    required:
    - projectRef
    type: object
required:
- spec
type: object

```
