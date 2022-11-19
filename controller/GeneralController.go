package controller

import (
	"github.com/gin-gonic/gin"
	"rackrock/model"
)

type GeneralController struct {
	BaseController
}

func (con GeneralController) CreateBrand(c *gin.Context) {
	var createBrandRequest model.CreateBrandRequest
	if err := c.ShouldBind(&createBrandRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con GeneralController) GetBrandList(c *gin.Context) {
}

func (con GeneralController) CreateTag(c *gin.Context) {
	var createTagRequest model.CreateTagRequest
	if err := c.ShouldBind(&createTagRequest); err != nil {
		con.Error(c, model.RequestBodyError)
	}
}

func (con GeneralController) GetTagList(c *gin.Context) {

}

func (con GeneralController) GetCities(c *gin.Context) {

}
