package web

import (
	"github.com/pcavezzan/sample-apiserver/app"
)

func Start(config *app.Config) {

	appFactory := app.NewAppFactory(config)

	// startChi(appFactory, config)
	startGin(appFactory, config)
}