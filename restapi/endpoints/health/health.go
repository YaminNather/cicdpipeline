package health 

import (
    "fmt"
    "net/http"
    "go.uber.org/fx"
)

type PingRouteHandler struct {}

func newPingRouteHandler() *PingRouteHandler {
    return &PingRouteHandler{}
}

func (PingRouteHandler) ServeHTTP(responseWriter http.ResponseWriter, _ *http.Request) {
    fmt.Fprintln(responseWriter, "pong")
}




type RoutesAdder struct {
    pingRouteHandler *PingRouteHandler
}

func newHealthRoutesAdder(pingRouteHandler *PingRouteHandler) *RoutesAdder {
    return &RoutesAdder{
        pingRouteHandler: pingRouteHandler,
    }
}

func (healthRoutesAdder RoutesAdder) AddRoutes(serveMux *http.ServeMux) {
    serveMux.Handle("/health/ping", healthRoutesAdder.pingRouteHandler)
}




func FxOptions() []fx.Option {
    return []fx.Option{
        fx.Provide(newPingRouteHandler),
        fx.Provide(newHealthRoutesAdder),
    }
}
