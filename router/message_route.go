package controller_registrator

import (
	
	"simpleapi/utils"
	"simpleapi/controllers"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	app := utils.GetIrisApplication()

	// Prepare our repositories and services.
	
	mvcBase := mvc.New(app.Party("/message"))
	
	mvcBase.Handle(new(controllers.MessageController))
}
