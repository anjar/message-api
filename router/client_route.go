package router

import (
	"simpleapi/utils"
	"github.com/kataras/iris/v12"
)

func init() {
	// these route just for example websocket implement
	app := utils.GetIrisApplication()
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	app.Get("/client", func(ctx iris.Context) {
        ctx.View("client.html")
    })

}
