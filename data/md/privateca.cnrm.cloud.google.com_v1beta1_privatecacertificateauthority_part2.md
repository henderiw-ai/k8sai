# openAPI schema part2 for privateca.cnrm.cloud.google.com.v1beta1.PrivateCACertificateAuthority

## schema

```yaml
properties:
  status:
    properties:
      accessUrls:
        description: Output only. URLs for accessing content published by this CA,
          such as the CA certificate and CRLs.
        properties:
          caCertificateAccessUrl:
            description: The URL where this CertificateAuthority's CA certificate
              is published. This will only be set for CAs that have been activated.
            type: string
          crlAccessUrls:
            description: The URLs where this CertificateAuthority's CRLs are published.
              This will only be set for CAs that have been activated.
            items:
              type: string
            type: array
        type: object
      caCertificateDescriptions:
        description: Output only. A structured description of this CertificateAuthority's
          CA certificate and its issuers. Ordered as self-to-root.
        items:
          properties:
            aiaIssuingCertificateUrls:
              description: Describes lists of issuer CA certificate URLs that appear
                in the "Authority Information Access" extension in the certificate.
              items:
                type: string
              type: array
            authorityKeyId:
              description: Identifies the subject_key_id of the parent certificate,
                per https://tools.ietf.org/html/rfc5280#section-4.2.1.1
              properties:
                keyId:
                  description: Optional. The value of this KeyId encoded in lowercase
                    hexadecimal. This is most likely the 160 bit SHA-1 hash of the
                    public key.
                  type: string
              type: object
            certFingerprint:
              description: The hash of the x.509 certificate.
              properties:
                sha256Hash:
                  description: The SHA 256 hash, encoded in hexadecimal, of the DER
                    x509 certificate.
                  type: string
              type: object
            crlDistributionPoints:
              description: Describes a list of locations to obtain CRL information,
                i.e. the DistributionPoint.fullName described by https://tools.ietf.org/html/rfc5280#section-4.2.1.13
              items:
                type: string
              type: array
            publicKey:
              description: The public key that corresponds to an issued certificate.
              properties:
                format:
                  description: 'Required. The format of the public key. Possible values:
                    PEM'
                  type: string
                key:
                  description: Required. A public key. The padding and encoding must
                    match with the `KeyFormat` value specified for the `format` field.
                  type: string
              type: object
            subjectDescription:
              description: Describes some of the values in a certificate that are
                related to the subject and lifetime.
              properties:
                hexSerialNumber:
                  description: The serial number encoded in lowercase hexadecimal.
                  type: string
                lifetime:
                  description: For convenience, the actual lifetime of an issued certificate.
                  type: string
                notAfterTime:
                  description: The time after which the certificate is expired. Per
                    RFC 5280, the validity period for a certificate is the period
                    of time from not_before_time through not_after_time, inclusive.
                    Corresponds to 'not_before_time' + 'lifetime' - 1 second.
                  format: date-time
                  type: string
                notBeforeTime:
                  description: The time at which the certificate becomes valid.
                  format: date-time
                  type: string
                subject:
                  description: Contains distinguished name fields such as the common
                    name, location and organization.
                  properties:
                    commonName:
                      description: The "common name" of the subject.
                      type: string
                    countryCode:
                      description: The country code of the subject.
                      type: string
                    locality:
                      description: The locality or city of the subject.
                      type: string
                    organization:
                      description: The organization of the subject.
                      type: string
                    organizationalUnit:
                      description: The organizational_unit of the subject.
                      type: string
                    postalCode:
                      description: The postal code of the subject.
                      type: string
                    province:
                      description: The province, territory, or regional state of the
                        subject.
                      type: string
                    streetAddress:
                      description: The street address of the subject.
                      type: string
                  type: object
                subjectAltName:
                  description: The subject alternative name fields.
                  properties:
                    customSans:
                      description: Contains additional subject alternative name values.
                      items:
                        properties:
                          critical:
                            description: Optional. Indicates whether or not this extension
                              is critical (i.e., if the client does not know how to
                              handle this extension, the client should consider this
                              to be an error).
                            type: boolean
                          objectId:
                            description: Required. The OID for this X.509 extension.
                            properties:
                              objectIdPath:
                                description: Required. The parts of an OID path. The
                                  most significant parts of the path come first.
                                items:
                                  format: int64
                                  type: integer
                                type: array
                            type: object
                          value:
                            description: Required. The value of this X.509 extension.
                            type: string
                        type: object
                      type: array
                    dnsNames:
                      description: Contains only valid, fully-qualified host names.
                      items:
                        type: string
                      type: array
                    emailAddresses:
                      description: Contains only valid RFC 2822 E-mail addresses.
                      items:
                        type: string
                      type: array
                    ipAddresses:
                      description: Contains only valid 32-bit IPv4 addresses or RFC
                        4291 IPv6 addresses.
                      items:
                        type: string
                      type: array
                    uris:
                      description: Contains only valid RFC 3986 URIs.
                      items:
                        type: string
                      type: array
                  type: object
              type: object
            subjectKeyId:
              description: Provides a means of identifiying certificates that contain
                a particular public key, per https://tools.ietf.org/html/rfc5280#section-4.2.1.2.
              properties:
                keyId:
                  description: Optional. The value of this KeyId encoded in lowercase
                    hexadecimal. This is most likely the 160 bit SHA-1 hash of the
                    public key.
                  type: string
              type: object
            x509Description:
              description: Describes some of the technical X.509 fields in a certificate.
              properties:
                additionalExtensions:
                  description: Optional. Describes custom X.509 extensions.
                  items:
                    properties:
                      critical:
                        description: Optional. Indicates whether or not this extension
                          is critical (i.e., if the client does not know how to handle
                          this extension, the client should consider this to be an
                          error).
                        type: boolean
                      objectId:
                        description: Required. The OID for this X.509 extension.
                        properties:
                          objectIdPath:
                            description: Required. The parts of an OID path. The most
                              significant parts of the path come first.
                            items:
                              format: int64
                              type: integer
                            type: array
                        type: object
                      value:
                        description: Required. The value of this X.509 extension.
                        type: string
                    type: object
                  type: array
                aiaOcspServers:
                  description: Optional. Describes Online Certificate Status Protocol
                    (OCSP) endpoint addresses that appear in the "Authority Information
                    Access" extension in the certificate.
                  items:
                    type: string
                  type: array
                caOptions:
                  description: Optional. Describes options in this X509Parameters
                    that are relevant in a CA certificate.
                  properties:
                    isCa:
                      description: Optional. Refers to the "CA" X.509 extension, which
                        is a boolean value. When this value is missing, the extension
                        will be omitted from the CA certificate.
                      type: boolean
                    maxIssuerPathLength:
                      description: Optional. Refers to the path length restriction
                        X.509 extension. For a CA certificate, this value describes
                        the depth of subordinate CA certificates that are allowed.
                        If this value is less than 0, the request will fail. If this
                        value is missing, the max path length will be omitted from
                        the CA certificate.
                      format: int64
                      type: integer
                  type: object
                keyUsage:
                  description: Optional. Indicates the intended use for keys that
                    correspond to a certificate.
                  properties:
                    baseKeyUsage:
                      description: Describes high-level ways in which a key may be
                        used.
                      properties:
                        certSign:
                          description: The key may be used to sign certificates.
                          type: boolean
                        contentCommitment:
                          description: The key may be used for cryptographic commitments.
                            Note that this may also be referred to as "non-repudiation".
                          type: boolean
                        crlSign:
                          description: The key may be used sign certificate revocation
                            lists.
                          type: boolean
                        dataEncipherment:
                          description: The key may be used to encipher data.
                          type: boolean
                        decipherOnly:
                          description: The key may be used to decipher only.
                          type: boolean
                        digitalSignature:
                          description: The key may be used for digital signatures.
                          type: boolean
                        encipherOnly:
                          description: The key may be used to encipher only.
                          type: boolean
                        keyAgreement:
                          description: The key may be used in a key agreement protocol.
                          type: boolean
                        keyEncipherment:
                          description: The key may be used to encipher other keys.
                          type: boolean
                      type: object
                    extendedKeyUsage:
                      description: Detailed scenarios in which a key may be used.
                      properties:
                        clientAuth:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.2. Officially
                            described as "TLS WWW client authentication", though regularly
                            used for non-WWW TLS.
                          type: boolean
                        codeSigning:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.3. Officially
                            described as "Signing of downloadable executable code
                            client authentication".
                          type: boolean
                        emailProtection:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.4. Officially
                            described as "Email protection".
                          type: boolean
                        ocspSigning:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.9. Officially
                            described as "Signing OCSP responses".
                          type: boolean
                        serverAuth:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.1. Officially
                            described as "TLS WWW server authentication", though regularly
                            used for non-WWW TLS.
                          type: boolean
                        timeStamping:
                          description: Corresponds to OID 1.3.6.1.5.5.7.3.8. Officially
                            described as "Binding the hash of an object to a time".
                          type: boolean
                      type: object
                    unknownExtendedKeyUsages:
                      description: Used to describe extended key usages that are not
                        listed in the KeyUsage.ExtendedKeyUsageOptions message.
                      items:
                        properties:
                          objectIdPath:
                            description: Required. The parts of an OID path. The most
                              significant parts of the path come first.
                            items:
                              format: int64
                              type: integer
                            type: array
                        type: object
                      type: array
                  type: object
                policyIds:
                  description: Optional. Describes the X.509 certificate policy object
                    identifiers, per https://tools.ietf.org/html/rfc5280#section-4.2.1.4.
                  items:
                    properties:
                      objectIdPath:
                        description: Required. The parts of an OID path. The most
                          significant parts of the path come first.
                        items:
                          format: int64
                          type: integer
                        type: array
                    type: object
                  type: array
              type: object
          type: object
        type: array
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
      config:
        properties:
          publicKey:
            description: Optional. The public key that corresponds to this config.
              This is, for example, used when issuing Certificates, but not when creating
              a self-signed CertificateAuthority or CertificateAuthority CSR.
            properties:
              format:
                description: 'Required. The format of the public key. Possible values:
                  PEM'
                type: string
              key:
                description: Required. A public key. The padding and encoding must
                  match with the `KeyFormat` value specified for the `format` field.
                type: string
            type: object
          x509Config:
            properties:
              aiaOcspServers:
                description: Optional. Describes Online Certificate Status Protocol
                  (OCSP) endpoint addresses that appear in the "Authority Information
                  Access" extension in the certificate.
                items:
                  type: string
                type: array
            type: object
        type: object
      createTime:
        description: Output only. The time at which this CertificateAuthority was
          created.
        format: date-time
        type: string
      deleteTime:
        description: Output only. The time at which this CertificateAuthority was
          soft deleted, if it is in the DELETED state.
        format: date-time
        type: string
      expireTime:
        description: Output only. The time at which this CertificateAuthority will
          be permanently purged, if it is in the DELETED state.
        format: date-time
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      pemCaCertificates:
        description: Output only. This CertificateAuthority's certificate chain, including
          the current CertificateAuthority's certificate. Ordered such that the root
          issuer is the final element (consistent with RFC 5246). For a self-signed
          CA, this will only list the current CertificateAuthority's certificate.
        items:
          type: string
        type: array
      state:
        description: 'Output only. The State for this CertificateAuthority. Possible
          values: ENABLED, DISABLED, STAGED, AWAITING_USER_ACTIVATION, DELETED'
        type: string
      status:
        type: object
      subordinateConfig:
        description: Optional. If this is a subordinate CertificateAuthority, this
          field will be set with the subordinate configuration, which describes its
          issuers. This may be updated, but this CertificateAuthority must continue
          to validate.
        properties:
          certificateAuthority:
            description: Required. This can refer to a CertificateAuthority in the
              same project that was used to create a subordinate CertificateAuthority.
              This field is used for information and usability purposes only. The
              resource name is in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
            type: string
          pemIssuerChain:
            description: Required. Contains the PEM certificate chain for the issuers
              of this CertificateAuthority, but not pem certificate for this CA itself.
            properties:
              pemCertificates:
                description: Required. Expected to be in leaf-to-root order according
                  to RFC 5246.
                items:
                  type: string
                type: array
            type: object
        type: object
      tier:
        description: 'Output only. The CaPool.Tier of the CaPool that includes this
          CertificateAuthority. Possible values: ENTERPRISE, DEVOPS'
        type: string
      updateTime:
        description: Output only. The time at which this CertificateAuthority was
          last updated.
        format: date-time
        type: string
required:
- spec
type: object

```
