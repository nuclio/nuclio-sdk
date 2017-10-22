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

Finally, add $(minikube ip):5000 to the list of insecure registries in Docker (e.g. in Docker for Mac this is in Preferences -> Daemon -> Insecure registries -> Apply & Restart).

Once you are done, you can resume the [getting started guide](/README.md) to install nuclio on this cluster.