# openAPI schema part4 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      inspectJob:
        properties:
          inspectConfig:
            properties:
              ruleSet:
                properties:
                  ruleSet:
                    description: Set of rules to apply to the findings for this InspectConfig.
                      Exclusion rules, contained in the set are executed in the end,
                      other rules are executed in the order they are specified for
                      each info type.
                    items:
                      properties:
                        infoTypes:
                          description: List of infoTypes this rule set is applied
                            to.
                          items:
                            properties:
                              name:
                                description: Name of the information type. Either
                                  a name of your choosing when creating a CustomInfoType,
                                  or one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference
                                  when specifying a built-in type. When sending Cloud
                                  DLP results to Data Catalog, infoType names should
                                  conform to the pattern `[A-Za-z0-9$-_]{1,64}`.
                                type: string
                              version:
                                description: Optional version name for this InfoType.
                                type: string
                            type: object
                          type: array
                        rules:
                          description: Set of rules to be applied to infoTypes. The
                            rules are applied in order.
                          items:
                            properties:
                              exclusionRule:
                                description: Exclusion rule.
                                properties:
                                  dictionary:
                                    description: Dictionary which defines the rule.
                                    properties:
                                      cloudStoragePath:
                                        description: Newline-delimited file of words
                                          in Cloud Storage. Only a single file is
                                          accepted.
                                        properties:
                                          path:
                                            description: 'A url representing a file
                                              or path (no wildcards) in Cloud Storage.
                                              Example: gs://[BUCKET_NAME]/dictionary.txt'
                                            type: string
                                        type: object
                                      wordList:
                                        description: List of words or phrases to search
                                          for.
                                        properties:
                                          words:
                                            description: Words or phrases defining
                                              the dictionary. The dictionary must
                                              contain at least one phrase and every
                                              phrase must contain at least 2 characters
                                              that are letters or digits. [required]
                                            items:
                                              type: string
                                            type: array
                                        type: object
                                    type: object
                                  excludeInfoTypes:
                                    description: Set of infoTypes for which findings
                                      would affect this rule.
                                    properties:
                                      infoTypes:
                                        description: InfoType list in ExclusionRule
                                          rule drops a finding when it overlaps or
                                          contained within with a finding of an infoType
                                          from this list. For example, for `InspectionRuleSet.info_types`
                                          containing "PHONE_NUMBER"` and `exclusion_rule`
                                          containing `exclude_info_types.info_types`
                                          with "EMAIL_ADDRESS" the phone number findings
                                          are dropped if they overlap with EMAIL_ADDRESS
                                          finding. That leads to "555-222-2222@example.org"
                                          to generate only a single finding, namely
                                          email address.
                                        items:
                                          properties:
                                            name:
                                              description: Name of the information
                                                type. Either a name of your choosing
                                                when creating a CustomInfoType, or
                                                one of the names listed at https://cloud.google.com/dlp/docs/infotypes-reference
                                                when specifying a built-in type. When
                                                sending Cloud DLP results to Data
                                                Catalog, infoType names should conform
                                                to the pattern `[A-Za-z0-9$-_]{1,64}`.
                                              type: string
                                            version:
                                              description: Optional version name for
                                                this InfoType.
                                              type: string
                                          type: object
                                        type: array
                                    type: object
                                  matchingType:
                                    description: 'How the rule is applied, see MatchingType
                                      documentation for details. Possible values:
                                      MATCHING_TYPE_UNSPECIFIED, MATCHING_TYPE_FULL_MATCH,
                                      MATCHING_TYPE_PARTIAL_MATCH, MATCHING_TYPE_INVERSE_MATCH'
                                    type: string
                                  regex:
                                    description: Regular expression which defines
                                      the rule.
                                    properties:
                                      groupIndexes:
                                        description: The index of the submatch to
                                          extract as findings. When not specified,
                                          the entire match is returned. No more than
                                          3 may be included.
                                        items:
                                          format: int64
                                          type: integer
                                        type: array
                                      pattern:
                                        description: Pattern defining the regular
                                          expression. Its syntax (https://github.com/google/re2/wiki/Syntax)
                                          can be found under the google/re2 repository
                                          on GitHub.
                                        type: string
                                    type: object
                                type: object
                              hotwordRule:
                                properties:
                                  hotwordRegex:
                                    description: Regular expression pattern defining
                                      what qualifies as a hotword.
                                    properties:
                                      groupIndexes:
                                        description: The index of the submatch to
                                          extract as findings. When not specified,
                                          the entire match is returned. No more than
                                          3 may be included.
                                        items:
                                          format: int64
                                          type: integer
                                        type: array
                                      pattern:
                                        description: Pattern defining the regular
                                          expression. Its syntax (https://github.com/google/re2/wiki/Syntax)
                                          can be found under the google/re2 repository
                                          on GitHub.
                                        type: string
                                    type: object
                                  likelihoodAdjustment:
                                    description: Likelihood adjustment to apply to
                                      all matching findings.
                                    properties:
                                      fixedLikelihood:
                                        description: 'Set the likelihood of a finding
                                          to a fixed value. Possible values: LIKELIHOOD_UNSPECIFIED,
                                          VERY_UNLIKELY, UNLIKELY, POSSIBLE, LIKELY,
                                          VERY_LIKELY'
                                        type: string
                                      relativeLikelihood:
                                        description: Increase or decrease the likelihood
                                          by the specified number of levels. For example,
                                          if a finding would be `POSSIBLE` without
                                          the detection rule and `relative_likelihood`
                                          is 1, then it is upgraded to `LIKELY`, while
                                          a value of -1 would downgrade it to `UNLIKELY`.
                                          Likelihood may never drop below `VERY_UNLIKELY`
                                          or exceed `VERY_LIKELY`, so applying an
                                          adjustment of 1 followed by an adjustment
                                          of -1 when base likelihood is `VERY_LIKELY`
                                          will result in a final likelihood of `LIKELY`.
                                        format: int64
                                        type: integer
                                    type: object
                                  proximity:
                                    description: Proximity of the finding within which
                                      the entire hotword must reside. The total length
                                      of the window cannot exceed 1000 characters.
                                      Note that the finding itself will be included
                                      in the window, so that hotwords may be used
                                      to match substrings of the finding itself. For
                                      example, the certainty of a phone number regex
                                      "(d{3}) d{3}-d{4}" could be adjusted upwards
                                      if the area code is known to be the local area
                                      code of a company office using the hotword regex
                                      "(xxx)", where "xxx" is the area code in question.
                                    properties:
                                      windowAfter:
                                        description: Number of characters after the
                                          finding to consider.
                                        format: int64
                                        type: integer
                                      windowBefore:
                                        description: Number of characters before the
                                          finding to consider.
                                        format: int64
                                        type: integer
                                    type: object
                                type: object
                            type: object
                          type: array
                      type: object
                    type: array
          inspectTemplateRef:
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
                  If provided, will be used as the default for all values in InspectConfig. `inspect_config` will be merged into the values persisted as part of the template.

                  Allowed value: The Google Cloud resource name of a `DLPInspectTemplate` resource (format: `{{parent}}/inspectTemplates/{{name}}`).
                type: string
              name:
                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                type: string
              namespace:
                description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                type: string
            type: object
required:
- spec
type: object

```
