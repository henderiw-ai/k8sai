# openAPI schema status for io.k8s.apimachinery.pkg.version.Info

## description

Info contains versioning information. how we'll want to distribute that information.

## schema

```yaml
|
  description: Info contains versioning information. how we'll want to distribute that
    information.
  properties:
    buildDate:
      type: string
    compiler:
      type: string
    gitCommit:
      type: string
    gitTreeState:
      type: string
    gitVersion:
      type: string
    goVersion:
      type: string
    major:
      type: string
    minor:
      type: string
    platform:
      type: string
  required:
  - major
  - minor
  - gitVersion
  - gitCommit
  - gitTreeState
  - buildDate
  - goVersion
  - compiler
  - platform
  type: object

```
