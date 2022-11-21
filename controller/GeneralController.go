package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rackrock/context"
	"rackrock/model"
	"rackrock/service"
)

type GeneralController struct {
	BaseController
}

func (con GeneralController) CreateBrand(c *gin.Context) (res model.RockResp) {
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

	var createBrandRequest model.CreateBrandRequest
	if err := c.ShouldBind(&createBrandRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	id, err := service.CreatBrand(createBrandRequest)
	if err != nil {
		con.Error(c, model.SqlInsertionErrorCode, model.SqlInsertionError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return
}

func (con GeneralController) GetBrandList(c *gin.Context) (res model.RockResp) {
	return
}

func (con GeneralController) CreateTag(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	var createTagRequest model.CreateTagRequest
	if err := c.ShouldBind(&createTagRequest); err != nil {
		con.Error(c, model.RequestBodyErrorCode, model.RequestBodyError)
		return
	}

	id, err := service.CreateTag(createTagRequest, userId)
	if err != nil {
		con.Error(c, model.SqlInsertionErrorCode, model.SqlInsertionError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
	return
}

func (con GeneralController) GetTagList(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	tags, err := service.GetTagList(userId)
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, tags)
	return
}

func (con GeneralController) GetCities(c *gin.Context) (res model.RockResp) {
	var cityList model.CityResponse
	cityList.Cities = model.Cities

	con.Success(c, model.RequestSuccessMsg, cityList)
	return
}

func (con GeneralController) GetIndustries(c *gin.Context) (res model.RockResp) {
	industries, err := service.GetIndustryList()
	if err != nil {
		con.Error(c, model.SqlQueryErrorCode, model.SqlQueryError)
		return
	}

	con.Success(c, model.RequestSuccessMsg, industries)
	return
}
