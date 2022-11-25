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

	var createBrandRequest model.CreateBrandRequest
	if err := c.ShouldBind(&createBrandRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	id, err := service.CreatBrand(createBrandRequest)
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

func (con GeneralController) GetBrandList(c *gin.Context) (res model.RockResp) {
	brandResponse, err := service.GetBrands()
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    brandResponse,
	}
}

//func (con GeneralController) GetFilterBrandList(c *gin.Context) (res model.RockResp) {
//	loginUser := context.GetLoginUser(c)
//	userId := loginUser.ID
//
//	brandResponse, err := service.GetFilterBrands(userId)
//	if err != nil {
//		return model.RockResp{
//			Code:    model.SqlQueryErrorCode,
//			Message: model.SqlQueryError,
//			Data:    nil,
//		}
//	}
//
//	return model.RockResp{
//		Code:    model.REQUEST_OK,
//		Message: model.RequestSuccessMsg,
//		Data:    brandResponse,
//	}
//}

func (con GeneralController) CreateTag(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	var createTagRequest model.CreateTagRequest
	if err := c.ShouldBind(&createTagRequest); err != nil {
		return model.RockResp{
			Code:    model.RequestBodyErrorCode,
			Message: model.RequestBodyError,
			Data:    nil,
		}
	}

	id, err := service.CreateTag(createTagRequest, userId)
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

func (con GeneralController) GetTagList(c *gin.Context) (res model.RockResp) {
	loginUser := context.GetLoginUser(c)
	userId := loginUser.ID

	tags, err := service.GetTagList(userId)
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    tags,
	}
}

func (con GeneralController) GetCities(c *gin.Context) (res model.RockResp) {
	var cityList model.CityResponse
	cityList.Cities = model.Cities

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    cityList,
	}
}

func (con GeneralController) GetIndustries(c *gin.Context) model.RockResp {
	industries, err := service.GetIndustryList()
	if err != nil {
		return model.RockResp{
			Code:    model.SqlQueryErrorCode,
			Message: model.SqlQueryError,
			Data:    nil,
		}
	}

	return model.RockResp{
		Code:    model.OK,
		Message: model.RequestSuccessMsg,
		Data:    industries,
	}
}
