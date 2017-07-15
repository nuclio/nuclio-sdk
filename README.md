# nuclio-sdk
SDK for working with Nuclio

### Getting started
Get the stuff
```
go get -d github.com/nuclio/nuclio-sdk/...
go get github.com/nuclio/nuclio/cmd/nuclio-build
go get github.com/nuclio/nuclio/cmd/nuclio-deploy
```

Add the bin to path, so we can run nuclio-build
```
PATH=$PATH:$GOPATH/bin
```
Build the example from SDK example (function.yml from the path is used as well)
```
nuclio-build -n example $GOPATH/src/github.com/nuclio/nuclio-sdk/examples/golang
```

Run the processor locally and then access port 8080:
```
docker run -p 8080:8080 example:latest
```

Now deploy it to the kubernetes cluster:
```
nuclio-deploy -k ~/.kube/config -r <external IP address>:31276 -p 31010 example:latest
```
