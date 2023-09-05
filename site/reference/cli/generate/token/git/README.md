---
title: "`git`"
linkTitle: "git"
type: docs
description: >
  generate a git token for the respective git provider
---

<!--mdtogo:Short
    generate a git token for the respective git provider
-->

`git` adds a set of configuration specifed in the fn-config files to the package.

### Synopsis

<!--mdtogo:Long-->

```
k8sai generate token git PROVIDER
```

#### Args

```
PROVIDER:
  The git provider such as github, gitlab, etc
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
$ k8sai generate token git github
```

```shell
# generate a gitlab token
$ k8sai generate token git gitlab
```

<!--mdtogo-->
