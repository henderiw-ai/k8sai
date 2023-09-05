# openAPI schema part2 for dlp.cnrm.cloud.google.com.v1beta1.DLPJobTrigger

## schema

```yaml
properties:
  spec:
    properties:
      inspectJob:
        properties:
          actions:
            description: Actions to execute at the completion of the job.
            items:
              properties:
                jobNotificationEmails:
                  description: Enable email notification for project owners and editors
                    on job's completion/failure.
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                pubSub:
                  description: Publish a notification to a pubsub topic.
                  properties:
                    topicRef:
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
                            Cloud Pub/Sub topic to send notifications to. The topic must have given publishing access rights to the DLP API service account executing the long running DlpJob sending the notifications. Format is projects/{project}/topics/{topic}.

                            Allowed value: The Google Cloud resource name of a `PubSubTopic` resource (format: `projects/{{project}}/topics/{{name}}`).
                          type: string
                        name:
                          description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                          type: string
                        namespace:
                          description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                          type: string
                      type: object
                  type: object
                publishFindingsToCloudDataCatalog:
                  description: Publish findings to Cloud Datahub.
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                publishSummaryToCscc:
                  description: Publish summary to Cloud Security Command Center (Alpha).
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                publishToStackdriver:
                  description: Enable Stackdriver metric dlp.googleapis.com/finding_count.
                  type: object
                  x-kubernetes-preserve-unknown-fields: true
                saveFindings:
                  description: Save resulting findings in a provided location.
                  properties:
                    outputConfig:
                      description: Location to store findings outside of DLP.
                      properties:
                        dlpStorage:
                          description: Store findings directly to DLP. If neither
                            this or bigquery is chosen only summary stats of total
                            infotype count will be stored. Quotes will not be stored
                            to dlp findings. If quotes are needed, store to BigQuery.
                            Currently only for inspect jobs.
                          type: object
                          x-kubernetes-preserve-unknown-fields: true
                        outputSchema:
                          description: 'Schema used for writing the findings for Inspect
                            jobs. This field is only used for Inspect and must be
                            unspecified for Risk jobs. Columns are derived from the
                            `Finding` object. If appending to an existing table, any
                            columns from the predefined schema that are missing will
                            be added. No columns in the existing table will be deleted.
                            If unspecified, then all available columns will be used
                            for a new table or an (existing) table with no schema,
                            and no changes will be made to an existing table that
                            has a schema. Only for use with external storage. Possible
                            values: OUTPUT_SCHEMA_UNSPECIFIED, BASIC_COLUMNS, GCS_COLUMNS,
                            DATASTORE_COLUMNS, BIG_QUERY_COLUMNS, ALL_COLUMNS'
                          type: string
                        table:
                          description: 'Store findings in an existing table or a new
                            table in an existing dataset. If table_id is not set a
                            new one will be generated for you with the following format:
                            dlp_googleapis_yyyy_mm_dd_[dlp_job_id]. Pacific timezone
                            will be used for generating the date details. For Inspect,
                            each column in an existing output table must have the
                            same name, type, and mode of a field in the `Finding`
                            object. For Risk, an existing output table should be the
                            output of a previous Risk analysis job run on the same
                            source table, with the same privacy metric and quasi-identifiers.
                            Risk jobs that analyze the same table but compute a different
                            privacy metric, or use different sets of quasi-identifiers,
                            cannot store their results in the same table.'
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
                      type: object
                  type: object
              type: object
            type: array
          inspectJob:
            description: For inspect jobs, a snapshot of the configuration.
            required:
            - storageConfig
            type: object
required:
- spec
type: object

```
