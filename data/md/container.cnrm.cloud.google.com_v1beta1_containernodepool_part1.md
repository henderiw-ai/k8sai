# openAPI schema part1 for container.cnrm.cloud.google.com.v1beta1.ContainerNodePool

## schema

```yaml
properties:
  spec:
    properties:
      autoscaling:
        description: Configuration required by cluster autoscaler to adjust the size
          of the node pool to the current cluster usage. To disable autoscaling, set
          minNodeCount and maxNodeCount to 0.
        properties:
          locationPolicy:
            description: Location policy specifies the algorithm used when scaling-up
              the node pool. "BALANCED" - Is a best effort policy that aims to balance
              the sizes of available zones. "ANY" - Instructs the cluster autoscaler
              to prioritize utilization of unused reservations, and reduces preemption
              risk for Spot VMs.
            type: string
          maxNodeCount:
            description: Maximum number of nodes per zone in the node pool. Must be
              >= min_node_count. Cannot be used with total limits.
            type: integer
          minNodeCount:
            description: Minimum number of nodes per zone in the node pool. Must be
              >=0 and <= max_node_count. Cannot be used with total limits.
            type: integer
          totalMaxNodeCount:
            description: Maximum number of all nodes in the node pool. Must be >=
              total_min_node_count. Cannot be used with per zone limits.
            type: integer
          totalMinNodeCount:
            description: Minimum number of all nodes in the node pool. Must be >=0
              and <= total_max_node_count. Cannot be used with per zone limits.
            type: integer
        type: object
      clusterRef:
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
            description: 'Allowed value: The `name` field of a `ContainerCluster`
              resource.'
            type: string
          name:
            description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
            type: string
          namespace:
            description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
            type: string
        type: object
      initialNodeCount:
        description: Immutable. The initial number of nodes for the pool. In regional
          or multi-zonal clusters, this is the number of nodes per zone. Changing
          this will force recreation of the resource.
        type: integer
      location:
        description: Immutable. The location (region or zone) of the cluster.
        type: string
      management:
        description: Node management configuration, wherein auto-repair and auto-upgrade
          is configured.
        properties:
          autoRepair:
            description: Whether the nodes will be automatically repaired.
            type: boolean
          autoUpgrade:
            description: Whether the nodes will be automatically upgraded.
            type: boolean
        type: object
      maxPodsPerNode:
        description: Immutable. The maximum number of pods per node in this node pool.
          Note that this does not work on node pools which are "route-based" - that
          is, node pools belonging to clusters that do not have IP Aliasing enabled.
        type: integer
      namePrefix:
        description: Immutable. Creates a unique name for the node pool beginning
          with the specified prefix. Conflicts with name.
        type: string
      networkConfig:
        description: Networking configuration for this NodePool. If specified, it
          overrides the cluster-level defaults.
        properties:
          createPodRange:
            description: Immutable. Whether to create a new range for pod IPs in this
              node pool. Defaults are provided for pod_range and pod_ipv4_cidr_block
              if they are not specified.
            type: boolean
          enablePrivateNodes:
            description: Whether nodes have internal IP addresses only.
            type: boolean
          podCidrOverprovisionConfig:
            description: Immutable. Configuration for node-pool level pod cidr overprovision.
              If not set, the cluster level setting will be inherited.
            properties:
              disabled:
                type: boolean
            required:
            - disabled
            type: object
          podIpv4CidrBlock:
            description: Immutable. The IP address range for pod IPs in this node
              pool. Only applicable if create_pod_range is true. Set to blank to have
              a range chosen with the default size. Set to /netmask (e.g. /14) to
              have a range chosen with a specific netmask. Set to a CIDR notation
              (e.g. 10.96.0.0/14) to pick a specific range to use.
            type: string
          podRange:
            description: Immutable. The ID of the secondary range for pod IPs. If
              create_pod_range is true, this ID is used for the new range. If create_pod_range
              is false, uses an existing secondary range with this ID.
            type: string
        type: object
      spec:
        required:
        - clusterRef
        - location
        type: object
required:
- spec
type: object

```
