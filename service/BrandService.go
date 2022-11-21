package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
)

func CreatBrand(brandInfo model.CreateBrandRequest) (uint64, error) {
	_, err := repo.GetBrandByBrandInfo(setting.DB, brandInfo.Brand, brandInfo.IndustryCode, brandInfo.SubindustryCode)
	if err == nil {
		fmt.Println(fmt.Sprintf("Error: 品牌已存在，%s", err.Error()))
		return -1, errors.New(model.RecordExistError)
	}

	brand := model.Brand{}
	brand.Brand = brandInfo.Brand
	brand.IndustryCode = brandInfo.IndustryCode
	brand.SubindustryCode = brandInfo.SubindustryCode

	id, err := repo.InsertBrand(setting.DB, brand)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return -1, errors.New(model.SqlInsertionError)
	}

	return id, nil
}

func GetIndustryList() (model.IndustryResponse, error) {
	var industryResponse model.IndustryResponse
	var industryList = make([]model.IndustryInfo, 0)

	industries, err := repo.GetIndustries(setting.DB)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.IndustryResponse{}, errors.New(model.SqlQueryError)
	}

	for _, industry := range industries {
		var industryInfo model.IndustryInfo
		industryInfo.Industry = industry.Industry
		industryInfo.IndustryCode = industry.IndustryCode
		industryInfo.Subindustries = make([]model.SubindustryInfo, 0)

		subindustries, err := repo.GetSunindustryByParentIndustryCode(setting.DB, industry.IndustryCode)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
			continue
		}

		for _, subindustry := range subindustries {
			var subindustryInfo model.SubindustryInfo
			subindustryInfo.SubIndustry = subindustry.Industry
			subindustryInfo.SubIndustryCode = subindustry.IndustryCode
			industryInfo.Subindustries = append(industryInfo.Subindustries, subindustryInfo)
		}

		industryList = append(industryList, industryInfo)
	}

	industryResponse.Industries = industryList
	return industryResponse, nil
}

func GetIndustryByIndustryCode(industryCode int) (string, error) {
	industry, err := repo.GetIndustryByIndustryCode(setting.DB, industryCode)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return "", errors.New(model.SqlQueryError)
	}

	return industry.Industry, nil
}

func ConvertBrandToBrandInfo(brand model.Brand) (model.BrandInfo, error) {
	var brandInfo = model.BrandInfo{}
	brandInfo.Id = fmt.Sprintf("%s", brand.Id)
	brandInfo.Brand = brand.Brand
	industry, err := repo.GetIndustryByIndustryCode(setting.DB, brand.IndustryCode)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query Brand Industry Code Error. %s", err.Error()))
		return model.BrandInfo{}, errors.New(model.SqlQueryError)
	}

	subindustry, err := repo.GetIndustryByIndustryCode(setting.DB, brand.SubindustryCode)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query Brand Subindustry Code Error. %s", err.Error()))
		return model.BrandInfo{}, errors.New(model.SqlQueryError)
	}

	brandInfo.Industry = industry.Industry
	brandInfo.Subindustry = subindustry.Industry

	return brandInfo, nil
}
