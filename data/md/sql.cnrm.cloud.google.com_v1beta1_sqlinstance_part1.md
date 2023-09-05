# openAPI schema part1 for sql.cnrm.cloud.google.com.v1beta1.SQLInstance

## schema

```yaml
properties:
  spec:
    properties:
      databaseVersion:
        default: MYSQL_5_6
        description: The MySQL, PostgreSQL or SQL Server (beta) version to use. Supported
          values include MYSQL_5_6, MYSQL_5_7, MYSQL_8_0, POSTGRES_9_6, POSTGRES_10,
          POSTGRES_11, POSTGRES_12, POSTGRES_13, POSTGRES_14, POSTGRES_15, SQLSERVER_2017_STANDARD,
          SQLSERVER_2017_ENTERPRISE, SQLSERVER_2017_EXPRESS, SQLSERVER_2017_WEB. Database
          Version Policies includes an up-to-date reference of supported versions.
        type: string
      encryptionKMSCryptoKeyRef:
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
      instanceType:
        description: The type of the instance. The valid values are:- 'SQL_INSTANCE_TYPE_UNSPECIFIED',
          'CLOUD_SQL_INSTANCE', 'ON_PREMISES_INSTANCE' and 'READ_REPLICA_INSTANCE'.
        type: string
      maintenanceVersion:
        description: Maintenance version.
        type: string
      masterInstanceRef:
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
            description: 'Allowed value: The `name` field of a `SQLInstance` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      region:
        description: Immutable. The region the instance will sit in. Note, Cloud SQL
          is not available in all regions. A valid region must be provided to use
          this resource. If a region is not provided in the resource definition, the
          provider region will be used instead, but this will be an apply-time error
          for instances if the provider region is not supported with Cloud SQL. If
          you choose not to provide the region argument for this resource, make sure
          you understand this.
        type: string
      replicaConfiguration:
        description: The configuration for replication.
        properties:
          caCertificate:
            description: Immutable. PEM representation of the trusted CA's x509 certificate.
            type: string
          clientCertificate:
            description: Immutable. PEM representation of the replica's x509 certificate.
            type: string
          clientKey:
            description: Immutable. PEM representation of the replica's private key.
              The corresponding public key in encoded in the client_certificate.
            type: string
          connectRetryInterval:
            description: Immutable. The number of seconds between connect retries.
              MySQL's default is 60 seconds.
            type: integer
          dumpFilePath:
            description: Immutable. Path to a SQL file in Google Cloud Storage from
              which replica instances are created. Format is gs://bucket/filename.
            type: string
          failoverTarget:
            description: Immutable. Specifies if the replica is the failover target.
              If the field is set to true the replica will be designated as a failover
              replica. If the master instance fails, the replica instance will be
              promoted as the new master instance.
            type: boolean
          masterHeartbeatPeriod:
            description: Immutable. Time in ms between replication heartbeats.
            type: integer
          password:
            description: Immutable. Password for the replication connection.
            oneOf:
            - not:
                required:
                - valueFrom
              required:
              - value
            - not:
                required:
                - value
              required:
              - valueFrom
            properties:
              value:
                description: Value of the field. Cannot be used if 'valueFrom' is
                  specified.
                type: string
              valueFrom:
                description: Source for the field's value. Cannot be used if 'value'
                  is specified.
                properties:
                  secretKeyRef:
                    description: Reference to a value with the given key in the given
                      Secret in the resource's namespace.
                    properties:
                      key:
                        description: Key that identifies the value to be extracted.
                        type: string
                      name:
                        description: Name of the Secret to extract a value from.
                        type: string
                    required:
                    - name
                    - key
                    type: object
                type: object
            type: object
          sslCipher:
            description: Immutable. Permissible ciphers for use in SSL encryption.
            type: string
          username:
            description: Immutable. Username for replication connection.
            type: string
          verifyServerCertificate:
            description: Immutable. True if the master's common name value is checked
              during the SSL handshake.
            type: boolean
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
      rootPassword:
        description: Initial root password. Required for MS SQL Server.
        oneOf:
        - not:
            required:
            - valueFrom
          required:
          - value
        - not:
            required:
            - value
          required:
          - valueFrom
        properties:
          value:
            description: Value of the field. Cannot be used if 'valueFrom' is specified.
            type: string
          valueFrom:
            description: Source for the field's value. Cannot be used if 'value' is
              specified.
            properties:
              secretKeyRef:
                description: Reference to a value with the given key in the given
                  Secret in the resource's namespace.
                properties:
                  key:
                    description: Key that identifies the value to be extracted.
                    type: string
                  name:
                    description: Name of the Secret to extract a value from.
                    type: string
                required:
                - name
                - key
                type: object
            type: object
        type: object
      spec:
        required:
        - settings
        type: object
required:
- spec
type: object

```
