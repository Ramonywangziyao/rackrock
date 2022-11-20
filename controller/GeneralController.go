package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
	"rackrock/service"
)

type GeneralController struct {
	BaseController
}

func (con GeneralController) CreateBrand(c *gin.Context) {
	var createBrandRequest model.CreateBrandRequest
	if err := c.ShouldBind(&createBrandRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	id, err := service.CreatBrand(createBrandRequest)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
}

func (con GeneralController) GetBrandList(c *gin.Context) {
}

func (con GeneralController) CreateTag(c *gin.Context) {
	var createTagRequest model.CreateTagRequest
	if err := c.ShouldBind(&createTagRequest); err != nil {
		con.Error(c, model.RequestBodyError)
		return
	}

	id, err := service.CreateTag(createTagRequest)
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, id)
}

func (con GeneralController) GetTagList(c *gin.Context) {

}

func (con GeneralController) GetCities(c *gin.Context) {
	var cityList model.CityResponse
	cityList.Cities = model.Cities

	con.Success(c, model.RequestSuccessMsg, cityList)
}

func (con GeneralController) GetIndustries(c *gin.Context) {
	industries, err := service.GetIndustryList()
	if err != nil {
		con.Error(c, err.Error())
		return
	}

	con.Success(c, model.RequestSuccessMsg, industries)
}
