# openAPI schema part2 for dataproc.cnrm.cloud.google.com.v1beta1.DataprocWorkflowTemplate

## schema

```yaml
properties:
  spec:
    properties:
      jobs:
        properties:
          jobs:
            description: Immutable. Required. The Directed Acyclic Graph of Jobs to
              submit.
            items:
              properties:
                hadoopJob:
                  description: Immutable. Optional. Job is a Hadoop job.
                  properties:
                    archiveUris:
                      description: 'Immutable. Optional. HCFS URIs of archives to
                        be extracted in the working directory of Hadoop drivers and
                        tasks. Supported file types: .jar, .tar, .tar.gz, .tgz, or
                        .zip.'
                      items:
                        type: string
                      type: array
                    args:
                      description: Immutable. Optional. The arguments to pass to the
                        driver. Do not include arguments, such as `-libjars` or `-Dfoo=bar`,
                        that can be set as job properties, since a collision may occur
                        that causes an incorrect job submission.
                      items:
                        type: string
                      type: array
                    fileUris:
                      description: Immutable. Optional. HCFS (Hadoop Compatible Filesystem)
                        URIs of files to be copied to the working directory of Hadoop
                        drivers and distributed tasks. Useful for naively parallel
                        tasks.
                      items:
                        type: string
                      type: array
                    jarFileUris:
                      description: Immutable. Optional. Jar file URIs to add to the
                        CLASSPATHs of the Hadoop driver and tasks.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    mainClass:
                      description: Immutable. The name of the driver's main class.
                        The jar file containing the class must be in the default CLASSPATH
                        or specified in `jar_file_uris`.
                      type: string
                    mainJarFileUri:
                      description: 'Immutable. The HCFS URI of the jar file containing
                        the main class. Examples: ''gs://foo-bucket/analytics-binaries/extract-useful-metrics-mr.jar''
                        ''hdfs:/tmp/test-samples/custom-wordcount.jar'' ''file:///home/usr/lib/hadoop-mapreduce/hadoop-mapreduce-examples.jar'''
                      type: string
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure Hadoop. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/hadoop/conf/*-site and classes
                        in user code.
                      type: object
                  type: object
                hiveJob:
                  description: Immutable. Optional. Job is a Hive job.
                  properties:
                    continueOnFailure:
                      description: Immutable. Optional. Whether to continue executing
                        queries if a query fails. The default value is `false`. Setting
                        to `true` can be useful when executing independent parallel
                        queries.
                      type: boolean
                    jarFileUris:
                      description: Immutable. Optional. HCFS URIs of jar files to
                        add to the CLASSPATH of the Hive server and Hadoop MapReduce
                        (MR) tasks. Can contain Hive SerDes and UDFs.
                      items:
                        type: string
                      type: array
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        and values, used to configure Hive. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/hadoop/conf/*-site.xml, /etc/hive/conf/hive-site.xml,
                        and classes in user code.
                      type: object
                    queryFileUri:
                      description: Immutable. The HCFS URI of the script that contains
                        Hive queries.
                      type: string
                    queryList:
                      description: Immutable. A list of queries.
                      properties:
                        queries:
                          description: 'Immutable. Required. The queries to execute.
                            You do not need to end a query expression with a semicolon.
                            Multiple queries can be specified in one string by separating
                            each with a semicolon. Here is an example of a Dataproc
                            API snippet that uses a QueryList to specify a HiveJob:
                            "hiveJob": { "queryList": { "queries": [ "query1", "query2",
                            "query3;query4", ] } }'
                          items:
                            type: string
                          type: array
                      required:
                      - queries
                      type: object
                    scriptVariables:
                      additionalProperties:
                        type: string
                      description: 'Immutable. Optional. Mapping of query variable
                        names to values (equivalent to the Hive command: `SET name="value";`).'
                      type: object
                  type: object
                labels:
                  additionalProperties:
                    type: string
                  description: 'Immutable. Optional. The labels to associate with
                    this job. Label keys must be between 1 and 63 characters long,
                    and must conform to the following regular expression: p{Ll}p{Lo}{0,62}
                    Label values must be between 1 and 63 characters long, and must
                    conform to the following regular expression: [p{Ll}p{Lo}p{N}_-]{0,63}
                    No more than 32 labels can be associated with a given job.'
                  type: object
                pigJob:
                  description: Immutable. Optional. Job is a Pig job.
                  properties:
                    continueOnFailure:
                      description: Immutable. Optional. Whether to continue executing
                        queries if a query fails. The default value is `false`. Setting
                        to `true` can be useful when executing independent parallel
                        queries.
                      type: boolean
                    jarFileUris:
                      description: Immutable. Optional. HCFS URIs of jar files to
                        add to the CLASSPATH of the Pig Client and Hadoop MapReduce
                        (MR) tasks. Can contain Pig UDFs.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure Pig. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/hadoop/conf/*-site.xml, /etc/pig/conf/pig.properties,
                        and classes in user code.
                      type: object
                    queryFileUri:
                      description: Immutable. The HCFS URI of the script that contains
                        the Pig queries.
                      type: string
                    queryList:
                      description: Immutable. A list of queries.
                      properties:
                        queries:
                          description: 'Immutable. Required. The queries to execute.
                            You do not need to end a query expression with a semicolon.
                            Multiple queries can be specified in one string by separating
                            each with a semicolon. Here is an example of a Dataproc
                            API snippet that uses a QueryList to specify a HiveJob:
                            "hiveJob": { "queryList": { "queries": [ "query1", "query2",
                            "query3;query4", ] } }'
                          items:
                            type: string
                          type: array
                      required:
                      - queries
                      type: object
                    scriptVariables:
                      additionalProperties:
                        type: string
                      description: 'Immutable. Optional. Mapping of query variable
                        names to values (equivalent to the Pig command: `name=[value]`).'
                      type: object
                  type: object
                prerequisiteStepIds:
                  description: Immutable. Optional. The optional list of prerequisite
                    job step_ids. If not specified, the job will start at the beginning
                    of workflow.
                  items:
                    type: string
                  type: array
                prestoJob:
                  description: Immutable. Optional. Job is a Presto job.
                  properties:
                    clientTags:
                      description: Immutable. Optional. Presto client tags to attach
                        to this query
                      items:
                        type: string
                      type: array
                    continueOnFailure:
                      description: Immutable. Optional. Whether to continue executing
                        queries if a query fails. The default value is `false`. Setting
                        to `true` can be useful when executing independent parallel
                        queries.
                      type: boolean
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    outputFormat:
                      description: Immutable. Optional. The format in which query
                        output will be displayed. See the Presto documentation for
                        supported output formats
                      type: string
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values. Used to set Presto [session properties](https://prestodb.io/docs/current/sql/set-session.html)
                        Equivalent to using the --session flag in the Presto CLI
                      type: object
                    queryFileUri:
                      description: Immutable. The HCFS URI of the script that contains
                        SQL queries.
                      type: string
                    queryList:
                      description: Immutable. A list of queries.
                      properties:
                        queries:
                          description: 'Immutable. Required. The queries to execute.
                            You do not need to end a query expression with a semicolon.
                            Multiple queries can be specified in one string by separating
                            each with a semicolon. Here is an example of a Dataproc
                            API snippet that uses a QueryList to specify a HiveJob:
                            "hiveJob": { "queryList": { "queries": [ "query1", "query2",
                            "query3;query4", ] } }'
                          items:
                            type: string
                          type: array
                      required:
                      - queries
                      type: object
                  type: object
                pysparkJob:
                  description: Immutable. Optional. Job is a PySpark job.
                  properties:
                    archiveUris:
                      description: 'Immutable. Optional. HCFS URIs of archives to
                        be extracted into the working directory of each executor.
                        Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.'
                      items:
                        type: string
                      type: array
                    args:
                      description: Immutable. Optional. The arguments to pass to the
                        driver. Do not include arguments, such as `--conf`, that can
                        be set as job properties, since a collision may occur that
                        causes an incorrect job submission.
                      items:
                        type: string
                      type: array
                    fileUris:
                      description: Immutable. Optional. HCFS URIs of files to be placed
                        in the working directory of each executor. Useful for naively
                        parallel tasks.
                      items:
                        type: string
                      type: array
                    jarFileUris:
                      description: Immutable. Optional. HCFS URIs of jar files to
                        add to the CLASSPATHs of the Python driver and tasks.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    mainPythonFileUri:
                      description: Immutable. Required. The HCFS URI of the main Python
                        file to use as the driver. Must be a .py file.
                      type: string
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure PySpark. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/spark/conf/spark-defaults.conf
                        and classes in user code.
                      type: object
                    pythonFileUris:
                      description: 'Immutable. Optional. HCFS file URIs of Python
                        files to pass to the PySpark framework. Supported file types:
                        .py, .egg, and .zip.'
                      items:
                        type: string
                      type: array
                  required:
                  - mainPythonFileUri
                  type: object
                scheduling:
                  description: Immutable. Optional. Job scheduling configuration.
                  properties:
                    maxFailuresPerHour:
                      description: Immutable. Optional. Maximum number of times per
                        hour a driver may be restarted as a result of driver exiting
                        with non-zero code before job is reported failed. A job may
                        be reported as thrashing if driver exits with non-zero code
                        4 times within 10 minute window. Maximum value is 10.
                      format: int64
                      type: integer
                    maxFailuresTotal:
                      description: Immutable. Optional. Maximum number of times in
                        total a driver may be restarted as a result of driver exiting
                        with non-zero code before job is reported failed. Maximum
                        value is 240.
                      format: int64
                      type: integer
                  type: object
                sparkJob:
                  description: Immutable. Optional. Job is a Spark job.
                  properties:
                    archiveUris:
                      description: 'Immutable. Optional. HCFS URIs of archives to
                        be extracted into the working directory of each executor.
                        Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.'
                      items:
                        type: string
                      type: array
                    args:
                      description: Immutable. Optional. The arguments to pass to the
                        driver. Do not include arguments, such as `--conf`, that can
                        be set as job properties, since a collision may occur that
                        causes an incorrect job submission.
                      items:
                        type: string
                      type: array
                    fileUris:
                      description: Immutable. Optional. HCFS URIs of files to be placed
                        in the working directory of each executor. Useful for naively
                        parallel tasks.
                      items:
                        type: string
                      type: array
                    jarFileUris:
                      description: Immutable. Optional. HCFS URIs of jar files to
                        add to the CLASSPATHs of the Spark driver and tasks.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    mainClass:
                      description: Immutable. The name of the driver's main class.
                        The jar file that contains the class must be in the default
                        CLASSPATH or specified in `jar_file_uris`.
                      type: string
                    mainJarFileUri:
                      description: Immutable. The HCFS URI of the jar file that contains
                        the main class.
                      type: string
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure Spark. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/spark/conf/spark-defaults.conf
                        and classes in user code.
                      type: object
                  type: object
                sparkRJob:
                  description: Immutable. Optional. Job is a SparkR job.
                  properties:
                    archiveUris:
                      description: 'Immutable. Optional. HCFS URIs of archives to
                        be extracted into the working directory of each executor.
                        Supported file types: .jar, .tar, .tar.gz, .tgz, and .zip.'
                      items:
                        type: string
                      type: array
                    args:
                      description: Immutable. Optional. The arguments to pass to the
                        driver. Do not include arguments, such as `--conf`, that can
                        be set as job properties, since a collision may occur that
                        causes an incorrect job submission.
                      items:
                        type: string
                      type: array
                    fileUris:
                      description: Immutable. Optional. HCFS URIs of files to be placed
                        in the working directory of each executor. Useful for naively
                        parallel tasks.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    mainRFileUri:
                      description: Immutable. Required. The HCFS URI of the main R
                        file to use as the driver. Must be a .R file.
                      type: string
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure SparkR. Properties that conflict
                        with values set by the Dataproc API may be overwritten. Can
                        include properties set in /etc/spark/conf/spark-defaults.conf
                        and classes in user code.
                      type: object
                  required:
                  - mainRFileUri
                  type: object
                sparkSqlJob:
                  description: Immutable. Optional. Job is a SparkSql job.
                  properties:
                    jarFileUris:
                      description: Immutable. Optional. HCFS URIs of jar files to
                        be added to the Spark CLASSPATH.
                      items:
                        type: string
                      type: array
                    loggingConfig:
                      description: Immutable. Optional. The runtime log config for
                        job execution.
                      properties:
                        driverLogLevels:
                          additionalProperties:
                            type: string
                          description: 'Immutable. The per-package log levels for
                            the driver. This may include "root" package name to configure
                            rootLogger. Examples: ''com.google = FATAL'', ''root =
                            INFO'', ''org.apache = DEBUG'''
                          type: object
                      type: object
                    properties:
                      additionalProperties:
                        type: string
                      description: Immutable. Optional. A mapping of property names
                        to values, used to configure Spark SQL's SparkConf. Properties
                        that conflict with values set by the Dataproc API may be overwritten.
                      type: object
                    queryFileUri:
                      description: Immutable. The HCFS URI of the script that contains
                        SQL queries.
                      type: string
                    queryList:
                      description: Immutable. A list of queries.
                      properties:
                        queries:
                          description: 'Immutable. Required. The queries to execute.
                            You do not need to end a query expression with a semicolon.
                            Multiple queries can be specified in one string by separating
                            each with a semicolon. Here is an example of a Dataproc
                            API snippet that uses a QueryList to specify a HiveJob:
                            "hiveJob": { "queryList": { "queries": [ "query1", "query2",
                            "query3;query4", ] } }'
                          items:
                            type: string
                          type: array
                      required:
                      - queries
                      type: object
                    scriptVariables:
                      additionalProperties:
                        type: string
                      description: 'Immutable. Optional. Mapping of query variable
                        names to values (equivalent to the Spark SQL command: SET
                        `name="value";`).'
                      type: object
                  type: object
                stepId:
                  description: Immutable. Required. The step id. The id must be unique
                    among all jobs within the template. The step id is used as prefix
                    for job id, as job `goog-dataproc-workflow-step-id` label, and
                    in prerequisiteStepIds field from other steps. The id must contain
                    only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens
                    (-). Cannot begin or end with underscore or hyphen. Must consist
                    of between 3 and 50 characters.
                  type: string
              required:
              - stepId
              type: object
            type: array
      location:
        description: Immutable. The location for the resource
        type: string
      parameters:
        description: Immutable. Optional. Template parameters whose values are substituted
          into the template. Values for parameters must be provided when the template
          is instantiated.
        items:
          properties:
            description:
              description: Immutable. Optional. Brief description of the parameter.
                Must not exceed 1024 characters.
              type: string
            fields:
              description: 'Immutable. Required. Paths to all fields that the parameter
                replaces. A field is allowed to appear in at most one parameter''s
                list of field paths. A field path is similar in syntax to a google.protobuf.FieldMask.
                For example, a field path that references the zone field of a workflow
                template''s cluster selector would be specified as `placement.clusterSelector.zone`.
                Also, field paths can reference fields using the following syntax:
                * Values in maps can be referenced by key: * labels[''key''] * placement.clusterSelector.clusterLabels[''key'']
                * placement.managedCluster.labels[''key''] * placement.clusterSelector.clusterLabels[''key'']
                * jobs[''step-id''].labels[''key''] * Jobs in the jobs list can be
                referenced by step-id: * jobs[''step-id''].hadoopJob.mainJarFileUri
                * jobs[''step-id''].hiveJob.queryFileUri * jobs[''step-id''].pySparkJob.mainPythonFileUri
                * jobs[''step-id''].hadoopJob.jarFileUris[0] * jobs[''step-id''].hadoopJob.archiveUris[0]
                * jobs[''step-id''].hadoopJob.fileUris[0] * jobs[''step-id''].pySparkJob.pythonFileUris[0]
                * Items in repeated fields can be referenced by a zero-based index:
                * jobs[''step-id''].sparkJob.args[0] * Other examples: * jobs[''step-id''].hadoopJob.properties[''key'']
                * jobs[''step-id''].hadoopJob.args[0] * jobs[''step-id''].hiveJob.scriptVariables[''key'']
                * jobs[''step-id''].hadoopJob.mainJarFileUri * placement.clusterSelector.zone
                It may not be possible to parameterize maps and repeated fields in
                their entirety since only individual map values and individual items
                in repeated fields can be referenced. For example, the following field
                paths are invalid: - placement.clusterSelector.clusterLabels - jobs[''step-id''].sparkJob.args'
              items:
                type: string
              type: array
            name:
              description: Immutable. Required. Parameter name. The parameter name
                is used as the key, and paired with the parameter value, which are
                passed to the template when the template is instantiated. The name
                must contain only capital letters (A-Z), numbers (0-9), and underscores
                (_), and must not start with a number. The maximum length is 40 characters.
              type: string
            validation:
              description: Immutable. Optional. Validation rules to be applied to
                this parameter's value.
              properties:
                regex:
                  description: Immutable. Validation based on regular expressions.
                  properties:
                    regexes:
                      description: Immutable. Required. RE2 regular expressions used
                        to validate the parameter's value. The value must match the
                        regex in its entirety (substring matches are not sufficient).
                      items:
                        type: string
                      type: array
                  required:
                  - regexes
                  type: object
                values:
                  description: Immutable. Validation based on a list of allowed values.
                  properties:
                    values:
                      description: Immutable. Required. List of allowed values for
                        the parameter.
                      items:
                        type: string
                      type: array
                  required:
                  - values
                  type: object
              type: object
          required:
          - fields
          - name
          type: object
        type: array
required:
- spec
type: object

```
