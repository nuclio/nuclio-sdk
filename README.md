# nuclio-sdk
SDK for working with nuclio

### Getting started
First, let's prepare a GOPATH directory (or use your own) and get the nuclio SDK:
```
export GOPATH=~/nuclio && mkdir -p $GOPATH
go get -u github.com/nuclio/nuclio-sdk
```

To start deploying functions we'll need a remote kubernetes cluster which we can install in one of three ways:

1. [On a local VM with minikube](hack/k8s/install/minikube/README.md)
2. [On a local VM with Vagrant](hack/k8s/install/vagrant/README.md)
3. [From scratch with kubeadm on Ubuntu](hack/k8s/install/scratch/README.md)

With a functioning kuberenetes cluster (with built-in docker registry) and a working kubectl, we can go ahead and install the nuclio services on the cluster:

```
cd $GOPATH/src/github.com/nuclio/nuclio-sdk/hack/k8s/install/scratch/resources && kubectl create -f controller.yaml,playground.yaml && cd -
```

Use `kubectl get pods` to verify both controller and playground have a status of `RUNNING`.

#### Using the nuclio playground
Browse to `http://<cluster-ip>:32050` - you should be greeted by the nuclio playground. Choose one of the built in examples and click deploy. The first build will populate the local docker cache with base images and such, so it might take a while depending on your network. Once the function has been deployed, you can invoke it with a body by clicking "Invoke".

#### Using nuctl, the nuclio command line tool

First, make sure you have Golang 1.8+ (https://golang.org/doc/install) and Docker (https://docs.docker.com/engine/installation). Now build nuctl, the nuclio command line tool and add `$GOPATH/bin` to path for this session:
```
go get -u github.com/nuclio/nuclio/cmd/nuctl
PATH=$PATH:$GOPATH/bin
```

Before docker images can be pushed to our built in registry, we need to add `<cluster-ip>:31276` (e.g. `10.100.100.10:31276` if you're using Vagrant) to the list of insecure registries. If you're using Docker for Mac you can find this under `Preferences -> Daemon`.

Deploy the hello world example:
```
nuctl deploy -p $GOPATH/src/github.com/nuclio/nuclio-sdk/examples/hello-world --registry [registry address] helloworld --run-registry localhost:5000
```

If you're using `minikube`, the registry address is `$(minikube ip):5000`. If you used `kubeadm` or `Vagrant`, the registry address is `10.100.100.10:31276`.

And finally execute it:
```
nuctl invoke helloworld
```
