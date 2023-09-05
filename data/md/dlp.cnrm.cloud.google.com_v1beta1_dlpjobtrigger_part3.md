# openAPI schema part3 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      inspectJob:
        properties:
          inspectConfig:
            properties:
              customInfoTypes:
                description: CustomInfoTypes provided by the user. See https://cloud.google.com/dlp/docs/creating-custom-infotypes
                  to learn more.
                items:
                  properties:
                    detectionRules:
                      description: Set of detection rules to apply to all findings
                        of this CustomInfoType. Rules are applied in order that they
                        are specified. Not supported for the `surrogate_type` CustomInfoType.
                      items:
                        properties:
                          hotwordRule:
                            description: Hotword-based detection rule.
                            properties:
                              hotwordRegex:
                                description: Regular expression pattern defining what
                                  qualifies as a hotword.
                                properties:
                                  groupIndexes:
                                    description: The index of the submatch to extract
                                      as findings. When not specified, the entire
                                      match is returned. No more than 3 may be included.
                                    items:
                                      format: int64
                                      type: integer
                                    type: array
                                  pattern:
                                    description: Pattern defining the regular expression.
                                      Its syntax (https://github.com/google/re2/wiki/Syntax)
                                      can be found under the google/re2 repository
                                      on GitHub.
                                    type: string
                                type: object
                              likelihoodAdjustment:
                                description: Likelihood adjustment to apply to all
                                  matching findings.
                                properties:
                                  fixedLikelihood:
                                    description: 'Set the likelihood of a finding
                                      to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED,
                                      VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'
                                    type: string
                                  relativeLikelihood:
                                    description: Increase or decrease the likelihood
                                      by the specified number of levels. For example,
                                      if a finding would be `POSSIBLE` without the
                                      detection rule and `relative_likelihood` is
                                      1, then it is upgraded to `LIKELY`, while a
                                      value of -1 would downgrade it to `UNLIKELY`.
                                      Likelihood may never drop below `VERY_UNLIKELY`
                                      or exceed `VERY_LIKELY`, so applying an adjustment
                                      of 1 followed by an adjustment of -1 when base
                                      likelihood is `VERY_LIKELY` will result in a
                                      final likelihood of `LIKELY`.
                                    format: int64
                                    type: integer
                                type: object
                              proximity:
                                description: Proximity of the finding within which
                                  the entire hotword must reside. The total length
                                  of the window cannot exceed 1000 characters. Note
                                  that the finding itself will be included in the
                                  window, so that hotwords may be used to match substrings
                                  of the finding itself. For example, the certainty
                                  of a phone number regex "(d{3}) d{3}-d{4}" could
                                  be adjusted upwards if the area code is known to
                                  be the local area code of a company office using
                                  the hotword regex "(xxx)", where "xxx" is the area
                                  code in question.
                                properties:
                                  windowAfter:
                                    description: Number of characters after the finding
                                      to consider.
                                    format: int64
                                    type: integer
                                  windowBefore:
                                    description: Number of characters before the finding
                                      to consider.
                                    format: int64
                                    type: integer
                                type: object
                            type: object
                        type: object
                      type: array
                    dictionary:
                      description: A list of phrases to detect as a CustomInfoType.
                      properties:
                        cloudStoragePath:
                          description: Newline-delimited file of words in Cloud Storage.
                            Only a single file is accepted.
                          properties:
                            path:
                              description: 'A url representing a file or path (no
                                wildcards) in Cloud Storage. Example: gs://[BUCKET_NAME]/dictionary.txt'
                              type: string
                          type: object
                        wordList:
                          description: List of words or phrases to search for.
                          properties:
                            words:
                              description: Words or phrases defining the dictionary.
                                The dictionary must contain at least one phrase and
                                every phrase must contain at least 2 characters that
                                are letters or digits. [required]
                              items:
                                type: string
                              type: array
                          type: object
                      type: object
                    exclusionType:
                      description: 'If set to EXCLUSION_TYPE_EXCLUDE this infoType
                        will not cause a finding to be returned. It still can be used
                        for rules matching. Possible values: EXCLUSION_TYPE_UNSPECIFIED,
                        EXCLUSION_TYPE_EXCLUDE'
                      type: string
                    infoType:
                      description: CustomInfoType can either be a new infoType, or
                        an extension of built-in infoType, when the name matches one
                        of existing infoTypes and that infoType is specified in `InspectContent.info_types`
                        field. Specifying the latter adds findings to the one detected
                        by the system. If built-in info type is not specified in `InspectContent.info_types`
                        list then the name is treated as a custom info type.
                      properties:
                        name:
                          description: Name of the information type. Either a name
                            of your choosing when creating a CustomInfoType, or one
                            of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference
                            when specifying a built-in type. When sending Cloud DLP
                            results to Data Catalog, infoType names should conform
                            to the pattern `[A-Za-z0-9$-_]{1,64}`.
                          type: string
                        version:
                          description: Optional version name for this InfoType.
                          type: string
                      type: object
                    likelihood:
                      description: 'Likelihood to return for this CustomInfoType.
                        This base value can be altered by a detection rule if the
                        finding meets the criteria specified by the rule. Defaults
                        to `VERY_LIKELY` if not specified. Possible values: LIKELIHOOD_UNSPECIFIED,
                        VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'
                      type: string
                    regex:
                      description: Regular expression based CustomInfoType.
                      properties:
                        groupIndexes:
                          description: The index of the submatch to extract as findings.
                            When not specified, the entire match is returned. No more
                            than 3 may be included.
                          items:
                            format: int64
                            type: integer
                          type: array
                        pattern:
                          description: Pattern defining the regular expression. Its
                            syntax (https://github.com/google/re2/wiki/Syntax) can
                            be found under the google/re2 repository on GitHub.
                          type: string
                      type: object
                    storedType:
                      description: Load an existing `StoredInfoType` resource for
                        use in `InspectDataSource`. Not currently supported in `InspectContent`.
                      properties:
                        createTime:
                          description: Timestamp indicating when the version of the
                            `StoredInfoType` used for inspection was created. Output-only
                            field, populated by the system.
                          format: date-time
                          type: string
                        nameRef:
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
                                Resource name of the requested `StoredInfoType`, for example `organizations/433245324/storedInfoTypes/432452342` or `projects/project-id/storedInfoTypes/432452342`.

                                Allowed value: The Google Cloud resource name of a `DLPStoredInfoType` resource (format: `{{parent}}/storedInfoTypes/{{name}}`).
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
                    surrogateType:
                      description: Message for detecting output from deidentification
                        transformations that support reversing.
                      type: object
                      x-kubernetes-preserve-unknown-fields: true
                  type: object
                type: array
              excludeInfoTypes:
                description: When true, excludes type information of the findings.
                  This is not used for data profiling.
                type: boolean
              includeQuote:
                description: When true, a contextual quote from the data that triggered
                  a finding is included in the response; see Finding.quote. This is
                  not used for data profiling.
                type: boolean
              infoTypes:
                description: Restricts what info_types to look for. The values must
                  correspond to InfoType values returned by ListInfoTypes or listed
                  at https://cloud.google.com/dlp/docs/infotypes-reference. When no
                  InfoTypes or CustomInfoTypes are specified in a request, the system
                  may automatically choose what detectors to run. By default this
                  may be all types, but may change over time as detectors are updated.
                  If you need precise control and predictability as to what detectors
                  are run you should specify specific InfoTypes listed in the reference,
                  otherwise a default list will be used, which may change over time.
                items:
                  properties:
                    name:
                      description: Name of the information type. Either a name of
                        your choosing when creating a CustomInfoType, or one of the
                        names listed at https://cloud.google.com/dlp/docs/infotypes-reference
                        when specifying a built-in type. When sending Cloud DLP results
                        to Data Catalog, infoType names should conform to the pattern
                        `[A-Za-z0-9$-_]{1,64}`.
                      type: string
                  type: object
                type: array
              inspectConfig:
                description: How and what to scan for.
                type: object
              limits:
                description: Configuration to control the number of findings returned.
                  This is not used for data profiling.
                properties:
                  maxFindingsPerInfoType:
                    description: Configuration of findings limit given for specified
                      infoTypes.
                    items:
                      properties:
                        infoType:
                          description: Type of information the findings limit applies
                            to. Only one limit per info_type should be provided. If
                            InfoTypeLimit does not have an info_type, the DLP API
                            applies the limit against all info_types that are found
                            but not specified in another InfoTypeLimit.
                          properties:
                            name:
                              description: Name of the information type. Either a
                                name of your choosing when creating a CustomInfoType,
                                or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference
                                when specifying a built-in type. When sending Cloud
                                DLP results to Data Catalog, infoType names should
                                conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
                              type: string
                            version:
                              description: Optional version name for this InfoType.
                              type: string
                          type: object
                        maxFindings:
                          description: Max findings limit for the given infoType.
                          format: int64
                          type: integer
                      type: object
                    type: array
                  maxFindingsPerItem:
                    description: Max number of findings that will be returned for
                      each item scanned. When set within `InspectJobConfig`, the maximum
                      returned is 2000 regardless if this is set higher. When set
                      within `InspectContentRequest`, this field is ignored.
                    format: int64
                    type: integer
                  maxFindingsPerRequest:
                    description: Max number of findings that will be returned per
                      request/job. When set within `InspectContentRequest`, the maximum
                      returned is 2000 regardless if this is set higher.
                    format: int64
                    type: integer
                type: object
              minLikelihood:
                description: 'Only returns findings equal or above this threshold.
                  The default is POSSIBLE. See https://cloud.google.com/dlp/docs/likelihood
                  to learn more. Possible values: LIKELIHOOD_UNSPECIFIED, VERY_UNLIKELY,
                  UNLIKELY, POSSIBLE, LIKELY, VERY_LIKELY'
                type: string
required:
- spec
type: object

```
