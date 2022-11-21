package service

import (
	"errors"
	"fmt"
	"github.com/farmerx/gorsa"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
	"rackrock/utils"
)

func CreateUser(registerRequest model.RegisterRequest) (uint64, error) {
	var decodedPassword, err = gorsa.RSA.PriKeyDECRYPT([]byte(registerRequest.Password))
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Decode Password Error. %s", err.Error()))
		return 0, err
	}

	var user = model.User{}
	user.Password = string(decodedPassword)
	user.Account = registerRequest.Account
	user.Nickname = registerRequest.NickName
	user.AccessLevel = model.VISITOR
	user.BrandId, _ = utils.ConvertStringToUint64(registerRequest.BrandId)

	// get unique id
	tempId := utils.GenerateRandomId()
	_, err = repo.GetUserByUserId(setting.DB, tempId)
	for err == nil {
		tempId = utils.GenerateRandomId()
		_, err = repo.GetUserByUserId(setting.DB, tempId)
	}

	user.Id = tempId

	err = repo.InsertUser(setting.DB, user)
	if err != nil {
		fmt.Println(fmt.Sprintf("注册用户失败"))
		return 0, err
	}
	return user.Id, nil
}

func SetLoginTime(userId uint64) error {
	err := repo.UpdateLoginTimeByUserId(setting.DB, userId)
	return err
}

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

func GetUserByAccount(account string) (model.User, error) {
	return repo.GetUserByAccount(setting.DB, account)
}
