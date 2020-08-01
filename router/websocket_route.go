package router

import (
	"fmt"
	"simpleapi/repositories"
	"simpleapi/services"
	"simpleapi/utils"
	"simpleapi/controllers"

	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/websocket"
	// "github.com/kataras/neffos"
)

type prefixedLogger struct {
	prefix string
}

func (s *prefixedLogger) Log(msg string) {
	fmt.Printf("%s: %s\n", s.prefix, msg)
}

func init() {
	app := utils.GetIrisApplication()
	websocketAPI := app.Party("/websocket")
	repo := repositories.NewMessageRepository()
	service := services.NewMessageService(repo)

	mvcBase := mvc.New(websocketAPI)
	mvcBase.Register(
		&prefixedLogger{prefix: "DEV"},
	)
	mvcBase.Register(service)
	mvcBase.HandleWebsocket(&controllers.WebsocketController{Namespace: "default"})

	websocketServer := websocket.New(websocket.DefaultGorillaUpgrader, mvcBase)
	websocketAPI.Get("/", websocket.Handler(websocketServer))

}
