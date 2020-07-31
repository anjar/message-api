package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"log"
	"sync"
)

type irisAppUtil struct {
	app *iris.Application
}

var irisAppInstance *irisAppUtil
var onceIrisApp sync.Once

// GetIrisApplication get iris Application instance
func GetIrisApplication() *iris.Application {
	onceIrisApp.Do(func() {
		log.Println("Initialize iris application instance...")

		app := iris.New()

		// recover from any http-relative panics
		app.Use(recover.New())

		// Log everything to terminal
		app.Use(logger.New())

		// You got full debug messages, useful when using MVC and you want to make
		// sure that your code is aligned with the Iris' MVC Architecture.
		app.Logger().SetLevel("debug")

		irisAppInstance = &irisAppUtil{
			app: app,
		}
	})

	return irisAppInstance.app
}
