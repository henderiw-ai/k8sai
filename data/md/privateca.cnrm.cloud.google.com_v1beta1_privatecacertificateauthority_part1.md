# openAPI schema part1 for privateca.cnrm.cloud.google.com.v1beta1.PrivateCACertificateAuthority

## schema

```yaml
properties:
  spec:
    properties:
      spec:
        properties:
          caPoolRef:
            description: Immutable.
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
                  The caPool for the resource

                  Allowed value: The Google Cloud resource name of a `PrivateCACAPool` resource (format: `projects/{{project}}/locations/{{location}}/caPools/{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          config:
            description: Immutable. Required. Immutable. The config used to create
              a self-signed X.509 certificate or CSR.
            properties:
              subjectConfig:
                description: Immutable. Required. Specifies some of the values in
                  a certificate that are related to the subject.
                properties:
                  subject:
                    description: Immutable. Required. Contains distinguished name
                      fields such as the common name, location and organization.
                    properties:
                      commonName:
                        description: Immutable. The "common name" of the subject.
                        type: string
                      countryCode:
                        description: Immutable. The country code of the subject.
                        type: string
                      locality:
                        description: Immutable. The locality or city of the subject.
                        type: string
                      organization:
                        description: Immutable. The organization of the subject.
                        type: string
                      organizationalUnit:
                        description: Immutable. The organizational_unit of the subject.
                        type: string
                      postalCode:
                        description: Immutable. The postal code of the subject.
                        type: string
                      province:
                        description: Immutable. The province, territory, or regional
                          state of the subject.
                        type: string
                      streetAddress:
                        description: Immutable. The street address of the subject.
                        type: string
                    type: object
                  subjectAltName:
                    description: Immutable. Optional. The subject alternative name
                      fields.
                    properties:
                      customSans:
                        description: Immutable. Contains additional subject alternative
                          name values.
                        items:
                          properties:
                            critical:
                              description: Immutable. Optional. Indicates whether
                                or not this extension is critical (i.e., if the client
                                does not know how to handle this extension, the client
                                should consider this to be an error).
                              type: boolean
                            objectId:
                              description: Immutable. Required. The OID for this X.509
                                extension.
                              properties:
                                objectIdPath:
                                  description: Immutable. Required. The parts of an
                                    OID path. The most significant parts of the path
                                    come first.
                                  items:
                                    format: int64
                                    type: integer
                                  type: array
                              required:
                              - objectIdPath
                              type: object
                            value:
                              description: Immutable. Required. The value of this
                                X.509 extension.
                              type: string
                          required:
                          - objectId
                          - value
                          type: object
                        type: array
                      dnsNames:
                        description: Immutable. Contains only valid, fully-qualified
                          host names.
                        items:
                          type: string
                        type: array
                      emailAddresses:
                        description: Immutable. Contains only valid RFC 2822 E-mail
                          addresses.
                        items:
                          type: string
                        type: array
                      ipAddresses:
                        description: Immutable. Contains only valid 32-bit IPv4 addresses
                          or RFC 4291 IPv6 addresses.
                        items:
                          type: string
                        type: array
                      uris:
                        description: Immutable. Contains only valid RFC 3986 URIs.
                        items:
                          type: string
                        type: array
                    type: object
                required:
                - subject
                type: object
              x509Config:
                description: Immutable. Required. Describes how some of the technical
                  X.509 fields in a certificate should be populated.
                properties:
                  additionalExtensions:
                    description: Immutable. Optional. Describes custom X.509 extensions.
                    items:
                      properties:
                        critical:
                          description: Immutable. Optional. Indicates whether or not
                            this extension is critical (i.e., if the client does not
                            know how to handle this extension, the client should consider
                            this to be an error).
                          type: boolean
                        objectId:
                          description: Immutable. Required. The OID for this X.509
                            extension.
                          properties:
                            objectIdPath:
                              description: Immutable. Required. The parts of an OID
                                path. The most significant parts of the path come
                                first.
                              items:
                                format: int64
                                type: integer
                              type: array
                          required:
                          - objectIdPath
                          type: object
                        value:
                          description: Immutable. Required. The value of this X.509
                            extension.
                          type: string
                      required:
                      - objectId
                      - value
                      type: object
                    type: array
                  caOptions:
                    description: Immutable. Optional. Describes options in this X509Parameters
                      that are relevant in a CA certificate.
                    properties:
                      isCa:
                        description: Immutable. Optional. Refers to the "CA" X.509
                          extension, which is a boolean value. When this value is
                          missing, the extension will be omitted from the CA certificate.
                        type: boolean
                      maxIssuerPathLength:
                        description: Immutable. Optional. Refers to the path length
                          restriction X.509 extension. For a CA certificate, this
                          value describes the depth of subordinate CA certificates
                          that are allowed. If this value is less than 0, the request
                          will fail. If this value is missing, the max path length
                          will be omitted from the CA certificate.
                        format: int64
                        type: integer
                      zeroMaxIssuerPathLength:
                        description: Immutable. Optional. When true, the "path length
                          constraint" in Basic Constraints extension will be set to
                          0. if both max_issuer_path_length and zero_max_issuer_path_length
                          are unset, the max path length will be omitted from the
                          CA certificate.
                        type: boolean
                    type: object
                  keyUsage:
                    description: Immutable. Optional. Indicates the intended use for
                      keys that correspond to a certificate.
                    properties:
                      baseKeyUsage:
                        description: Immutable. Describes high-level ways in which
                          a key may be used.
                        properties:
                          certSign:
                            description: Immutable. The key may be used to sign certificates.
                            type: boolean
                          contentCommitment:
                            description: Immutable. The key may be used for cryptographic
                              commitments. Note that this may also be referred to
                              as "non-repudiation".
                            type: boolean
                          crlSign:
                            description: Immutable. The key may be used sign certificate
                              revocation lists.
                            type: boolean
                          dataEncipherment:
                            description: Immutable. The key may be used to encipher
                              data.
                            type: boolean
                          decipherOnly:
                            description: Immutable. The key may be used to decipher
                              only.
                            type: boolean
                          digitalSignature:
                            description: Immutable. The key may be used for digital
                              signatures.
                            type: boolean
                          encipherOnly:
                            description: Immutable. The key may be used to encipher
                              only.
                            type: boolean
                          keyAgreement:
                            description: Immutable. The key may be used in a key agreement
                              protocol.
                            type: boolean
                          keyEncipherment:
                            description: Immutable. The key may be used to encipher
                              other keys.
                            type: boolean
                        type: object
                      extendedKeyUsage:
                        description: Immutable. Detailed scenarios in which a key
                          may be used.
                        properties:
                          clientAuth:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.2.
                              Officially described as "TLS WWW client authentication",
                              though regularly used for non-WWW TLS.
                            type: boolean
                          codeSigning:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.3.
                              Officially described as "Signing of downloadable executable
                              code client authentication".
                            type: boolean
                          emailProtection:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.4.
                              Officially described as "Email protection".
                            type: boolean
                          ocspSigning:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.9.
                              Officially described as "Signing OCSP responses".
                            type: boolean
                          serverAuth:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.1.
                              Officially described as "TLS WWW server authentication",
                              though regularly used for non-WWW TLS.
                            type: boolean
                          timeStamping:
                            description: Immutable. Corresponds to OID 1.3.6.1.5.5.7.3.8.
                              Officially described as "Binding the hash of an object
                              to a time".
                            type: boolean
                        type: object
                      unknownExtendedKeyUsages:
                        description: Immutable. Used to describe extended key usages
                          that are not listed in the KeyUsage.ExtendedKeyUsageOptions
                          message.
                        items:
                          properties:
                            objectIdPath:
                              description: Immutable. Required. The parts of an OID
                                path. The most significant parts of the path come
                                first.
                              items:
                                format: int64
                                type: integer
                              type: array
                          required:
                          - objectIdPath
                          type: object
                        type: array
                    type: object
                  policyIds:
                    description: Immutable. Optional. Describes the X.509 certificate
                      policy object identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
                    items:
                      properties:
                        objectIdPath:
                          description: Immutable. Required. The parts of an OID path.
                            The most significant parts of the path come first.
                          items:
                            format: int64
                            type: integer
                          type: array
                      required:
                      - objectIdPath
                      type: object
                    type: array
                type: object
            required:
            - subjectConfig
            - x509Config
            type: object
          gcsBucketRef:
            description: Immutable.
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
                  Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.

                  Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
          keySpec:
            description: Immutable. Required. Immutable. Used when issuing certificates
              for this CertificateAuthority. If this CertificateAuthority is a self-signed
              CertificateAuthority, this key is also used to sign the self-signed
              CA certificate. Otherwise, it is used to sign a CSR.
            properties:
              algorithm:
                description: 'Immutable. The algorithm to use for creating a managed
                  Cloud KMS key for a for a simplified experience. All managed keys
                  will be have their ProtectionLevel as `HSM`. Possible values: RSA_PSS_2048_SHA256,
                  RSA_PSS_3072_SHA256, RSA_PSS_4096_SHA256, RSA_PKCS1_2048_SHA256,
                  RSA_PKCS1_3072_SHA256, RSA_PKCS1_4096_SHA256, EC_P256_SHA256, EC_P384_SHA384'
                type: string
              cloudKmsKeyVersionRef:
                description: Immutable.
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
                    description: The resource name for an existing Cloud KMS CryptoKeyVersion
                      in the format `projects/*/locations/*/keyRings/*/cryptoKeys/*/cryptoKeyVersions/*`.
                      This option enables full flexibility in the key's capabilities
                      and properties.
                    type: string
                  name:
                    description: |-
                      [WARNING] KMSCryptoKeyVersion not yet supported in Config Connector, use 'external' field to reference existing resources.
                      Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
                    type: string
                  namespace:
                    description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                    type: string
                type: object
            type: object
          lifetime:
            description: Immutable. Required. The desired lifetime of the CA certificate.
              Used to create the "not_before_time" and "not_after_time" fields inside
              an X.509 certificate.
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
              and acquisition. When unset, the value of `metadata.name` is used as
              the default.
            type: string
          type:
            description: 'Immutable. Required. Immutable. The Type of this CertificateAuthority.
              Possible values: SELF_SIGNED, SUBORDINATE'
            type: string
        required:
        - caPoolRef
        - config
        - keySpec
        - lifetime
        - location
        - projectRef
        - type
        type: object
required:
- spec
type: object

```
