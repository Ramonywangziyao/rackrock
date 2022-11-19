package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type EventController struct {
	BaseController
}

func (con EventController) CreateEvent(c *gin.Context) {
	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

}

func (con EventController) ImportItems(c *gin.Context) {
	var importItemRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importItemRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con EventController) ImportSold(c *gin.Context) {
	var importSoldRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importSoldRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con EventController) ImportReturn(c *gin.Context) {
	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con EventController) GetEventList(c *gin.Context) {

}
