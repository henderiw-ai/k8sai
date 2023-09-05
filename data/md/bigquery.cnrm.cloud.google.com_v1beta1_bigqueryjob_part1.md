# openAPI schema part1 for bigquery.cnrm.cloud.google.com.v1beta1.BigQueryJob

## schema

```yaml
properties:
  spec:
    properties:
      copy:
        description: Immutable. Copies a table.
        properties:
          createDisposition:
            description: |-
              Immutable. Specifies whether the job is allowed to create new tables. The following values are supported:
              CREATE_IF_NEEDED: If the table does not exist, BigQuery creates the table.
              CREATE_NEVER: The table must already exist. If it does not, a 'notFound' error is returned in the job result.
              Creation, truncation and append actions occur as one atomic update upon job completion Default value: "CREATE_NEVER" Possible values: ["CREATE_IF_NEEDED", "CREATE_NEVER"].
            type: string
          destinationEncryptionConfiguration:
            description: Immutable. Custom encryption configuration (e.g., Cloud KMS
              keys).
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
                description: Describes the Cloud KMS encryption key version used to
                  protect destination BigQuery table.
                type: string
            required:
            - kmsKeyRef
            type: object
          destinationTable:
            description: Immutable. The destination table.
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
          sourceTables:
            description: Immutable. Source tables to copy.
            items:
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
        - sourceTables
        type: object
      extract:
        description: Immutable. Configures an extract job.
        properties:
          compression:
            description: |-
              Immutable. The compression type to use for exported files. Possible values include GZIP, DEFLATE, SNAPPY, and NONE.
              The default value is NONE. DEFLATE and SNAPPY are only supported for Avro.
            type: string
          destinationFormat:
            description: |-
              Immutable. The exported file format. Possible values include CSV, NEWLINE_DELIMITED_JSON and AVRO for tables and SAVED_MODEL for models.
              The default value for tables is CSV. Tables with nested or repeated fields cannot be exported as CSV.
              The default value for models is SAVED_MODEL.
            type: string
          destinationUris:
            description: Immutable. A list of fully-qualified Google Cloud Storage
              URIs where the extracted table should be written.
            items:
              type: string
            type: array
          fieldDelimiter:
            description: |-
              Immutable. When extracting data in CSV format, this defines the delimiter to use between fields in the exported data.
              Default is ','.
            type: string
          printHeader:
            description: Immutable. Whether to print out a header row in the results.
              Default is true.
            type: boolean
          sourceTable:
            description: Immutable. A reference to the table being exported.
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
          useAvroLogicalTypes:
            description: Immutable. Whether to use logical types when extracting to
              AVRO format.
            type: boolean
        required:
        - destinationUris
        type: object
      jobTimeoutMs:
        description: Immutable. Job timeout in milliseconds. If this time limit is
          exceeded, BigQuery may attempt to terminate the job.
        type: string
      spec:
        type: object
type: object

```
