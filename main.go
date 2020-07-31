package main

import (
	"flag"
	"fmt"
	"os"
	"simpleapi/utils"
	_ "simpleapi/router"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris/v12"
)


func main() {
	var host, port string
	flag.StringVar(&host, "host", os.Getenv("HOST"), "host of the service")
	flag.StringVar(&port, "port", os.Getenv("PORT"), "port of the service")
	flag.Parse()

	app := utils.GetIrisApplication()

	// homepage URL http://localhost:8080
	app.Get("/", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(iris.Map{
			"message": "You've been mixing with the wrong crowd.",
		})
		return
	})

	// http://localhost:8080/noexist
	// and all controller's methods like
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)
		_, _ = ctx.JSON(iris.Map{
			"message": "Page not found",
		})
		return
	})

	if rval := recover(); rval != nil {
		fmt.Printf("Rval: %+v\n", rval)
	}

	_ = app.Run(
		// Starts the web server at host and port
		iris.Addr(host+":"+port),
		// Ignores err server closed log when CTRL/CMD+C pressed.
		iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}