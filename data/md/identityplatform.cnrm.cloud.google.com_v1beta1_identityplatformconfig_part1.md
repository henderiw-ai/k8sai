# openAPI schema part1 for identityplatform.cnrm.cloud.google.com.v1beta1.IdentityPlatformConfig

## schema

```yaml
properties:
  status:
    properties:
      status:
        properties:
          client:
            properties:
              apiKey:
                description: Output only. API key that can be used when making requests
                  for this project.
                type: string
              firebaseSubdomain:
                description: Output only. Firebase subdomain.
                type: string
            type: object
          conditions:
            description: Conditions represent the latest available observation of
              the resource's current state.
            items:
              properties:
                lastTransitionTime:
                  description: Last time the condition transitioned from one status
                    to another.
                  type: string
                message:
                  description: Human-readable message indicating details about last
                    transition.
                  type: string
                reason:
                  description: Unique, one-word, CamelCase reason for the condition's
                    last transition.
                  type: string
                status:
                  description: Status is the status of the condition. Can be True,
                    False, Unknown.
                  type: string
                type:
                  description: Type is the type of the condition.
                  type: string
              type: object
            type: array
          notification:
            properties:
              sendEmail:
                properties:
                  changeEmailTemplate:
                    properties:
                      customized:
                        description: Output only. Whether the body or subject of the
                          email is customized.
                        type: boolean
                    type: object
                  dnsInfo:
                    properties:
                      customDomain:
                        description: Output only. The applied verified custom domain.
                        type: string
                      customDomainState:
                        description: 'Output only. The current verification state
                          of the custom domain. The custom domain will only be used
                          once the domain verification is successful. Possible values:
                          VERIFICATION_STATE_UNSPECIFIED, NOT_STARTED, IN_PROGRESS,
                          FAILED, SUCCEEDED'
                        type: string
                      domainVerificationRequestTime:
                        description: Output only. The timestamp of initial request
                          for the current domain verification.
                        format: date-time
                        type: string
                      pendingCustomDomain:
                        description: Output only. The custom domain that's to be verified.
                        type: string
                    type: object
                  resetPasswordTemplate:
                    properties:
                      customized:
                        description: Output only. Whether the body or subject of the
                          email is customized.
                        type: boolean
                    type: object
                  revertSecondFactorAdditionTemplate:
                    properties:
                      customized:
                        description: Output only. Whether the body or subject of the
                          email is customized.
                        type: boolean
                    type: object
                  verifyEmailTemplate:
                    properties:
                      customized:
                        description: Output only. Whether the body or subject of the
                          email is customized.
                        type: boolean
                    type: object
                type: object
              sendSms:
                properties:
                  smsTemplate:
                    description: Output only. The template to use when sending an
                      SMS.
                    properties:
                      content:
                        description: 'Output only. The SMS''s content. Can contain
                          the following placeholders which will be replaced with the
                          appropriate values: %APP_NAME% - For Android or iOS apps,
                          the app''s display name. For web apps, the domain hosting
                          the application. %LOGIN_CODE% - The OOB code being sent
                          in the SMS.'
                        type: string
                    type: object
                type: object
            type: object
          observedGeneration:
            description: ObservedGeneration is the generation of the resource that
              was most recently observed by the Config Connector controller. If this
              is equal to metadata.generation, then that means that the current reported
              status reflects the most recent desired state of the resource.
            type: integer
          signIn:
            properties:
              email:
                properties:
                  hashConfig:
                    description: Output only. Hash config information.
                    properties:
                      algorithm:
                        description: 'Output only. Different password hash algorithms
                          used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED,
                          HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5,
                          HMAC_SHA512, SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512,
                          STANDARD_SCRYPT'
                        type: string
                      memoryCost:
                        description: Output only. Memory cost for hash calculation.
                          Used by scrypt and other similar password derivation algorithms.
                          See https://tools.ietf.org/html/rfc7914 for explanation
                          of field.
                        format: int64
                        type: integer
                      rounds:
                        description: Output only. How many rounds for hash calculation.
                          Used by scrypt and other similar password derivation algorithms.
                        format: int64
                        type: integer
                      saltSeparator:
                        description: Output only. Non-printable character to be inserted
                          between the salt and plain text password in base64.
                        type: string
                      signerKey:
                        description: Output only. Signer key in base64.
                        type: string
                    type: object
                type: object
              hashConfig:
                description: Output only. Hash config information.
                properties:
                  algorithm:
                    description: 'Output only. Different password hash algorithms
                      used in Identity Toolkit. Possible values: HASH_ALGORITHM_UNSPECIFIED,
                      HMAC_SHA256, HMAC_SHA1, HMAC_MD5, SCRYPT, PBKDF_SHA1, MD5, HMAC_SHA512,
                      SHA1, BCRYPT, PBKDF2_SHA256, SHA256, SHA512, STANDARD_SCRYPT'
                    type: string
                  memoryCost:
                    description: Output only. Memory cost for hash calculation. Used
                      by scrypt and other similar password derivation algorithms.
                      See https://tools.ietf.org/html/rfc7914 for explanation of field.
                    format: int64
                    type: integer
                  rounds:
                    description: Output only. How many rounds for hash calculation.
                      Used by scrypt and other similar password derivation algorithms.
                    format: int64
                    type: integer
                  saltSeparator:
                    description: Output only. Non-printable character to be inserted
                      between the salt and plain text password in base64.
                    type: string
                  signerKey:
                    description: Output only. Signer key in base64.
                    type: string
                type: object
            type: object
          subtype:
            description: 'Output only. The subtype of this config. Possible values:
              SUBTYPE_UNSPECIFIED, IDENTITY_PLATFORM, FIREBASE_AUTH'
            type: string
        type: object
required:
- spec
type: object

```
