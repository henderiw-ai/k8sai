# openAPI schema part0 for certificatemanager.cnrm.cloud.google.com.v1alpha1.CertificateManagerCertificate

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
      description:
        description: A human-readable description of the resource.
        type: string
      location:
        description: The Certificate Manager location. If not specified, "global"
          is used.
        type: string
      managed:
        description: |-
          Immutable. Configuration and state of a Managed Certificate.
          Certificate Manager provisions and renews Managed Certificates
          automatically, for as long as it's authorized to do so.
        properties:
          authorizationAttemptInfo:
            description: |-
              Detailed state of the latest authorization attempt for each domain
              specified for this Managed Certificate.
            items:
              properties:
                details:
                  description: |-
                    Human readable explanation for reaching the state. Provided to help
                    address the configuration issues.
                    Not guaranteed to be stable. For programmatic access use 'failure_reason' field.
                  type: string
                domain:
                  description: Domain name of the authorization attempt.
                  type: string
                failureReason:
                  description: Reason for failure of the authorization attempt for
                    the domain.
                  type: string
                state:
                  description: State of the domain for managed certificate issuance.
                  type: string
              type: object
            type: array
          dnsAuthorizations:
            description: Immutable. Authorizations that will be used for performing
              domain authorization. Either issuanceConfig or dnsAuthorizations should
              be specificed, but not both.
            items:
              type: string
            type: array
          domains:
            description: |-
              Immutable. The domains for which a managed SSL certificate will be generated.
              Wildcard domains are only supported with DNS challenge resolution.
            items:
              type: string
            type: array
          issuanceConfig:
            description: |-
              Immutable. The resource name for a CertificateIssuanceConfig used to configure private PKI certificates in the format projects/*/locations/*/certificateIssuanceConfigs/*.
              If this field is not set, the certificates will instead be publicly signed as documented at https://cloud.google.com/load-balancing/docs/ssl-certificates/google-managed-certs#caa.
              Either issuanceConfig or dnsAuthorizations should be specificed, but not both.
            type: string
          provisioningIssue:
            description: Information about issues with provisioning this Managed Certificate.
            items:
              properties:
                details:
                  description: |-
                    Human readable explanation about the issue. Provided to help address
                    the configuration issues.
                    Not guaranteed to be stable. For programmatic access use 'reason' field.
                  type: string
                reason:
                  description: Reason for provisioning failures.
                  type: string
              type: object
            type: array
          state:
            description: A state of this Managed Certificate.
            type: string
        type: object
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
      scope:
        description: |-
          Immutable. The scope of the certificate.

          DEFAULT: Certificates with default scope are served from core Google data centers.
          If unsure, choose this option.

          EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
          served from non-core Google data centers.
          Currently allowed only for managed certificates.
        type: string
      selfManaged:
        description: |-
          Immutable. Certificate data for a SelfManaged Certificate.
          SelfManaged Certificates are uploaded by the user. Updating such
          certificates before they expire remains the user's responsibility.
        properties:
          certificatePem:
            description: |-
              DEPRECATED. Deprecated in favor of `pem_certificate`. Immutable. **Deprecated** The certificate chain in PEM-encoded form.

              Leaf certificate comes first, followed by intermediate ones if any.
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
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - name
                    - key
                    type: object
                type: object
            type: object
          pemCertificate:
            description: |-
              Immutable. The certificate chain in PEM-encoded form.

              Leaf certificate comes first, followed by intermediate ones if any.
            type: string
          pemPrivateKey:
            description: Immutable. The private key of the leaf certificate in PEM-encoded
              form.
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
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - name
                    - key
                    type: object
                type: object
            type: object
          privateKeyPem:
            description: DEPRECATED. Deprecated in favor of `pem_private_key`. Immutable.
              **Deprecated** The private key of the leaf certificate in PEM-encoded
              form.
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
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - name
                    - key
                    type: object
                type: object
            type: object
        type: object
    required:
    - projectRef
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
