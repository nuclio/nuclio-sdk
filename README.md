# nuclio-sdk
SDK for working with nuclio

### Getting started
First, let's prepare a GOPATH directory (or use your own) and get the nuclio SDK:
```
export GOPATH=~/nuclio && mkdir -p $GOPATH
go get -u github.com/nuclio/nuclio-sdk
```

To start deploying functions we'll need a remote kubernetes cluster which we can install in one of two ways:

1. [On a local VM with Vagrant](hack/k8s/install/vagrant/README.md) (recommended)
2. [From scratch on Ubuntu](hack/k8s/install/scratch/README.md)

With a functioning kuberenetes cluster (with built-in docker registry) and a working kubectl, we can go ahead and install the nuclio services on the cluster:

```
cd $GOPATH/src/github.com/nuclio/nuclio-sdk/hack/k8s/install/scratch/resources && kubectl create -f controller.yaml,playground.yaml && cd -
```

Use `kubectl get pods` to verify both controller and playground have a status of `RUNNING`.

#### Using the nuclio playground
Browse to http://10.100.100.10:32050 - you should be greeted by the nuclio playground. Paste the following into the editor, name it something like `helloworld.go` and click deploy. The first build will populate the local docker cache with base images and such, so it might take a while depending on your network.

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

#### Using nuctl, the nuclio command line tool

First, make sure you have Golang 1.8+ (https://golang.org/doc/install) and Docker (https://docs.docker.com/engine/installation). Now build nuctl, the nuclio command line tool and add `$GOPATH/bin` to path for this session:
```
go get -u github.com/nuclio/nuclio/cmd/nuctl
PATH=$PATH:$GOPATH/bin
```

Before docker images can be pushed to our built in registry, we need to add `<kubernetes cluster ip>:31276` (e.g. `10.100.100.10:31276` if you're using Vagrant) to the list of insecure registries. If you're using Docker for Mac you can find this under `Preferences -> Daemon`.

Deploy the hello world example:
```
nuctl deploy -p $GOPATH/src/github.com/nuclio/nuclio-sdk/examples/hello-world -r <kubernetes cluster ip>:31276 helloworld --run-registry localhost:5000
```

And finally execute it:
```
nuctl invoke helloworld
```
