package router

import (
	"log"
	// "simpleapi/repositories"
	// "simpleapi/services"
	"simpleapi/utils"
	// "simpleapi/controllers"

	"github.com/kataras/iris/v12"
)

func init() {
	app := utils.GetIrisApplication()
	tmpl := iris.HTML("./views", ".html")
	app.RegisterView(tmpl)

	app.Get("/client", func(ctx iris.Context) {
		log.Printf("here")
        // Bind: {{.message}} with "Hello world!"
        ctx.ViewData("message", "Hello world!")
        // Render template file: ./views/hi.html
        ctx.View("client.html")
    })

}
