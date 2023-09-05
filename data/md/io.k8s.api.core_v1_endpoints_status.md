# openAPI schema status for io.k8s.api.core.v1.Endpoints

## description

Endpoints is a collection of endpoints that implement the actual service. Example:

	 Name: "mysvc",
	 Subsets: [
	   {
	     Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
	     Ports: [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
	   },
	   {
	     Addresses: [{"ip": "10.10.3.3"}],
	     Ports: [{"name": "a", "port": 93}, {"name": "b", "port": 76}]
	   },
	]

## schema

```yaml
|
  description: "Endpoints is a collection of endpoints that implement the actual service.
    Example:\n\n\t Name: \"mysvc\",\n\t Subsets: [\n\t   {\n\t     Addresses: [{\"ip\":
    \"10.10.1.1\"}, {\"ip\": \"10.10.2.2\"}],\n\t     Ports: [{\"name\": \"a\", \"port\":
    8675}, {\"name\": \"b\", \"port\": 309}]\n\t   },\n\t   {\n\t     Addresses: [{\"ip\":
    \"10.10.3.3\"}],\n\t     Ports: [{\"name\": \"a\", \"port\": 93}, {\"name\": \"b\",
    \"port\": 76}]\n\t   },\n\t]"
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
      $ref: '#/definitions/io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta'
      description: 'Standard object''s metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata'
    subsets:
      description: The set of all endpoints is the union of all subsets. Addresses are
        placed into subsets according to the IPs they share. A single address with multiple
        ports, some of which are ready and some of which are not (because they come
        from different containers) will result in the address being displayed in different
        subsets for the different ports. No address will appear in both Addresses and
        NotReadyAddresses in the same subset. Sets of addresses and ports that comprise
        a service.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.EndpointSubset'
      type: array
  type: object
  x-kubernetes-group-version-kind:
  - group: ""
    kind: Endpoints
    version: v1

```