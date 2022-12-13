package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
	"strconv"
)

type ReportController struct {
	BaseController
}

func (con ReportController) GetBasic(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return model.RockResp{
			Code:    model.RequestParameterMissingErrorCode,
			Message: model.RequestParameterMissingError,
			Data:    nil,
		}
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return model.RockResp{
			Code:    model.DataTypeConversionErrorCode,
			Message: model.DataTypeConversionError,
			Data:    nil,
		}
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : event, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		reportResp := make(map[string]int, 0)
		reportResp["report_status"] = 0
		con.Success(c, model.ReportNotReadyError, reportResp)
		return model.RockResp{
			Code:    model.OK,
			Message: model.ReportNotReadyError,
			Data:    reportResp,
		}
	}

	// 获取筛选项
	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")

	// 根据筛选项查询数据，并开始计算
	reportResponse, err := service.GetReport(event, startTime, endTime, brand, source)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : report, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : report, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, reportResponse)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    reportResponse,
	}
}

func (con ReportController) GetRanking(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return model.RockResp{
			Code:    model.RequestParameterMissingErrorCode,
			Message: model.RequestParameterMissingError,
			Data:    nil,
		}
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return model.RockResp{
			Code:    model.DataTypeConversionErrorCode,
			Message: model.DataTypeConversionError,
			Data:    nil,
		}
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : event, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		reportResp := make(map[string]int, 0)
		reportResp["report_status"] = 0
		con.Success(c, model.ReportNotReadyError, reportResp)
		return model.RockResp{
			Code:    model.OK,
			Message: model.ReportNotReadyError,
			Data:    reportResp,
		}
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")
	dimension := c.Query("dimension")
	sortBy := c.Query("sortBy")
	order := c.Query("order")
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

	rankingResponse, err := service.GetReportRanking(event, startTime, endTime, brand, source, dimension, sortBy, order, page, pageSize)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : ranking, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : ranking, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, rankingResponse)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    rankingResponse,
	}
}

func (con ReportController) GetDailyDetail(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return model.RockResp{
			Code:    model.RequestParameterMissingErrorCode,
			Message: model.RequestParameterMissingError,
			Data:    nil,
		}
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return model.RockResp{
			Code:    model.DataTypeConversionErrorCode,
			Message: model.DataTypeConversionError,
			Data:    nil,
		}
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : event, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		reportResp := make(map[string]int, 0)
		reportResp["report_status"] = 0
		con.Success(c, model.ReportNotReadyError, reportResp)
		return model.RockResp{
			Code:    model.OK,
			Message: model.ReportNotReadyError,
			Data:    reportResp,
		}
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")

	dailyDetailResponse, err := service.GetReportDailyDetail(event, startTime, endTime, brand, source)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : daily detail, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : daily detail, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}

	con.Success(c, model.RequestSuccessMsg, dailyDetailResponse)
	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    dailyDetailResponse,
	}
}

func (con ReportController) GetShareLink(c *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) ExportSaleDetail(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return model.RockResp{
			Code:    model.RequestParameterMissingErrorCode,
			Message: model.RequestParameterMissingError,
			Data:    nil,
		}
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return model.RockResp{
			Code:    model.DataTypeConversionErrorCode,
			Message: model.DataTypeConversionError,
			Data:    nil,
		}
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: fmt.Sprintf("%s : event, %s", model.SqlQueryError, err),
			Data:    nil,
		}
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return model.RockResp{
			Code:    model.NotAuthorizedErrorCode,
			Message: model.NotAuthorizedError,
			Data:    nil,
		}
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		con.Error(c, model.ReportNotReadyErrorCode, model.ReportNotReadyError)
		return model.RockResp{
			Code:    model.ReportNotReadyErrorCode,
			Message: model.ReportNotReadyError,
			Data:    nil,
		}
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")

	xlsx := service.GetSaleDetailSheet(event, startTime, endTime, brand, source)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"sale.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    nil,
	}
}
