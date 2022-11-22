package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
	"rackrock/utils"
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
		return
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		con.Error(c, model.ReportNotReadyErrorCode, model.ReportNotReadyError)
		return
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
		return
	}

	con.Success(c, model.RequestSuccessMsg, reportResponse)
	return
}

func (con ReportController) GetRanking(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		con.Error(c, model.ReportNotReadyErrorCode, model.ReportNotReadyError)
		return
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")
	dimension := c.Query("dimension")
	sortBy := c.Query("sortBy")
	orderBy := c.Query("orderBy")
	page := c.Query("page")

	rankingResponse, err := service.GetReportRanking()
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : ranking, %s", model.SqlQueryError, err))
		return
	}

	con.Success(c, model.RequestSuccessMsg, rankingResponse)
	return
}

func (con ReportController) GetDailyDetail(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	eventIdStr := c.Query("eventId")
	if len(eventIdStr) == 0 {
		// 没有传场次
		con.Error(c, model.RequestParameterMissingErrorCode, model.RequestParameterMissingError)
		return
	}
	eventId, err := utils.ConvertStringToUint64(eventIdStr)
	if err != nil {
		con.Error(c, model.DataTypeConversionErrorCode, model.DataTypeConversionError)
		return
	}

	// 检查用户是否对该场次有权限
	accessLevel, err := service.GetUserAccessLevel(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return
	}

	event, err := service.GetEvent(eventId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : event, %s", model.SqlQueryError, err))
		return
	}
	// 管理员跳过
	if accessLevel != model.ADMIN && event.UserId != userId && event.CreatorId != userId {
		con.Error(c, model.NotAuthorizedErrorCode, model.NotAuthorizedError)
		return
	}

	// 检查报告页状态
	if event.ReportStatus == 0 {
		con.Error(c, model.ReportNotReadyErrorCode, model.ReportNotReadyError)
		return
	}

	startTime := c.Query("startTime")
	endTime := c.Query("endTime")
	brand := c.Query("brand")
	source := c.Query("source")
	page := c.Query("page")

	dailyDetailResponse, err := service.GetReportDailyDetail()
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, fmt.Sprintf("%s : daily detail, %s", model.SqlQueryError, err))
		return
	}

	con.Success(c, model.RequestSuccessMsg, dailyDetailResponse)
	return
}

func (con ReportController) GetShareLink(c *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) ExportSaleDetail(c *gin.Context) (res model.RockResp) {
	return
}
