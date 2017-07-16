// Currently, if processor.yaml is not provided, nuclio-build will look for
// a Handler() function in the handler package. Future implementations
// will alleviate this limitation

package handler

import (
    "github.com/nuclio/nuclio-sdk/event"
    "github.com/nuclio/nuclio/pkg/processor/runtime"
    "github.com/nuclio/nuclio/pkg/processor/eventsource/http"
)

func Handler(context *runtime.Context, event event.Event) (interface{}, error) {
    context.Logger.Info("Event received (packages)")

    return http.Response{
        StatusCode:  200,
        ContentType: "application/text",
        Body: []byte("Response from handler"),
    }, nil
}
