# openAPI schema part2 for datastream.cnrm.cloud.google.com.v1alpha1.DatastreamStream

## schema

```yaml
properties:
  spec:
    properties:
      sourceConfig:
        properties:
          sourceConfig:
            description: Source connection profile configuration.
            properties:
              mysqlSourceConfig:
                description: MySQL data source configuration.
                properties:
                  excludeObjects:
                    description: MySQL objects to exclude from the stream.
                    properties:
                      mysqlDatabases:
                        description: MySQL databases on the server.
                        items:
                          properties:
                            database:
                              description: Database name.
                              type: string
                            mysqlTables:
                              description: Tables in the database.
                              items:
                                properties:
                                  mysqlColumns:
                                    description: MySQL columns in the schema. When
                                      unspecified as part of include/exclude objects,
                                      includes/excludes everything.
                                    items:
                                      properties:
                                        collation:
                                          description: Column collation.
                                          type: string
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The MySQL data type. Full data types list can be found here:
                                            https://dev.mysql.com/doc/refman/8.0/en/data-types.html.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                          required:
                          - database
                          type: object
                        type: array
                    required:
                    - mysqlDatabases
                    type: object
                  includeObjects:
                    description: MySQL objects to retrieve from the source.
                    properties:
                      mysqlDatabases:
                        description: MySQL databases on the server.
                        items:
                          properties:
                            database:
                              description: Database name.
                              type: string
                            mysqlTables:
                              description: Tables in the database.
                              items:
                                properties:
                                  mysqlColumns:
                                    description: MySQL columns in the schema. When
                                      unspecified as part of include/exclude objects,
                                      includes/excludes everything.
                                    items:
                                      properties:
                                        collation:
                                          description: Column collation.
                                          type: string
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The MySQL data type. Full data types list can be found here:
                                            https://dev.mysql.com/doc/refman/8.0/en/data-types.html.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                          required:
                          - database
                          type: object
                        type: array
                    required:
                    - mysqlDatabases
                    type: object
                  maxConcurrentBackfillTasks:
                    description: |-
                      Maximum number of concurrent backfill tasks. The number should be non negative.
                      If not set (or set to 0), the system's default value will be used.
                    type: integer
                  maxConcurrentCdcTasks:
                    description: |-
                      Maximum number of concurrent CDC tasks. The number should be non negative.
                      If not set (or set to 0), the system's default value will be used.
                    type: integer
                type: object
              oracleSourceConfig:
                description: MySQL data source configuration.
                properties:
                  dropLargeObjects:
                    description: Configuration to drop large object values.
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                  excludeObjects:
                    description: Oracle objects to exclude from the stream.
                    properties:
                      oracleSchemas:
                        description: Oracle schemas/databases in the database server.
                        items:
                          properties:
                            oracleTables:
                              description: Tables in the database.
                              items:
                                properties:
                                  oracleColumns:
                                    description: Oracle columns in the schema. When
                                      unspecified as part of include/exclude objects,
                                      includes/excludes everything.
                                    items:
                                      properties:
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The Oracle data type. Full data types list can be found here:
                                            https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html.
                                          type: string
                                        encoding:
                                          description: Column encoding.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        precision:
                                          description: Column precision.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                        scale:
                                          description: Column scale.
                                          type: integer
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                            schema:
                              description: Schema name.
                              type: string
                          required:
                          - schema
                          type: object
                        type: array
                    required:
                    - oracleSchemas
                    type: object
                  includeObjects:
                    description: Oracle objects to retrieve from the source.
                    properties:
                      oracleSchemas:
                        description: Oracle schemas/databases in the database server.
                        items:
                          properties:
                            oracleTables:
                              description: Tables in the database.
                              items:
                                properties:
                                  oracleColumns:
                                    description: Oracle columns in the schema. When
                                      unspecified as part of include/exclude objects,
                                      includes/excludes everything.
                                    items:
                                      properties:
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The Oracle data type. Full data types list can be found here:
                                            https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/Data-Types.html.
                                          type: string
                                        encoding:
                                          description: Column encoding.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        precision:
                                          description: Column precision.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                        scale:
                                          description: Column scale.
                                          type: integer
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                            schema:
                              description: Schema name.
                              type: string
                          required:
                          - schema
                          type: object
                        type: array
                    required:
                    - oracleSchemas
                    type: object
                  maxConcurrentBackfillTasks:
                    description: |-
                      Maximum number of concurrent backfill tasks. The number should be non negative.
                      If not set (or set to 0), the system's default value will be used.
                    type: integer
                  maxConcurrentCdcTasks:
                    description: |-
                      Maximum number of concurrent CDC tasks. The number should be non negative.
                      If not set (or set to 0), the system's default value will be used.
                    type: integer
                  streamLargeObjects:
                    description: Configuration to drop large object values.
                    type: object
                    x-kubernetes-preserve-unknown-fields: true
                type: object
              postgresqlSourceConfig:
                description: PostgreSQL data source configuration.
                properties:
                  excludeObjects:
                    description: PostgreSQL objects to exclude from the stream.
                    properties:
                      postgresqlSchemas:
                        description: PostgreSQL schemas on the server.
                        items:
                          properties:
                            postgresqlTables:
                              description: Tables in the schema.
                              items:
                                properties:
                                  postgresqlColumns:
                                    description: PostgreSQL columns in the schema.
                                      When unspecified as part of include/exclude
                                      objects, includes/excludes everything.
                                    items:
                                      properties:
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The PostgreSQL data type. Full data types list can be found here:
                                            https://www.postgresql.org/docs/current/datatype.html.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        precision:
                                          description: Column precision.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                        scale:
                                          description: Column scale.
                                          type: integer
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                            schema:
                              description: Database name.
                              type: string
                          required:
                          - schema
                          type: object
                        type: array
                    required:
                    - postgresqlSchemas
                    type: object
                  includeObjects:
                    description: PostgreSQL objects to retrieve from the source.
                    properties:
                      postgresqlSchemas:
                        description: PostgreSQL schemas on the server.
                        items:
                          properties:
                            postgresqlTables:
                              description: Tables in the schema.
                              items:
                                properties:
                                  postgresqlColumns:
                                    description: PostgreSQL columns in the schema.
                                      When unspecified as part of include/exclude
                                      objects, includes/excludes everything.
                                    items:
                                      properties:
                                        column:
                                          description: Column name.
                                          type: string
                                        dataType:
                                          description: |-
                                            The PostgreSQL data type. Full data types list can be found here:
                                            https://www.postgresql.org/docs/current/datatype.html.
                                          type: string
                                        length:
                                          description: Column length.
                                          type: integer
                                        nullable:
                                          description: Whether or not the column can
                                            accept a null value.
                                          type: boolean
                                        ordinalPosition:
                                          description: The ordinal position of the
                                            column in the table.
                                          type: integer
                                        precision:
                                          description: Column precision.
                                          type: integer
                                        primaryKey:
                                          description: Whether or not the column represents
                                            a primary key.
                                          type: boolean
                                        scale:
                                          description: Column scale.
                                          type: integer
                                      type: object
                                    type: array
                                  table:
                                    description: Table name.
                                    type: string
                                required:
                                - table
                                type: object
                              type: array
                            schema:
                              description: Database name.
                              type: string
                          required:
                          - schema
                          type: object
                        type: array
                    required:
                    - postgresqlSchemas
                    type: object
                  maxConcurrentBackfillTasks:
                    description: |-
                      Maximum number of concurrent backfill tasks. The number should be non
                      negative. If not set (or set to 0), the system's default value will be used.
                    type: integer
                  publication:
                    description: |-
                      The name of the publication that includes the set of all tables
                      that are defined in the stream's include_objects.
                    type: string
                  replicationSlot:
                    description: |-
                      The name of the logical replication slot that's configured with
                      the pgoutput plugin.
                    type: string
                required:
                - publication
                - replicationSlot
                type: object
              sourceConnectionProfile:
                description: 'Immutable. Source connection profile resource. Format:
                  projects/{project}/locations/{location}/connectionProfiles/{name}.'
                type: string
            required:
            - sourceConnectionProfile
            type: object
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
      name:
        description: The stream's name.
        type: string
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      state:
        description: The state of the stream.
        type: string
    type: object
required:
- spec
type: object

```
