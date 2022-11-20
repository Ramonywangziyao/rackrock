package controller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
	"strconv"
)

type EventController struct {
	BaseController
}

func (con EventController) CreateEvent(c *gin.Context) (res model.RockResp) {
	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	creatorId := createEventRequest.CreatorId
	if creatorId != "0" {
		fmt.Errorf(fmt.Sprintf("用户 %s 无创建权限", creatorId))
		con.Error(c, model.NotAuthorizedError)
		return
	}

	id, err := service.CreateEvent(createEventRequest)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return
}

func (con EventController) ImportItems(c *gin.Context) (res model.RockResp) {
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
	return
}

func (con EventController) ImportSold(c *gin.Context) (res model.RockResp) {
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
	return
}

func (con EventController) ImportReturn(c *gin.Context) (res model.RockResp) {
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
	return
}

func (con EventController) GetEventList(c *gin.Context) (res model.RockResp) {
	userIdStr := c.Query("userId")
	userId, err := utils.ConvertStringToInt64(userIdStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	sortBy := c.Query("sortBy")
	orderBy := c.Query("orderBy")
	brand := c.Query("brand")
	tagIdStr := c.Query("tagId")
	tagId, err := utils.ConvertStringToInt64(tagIdStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}

	eventTypeStr := c.Query("type")
	eventType, err := strconv.Atoi(eventTypeStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}
	if page == 0 {
		page = 1
	}

	events, err := service.GetEventList(userId, tagId, startTime, endTime, sortBy, orderBy, brand, eventType, page)
	if err != nil {
		con.Error(c, model.RequestParameterError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, events)
	return
}
