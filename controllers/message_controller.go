package controllers

import (
	
	"simpleapi/services"
	"simpleapi/helpers"
	"simpleapi/models"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type MessageController struct {
	Ctx      iris.Context
	Services services.MessageService
}

func (controller *MessageController) Post() mvc.Result {
	ctx := controller.Ctx
	var inputForm models.Message
	err := ctx.ReadJSON(&inputForm)

	// check for valid json
	if err != nil {
		response := helpers.ResponseJson(iris.StatusBadRequest, iris.Map{
			"message": "input not valid",
		})
		return response
	}

	// save to DB
	resultDB, err := controller.Services.CreateMessage(inputForm, ctx)
	if err != nil {
		response := helpers.ResponseJson(iris.StatusBadRequest, iris.Map{
			"message": "failed to save db",
		})
		return response
	}

	return helpers.ResponseJson(iris.StatusOK, resultDB)
}

func (controller *MessageController) Get() mvc.Result {
	
	return mvc.Response{
		Code:        iris.StatusOK,
		ContentType: "application/json",
		Text:        "Message List Here",
	}
}
