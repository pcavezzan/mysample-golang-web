package cli

import (
	"github.com/pcavezzan/sample-apiserver/app"
)

var factory *app.AppFactory

type Command interface {
	Execute()
}

type cmdImpl struct {
	Run func()
}

func (c cmdImpl) Execute() {
	c.Run()
}

func NewInitCommand(config *app.Config) Command {
	factory = app.NewAppFactory(config)
	return cmdImpl{
		Run: initializeDb,
	}
}
