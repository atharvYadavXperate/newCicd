package applayer

import (
	"github.com/atharvYadavXperate/kapad-app/app"
	db "github.com/atharvYadavXperate/kapad-app/firestore"
)

var Db *db.Database

func Init() app.App {
	application := app.App{}
	application.Init()
	err := application.InitDatabase()
	if err != nil {
		panic(err)
	}
	return application
}
