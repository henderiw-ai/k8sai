---
title: "`add`"
linkTitle: "add"
type: docs
description: >
  add adds a git repo with the arguments repoURL and repoPATH to the config
---

<!--mdtogo:Short
    add adds a git repo with the arguments repoURL and repoPATH to the config
-->

`add` add adds a git repo with the arguments repoURL and repoPATH to the config

### Synopsis

<!--mdtogo:Long-->

```
k8sai repo add REPO_URL REPO_PATH
```

#### Args

```
REPO_URL:
  The repo url of the git repo

REPO_PATH:
  The path within the repo
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
# add adds a git repo with the REPO_URL k8sai "https://github.com/kubernetes/kubernetes"  and REPO_PATH "/api/openapi-spec"
$ k8sai repo add "https://github.com/kubernetes/kubernetes" "/api/openapi-spec"
```

```

<!--mdtogo-->