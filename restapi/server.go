package restapi

import (
	"context"
    "net"
	"net/http"

	"go.uber.org/fx"

	"example/user/cicdpipeline/restapi/endpoints/github"
	"example/user/cicdpipeline/restapi/endpoints/health"
)

func AddServerToFx() []fx.Option {
    var fxOptions = []fx.Option{
        fx.Provide(newServer),

        fx.Invoke(func(server *http.Server) {}),
    }

    fxOptions = append(fxOptions, health.FxOptions()...)
    fxOptions = append(fxOptions, github.FxOptions()...)

    return fxOptions
}

func newServer(healthRoutesAdder *health.RoutesAdder, githubWebhookRoutesAdder *github.RoutesAdder, fxLifecycle fx.Lifecycle) *http.Server {
    var serveMux *http.ServeMux = http.NewServeMux()
    healthRoutesAdder.AddRoutes(serveMux)
    githubWebhookRoutesAdder.AddRoutes(serveMux)

    var server = &http.Server{Addr: "0.0.0.0:8080", Handler: serveMux,}

    fxLifecycle.Append(
        fx.Hook{
            OnStart: func (context context.Context) error {
                var listener, err = net.Listen("tcp", server.Addr)
                if err != nil {
                    return err
                }

                go server.Serve(listener)

                return nil
            },
            OnStop: func (context context.Context) error {
                return server.Shutdown(context)
            },
        },
    )

    return server
}
