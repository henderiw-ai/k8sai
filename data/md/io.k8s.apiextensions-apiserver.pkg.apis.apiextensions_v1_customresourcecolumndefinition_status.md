# openAPI schema status for io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceColumnDefinition

## description

CustomResourceColumnDefinition specifies a column for server side printing.

## schema

```yaml
|
  description: CustomResourceColumnDefinition specifies a column for server side printing.
  properties:
    description:
      description: description is a human readable description of this column.
      type: string
    format:
      description: format is an optional OpenAPI type definition for this column. The
        'name' format is applied to the primary identifier column to assist in clients
        identifying column is the resource name. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
        for details.
      type: string
    jsonPath:
      description: jsonPath is a simple JSON path (i.e. with array notation) which is
        evaluated against each custom resource to produce the value for this column.
      type: string
    name:
      description: name is a human readable name for the column.
      type: string
    priority:
      description: priority is an integer defining the relative importance of this column
        compared to others. Lower numbers are considered higher priority. Columns that
        may be omitted in limited space scenarios should be given a priority greater
        than 0.
      format: int32
      type: integer
    type:
      description: type is an OpenAPI type definition for this column. See https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md#data-types
        for details.
      type: string
  required:
  - name
  - type
  - jsonPath
  type: object

```