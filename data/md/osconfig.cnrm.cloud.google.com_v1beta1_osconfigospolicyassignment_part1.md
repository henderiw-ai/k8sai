# openAPI schema part1 for osconfig.cnrm.cloud.google.com.v1beta1.OSConfigOSPolicyAssignment

## schema

```yaml
properties:
  spec:
    properties:
      description:
        description: OS policy assignment description. Length of the description is
          limited to 1024 characters.
        type: string
      instanceFilter:
        description: Required. Filter to select VMs.
        properties:
          all:
            description: Target all VMs in the project. If true, no other criteria
              is permitted.
            type: boolean
          exclusionLabels:
            description: List of label sets used for VM exclusion. If the list has
              more than one label set, the VM is excluded if any of the label sets
              are applicable for the VM.
            items:
              properties:
                labels:
                  additionalProperties:
                    type: string
                  description: Labels are identified by key/value pairs in this map.
                    A VM should contain all the key/value pairs specified in this
                    map to be selected.
                  type: object
              type: object
            type: array
          inclusionLabels:
            description: List of label sets used for VM inclusion. If the list has
              more than one `LabelSet`, the VM is included if any of the label sets
              are applicable for the VM.
            items:
              properties:
                labels:
                  additionalProperties:
                    type: string
                  description: Labels are identified by key/value pairs in this map.
                    A VM should contain all the key/value pairs specified in this
                    map to be selected.
                  type: object
              type: object
            type: array
          inventories:
            description: List of inventories to select VMs. A VM is selected if its
              inventory data matches at least one of the following inventories.
            items:
              properties:
                osShortName:
                  description: Required. The OS short name
                  type: string
                osVersion:
                  description: The OS version Prefix matches are supported if asterisk(*)
                    is provided as the last character. For example, to match all versions
                    with a major version of `7`, specify the following value for this
                    field `7.*` An empty string matches all OS versions.
                  type: string
              required:
              - osShortName
              type: object
            type: array
        type: object
      location:
        description: Immutable. The location for the resource
        type: string
      spec:
        required:
        - instanceFilter
        - location
        - osPolicies
        - projectRef
        - rollout
        type: object
required:
- spec
type: object

```
