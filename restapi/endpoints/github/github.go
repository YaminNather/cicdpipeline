package github

import (
    "io"
    "encoding/json"
    "fmt"
    "net/http"
    "log"
    "bytes"

    "go.uber.org/fx"
)

type PushHookRouteHandler struct {}

func NewPushHookRouteHandler() *PushHookRouteHandler {
    return &PushHookRouteHandler {}
}

func (PushHookRouteHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
    log.Println("Received request in route /github/hook")

    var requestBodyBytes, _ = io.ReadAll(request.Body)
    var prettyJsonRequestBodyByteBuffer bytes.Buffer
    var _ = json.Indent(&prettyJsonRequestBodyByteBuffer, requestBodyBytes, "", "\t")

    log.Printf("Request body:\n%s\n", prettyJsonRequestBodyByteBuffer.String())

    fmt.Fprintf(responseWriter, "Received request")
}




type RoutesAdder struct {
    pushHookRouteHandler *PushHookRouteHandler
}

func newRoutesAdder(pushHookRouteHandler *PushHookRouteHandler) *RoutesAdder {
    return &RoutesAdder{pushHookRouteHandler: pushHookRouteHandler}
}

func (routesAdder RoutesAdder) AddRoutes(serveMux *http.ServeMux) {
    serveMux.Handle("/github/push-hook", routesAdder.pushHookRouteHandler)
}



func FxOptions() []fx.Option {
    return []fx.Option{
        fx.Provide(NewPushHookRouteHandler),
        fx.Provide(newRoutesAdder),
    }
}
