# openAPI schema part3 for accesscontextmanager.cnrm.cloud.google.com.v1beta1.AccessContextManagerServicePerimeter

## schema

```yaml
properties:
  spec:
    properties:
      spec:
        properties:
          vpcAccessibleServices:
            properties:
              vpcAccessibleServices:
                description: |-
                  Specifies how APIs are allowed to communicate within the Service
                  Perimeter.
                properties:
                  allowedServices:
                    description: |-
                      The list of APIs usable within the Service Perimeter.
                      Must be empty unless 'enableRestriction' is True.
                    items:
                      type: string
                    type: array
                  enableRestriction:
                    description: |-
                      Whether to restrict API calls within the Service Perimeter to the
                      list of APIs specified in 'allowedServices'.
                    type: boolean
                type: object
required:
- spec
type: object

```
