# openAPI schema part0 for certificatemanager.cnrm.cloud.google.com.v1alpha1.CertificateManagerCertificateMap

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
      description:
        description: A human-readable description of the resource.
        type: string
      projectRef:
        description: The project that this resource belongs to.
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
            description: 'Allowed value: The `name` field of a `Project` resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      resourceID:
        description: Immutable. Optional. The name of the resource. Used for creation
          and acquisition. When unset, the value of `metadata.name` is used as the
          default.
        type: string
    required:
    - projectRef
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
      createTime:
        description: |-
          Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
          accurate to nanoseconds with up to nine fractional digits.
          Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        type: string
      gclbTargets:
        description: A list of target proxies that use this Certificate Map.
        items:
          properties:
            ipConfigs:
              description: An IP configuration where this Certificate Map is serving.
              items:
                properties:
                  ipAddress:
                    description: An external IP address.
                    type: string
                  ports:
                    description: A list of ports.
                    items:
                      type: integer
                    type: array
                type: object
              type: array
            targetHttpsProxy:
              description: |-
                Proxy name must be in the format projects/*/locations/*/targetHttpsProxies/*.
                This field is part of a union field 'target_proxy': Only one of 'targetHttpsProxy' or
                'targetSslProxy' may be set.
              type: string
            targetSslProxy:
              description: |-
                Proxy name must be in the format projects/*/locations/*/targetSslProxies/*.
                This field is part of a union field 'target_proxy': Only one of 'targetHttpsProxy' or
                'targetSslProxy' may be set.
              type: string
          type: object
        type: array
      observedGeneration:
        description: ObservedGeneration is the generation of the resource that was
          most recently observed by the Config Connector controller. If this is equal
          to metadata.generation, then that means that the current reported status
          reflects the most recent desired state of the resource.
        type: integer
      updateTime:
        description: |-
          Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
          accurate to nanoseconds with up to nine fractional digits.
          Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        type: string
    type: object
required:
- spec
type: object

```