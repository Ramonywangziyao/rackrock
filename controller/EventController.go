package controller

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"rackrock/model"
	"rackrock/service"
)

type EventController struct {
	BaseController
}

func (con EventController) CreateEvent(c *gin.Context) {
	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con EventController) ImportItems(c *gin.Context) {
	var importItemRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importItemRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseError)
		return
	}

	go service.ReadEventItemFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con EventController) ImportSold(c *gin.Context) {
	var importSoldRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importSoldRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseError)
		return
	}

	go service.ReadEventSoldFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con EventController) ImportReturn(c *gin.Context) {
	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseError)
		return
	}

	go service.ReadEventReturnFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
}

func (con EventController) GetEventList(c *gin.Context) {

}
