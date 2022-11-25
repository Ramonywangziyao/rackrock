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
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	id, err := service.CreateEvent(createEventRequest, loginUser.ID)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlInsertionErrorCode,
			Message: model.SqlInsertionError,
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    id,
	}
}

func (con EventController) ImportItems(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importItemRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importItemRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventItemFile(xlsx)

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    nil,
	}
}

func (con EventController) ImportSold(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importSoldRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importSoldRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventSoldFile(xlsx)

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    nil,
	}
}

func (con EventController) ImportReturn(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	accessLevel, err := service.GetUserAccessLevel(loginUser.ID)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventReturnFile(xlsx)

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    nil,
	}
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
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	eventTypeStr := c.Query("type")
	eventType, err := strconv.Atoi(eventTypeStr)
	if err != nil {
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}
	if page == 0 {
		page = 1
	}

	events, err := service.GetEventList(userId, tagId, startTime, endTime, sortBy, order, brand, eventType, page)
	if err != nil {
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    events,
	}
}
