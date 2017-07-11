# nuclio-sdk
SDK for working with Nuclio

Get the stuff
```
go get -d github.com/nuclio/nuclio-sdk
go get github.com/nuclio/nuclio-tools/cmd/...
```

Add the bin to path, so we can run nuclio-build

`PATH=$PATH:$GOPATH/bin`

Build the SDK example (function.yml from the path is used as well, configuring HTTP to listen at port 1968)

`nuclio-build processor -n example $GOPATH/src/github.com/nuclio/nuclio-sdk/examples/golang`

Run the processor locally and then access port 1968

`docker run -p 1968:1968 example:latest`
