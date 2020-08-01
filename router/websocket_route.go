package router

import (
	// "fmt"
	"simpleapi/repositories"
	"simpleapi/services"
	"simpleapi/utils"
	"simpleapi/controllers"

	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"

)

func init() {
	app := utils.GetIrisApplication()
	websocketAPI := app.Party("/websocket")
	repo := repositories.NewMessageRepository()
	service := services.NewMessageService(repo)

	mvcBase := mvc.New(websocketAPI)
	mvcBase.Register(service)
	mvcBase.HandleWebsocket(&controllers.WebsocketController{Namespace: "default"})

	websocketServer := websocket.New(websocket.DefaultGorillaUpgrader, mvcBase)
	websocketAPI.Get("/", websocket.Handler(websocketServer))

}
