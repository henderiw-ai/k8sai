# openAPI schema part0 for privateca.cnrm.cloud.google.com.v1beta1.PrivateCACertificate

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
              The ca_pool for the resource

              Allowed value: The Google Cloud resource name of a `PrivateCACAPool` resource (format: `projects/{{project}}/locations/{{location}}/caPools/{{name}}`).
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      certificateAuthorityRef:
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
              The certificate authority for the resource

              Allowed value: The Google Cloud resource name of a `PrivateCACertificateAuthority` resource (format: `projects/{{project}}/locations/{{location}}/caPools/{{ca_pool}}/certificateAuthorities/{{name}}`).
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      certificateTemplateRef:
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
              Immutable. The resource name for a CertificateTemplate used to issue this certificate, in the format `projects/*/locations/*/certificateTemplates/*`. If this is specified, the caller must have the necessary permission to use this template. If this is omitted, no template will be used. This template must be in the same location as the Certificate.

              Allowed value: The `selfLink` field of a `PrivateCACertificateTemplate` resource.
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      config:
        description: Immutable. Immutable. A description of the certificate and key
          that does not require X.509 or ASN.1.
        properties:
          publicKey:
            description: Immutable. Optional. The public key that corresponds to this
              config. This is, for example, used when issuing Certificates, but not
              when creating a self-signed CertificateAuthority or CertificateAuthority
              CSR.
            properties:
              format:
                description: 'Immutable. Required. The format of the public key. Possible
                  values: KEY_FORMAT_UNSPECIFIED, PEM'
                type: string
              key:
                description: Immutable. Required. A public key. The padding and encoding
                  must match with the `KeyFormat` value specified for the `format`
                  field.
                type: string
            required:
            - format
            - key
            type: object
          subjectConfig:
            description: Immutable. Required. Specifies some of the values in a certificate
              that are related to the subject.
            properties:
              subject:
                description: Immutable. Required. Contains distinguished name fields
                  such as the common name, location and organization.
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
                    description: Immutable. The province, territory, or regional state
                      of the subject.
                    type: string
                  streetAddress:
                    description: Immutable. The street address of the subject.
                    type: string
                type: object
              subjectAltName:
                description: Immutable. Optional. The subject alternative name fields.
                properties:
                  dnsNames:
                    description: Immutable. Contains only valid, fully-qualified host
                      names.
                    items:
                      type: string
                    type: array
                  emailAddresses:
                    description: Immutable. Contains only valid RFC 2822 E-mail addresses.
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
                      description: Immutable. Optional. Indicates whether or not this
                        extension is critical (i.e., if the client does not know how
                        to handle this extension, the client should consider this
                        to be an error).
                      type: boolean
                    objectId:
                      description: Immutable. Required. The OID for this X.509 extension.
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
                    value:
                      description: Immutable. Required. The value of this X.509 extension.
                      type: string
                  required:
                  - objectId
                  - value
                  type: object
                type: array
              aiaOcspServers:
                description: Immutable. Optional. Describes Online Certificate Status
                  Protocol (OCSP) endpoint addresses that appear in the "Authority
                  Information Access" extension in the certificate.
                items:
                  type: string
                type: array
              caOptions:
                description: Immutable. Optional. Describes options in this X509Parameters
                  that are relevant in a CA certificate.
                properties:
                  isCa:
                    description: Immutable. Optional. When true, the "CA" in Basic
                      Constraints extension will be set to true.
                    type: boolean
                  maxIssuerPathLength:
                    description: Immutable. Optional. Refers to the "path length constraint"
                      in Basic Constraints extension. For a CA certificate, this value
                      describes the depth of subordinate CA certificates that are
                      allowed. If this value is less than 0, the request will fail.
                    format: int64
                    type: integer
                  nonCa:
                    description: Immutable. Optional. When true, the "CA" in Basic
                      Constraints extension will be set to false. If both `is_ca`
                      and `non_ca` are unset, the extension will be omitted from the
                      CA certificate.
                    type: boolean
                  zeroMaxIssuerPathLength:
                    description: Immutable. Optional. When true, the "path length
                      constraint" in Basic Constraints extension will be set to 0.
                      if both max_issuer_path_length and zero_max_issuer_path_length
                      are unset, the max path length will be omitted from the CA certificate.
                    type: boolean
                type: object
              keyUsage:
                description: Immutable. Optional. Indicates the intended use for keys
                  that correspond to a certificate.
                properties:
                  baseKeyUsage:
                    description: Immutable. Describes high-level ways in which a key
                      may be used.
                    properties:
                      certSign:
                        description: Immutable. The key may be used to sign certificates.
                        type: boolean
                      contentCommitment:
                        description: Immutable. The key may be used for cryptographic
                          commitments. Note that this may also be referred to as "non-repudiation".
                        type: boolean
                      crlSign:
                        description: Immutable. The key may be used sign certificate
                          revocation lists.
                        type: boolean
                      dataEncipherment:
                        description: Immutable. The key may be used to encipher data.
                        type: boolean
                      decipherOnly:
                        description: Immutable. The key may be used to decipher only.
                        type: boolean
                      digitalSignature:
                        description: Immutable. The key may be used for digital signatures.
                        type: boolean
                      encipherOnly:
                        description: Immutable. The key may be used to encipher only.
                        type: boolean
                      keyAgreement:
                        description: Immutable. The key may be used in a key agreement
                          protocol.
                        type: boolean
                      keyEncipherment:
                        description: Immutable. The key may be used to encipher other
                          keys.
                        type: boolean
                    type: object
                  extendedKeyUsage:
                    description: Immutable. Detailed scenarios in which a key may
                      be used.
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
                          Officially described as "Binding the hash of an object to
                          a time".
                        type: boolean
                    type: object
                  unknownExtendedKeyUsages:
                    description: Immutable. Used to describe extended key usages that
                      are not listed in the KeyUsage.ExtendedKeyUsageOptions message.
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
      lifetime:
        description: Immutable. Required. Immutable. The desired lifetime of a certificate.
          Used to create the "not_before_time" and "not_after_time" fields inside
          an X.509 certificate. Note that the lifetime may be truncated if it would
          extend past the life of any certificate authority in the issuing chain.
        type: string
      location:
        description: Immutable. The location for the resource
        type: string
      pemCsr:
        description: Immutable. Immutable. A pem-encoded X.509 certificate signing
          request (CSR).
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
      subjectMode:
        description: 'Immutable. Immutable. Specifies how the Certificate''s identity
          fields are to be decided. If this is omitted, the `DEFAULT` subject mode
          will be used. Possible values: SUBJECT_REQUEST_MODE_UNSPECIFIED, DEFAULT,
          REFLECTED_SPIFFE'
        type: string
    required:
    - caPoolRef
    - lifetime
    - location
    - projectRef
    type: object
required:
- spec
type: object

```
