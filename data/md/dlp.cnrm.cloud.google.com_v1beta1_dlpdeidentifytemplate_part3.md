# openAPI schema part3 for dlp.cnrm.cloud.google.com.v1beta1.DLPDeidentifyTemplate

## schema

```yaml
properties:
  spec:
    properties:
      deidentifyConfig:
        properties:
          infoTypeTransformations:
            properties:
              infoTypeTransformations:
                description: Treat the dataset as free-form text and apply the same
                  free text transformation everywhere.
                required:
                - transformations
                type: object
type: object

```
