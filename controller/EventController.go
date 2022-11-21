package controller

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
	"strconv"
)

type EventController struct {
	BaseController
}

func (con EventController) CreateEvent(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	id, err := service.CreateEvent(createEventRequest, loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlInsertionErrorCode, model.SqlInsertionError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return
}

func (con EventController) ImportItems(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	var importItemRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importItemRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return
	}

	go service.ReadEventItemFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}

func (con EventController) ImportSold(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	var importSoldRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importSoldRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return
	}

	go service.ReadEventSoldFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}

func (con EventController) ImportReturn(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return
	}

	go service.ReadEventReturnFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
	return
}

func (con EventController) GetEventList(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	sortBy := c.Query("sortBy")
	order := c.Query("order")
	brand := c.Query("brand")
	tagIdStr := c.Query("tagId")
	tagId, err := utils.ConvertStringToUint64(tagIdStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return
	}

	eventTypeStr := c.Query("type")
	eventType, err := strconv.Atoi(eventTypeStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return
	}
	if page == 0 {
		page = 1
	}

	events, err := service.GetEventList(userId, tagId, startTime, endTime, sortBy, order, brand, eventType, page)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, events)
	return
}
