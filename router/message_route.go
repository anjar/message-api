package router

import (
	"simpleapi/repositories"
	"simpleapi/services"
	"simpleapi/utils"
	"simpleapi/controllers"

	"github.com/kataras/iris/v12/mvc"
)

func init() {
	app := utils.GetIrisApplication()

	// Prepare our repositories and services.
	repo := repositories.NewMessageRepository()
	service := services.NewMessageService(repo)

	mvcBase := mvc.New(app.Party("/message"))
	mvcBase.Register(service)
	mvcBase.Handle(new(controllers.MessageController))
}
