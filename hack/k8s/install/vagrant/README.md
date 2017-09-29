# Installing Kubernetes with Vagrant

## Prerequisites

Please make sure you have the following installed on your machine

- [VirtualBox](https://www.virtualbox.org/)
- [Vagrant](https://www.vagrantup.com/)

Make sure `vagrant` is working by running `vagrant version`

### MacOS: Installing prerequisites with Brew

```bash
$ brew install caskroom/cask/virtualbox
$ brew install caskroom/cask/vagrant
```

## Starting the Kubernetes VM:

Ask Vagrant to provision our VM. This will take a while, as it install Kubernetes and all of its dependencies from scratch.

```bash
cd $GOPATH/src/github.com/nuclio/nuclio-sdk/hack/k8s/install/vagrant && vagrant up
```

This will start an Ubuntu 16.04 VM and run each of the [required steps](../../../docs/k8s/README.md) to have a Kubernetes cluster running with Nuclio.

### Cluster Defaults

- Host IP: `10.100.100.10`
- Docker Registry: `10.100.100.10:31276`
- GOPATH: `/home/ubuntu/nuclio`

## Optional: Verify that everything is up and running

SSH into the machine and make sure the cluster is fully functioning by:

```bash
$GOPATH/src/github.com/nuclio-sdk/hack/k8s/install/vagrant && vagrant ssh
```

Ask kubectl to print out all of the Kubernetes pods running in all namespaces:

```bash
ubuntu@k8s:~$ kubectl get pods --all-namespaces
```
Output should be similar to:
```bash
  NAMESPACE     NAME                                 READY     STATUS    RESTARTS   AGE
  kube-system   etcd-k8s                             1/1       Running   0          8m
  kube-system   kube-apiserver-k8s                   1/1       Running   0          8m
  kube-system   kube-controller-manager-k8s          1/1       Running   0          8m
  kube-system   kube-dns-2425271678-nw37t            3/3       Running   0          8m
  kube-system   kube-proxy-44pr8                     1/1       Running   0          8m
  kube-system   kube-registry-proxy-d5mgc            1/1       Running   0          8m
  kube-system   kube-registry-v0-nswz4               1/1       Running   0          8m
  kube-system   kube-scheduler-k8s                   1/1       Running   0          8m
  kube-system   weave-net-jc601                      2/2       Running   0          8m
```

## Optional: Install kubectl locally

If you don't want to keep SSH'ing into the machine to run kubectl, you can install kubectl locally and configure it to work with the VM Kubernetes cluster. Start by installing kubectl:

https://kubernetes.io/docs/tasks/tools/install-kubectl/

Now get the kubeconfig from within the cluster and copy it to ~/.kube/config so that `kubectl` uses it by default.

```
cd $GOPATH/src/github.com/nuclio-sdk/hack/k8s/install/vagrant && vagrant ssh -c "sudo cat /home/ubuntu/.kube/config" > ~/.kube/config && cd -
```

Open up `~/.kube/config` in an editor and replace `server: https://<whatever>:6443` with `server: https://10.100.100.10:6443` so that kubectl uses the external IP address of the VM.

