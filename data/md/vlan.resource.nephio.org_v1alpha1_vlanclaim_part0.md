# openAPI schema part0 for vlan.resource.nephio.org.v1alpha1.VLANClaim

## description

VLANClaim is the Schema for the vlan claim API

## schema

```yaml
description: VLANClaim is the Schema for the vlan claim API
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
    description: VLANClaimSpec defines the desired state of VLANClaim
    properties:
      labels:
        additionalProperties:
          type: string
        description: Labels as user defined labels
        type: object
      range:
        description: VLANRange defines the vlan range for the VLAN claim
        type: string
      selector:
        description: Selector defines the selector criterias
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
      vlanID:
        description: VLANID defines the vlan for the VLAN claim
        type: integer
      vlanIndex:
        description: VLANIndex defines the vlan index for the VLAN Claim
        properties:
          apiVersion:
            description: API version of the referent.
            type: string
          fieldPath:
            description: 'If referring to a piece of an object instead of an entire
              object, this string should contain a valid JSON/Go field access statement,
              such as desiredState.manifest.containers[2]. For example, if the object
              reference is to a container within a pod, this would take on a value
              like: "spec.containers{name}" (where "name" refers to the name of the
              container that triggered the event) or if no container name is specified
              "spec.containers[2]" (container with index 2 in this pod). This syntax
              is chosen only to have some well-defined way of referencing a part of
              an object. TODO: this design is not final and this field is subject
              to change in the future.'
            type: string
          kind:
            description: 'Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
          resourceVersion:
            description: 'Specific resourceVersion to which this reference is made,
              if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency'
            type: string
          uid:
            description: 'UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids'
            type: string
        type: object
        x-kubernetes-map-type: atomic
    required:
    - vlanIndex
    type: object
  status:
    description: VLANClaimStatus defines the observed state of VLANClaim
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
      expiryTime:
        description: ExpiryTime indicated when the claim expires
        type: string
      vlanID:
        description: VLANID defines the vlan ID, claimed through the VLAN backend
        type: integer
      vlanRange:
        description: VLANRange defines the vlan range, claimed through the VLAN backend
        type: string
    type: object
type: object

```