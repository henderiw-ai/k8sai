# openAPI schema part5 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      inspectJob:
        properties:
          storageConfig:
            properties:
              storageConfig:
                description: The data to scan.
                properties:
                  bigQueryOptions:
                    description: BigQuery options.
                    properties:
                      excludedFields:
                        description: References to fields excluded from scanning.
                          This allows you to skip inspection of entire columns which
                          you know have no findings.
                        items:
                          properties:
                            name:
                              description: Name describing the field.
                              type: string
                          type: object
                        type: array
                      identifyingFields:
                        description: Table fields that may uniquely identify a row
                          within the table. When `actions.saveFindings.outputConfig.table`
                          is specified, the values of columns specified here are available
                          in the output table under `location.content_locations.record_location.record_key.id_values`.
                          Nested fields such as `person.birthdate.year` are allowed.
                        items:
                          properties:
                            name:
                              description: Name describing the field.
                              type: string
                          type: object
                        type: array
                      includedFields:
                        description: Limit scanning only to these fields.
                        items:
                          properties:
                            name:
                              description: Name describing the field.
                              type: string
                          type: object
                        type: array
                      rowsLimit:
                        description: Max number of rows to scan. If the table has
                          more rows than this value, the rest of the rows are omitted.
                          If not set, or if set to 0, all rows will be scanned. Only
                          one of rows_limit and rows_limit_percent can be specified.
                          Cannot be used in conjunction with TimespanConfig.
                        format: int64
                        type: integer
                      rowsLimitPercent:
                        description: Max percentage of rows to scan. The rest are
                          omitted. The number of rows scanned is rounded down. Must
                          be between 0 and 100, inclusively. Both 0 and 100 means
                          no limit. Defaults to 0. Only one of rows_limit and rows_limit_percent
                          can be specified. Cannot be used in conjunction with TimespanConfig.
                        format: int64
                        type: integer
                      sampleMethod:
                        description: ' Possible values: SAMPLE_METHOD_UNSPECIFIED,
                          TOP, RANDOM_START'
                        type: string
                      tableReference:
                        description: Complete BigQuery table reference.
                        properties:
                          datasetRef:
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
                                  Dataset ID of the table.

                                  Allowed value: The Google Cloud resource name of a `BigQueryDataset` resource (format: `projects/{{project}}/datasets/{{name}}`).
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          projectRef:
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
                                  The Google Cloud Platform project ID of the project containing the table. If omitted, project ID is inferred from the API call.

                                  Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
                                type: string
                              name:
                                description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                type: string
                              namespace:
                                description: 'Namespace of the referent. More info:
                                  https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                type: string
                            type: object
                          tableRef:
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
                                  Name of the table.

                                  Allowed value: The Google Cloud resource name of a `BigQueryTable` resource (format: `projects/{{project}}/datasets/{{dataset_id}}/tables/{{name}}`).
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
                    required:
                    - tableReference
                    type: object
                  cloudStorageOptions:
                    description: Google Cloud Storage options.
                    properties:
                      bytesLimitPerFile:
                        description: Max number of bytes to scan from a file. If a
                          scanned file's size is bigger than this value then the rest
                          of the bytes are omitted. Only one of bytes_limit_per_file
                          and bytes_limit_per_file_percent can be specified. Cannot
                          be set if de-identification is requested.
                        format: int64
                        type: integer
                      bytesLimitPerFilePercent:
                        description: Max percentage of bytes to scan from a file.
                          The rest are omitted. The number of bytes scanned is rounded
                          down. Must be between 0 and 100, inclusively. Both 0 and
                          100 means no limit. Defaults to 0. Only one of bytes_limit_per_file
                          and bytes_limit_per_file_percent can be specified. Cannot
                          be set if de-identification is requested.
                        format: int64
                        type: integer
                      fileSet:
                        description: The set of one or more files to scan.
                        properties:
                          regexFileSet:
                            description: The regex-filtered set of files to scan.
                              Exactly one of `url` or `regex_file_set` must be set.
                            properties:
                              bucketRef:
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
                                      The name of a Cloud Storage bucket. Required.

                                      Allowed value: The Google Cloud resource name of a `StorageBucket` resource (format: `{{name}}`).
                                    type: string
                                  name:
                                    description: 'Name of the referent. More info:
                                      https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                                    type: string
                                  namespace:
                                    description: 'Namespace of the referent. More
                                      info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                                    type: string
                                type: object
                              excludeRegex:
                                description: A list of regular expressions matching
                                  file paths to exclude. All files in the bucket that
                                  match at least one of these regular expressions
                                  will be excluded from the scan. Regular expressions
                                  use RE2 [syntax](https://github.com/google/re2/wiki/Syntax);
                                  a guide can be found under the google/re2 repository
                                  on GitHub.
                                items:
                                  type: string
                                type: array
                              includeRegex:
                                description: A list of regular expressions matching
                                  file paths to include. All files in the bucket that
                                  match at least one of these regular expressions
                                  will be included in the set of files, except for
                                  those that also match an item in `exclude_regex`.
                                  Leaving this field empty will match all files by
                                  default (this is equivalent to including `.*` in
                                  the list). Regular expressions use RE2 [syntax](https://github.com/google/re2/wiki/Syntax);
                                  a guide can be found under the google/re2 repository
                                  on GitHub.
                                items:
                                  type: string
                                type: array
                            required:
                            - bucketRef
                            type: object
                          url:
                            description: The Cloud Storage url of the file(s) to scan,
                              in the format `gs:///`. Trailing wildcard in the path
                              is allowed. If the url ends in a trailing slash, the
                              bucket or directory represented by the url will be scanned
                              non-recursively (content in sub-directories will not
                              be scanned). This means that `gs://mybucket/` is equivalent
                              to `gs://mybucket/*`, and `gs://mybucket/directory/`
                              is equivalent to `gs://mybucket/directory/*`. Exactly
                              one of `url` or `regex_file_set` must be set.
                            type: string
                        type: object
                      fileTypes:
                        description: List of file type groups to include in the scan.
                          If empty, all files are scanned and available data format
                          processors are applied. In addition, the binary content
                          of the selected files is always scanned as well. Images
                          are scanned only as binary if the specified region does
                          not support image inspection and no file_types were specified.
                          Image inspection is restricted to 'global', 'us', 'asia',
                          and 'europe'.
                        items:
                          type: string
                        type: array
                      filesLimitPercent:
                        description: Limits the number of files to scan to this percentage
                          of the input FileSet. Number of files scanned is rounded
                          down. Must be between 0 and 100, inclusively. Both 0 and
                          100 means no limit. Defaults to 0.
                        format: int64
                        type: integer
                      sampleMethod:
                        description: ' Possible values: SAMPLE_METHOD_UNSPECIFIED,
                          TOP, RANDOM_START'
                        type: string
                    type: object
                  datastoreOptions:
                    description: Google Cloud Datastore options.
                    properties:
                      kind:
                        description: The kind to process.
                        properties:
                          name:
                            description: The name of the kind.
                            type: string
                        type: object
                      partitionId:
                        description: A partition ID identifies a grouping of entities.
                          The grouping is always by project namespace ID may be empty.
                        properties:
                          namespaceId:
                            description: If not empty, the ID of the namespace to
                              which the entities belong.
                            type: string
                          projectRef:
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
                                  The ID of the project to which the entities belong.

                                  Allowed value: The Google Cloud resource name of a `Project` resource (format: `projects/{{name}}`).
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
                    type: object
                  hybridOptions:
                    description: Hybrid inspection options.
                    properties:
                      description:
                        description: A short description of where the data is coming
                          from. Will be stored once in the job. 256 max length.
                        type: string
                      labels:
                        additionalProperties:
                          type: string
                        description: 'To organize findings, these labels will be added
                          to each finding. Label keys must be between 1 and 63 characters
                          long and must conform to the following regular expression:
                          `[a-z]([-a-z0-9]*[a-z0-9])?`. Label values must be between
                          0 and 63 characters long and must conform to the regular
                          expression `([a-z]([-a-z0-9]*[a-z0-9])?)?`. No more than
                          10 labels can be associated with a given finding. Examples:
                          * `"environment" : "production"` * `"pipeline" : "etl"`'
                        type: object
                      requiredFindingLabelKeys:
                        description: 'These are labels that each inspection request
                          must include within their ''finding_labels'' map. Request
                          may contain others, but any missing one of these will be
                          rejected. Label keys must be between 1 and 63 characters
                          long and must conform to the following regular expression:
                          `[a-z]([-a-z0-9]*[a-z0-9])?`. No more than 10 keys can be
                          required.'
                        items:
                          type: string
                        type: array
                      tableOptions:
                        description: If the container is a table, additional information
                          to make findings meaningful such as the columns that are
                          primary keys.
                        properties:
                          identifyingFields:
                            description: The columns that are the primary keys for
                              table objects included in ContentItem. A copy of this
                              cell's value will stored alongside alongside each finding
                              so that the finding can be traced to the specific row
                              it came from. No more than 3 may be provided.
                            items:
                              properties:
                                name:
                                  description: Name describing the field.
                                  type: string
                              type: object
                            type: array
                        type: object
                    type: object
                  timespanConfig:
                    properties:
                      enableAutoPopulationOfTimespanConfig:
                        description: When the job is started by a JobTrigger we will
                          automatically figure out a valid start_time to avoid scanning
                          files that have not been modified since the last time the
                          JobTrigger executed. This will be based on the time of the
                          execution of the last run of the JobTrigger.
                        type: boolean
                      endTime:
                        description: Exclude files, tables, or rows newer than this
                          value. If not set, no upper time limit is applied.
                        format: date-time
                        type: string
                      startTime:
                        description: Exclude files, tables, or rows older than this
                          value. If not set, no lower time limit is applied.
                        format: date-time
                        type: string
                      timestampField:
                        description: 'Specification of the field containing the timestamp
                          of scanned items. Used for data sources like Datastore and
                          BigQuery. For BigQuery: If this value is not specified and
                          the table was modified between the given start and end times,
                          the entire table will be scanned. If this value is specified,
                          then rows are filtered based on the given start and end
                          times. Rows with a `NULL` value in the provided BigQuery
                          column are skipped. Valid data types of the provided BigQuery
                          column are: `INTEGER`, `DATE`, `TIMESTAMP`, and `DATETIME`.
                          For Datastore: If this value is specified, then entities
                          are filtered based on the given start and end times. If
                          an entity does not contain the provided timestamp property
                          or contains empty or invalid values, then it is included.
                          Valid data types of the provided timestamp property are:
                          `TIMESTAMP`.'
                        properties:
                          name:
                            description: Name describing the field.
                            type: string
                        type: object
                    type: object
                type: object
      location:
        description: Immutable. The location of the resource
        type: string
required:
- spec
type: object

```
