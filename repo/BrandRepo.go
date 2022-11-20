package repo

import (
	"gorm.io/gorm"
	"rackrock/model"
)

func GetBrandByBrandInfo(db *gorm.DB, brandName string, industry_code, subindustry_code int) (model.Brand, error) {
	var brand = model.Brand{}

	err := db.Table("brand").
		Where("brand = ? and industry_code = ? and subindustry_code = ?", brandName, industry_code, subindustry_code).
		First(&brand).
		Error

	return brand, err
}

func InsertBrand(db *gorm.DB, brand model.Brand) (int64, error) {
	err := db.Create(&brand).
		Error

	return brand.Id, err
}

func GetIndustries(db *gorm.DB) ([]model.Industry, error) {
	var industries []model.Industry

	err := db.Table("industry").
		Select("*").
		Where("industry_level = 0").
		Error

	return industries, err
}

func GetSunindustryByParentIndustryCode(db *gorm.DB, industryCode int) ([]model.Industry, error) {
	var industries []model.Industry

	err := db.Table("industry").
		Select("*").
		Where("industry_level = 1 and parent_industry_code = ?", industryCode).
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
