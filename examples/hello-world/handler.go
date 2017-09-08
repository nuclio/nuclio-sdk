// Currently, if processor.yaml is not provided, nuclio-build will look for
// a Handler() function in the handler package. Future implementations
// will alleviate this limitation

// To build the handler run
//	go build -buildmode=plugin -o handler.so handler.go

package main

import (
	"github.com/nuclio/nuclio-sdk"
)

var (
// Make sure our handler has the right signature
//	_ = nuclio.EventHandler(Handler)
)

// Handler is the event Handler
func Handler(context *nuclio.Context, event nuclio.Event) (interface{}, error) {
	context.Logger.Info("Event received")

	return nuclio.Response{
		StatusCode:  200,
		ContentType: "application/text",
		Body:        []byte("Response from handler"),
	}, nil
}

func main() {}
