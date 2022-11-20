package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type ReportController struct {
	BaseController
}

func (con ReportController) GetBasic(gin *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) GetShareLink(gin *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) GetRanking(gin *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) GetDailyDetail(gin *gin.Context) (res model.RockResp) {
	return
}

func (con ReportController) ExportSaleDetail(gin *gin.Context) (res model.RockResp) {
	return
}
