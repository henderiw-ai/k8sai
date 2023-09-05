# openAPI schema part3 for bigquery.cnrm.cloud.google.com.v1beta1.BigQueryJob

## schema

```yaml
properties:
  spec:
    properties:
      query:
        properties:
          query:
            description: Immutable. Configures a query job.
            properties:
              allowLargeResults:
                description: |-
                  Immutable. If true and query uses legacy SQL dialect, allows the query to produce arbitrarily large result tables at a slight cost in performance.
                  Requires destinationTable to be set. For standard SQL queries, this flag is ignored and large results are always allowed.
                  However, you must still set destinationTable when result size exceeds the allowed maximum response size.
                type: boolean
              createDisposition:
                description: |-
                  Immutable. Specifies whether the job is allowed to create new tables. The following values are supported:
                  CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
                  CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
                  Creation, truncation and append actions occur as one atomic update upon job completion Default value: "CREATE_NEVER" Possible values: ["CREATE_IF_NEEDED", "CREATE_NEVER"].
                type: string
              defaultDataset:
                description: Immutable. Specifies the default dataset to use for unqualified
                  table names in the query. Note that this does not alter behavior
                  of unqualified dataset names.
                properties:
                  datasetRef:
                    description: A reference to the dataset.
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
                        description: 'Allowed value: The `selfLink` field of a `BigQueryDataset`
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
                - datasetRef
                type: object
              destinationEncryptionConfiguration:
                description: Immutable. Custom encryption configuration (e.g., Cloud
                  KMS keys).
                properties:
                  kmsKeyRef:
                    description: |-
                      Describes the Cloud KMS encryption key that will be used to protect
                      destination BigQuery table. The BigQuery Service Account associated
                      with your project requires access to this encryption key.
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
                        description: 'Allowed value: The `selfLink` field of a `KMSCryptoKey`
                          resource.'
                        type: string
                      name:
                        description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                        type: string
                      namespace:
                        description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                        type: string
                    type: object
                  kmsKeyVersion:
                    description: Describes the Cloud KMS encryption key version used
                      to protect destination BigQuery table.
                    type: string
                required:
                - kmsKeyRef
                type: object
              destinationTable:
                description: |-
                  Immutable. Describes the table where the query results should be stored.
                  This property must be set for large results that exceed the maximum response size.
                  For queries that produce anonymous (cached) results, this field will be populated by BigQuery.
                properties:
                  tableRef:
                    description: A reference to the table.
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
                        description: 'Allowed value: The `selfLink` field of a `BigQueryTable`
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
                - tableRef
                type: object
              flattenResults:
                description: |-
                  Immutable. If true and query uses legacy SQL dialect, flattens all nested and repeated fields in the query results.
                  allowLargeResults must be true if this is set to false. For standard SQL queries, this flag is ignored and results are never flattened.
                type: boolean
              maximumBillingTier:
                description: |-
                  Immutable. Limits the billing tier for this job. Queries that have resource usage beyond this tier will fail (without incurring a charge).
                  If unspecified, this will be set to your project default.
                type: integer
              maximumBytesBilled:
                description: |-
                  Immutable. Limits the bytes billed for this job. Queries that will have bytes billed beyond this limit will fail (without incurring a charge).
                  If unspecified, this will be set to your project default.
                type: string
              parameterMode:
                description: Immutable. Standard SQL only. Set to POSITIONAL to use
                  positional (?) query parameters or to NAMED to use named (@myparam)
                  query parameters in this query.
                type: string
              priority:
                description: 'Immutable. Specifies a priority for the query. Default
                  value: "INTERACTIVE" Possible values: ["INTERACTIVE", "BATCH"].'
                type: string
              query:
                description: |-
                  Immutable. SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
                  *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
                  ('DELETE', 'UPDATE', 'MERGE', 'INSERT') must specify 'create_disposition = ""' and 'write_disposition = ""'.
                type: string
              schemaUpdateOptions:
                description: |-
                  Immutable. Allows the schema of the destination table to be updated as a side effect of the query job.
                  Schema update options are supported in two cases: when writeDisposition is WRITE_APPEND;
                  when writeDisposition is WRITE_TRUNCATE and the destination table is a partition of a table,
                  specified by partition decorators. For normal tables, WRITE_TRUNCATE will always overwrite the schema.
                  One or more of the following values are specified:
                  ALLOW_FIELD_ADDITION: allow adding a nullable field to the schema.
                  ALLOW_FIELD_RELAXATION: allow relaxing a required field in the original schema to nullable.
                items:
                  type: string
                type: array
              scriptOptions:
                description: Immutable. Options controlling the execution of scripts.
                properties:
                  keyResultStatement:
                    description: |-
                      Immutable. Determines which statement in the script represents the "key result",
                      used to populate the schema and query results of the script job. Possible values: ["LAST", "FIRST_SELECT"].
                    type: string
                  statementByteBudget:
                    description: Immutable. Limit on the number of bytes billed per
                      statement. Exceeding this budget results in an error.
                    type: string
                  statementTimeoutMs:
                    description: Immutable. Timeout period for each statement in a
                      script.
                    type: string
                type: object
              useLegacySql:
                description: |-
                  Immutable. Specifies whether to use BigQuery's legacy SQL dialect for this query. The default value is true.
                  If set to false, the query will use BigQuery's standard SQL.
                type: boolean
              useQueryCache:
                description: |-
                  Immutable. Whether to look for the result in the query cache. The query cache is a best-effort cache that will be flushed whenever
                  tables in the query are modified. Moreover, the query cache is only available when a query does not have a destination table specified.
                  The default value is true.
                type: boolean
              userDefinedFunctionResources:
                description: Immutable. Describes user-defined function resources
                  used in the query.
                items:
                  properties:
                    inlineCode:
                      description: |-
                        Immutable. An inline resource that contains code for a user-defined function (UDF).
                        Providing a inline code resource is equivalent to providing a URI for a file containing the same code.
                      type: string
                    resourceUri:
                      description: Immutable. A code resource to load from a Google
                        Cloud Storage URI (gs://bucket/path).
                      type: string
                  type: object
                type: array
              writeDisposition:
                description: |-
                  Immutable. Specifies the action that occurs if the destination table already exists. The following values are supported:
                  WRITE_TRUNCATE: If the table already exists, BigQuery overwrites the table data and uses the schema from the query result.
                  WRITE_APPEND: If the table already exists, BigQuery appends the data to the table.
                  WRITE_EMPTY: If the table already exists and contains data, a 'duplicate' error is returned in the job result.
                  Each action is atomic and only occurs if BigQuery is able to complete the job successfully.
                  Creation, truncation and append actions occur as one atomic update upon job completion. Default value: "WRITE_EMPTY" Possible values: ["WRITE_TRUNCATE", "WRITE_APPEND", "WRITE_EMPTY"].
                type: string
            required:
            - query
            type: object
      resourceID:
        description: Immutable. Optional. The jobId of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
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
      jobType:
        description: The type of the job.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      status:
        description: The status of this job. Examine this value when polling an asynchronous
          job to see if the job is complete.
        items:
          properties:
            errorResult:
              description: Final error result of the job. If present, indicates that
                the job has completed and was unsuccessful.
              items:
                properties:
                  location:
                    description: Specifies where the error occurred, if present.
                    type: string
                  message:
                    description: A human-readable description of the error.
                    type: string
                  reason:
                    description: A short error code that summarizes the error.
                    type: string
                type: object
              type: array
            errors:
              description: |-
                The first errors encountered during the running of the job. The final message
                includes the number of errors that caused the process to stop. Errors here do
                not necessarily mean that the job has not completed or was unsuccessful.
              items:
                properties:
                  location:
                    description: Specifies where the error occurred, if present.
                    type: string
                  message:
                    description: A human-readable description of the error.
                    type: string
                  reason:
                    description: A short error code that summarizes the error.
                    type: string
                type: object
              type: array
            state:
              description: Running state of the job. Valid states include 'PENDING',
                'RUNNING', and 'DONE'.
              type: string
          type: object
        type: array
      userEmail:
        description: Email address of the user who ran the job.
        type: string
    type: object
type: object

```
