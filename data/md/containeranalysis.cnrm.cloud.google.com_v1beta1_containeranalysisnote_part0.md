# openAPI schema part0 for containeranalysis.cnrm.cloud.google.com.v1beta1.ContainerAnalysisNote

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
      attestation:
        description: A note describing an attestation role.
        properties:
          hint:
            description: Hint hints at the purpose of the attestation authority.
            properties:
              humanReadableName:
                description: Required. The human readable name of this attestation
                  authority, for example "qa".
                type: string
            required:
            - humanReadableName
            type: object
        type: object
      build:
        description: A note describing build provenance for a verifiable build.
        properties:
          builderVersion:
            description: Required. Immutable. Version of the builder which produced
              this build.
            type: string
        required:
        - builderVersion
        type: object
      deployment:
        description: A note describing something that can be deployed.
        properties:
          resourceUri:
            description: Required. Resource URI for the artifact being deployed.
            items:
              type: string
            type: array
        required:
        - resourceUri
        type: object
      discovery:
        description: A note describing the initial analysis of a resource.
        properties:
          analysisKind:
            description: 'The kind of analysis that is handled by this discovery.
              Possible values: NOTE_KIND_UNSPECIFIED, VULNERABILITY, BUILD, IMAGE,
              PACKAGE, DEPLOYMENT, DISCOVERY, ATTESTATION, UPGRADE'
            type: string
        required:
        - analysisKind
        type: object
      expirationTime:
        description: Time of expiration for this note. Empty if note does not expire.
        format: date-time
        type: string
      image:
        description: A note describing a base image.
        properties:
          fingerprint:
            description: Required. Immutable. The fingerprint of the base image.
            properties:
              v1Name:
                description: Required. The layer ID of the final layer in the Docker
                  image's v1 representation.
                type: string
              v2Blob:
                description: Required. The ordered list of v2 blobs that represent
                  a given image.
                items:
                  type: string
                type: array
            required:
            - v1Name
            - v2Blob
            type: object
          resourceUrl:
            description: Required. Immutable. The resource_url for the resource representing
              the basis of associated occurrence images.
            type: string
        required:
        - fingerprint
        - resourceUrl
        type: object
      longDescription:
        description: A detailed description of this note.
        type: string
      package:
        description: Required for non-Windows OS. The package this Upgrade is for.
        properties:
          distribution:
            description: The various channels by which a package is distributed.
            items:
              properties:
                architecture:
                  description: 'The CPU architecture for which packages in this distribution
                    channel were built Possible values: ARCHITECTURE_UNSPECIFIED,
                    X86, X64'
                  type: string
                cpeUri:
                  description: The cpe_uri in [cpe format](https://cpe.mitre.org/specification/)
                    denoting the package manager version distributing a package.
                  type: string
                description:
                  description: The distribution channel-specific description of this
                    package.
                  type: string
                latestVersion:
                  description: The latest available version of this package in this
                    distribution channel.
                  properties:
                    epoch:
                      description: Used to correct mistakes in the version numbering
                        scheme.
                      format: int64
                      type: integer
                    fullName:
                      description: Human readable version string. This string is of
                        the form :- and is only set when kind is NORMAL.
                      type: string
                    kind:
                      description: 'Distinguish between sentinel MIN/MAX versions
                        and normal versions. If kind is not NORMAL, then the other
                        fields are ignored. Possible values: VERSION_KIND_UNSPECIFIED,
                        NORMAL, MINIMUM, MAXIMUM'
                      type: string
                    name:
                      description: The main part of the version name.
                      type: string
                    revision:
                      description: The iteration of the package build from the above
                        version.
                      type: string
                  required:
                  - kind
                  type: object
                maintainer:
                  description: A freeform string denoting the maintainer of this package.
                  type: string
                url:
                  description: The distribution channel-specific homepage for this
                    package.
                  type: string
              required:
              - cpeUri
              type: object
            type: array
          name:
            description: The name of the package.
            type: string
        required:
        - name
        type: object
      relatedNoteNames:
        items:
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
              description: 'Allowed value: The Google Cloud resource name of a `ContainerAnalysisNote`
                resource (format: `projects/{{project}}/notes/{{name}}`).'
              type: string
            name:
              description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
              type: string
            namespace:
              description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
              type: string
          type: object
        type: array
      relatedUrl:
        description: URLs associated with this note.
        items:
          properties:
            label:
              description: Label to describe usage of the URL
              type: string
            url:
              description: Specific URL to associate with the note
              type: string
          type: object
        type: array
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      shortDescription:
        description: A one sentence description of this note.
        type: string
      vulnerability:
        description: A note describing a package vulnerability.
        properties:
          cvssScore:
            description: The CVSS score of this vulnerability. CVSS score is on a
              scale of 0 - 10 where 0 indicates low severity and 10 indicates high
              severity.
            format: double
            type: number
          cvssV3:
            description: The full description of the CVSSv3 for this vulnerability.
            properties:
              attackComplexity:
                description: ' Possible values: ATTACK_COMPLEXITY_UNSPECIFIED, ATTACK_COMPLEXITY_LOW,
                  ATTACK_COMPLEXITY_HIGH'
                type: string
              attackVector:
                description: 'Base Metrics Represents the intrinsic characteristics
                  of a vulnerability that are constant over time and across user environments.
                  Possible values: ATTACK_VECTOR_UNSPECIFIED, ATTACK_VECTOR_NETWORK,
                  ATTACK_VECTOR_ADJACENT, ATTACK_VECTOR_LOCAL, ATTACK_VECTOR_PHYSICAL'
                type: string
              availabilityImpact:
                description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW,
                  IMPACT_NONE'
                type: string
              baseScore:
                description: The base score is a function of the base metric scores.
                format: double
                type: number
              confidentialityImpact:
                description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW,
                  IMPACT_NONE'
                type: string
              exploitabilityScore:
                format: double
                type: number
              impactScore:
                format: double
                type: number
              integrityImpact:
                description: ' Possible values: IMPACT_UNSPECIFIED, IMPACT_HIGH, IMPACT_LOW,
                  IMPACT_NONE'
                type: string
              privilegesRequired:
                description: ' Possible values: PRIVILEGES_REQUIRED_UNSPECIFIED, PRIVILEGES_REQUIRED_NONE,
                  PRIVILEGES_REQUIRED_LOW, PRIVILEGES_REQUIRED_HIGH'
                type: string
              scope:
                description: ' Possible values: SCOPE_UNSPECIFIED, SCOPE_UNCHANGED,
                  SCOPE_CHANGED'
                type: string
              userInteraction:
                description: ' Possible values: USER_INTERACTION_UNSPECIFIED, USER_INTERACTION_NONE,
                  USER_INTERACTION_REQUIRED'
                type: string
            type: object
          details:
            description: Details of all known distros and packages affected by this
              vulnerability.
            items:
              properties:
                affectedCpeUri:
                  description: Required. The (https://cpe.mitre.org/specification/)
                    this vulnerability affects.
                  type: string
                affectedPackage:
                  description: Required. The package this vulnerability affects.
                  type: string
                affectedVersionEnd:
                  description: 'The version number at the end of an interval in which
                    this vulnerability exists. A vulnerability can affect a package
                    between version numbers that are disjoint sets of intervals (example:
                    ) each of which will be represented in its own Detail. If a specific
                    affected version is provided by a vulnerability database, affected_version_start
                    and affected_version_end will be the same in that Detail.'
                  properties:
                    epoch:
                      description: Used to correct mistakes in the version numbering
                        scheme.
                      format: int64
                      type: integer
                    fullName:
                      description: Human readable version string. This string is of
                        the form :- and is only set when kind is NORMAL.
                      type: string
                    kind:
                      description: 'Required. Distinguishes between sentinel MIN/MAX
                        versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,
                        VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,
                        ATTESTATION, UPGRADE'
                      type: string
                    name:
                      description: Required only when version kind is NORMAL. The
                        main part of the version name.
                      type: string
                    revision:
                      description: The iteration of the package build from the above
                        version.
                      type: string
                  required:
                  - kind
                  type: object
                affectedVersionStart:
                  description: 'The version number at the start of an interval in
                    which this vulnerability exists. A vulnerability can affect a
                    package between version numbers that are disjoint sets of intervals
                    (example: ) each of which will be represented in its own Detail.
                    If a specific affected version is provided by a vulnerability
                    database, affected_version_start and affected_version_end will
                    be the same in that Detail.'
                  properties:
                    epoch:
                      description: Used to correct mistakes in the version numbering
                        scheme.
                      format: int64
                      type: integer
                    fullName:
                      description: Human readable version string. This string is of
                        the form :- and is only set when kind is NORMAL.
                      type: string
                    kind:
                      description: 'Required. Distinguishes between sentinel MIN/MAX
                        versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,
                        VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,
                        ATTESTATION, UPGRADE'
                      type: string
                    name:
                      description: Required only when version kind is NORMAL. The
                        main part of the version name.
                      type: string
                    revision:
                      description: The iteration of the package build from the above
                        version.
                      type: string
                  required:
                  - kind
                  type: object
                description:
                  description: A vendor-specific description of this vulnerability.
                  type: string
                fixedCpeUri:
                  description: The distro recommended (https://cpe.mitre.org/specification/)
                    to update to that contains a fix for this vulnerability. It is
                    possible for this to be different from the affected_cpe_uri.
                  type: string
                fixedPackage:
                  description: The distro recommended package to update to that contains
                    a fix for this vulnerability. It is possible for this to be different
                    from the affected_package.
                  type: string
                fixedVersion:
                  description: The distro recommended version to update to that contains
                    a fix for this vulnerability. Setting this to VersionKind.MAXIMUM
                    means no such version is yet available.
                  properties:
                    epoch:
                      description: Used to correct mistakes in the version numbering
                        scheme.
                      format: int64
                      type: integer
                    fullName:
                      description: Human readable version string. This string is of
                        the form :- and is only set when kind is NORMAL.
                      type: string
                    kind:
                      description: 'Required. Distinguishes between sentinel MIN/MAX
                        versions and normal versions. Possible values: NOTE_KIND_UNSPECIFIED,
                        VULNERABILITY, BUILD, IMAGE, PACKAGE, DEPLOYMENT, DISCOVERY,
                        ATTESTATION, UPGRADE'
                      type: string
                    name:
                      description: Required only when version kind is NORMAL. The
                        main part of the version name.
                      type: string
                    revision:
                      description: The iteration of the package build from the above
                        version.
                      type: string
                  required:
                  - kind
                  type: object
                isObsolete:
                  description: Whether this detail is obsolete. Occurrences are expected
                    not to point to obsolete details.
                  type: boolean
                packageType:
                  description: The type of package; whether native or non native (e.g.,
                    ruby gems, node.js packages, etc.).
                  type: string
                severityName:
                  description: The distro assigned severity of this vulnerability.
                  type: string
                sourceUpdateTime:
                  description: The time this information was last changed at the source.
                    This is an upstream timestamp from the underlying information
                    source - e.g. Ubuntu security tracker.
                  format: date-time
                  type: string
              required:
              - affectedCpeUri
              - affectedPackage
              type: object
            type: array
          severity:
            description: 'The note provider assigned severity of this vulnerability.
              Possible values: SEVERITY_UNSPECIFIED, MINIMAL, LOW, MEDIUM, HIGH, CRITICAL'
            type: string
          sourceUpdateTime:
            description: The time this information was last changed at the source.
              This is an upstream timestamp from the underlying information source
              - e.g. Ubuntu security tracker.
            format: date-time
            type: string
          windowsDetails:
            description: Windows details get their own format because the information
              format and model don't match a normal detail. Specifically Windows updates
              are done as patches, thus Windows vulnerabilities really are a missing
              package, rather than a package being at an incorrect version.
            items:
              properties:
                cpeUri:
                  description: Required. The (https://cpe.mitre.org/specification/)
                    this vulnerability affects.
                  type: string
                description:
                  description: The description of this vulnerability.
                  type: string
                fixingKbs:
                  description: Required. The names of the KBs which have hotfixes
                    to mitigate this vulnerability. Note that there may be multiple
                    hotfixes (and thus multiple KBs) that mitigate a given vulnerability.
                    Currently any listed KBs presence is considered a fix.
                  items:
                    properties:
                      name:
                        description: The KB name (generally of the form KB+ (e.g.,
                          KB123456)).
                        type: string
                      url:
                        description: A link to the KB in the (https://www.catalog.update.microsoft.com/).
                        type: string
                    type: object
                  type: array
                name:
                  description: Required. The name of this vulnerability.
                  type: string
              required:
              - cpeUri
              - fixingKbs
              - name
              type: object
            type: array
        type: object
    type: object
type: object

```
