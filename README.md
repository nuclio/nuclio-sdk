# nuclio-sdk
SDK for working with nuclio

### Getting started
First, let's prepare a GOPATH directory (or use your own) and get the nuclio SDK:
```
export GOPATH=~/nuclio && mkdir -p $GOPATH

go get github.com/nuclio/nuclio-sdk
```

To start deploying functions we'll need a few local tools and a remote kubernetes cluster. Let's start by preparing a kuberenetes cluster with one of two options:

1. Installation from scratch on Ubuntu: [hack/k8s/install/scratch/README.md]
2. Local VM with Vagrant: [hack/k8s/install/vagrant/README.md]

With a functioning kuberenetes cluster (with built-in docker registry) and a working kubectl, we can go ahead and install the nuclio services on the cluster:

```
cd $GOPATH/src/github.com/nuclio/nuclio-sdk/hack/k8s/resources && kubectl create -f controller.yaml,playground.yaml && cd -
```

Use `kubectl get pods` to verify both controller and playground have a status of `RUNNING`.

#### Using the nuclio playground
Browse to http://10.100.100.10:32050 - you should be greeted by the nuclio playground. Paste the following into the editor, name it "helloworld.go" and click deploy. The first build will populate the local docker cache with base images and such, so it might take a while depending on your network.

```
package helloworld

import (
    "github.com/nuclio/nuclio-sdk"
)

func HelloWorld(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
    return "Hello, World", nil
}
```

Once the playground indicates that the function was deployed successfully, head over to the "Invoke" tab and invoke your first nuclio function.


