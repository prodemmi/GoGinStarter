//go:build wireinject
// +build wireinject

package main

import (
	"GoGinStarter/app/console"
	"GoGinStarter/wire"
)

func main() {
	container := wire.InitializeContainer()
	if cmdErr := console.Run(container); cmdErr != nil {
		container.Log.Error(cmdErr.Error())
	}
}
