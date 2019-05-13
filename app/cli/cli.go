package cli

import (
	"github.com/pcavezzan/sample-apiserver/app"
	"github.com/pcavezzan/sample-apiserver/log"
)

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
	factory := app.NewAppFactory(config)
	return cmdImpl{
		Run: func() {
			db := factory.DataSource()
			result, err := db.Exec(`
CREATE TABLE parking(
 id serial PRIMARY KEY,
 number SMALLINT,
 owner VARCHAR,
 car VARCHAR
)
`)
			if err != nil {
				log.Fatal(err)
			}
			nbRowsAffected, err := result.RowsAffected()

			if err != nil {
				log.Fatal(err)
			}
			log.Info("Init completed", nbRowsAffected)
		},
	}
}
