# openAPI schema part1 for compute.cnrm.cloud.google.com.v1beta1.ComputeSecurityPolicy

## schema

```yaml
properties:
  spec:
    properties:
      adaptiveProtectionConfig:
        description: Adaptive Protection Config of this security policy.
        properties:
          autoDeployConfig:
            description: Auto Deploy Config of this security policy.
            properties:
              confidenceThreshold:
                description: Rules are only automatically deployed for alerts on potential
                  attacks with confidence scores greater than this threshold.
                type: number
              expirationSec:
                description: Google Cloud Armor stops applying the action in the automatically
                  deployed rule to an identified attacker after this duration. The
                  rule continues to operate against new requests.
                type: integer
              impactedBaselineThreshold:
                description: Rules are only automatically deployed when the estimated
                  impact to baseline traffic from the suggested mitigation is below
                  this threshold.
                type: number
              loadThreshold:
                description: Identifies new attackers only when the load to the backend
                  service that is under attack exceeds this threshold.
                type: number
            type: object
          layer7DdosDefenseConfig:
            description: Layer 7 DDoS Defense Config of this security policy.
            properties:
              enable:
                description: If set to true, enables CAAP for L7 DDoS detection.
                type: boolean
              ruleVisibility:
                description: 'Rule visibility. Supported values include: "STANDARD",
                  "PREMIUM".'
                type: string
            type: object
        type: object
      advancedOptionsConfig:
        description: Advanced Options Config of this security policy.
        properties:
          jsonCustomConfig:
            description: Custom configuration to apply the JSON parsing. Only applicable
              when JSON parsing is set to STANDARD.
            properties:
              contentTypes:
                description: A list of custom Content-Type header values to apply
                  the JSON parsing.
                items:
                  type: string
                type: array
            required:
            - contentTypes
            type: object
          jsonParsing:
            description: 'JSON body parsing. Supported values include: "DISABLED",
              "STANDARD".'
            type: string
          logLevel:
            description: 'Logging level. Supported values include: "NORMAL", "VERBOSE".'
            type: string
        type: object
      description:
        description: An optional description of this security policy. Max size is
          2048.
        type: string
      recaptchaOptionsConfig:
        description: reCAPTCHA configuration options to be applied for the security
          policy.
        properties:
          redirectSiteKeyRef:
            description: |-
              A field to supply a reCAPTCHA site key to be used for all the rules
              using the redirect action with the type of GOOGLE_RECAPTCHA under
              the security policy. The specified site key needs to be created from
              the reCAPTCHA API. The user is responsible for the validity of the
              specified site key. If not specified, a Google-managed site key is
              used.
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
                description: 'Allowed value: The `name` field of a `RecaptchaEnterpriseKey`
                  resource.'
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
        required:
        - redirectSiteKeyRef
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      spec:
        type: object
type: object

```
