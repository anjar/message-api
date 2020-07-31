package controllers

import (
	"simpleapi/services"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type MessageController struct {
	Ctx      iris.Context
	Services services.MessageService
}

func (controller *MessageController) Get() mvc.Result {
	
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: "application/json",
		Text:        "string(tokenMarshal)",
	}
}
