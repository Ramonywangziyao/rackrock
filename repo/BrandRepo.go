package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
	"rackrock/starter/component"
)

func GetBrandByBrandInfo(db *gorm.DB, brandName string, industry_code, subindustry_code int) (model.Brand, error) {
	var brand = model.Brand{}

	err := db.Table("brand").
		Where("brand = ? and industry_code = ? and subindustry_code = ?", brandName, industry_code, subindustry_code).
		First(&brand).
		Error

	return brand, err
}

func GetBrands(db *gorm.DB) ([]model.Brand, error) {
	var brands = make([]model.Brand, 0)

	err := db.Debug().Table("brand").
		Find(&brands).
		Error

	return brands, err
}

func GetBrandByBrandId(db *gorm.DB, brandId uint64) (model.Brand, error) {
	var brand = model.Brand{}

	err := db.Table("brand").
		Where("id = ?", brandId).
		First(&brand).
		Error

	return brand, err
}

func InsertBrand(db *gorm.DB, brand model.Brand) (uint64, error) {
	err := db.Table("brand").
		Create(&brand).
		Error

	return brand.Id, err
}

func GetIndustries(db *gorm.DB) ([]model.Industry, error) {
	var industries []model.Industry
	err := component.DB.Debug().Table("industry").
		Select("*").
		Where("industry_level = 0").
		Find(&industries).
		Error

	return industries, err
}

func GetSubindustryByParentIndustryCode(db *gorm.DB, industryCode int) ([]model.Industry, error) {
	var industries []model.Industry

	err := db.Table("industry").
		Select("*").
		Where("industry_level = 1 and parent_industry_code = ?", industryCode).
		Find(&industries).
		Error

	return industries, err
}

func GetIndustryByIndustryCode(db *gorm.DB, industryCode int) (model.Industry, error) {
	var industry model.Industry

	err := db.Table("industry").
		Where("industry_code = ?", industryCode).
		First(&industry).
		Error

	return industry, err
}
