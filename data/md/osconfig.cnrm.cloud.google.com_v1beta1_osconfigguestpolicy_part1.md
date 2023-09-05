# openAPI schema part1 for osconfig.cnrm.cloud.google.com.v1beta1.OSConfigGuestPolicy

## schema

```yaml
properties:
  spec:
    properties:
      assignment:
        description: Specifies the VMs that are assigned this policy. This allows
          you to target sets or groups of VMs by different parameters such as labels,
          names, OS, or zones. Empty assignments will target ALL VMs underneath this
          policy. Conflict Management Policies that exist higher up in the resource
          hierarchy (closer to the Org) will override those lower down if there is
          a conflict. At the same level in the resource hierarchy (ie. within a project),
          the service will prevent the creation of multiple policies that conflict
          with each other. If there are multiple policies that specify the same config
          (eg. package, software recipe, repository, etc.), the service will ensure
          that no VM could potentially receive instructions from both policies. To
          create multiple policies that specify different versions of a package or
          different configs for different Operating Systems, each policy must be mutually
          exclusive in their targeting according to labels, OS, or other criteria.
          Different configs are identified for conflicts in different ways. Packages
          are identified by their name and the package manager(s) they target. Package
          repositories are identified by their unique id where applicable. Some package
          managers don't have a unique identifier for repositories and where that's
          the case, no uniqueness is validated by the service. Note that if OS Inventory
          is disabled, a VM will not be assigned a policy that targets by OS because
          the service will see this VM's OS as unknown.
        properties:
          groupLabels:
            description: Targets instances matching at least one of these label sets.
              This allows an assignment to target disparate groups, for example "env=prod
              or env=staging".
            items:
              properties:
                labels:
                  additionalProperties:
                    type: string
                  description: Google Compute Engine instance labels that must be
                    present for an instance to be included in this assignment group.
                  type: object
              type: object
            type: array
          instanceNamePrefixes:
            description: Targets VM instances whose name starts with one of these
              prefixes. Like labels, this is another way to group VM instances when
              targeting configs, for example prefix="prod-". Only supported for project-level
              policies.
            items:
              type: string
            type: array
          instances:
            items:
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
                  description: 'Allowed value: The `selfLink` field of a `ComputeInstance`
                    resource.'
                  type: string
                name:
                  description: 'Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names'
                  type: string
                namespace:
                  description: 'Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/'
                  type: string
              type: object
            type: array
          osTypes:
            description: Targets VM instances matching at least one of the following
              OS types. VM instances must match all supplied criteria for a given
              OsType to be included.
            items:
              properties:
                osArchitecture:
                  description: Targets VM instances with OS Inventory enabled and
                    having the following OS architecture.
                  type: string
                osShortName:
                  description: Targets VM instances with OS Inventory enabled and
                    having the following OS short name, for example "debian" or "windows".
                  type: string
                osVersion:
                  description: Targets VM instances with OS Inventory enabled and
                    having the following following OS version.
                  type: string
              type: object
            type: array
          zones:
            description: Targets instances in any of these zones. Leave empty to target
              instances in any zone. Zonal targeting is uncommon and is supported
              to facilitate the management of changes by zone.
            items:
              type: string
            type: array
        type: object
      description:
        description: Description of the GuestPolicy. Length of the description is
          limited to 1024 characters.
        type: string
      packageRepositories:
        description: List of package repository configurations assigned to the VM
          instance.
        items:
          properties:
            apt:
              description: An Apt Repository.
              properties:
                archiveType:
                  description: 'Type of archive files in this repository. The default
                    behavior is DEB. Possible values: ARCHIVE_TYPE_UNSPECIFIED, DEB,
                    DEB_SRC'
                  type: string
                components:
                  description: Required. List of components for this repository. Must
                    contain at least one item.
                  items:
                    type: string
                  type: array
                distribution:
                  description: Required. Distribution of this repository.
                  type: string
                gpgKey:
                  description: URI of the key file for this repository. The agent
                    maintains a keyring at `/etc/apt/trusted.gpg.d/osconfig_agent_managed.gpg`
                    containing all the keys in any applied guest policy.
                  type: string
                uri:
                  description: Required. URI for this repository.
                  type: string
              required:
              - distribution
              - uri
              type: object
            goo:
              description: A Goo Repository.
              properties:
                name:
                  description: Required. The name of the repository.
                  type: string
                url:
                  description: Required. The url of the repository.
                  type: string
              required:
              - name
              - url
              type: object
            yum:
              description: A Yum Repository.
              properties:
                baseUrl:
                  description: Required. The location of the repository directory.
                  type: string
                displayName:
                  description: The display name of the repository.
                  type: string
                gpgKeys:
                  description: URIs of GPG keys.
                  items:
                    type: string
                  type: array
                id:
                  description: Required. A one word, unique name for this repository.
                    This is the `repo id` in the Yum config file and also the `display_name`
                    if `display_name` is omitted. This id is also used as the unique
                    identifier when checking for guest policy conflicts.
                  type: string
              required:
              - baseUrl
              - id
              type: object
            zypper:
              description: A Zypper Repository.
              properties:
                baseUrl:
                  description: Required. The location of the repository directory.
                  type: string
                displayName:
                  description: The display name of the repository.
                  type: string
                gpgKeys:
                  description: URIs of GPG keys.
                  items:
                    type: string
                  type: array
                id:
                  description: Required. A one word, unique name for this repository.
                    This is the `repo id` in the zypper config file and also the `display_name`
                    if `display_name` is omitted. This id is also used as the unique
                    identifier when checking for guest policy conflicts.
                  type: string
              required:
              - baseUrl
              - id
              type: object
          type: object
        type: array
      packages:
        description: List of package configurations assigned to the VM instance.
        items:
          properties:
            desiredState:
              description: 'The desired_state the agent should maintain for this package.
                The default is to ensure the package is installed. Possible values:
                DESIRED_STATE_UNSPECIFIED, INSTALLED, REMOVED'
              type: string
            manager:
              description: 'Type of package manager that can be used to install this
                package. If a system does not have the package manager, the package
                is not installed or removed no error message is returned. By default,
                or if you specify `ANY`, the agent attempts to install and remove
                this package using the default package manager. This is useful when
                creating a policy that applies to different types of systems. The
                default behavior is ANY. Possible values: MANAGER_UNSPECIFIED, ANY,
                APT, YUM, ZYPPER, GOO'
              type: string
            name:
              description: Required. The name of the package. A package is uniquely
                identified for conflict validation by checking the package name and
                the manager(s) that the package targets.
              type: string
          type: object
        type: array
      spec:
        type: object
type: object

```
