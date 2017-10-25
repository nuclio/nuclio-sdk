# Installing Kubernetes with minikube

Note: This is alpha level documentation - if you encounter any issues, feedback is welcomed.

## OSX

### Prerequisites

Please make sure you have the following installed on your machine:
- [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- [xhyve driver](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#xhyve-driver)

### Installation

Start minikube as you normally would:
```
minikube start --vm-driver=xhyve
```

Bring up a docker registry inside minikube so that we can push our functions to it:
```
minikube ssh -- docker run -d -p 5000:5000 registry:2
```

Finally, bind the default service account to the cluster-admin role (in the future, the hole punched into RBAC will be smaller):
```
cd $GOPATH/src/github.com/nuclio/nuclio-sdk/hack/k8s/install/scratch/resources && kubectl create -f default-cluster-admin.yaml && cd -
```

Once that completes, you can resume the [getting started guide](/README.md) to install nuclio on this cluster.
