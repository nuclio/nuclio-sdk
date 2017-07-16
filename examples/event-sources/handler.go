package example

import (
    "github.com/nuclio/nuclio-sdk/event"
    "github.com/nuclio/nuclio/pkg/processor/runtime"
    "github.com/nuclio/nuclio/pkg/processor/eventsource/http"
)

func GolangExample(context *runtime.Context, event event.Event) (interface{}, error) {
    context.Logger.InfoWith("Got event",
        "url", event.GetURL(),
        "size", event.GetSize(),
        "timestamp", event.GetTimestamp())

    return http.Response{
        StatusCode:  201,
        ContentType: "application/text",
        Header: map[string]string{
            "x-v3io-something": "30",
        },
        Body: []byte("Response from golang"),
    }, nil
}
