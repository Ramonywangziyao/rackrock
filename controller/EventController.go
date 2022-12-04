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
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var createEventRequest model.CreateEventRequest
	if err := c.ShouldBind(&createEventRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	id, err := service.CreateEvent(createEventRequest, loginUser.ID)
	if err != nil {
		con.Error(c, model.SqlInsertionErrorCode, model.SqlInsertionError)
		return model.RockResp{
			Code:    model.SqlInsertionErrorCode,
			Message: model.SqlInsertionError,
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, id)
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
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importItemRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importItemRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventItemFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
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
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importSoldRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importSoldRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventSoldFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
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
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : %s", model.SqlQueryError, "access_level"),
			Data:    nil,
		}
	}
	if accessLevel != model.ADMIN {
		fmt.Errorf(fmt.Sprintf("用户 %d 无创建权限", loginUser.ID))
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	var importReturnRequest model.ImportEventDataRequest
	if err := c.ShouldBind(&importReturnRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		con.Error(c, model.ImportFileErrorCode, model.ImportFileError)
		return model.RockResp{
			Code:    model.ImportFileErrorCode,
			Message: model.ImportFileError,
			Data:    nil,
		}
	}

	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		con.Error(c, model.ExcelParseErrorCode, model.ExcelParseError)
		return model.RockResp{
			Code:    model.ExcelParseErrorCode,
			Message: model.ExcelParseError,
			Data:    nil,
		}
	}

	go service.ReadEventReturnFile(xlsx)

	con.Success(c, model.RequestSuccessMsg, nil)
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
	user := c.Query("user")
	tagIdStr := c.Query("tagId")
	var tagId uint64
	var err error
	if len(tagIdStr) > 0 {
		tagId, err = utils.ConvertStringToUint64(tagIdStr)
		if err != nil {
			con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
			return model.RockResp{
				Code:    model.RequestParameterErrorCode,
				Message: model.RequestParameterError,
				Data:    nil,
			}
		}
	}

	eventTypeStr := c.Query("type")
	eventType, err := strconv.Atoi(eventTypeStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	pageStr := c.Query("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}
	if page == 0 {
		page = 1
	}

	pageSizeStr := c.Query("pageSize")
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	events, err := service.GetEventList(userId, tagId, startTime, endTime, sortBy, order, user, eventType, page, pageSize)
	if err != nil {
		con.Error(c, model.RequestParameterErrorCode, model.RequestParameterError)
		return model.RockResp{
			Code:    model.RequestParameterErrorCode,
			Message: model.RequestParameterError,
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, events)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    events,
	}
}
