package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/wenzhenxi/gorsa"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/starter/component"
	"rackrock/utils"
	"strconv"
)

func CreateUser(registerRequest model.RegisterRequest) (uint64, error) {
	//encrypt, err := gorsa.PublicEncrypt(`Love2013+`, model.Publickey)
	//if err != nil {
	//	fmt.Println(fmt.Sprintf("Error: Encode Password Error. %s", err.Error()))
	//	return 0, err
	//}

	decodedPassword, err := gorsa.PriKeyDecrypt(registerRequest.Password, model.Pirvatekey)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Decode Password Error. %s", err.Error()))
		return 0, err
	}
	data, err := hex.DecodeString(decodedPassword)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Decode Password Error. %s", err.Error()))
		return 0, err
	}

	var user = model.User{}
	user.Password = fmt.Sprintf("%s", data)
	user.Account = registerRequest.Account
	user.Nickname = registerRequest.NickName
	user.AccessLevel = model.VISITOR
	user.BrandId, _ = utils.ConvertStringToUint64(registerRequest.BrandId)

	// get unique id
	tempId := utils.GenerateRandomId()
	_, err = repo.GetUserByUserId(component.DB, tempId)
	for err == nil {
		tempId = utils.GenerateRandomId()
		_, err = repo.GetUserByUserId(component.DB, tempId)
	}

	user.Id = tempId

	err = repo.InsertUser(component.DB, user)
	if err != nil {
		fmt.Println(fmt.Sprintf("注册用户失败"))
		return 0, err
	}
	return user.Id, nil
}

func SetLoginTime(userId uint64) error {
	err := repo.UpdateLoginTimeByUserId(component.DB, userId)
	return err
}

func GetUserDetail(userId uint64) (model.UserInfo, error) {
	user, err := repo.GetUserByUserId(component.DB, userId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Query User Error. %s", err.Error()))
		return model.UserInfo{}, errors.New(model.SqlQueryError)
	}

	return generateUserInfoByUser(user)
}

func generateUserInfoByUser(user model.User) (model.UserInfo, error) {
	var userInfo = model.UserInfo{}
	userInfo.Id = strconv.FormatUint(user.Id, 10)
	brand, err := repo.GetBrandByBrandId(component.DB, user.BrandId)
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
	return repo.GetUserAccessLevelByUserId(component.DB, userId)
}

func GetUserListResponse() (model.UserListResponse, error) {
	var res = model.UserListResponse{}

	users, err := repo.GetUserList(component.DB)
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
	return repo.GetUserByAccount(component.DB, account)
}
