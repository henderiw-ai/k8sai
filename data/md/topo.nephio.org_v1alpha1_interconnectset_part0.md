# openAPI schema part0 for topo.nephio.org.v1alpha1.InterconnectSet

## description

Interconnect is the Schema for the interconnect API

## schema

```yaml
description: Interconnect is the Schema for the interconnect API
properties:
  apiVersion:
    description: 'APIVersion defines the versioned schema of this representation of
      an object. Servers should convert recognized schemas to the latest internal
      value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources'
    type: string
  kind:
    description: 'Kind is a string value representing the REST resource this object
      represents. Servers may infer this from the endpoint the client submits requests
      to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
    type: string
  metadata:
    type: object
  spec:
    description: InterconnectSetSpec defines the desired state of InterconnectSet
    properties:
      replicas:
        description: Number of desired interconnects. This is a pointer to distinguish
          between explicit zero and not specified. Defaults to 1.
        format: int32
        type: integer
      selector:
        description: Label selector for interconnects. Existing ReplicaSets whose
          pods are selected by this will be the ones affected by this deployment.
          It must match the pod template's labels.
        properties:
          matchExpressions:
            description: matchExpressions is a list of label selector requirements.
              The requirements are ANDed.
            items:
              description: A label selector requirement is a selector that contains
                values, a key, and an operator that relates the key and values.
              properties:
                key:
                  description: key is the label key that the selector applies to.
                  type: string
                operator:
                  description: operator represents a key's relationship to a set of
                    values. Valid operators are In, NotIn, Exists and DoesNotExist.
                  type: string
                values:
                  description: values is an array of string values. If the operator
                    is In or NotIn, the values array must be non-empty. If the operator
                    is Exists or DoesNotExist, the values array must be empty. This
                    array is replaced during a strategic merge patch.
                  items:
                    type: string
                  type: array
              required:
              - key
              - operator
              type: object
            type: array
          matchLabels:
            additionalProperties:
              type: string
            description: matchLabels is a map of {key,value} pairs. A single {key,value}
              in the matchLabels map is equivalent to an element of matchExpressions,
              whose key field is "key", the operator is "In", and the values array
              contains only "value". The requirements are ANDed.
            type: object
        type: object
        x-kubernetes-map-type: atomic
      template:
        description: Template describes the interconnects that will be created.
        properties:
          metadata:
            description: Standard object's metadata.
            type: object
          spec:
            description: Specification of the desired behavior of the interconnect.
            properties:
              links:
                description: Links define the links part of the interconnect
                items:
                  properties:
                    endpoints:
                      description: Endpoints defines exactly 2 endpoints, the first
                        entry is the local endpoint, the 2nd entry is the remote endpoint
                      items:
                        description: InterconnectLinkEndpoint
                        properties:
                          interfaceName:
                            type: string
                          nodeName:
                            description: NodeName provide the name of the node on
                              which this interconnect originates NodeName allows for
                              multi-homing if multiple endpoints of a InterconnectLink
                              reside on different nodes
                            type: string
                          selector:
                            description: A label selector is a label query over a
                              set of resources. The result of matchLabels and matchExpressions
                              are ANDed. An empty label selector matches all objects.
                              A null label selector matches no objects.
                            properties:
                              matchExpressions:
                                description: matchExpressions is a list of label selector
                                  requirements. The requirements are ANDed.
                                items:
                                  description: A label selector requirement is a selector
                                    that contains values, a key, and an operator that
                                    relates the key and values.
                                  properties:
                                    key:
                                      description: key is the label key that the selector
                                        applies to.
                                      type: string
                                    operator:
                                      description: operator represents a key's relationship
                                        to a set of values. Valid operators are In,
                                        NotIn, Exists and DoesNotExist.
                                      type: string
                                    values:
                                      description: values is an array of string values.
                                        If the operator is In or NotIn, the values
                                        array must be non-empty. If the operator is
                                        Exists or DoesNotExist, the values array must
                                        be empty. This array is replaced during a
                                        strategic merge patch.
                                      items:
                                        type: string
                                      type: array
                                  required:
                                  - key
                                  - operator
                                  type: object
                                type: array
                              matchLabels:
                                additionalProperties:
                                  type: string
                                description: matchLabels is a map of {key,value} pairs.
                                  A single {key,value} in the matchLabels map is equivalent
                                  to an element of matchExpressions, whose key field
                                  is "key", the operator is "In", and the values array
                                  contains only "value". The requirements are ANDed.
                                type: object
                            type: object
                            x-kubernetes-map-type: atomic
                        type: object
                      type: array
                    lag:
                      description: Lag defines if the interconnect link is part of
                        a link aggregation group A link aggregation group can be single-homed
                        ot multi-homed based on the endpoint node name.
                      type: boolean
                    name:
                      type: string
                  type: object
                type: array
              topology:
                description: Topology defines the topology to which this interconnect
                  applies
                type: string
            required:
            - topology
            type: object
        type: object
    required:
    - selector
    - template
    type: object
  status:
    description: InterconnectSetStatus defines the observed state of Interconnect
    properties:
      conditions:
        description: Conditions of the resource.
        items:
          properties:
            lastTransitionTime:
              description: lastTransitionTime is the last time the condition transitioned
                from one status to another. This should be when the underlying condition
                changed.  If that is not known, then using the time when the API field
                changed is acceptable.
              format: date-time
              type: string
            message:
              description: message is a human readable message indicating details
                about the transition. This may be an empty string.
              maxLength: 32768
              type: string
            observedGeneration:
              description: observedGeneration represents the .metadata.generation
                that the condition was set based upon. For instance, if .metadata.generation
                is currently 12, but the .status.conditions[x].observedGeneration
                is 9, the condition is out of date with respect to the current state
                of the instance.
              format: int64
              minimum: 0
              type: integer
            reason:
              description: reason contains a programmatic identifier indicating the
                reason for the condition's last transition. Producers of specific
                condition types may define expected values and meanings for this field,
                and whether the values are considered a guaranteed API. The value
                should be a CamelCase string. This field may not be empty.
              maxLength: 1024
              minLength: 1
              pattern: ^[A-Za-z]([A-Za-z0-9_,:]*[A-Za-z0-9_])?$
              type: string
            status:
              description: status of the condition, one of True, False, Unknown.
              enum:
              - "True"
              - "False"
              - Unknown
              type: string
            type:
              description: type of condition in CamelCase or in foo.example.com/CamelCase.
                --- Many .condition.type values are consistent across resources like
                Available, but because arbitrary conditions can be useful (see .node.status.conditions),
                the ability to deconflict is important. The regex it matches is (dns1123SubdomainFmt/)?(qualifiedNameFmt)
              maxLength: 316
              pattern: ^([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])$
              type: string
          required:
          - lastTransitionTime
          - message
          - reason
          - status
          - type
          type: object
        type: array
    type: object
type: object

```
