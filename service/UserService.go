package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
)

func GetUserDetail(userId uint64) (model.UserInfo, error) {
	user, err := repo.GetUserByUserId(setting.DB, userId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query User Error. %s", err.Error()))
		return model.UserInfo{}, errors.New(model.SqlQueryError)
	}

	return generateUserInfoByUser(user)
}

func generateUserInfoByUser(user model.User) (model.UserInfo, error) {
	var userInfo = model.UserInfo{}
	userInfo.Id = fmt.Sprintf("%s", user.Id)
	brand, err := repo.GetBrandByBrandId(setting.DB, user.BrandId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query Brand Error. %s", err.Error()))
		return model.UserInfo{}, errors.New(model.SqlQueryError)
	}
	userInfo.Brand, err = ConvertBrandToBrandInfo(brand)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Convert Brand Error. %s", err.Error()))
		return model.UserInfo{}, errors.New(model.DataTypeConversionError)
	}

	userInfo.AccessLevel = user.AccessLevel
	userInfo.Account = user.Account
	userInfo.Nickname = user.Nickname
	userInfo.CombinedName = fmt.Sprintf("%s-%s", userInfo.Brand.Brand, userInfo.Nickname)

	return userInfo, nil
}

func GetUserAccessLevel(userId uint64) (int, error) {
	return repo.GetUserAccessLevelByUserId(setting.DB, userId)
}

func GetUserListResponse() (model.UserListResponse, error) {
	var res = model.UserListResponse{}

	users, err := repo.GetUserList(setting.DB)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query Brand Error. %s", err.Error()))
		return model.UserListResponse{}, errors.New(model.SqlQueryError)
	}

	var userInfoList = make([]model.UserInfo, 0)
	for _, user := range users {
		userInfo, err := generateUserInfoByUser(user)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: Convert User Error. %s", err.Error()))
			continue
		}
		userInfoList = append(userInfoList, userInfo)
	}

	res.Users = userInfoList

	return res, nil
}
