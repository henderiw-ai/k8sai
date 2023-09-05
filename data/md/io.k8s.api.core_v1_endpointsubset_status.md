# openAPI schema status for io.k8s.api.core.v1.EndpointSubset

## description

EndpointSubset is a group of addresses with a common set of ports. The expanded set of endpoints is the Cartesian product of Addresses x Ports. For example, given:

	{
	  Addresses: [{"ip": "10.10.1.1"}, {"ip": "10.10.2.2"}],
	  Ports:     [{"name": "a", "port": 8675}, {"name": "b", "port": 309}]
	}

The resulting set of endpoints can be viewed as:

	a: [ 10.10.1.1:8675, 10.10.2.2:8675 ],
	b: [ 10.10.1.1:309, 10.10.2.2:309 ]

## schema

```yaml
|
  description: "EndpointSubset is a group of addresses with a common set of ports. The
    expanded set of endpoints is the Cartesian product of Addresses x Ports. For example,
    given:\n\n\t{\n\t  Addresses: [{\"ip\": \"10.10.1.1\"}, {\"ip\": \"10.10.2.2\"}],\n\t
    \ Ports:     [{\"name\": \"a\", \"port\": 8675}, {\"name\": \"b\", \"port\": 309}]\n\t}\n\nThe
    resulting set of endpoints can be viewed as:\n\n\ta: [ 10.10.1.1:8675, 10.10.2.2:8675
    ],\n\tb: [ 10.10.1.1:309, 10.10.2.2:309 ]"
  properties:
    addresses:
      description: IP addresses which offer the related ports that are marked as ready.
        These endpoints should be considered safe for load balancers and clients to
        utilize.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.EndpointAddress'
      type: array
    notReadyAddresses:
      description: IP addresses which offer the related ports but are not currently
        marked as ready because they have not yet finished starting, have recently failed
        a readiness check, or have recently failed a liveness check.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.EndpointAddress'
      type: array
    ports:
      description: Port numbers available on the related IP addresses.
      items:
        $ref: '#/definitions/io.k8s.api.core.v1.EndpointPort'
      type: array
  type: object

```