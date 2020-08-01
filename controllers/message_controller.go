package controllers

import (
	// "log"
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

	// save data json to Database
	resultDB, err := controller.Services.CreateMessage(inputForm, ctx)
	if err != nil {
		response := helpers.ResponseJson(iris.StatusBadRequest, iris.Map{
			"message": "failed to save db",
		})
		return response
	}

	// return json with new created result data from db
	return helpers.ResponseJson(iris.StatusOK, resultDB)
}

func (controller *MessageController) GetList() mvc.Result {

	ctx 	:= controller.Ctx
	limit 	:= ctx.URLParamIntDefault("limit", 10) // for limit, default 10
	page 	:= ctx.URLParamIntDefault("page", 1) // for pagination
	orderBy := ctx.URLParamDefault("order", "id DESC") // for ordering data

	input := services.InputPagination {
		Limit : limit,
		Page : page,
		OrderBy: orderBy,
	}
	
	// Query drom db
	results, pagination, err := controller.Services.GetMessageList(input)
	if err != nil {
		response := helpers.ResponseJson(iris.StatusBadRequest, iris.Map {
			"message": err.Error(),
		})
		return response
	}

	// return data & its pagination 
	response := helpers.ResponseJson(iris.StatusOK, iris.Map{
		"items":      results,
		"pagination": pagination,
	})
	return response
}
