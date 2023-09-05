---
title: "`add`"
linkTitle: "add"
type: docs
description: >
  add adds a token for a respective provider
---

<!--mdtogo:Short
    add adds a token for a respective provider
-->

`add` adds a token for a respective provider

### Synopsis

<!--mdtogo:Long-->

```
k8sai auth add PROVIDER TOKEN
```

#### Args

```
PROVIDER:
  The provider such as github, gitlab, openai, etc

TOKEN:
  The token use to authenticate to the respective provider
```

<!--mdtogo-->

### Examples

{{% hide %}}

<!-- @makeWorkplace @verifyExamples-->

```
# Set up workspace for the test.
TEST_HOME=$(mktemp -d)
cd $TEST_HOME
```

{{% /hide %}}

<!--mdtogo:Examples-->

<!-- @ @verifyStaleExamples-->

```shell
# generate a github token
$ k8sai auth add github <TOKEN>
```

```shell
# generate a gitlab token
$ k8sai auth add openai <TOKEN>
```

<!--mdtogo-->