package main

import (
    "go.uber.org/fx"

    "example/user/cicdpipeline/restapi"
)

func main() {
    var fxOption []fx.Option = []fx.Option{}
    fxOption = append(fxOption, restapi.AddServerToFx()...)
    var fxApplication = fx.New(fxOption...)

    fxApplication.Run()
}
