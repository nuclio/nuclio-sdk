# nuclio-sdk
SDK for working with Nuclio

### Getting started
Get the stuff
```
go get -d github.com/nuclio/nuclio-sdk
go get github.com/nuclio/nuclio-tools/cmd/...
```

Add the bin to path, so we can run nuclio-build
```
PATH=$PATH:$GOPATH/bin
```
Build the example from SDK example (function.yml from the path is used as well)
```
nuclio-build processor -n example $GOPATH/src/github.com/nuclio/nuclio-sdk/examples/golang
```

Run the processor locally and then access port 8080:
```
docker run -p 8080:8080 example:latest
```

Now deploy it to the kubernetes cluster:
```
nuclio-deploy processor -k ~/.kube/config -r <registry-url:port> -n example -p 31010 example:latest
```
