# openAPI schema status for io.k8s.api.core.v1.Toleration

## description

The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>.

## schema

```yaml
|
  description: The pod this Toleration is attached to tolerates any taint that matches
    the triple <key,value,effect> using the matching operator <operator>.
  properties:
    effect:
      description: Effect indicates the taint effect to match. Empty means match all
        taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule
        and NoExecute.
      type: string
    key:
      description: Key is the taint key that the toleration applies to. Empty means
        match all taint keys. If the key is empty, operator must be Exists; this combination
        means to match all values and all keys.
      type: string
    operator:
      description: Operator represents a key's relationship to the value. Valid operators
        are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for
        value, so that a pod can tolerate all taints of a particular category.
      type: string
    tolerationSeconds:
      description: TolerationSeconds represents the period of time the toleration (which
        must be of effect NoExecute, otherwise this field is ignored) tolerates the
        taint. By default, it is not set, which means tolerate the taint forever (do
        not evict). Zero and negative values will be treated as 0 (evict immediately)
        by the system.
      format: int64
      type: integer
    value:
      description: Value is the taint value the toleration matches to. If the operator
        is Exists, the value should be empty, otherwise just a regular string.
      type: string
  type: object

```