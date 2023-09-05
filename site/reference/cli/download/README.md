---
title: "`download`"
linkTitle: "download"
type: docs
description: >
  download downloads the git documents to a local directory
---

<!--mdtogo:Short
    download downloads the git documents to a local directory
-->

`download` downloads the git documents to a local directory

### Synopsis

<!--mdtogo:Long-->

```
k8sai download
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
# download downloads the git repo's to a local directory
$ k8sai download
```
<!--mdtogo-->