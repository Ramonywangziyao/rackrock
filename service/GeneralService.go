package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/starter/component"
	"rackrock/utils"
)

func CreateTag(tagRequest model.CreateTagRequest, userId uint64) (uint64, error) {
	var tag = model.Tag{}
	tag.Tag = tagRequest.Tag
	tag.UserId, _ = utils.ConvertStringToUint64(tagRequest.UserId)

	tags, err := repo.GetTagIdsByTag(component.DB, tag.Tag, tag.UserId)
	if len(tags) > 0 {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return 0, errors.New(model.RecordExistError)
	}

	id, err := repo.InsertTag(component.DB, tag)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return 0, errors.New(model.SqlInsertionError)
	}

	return id, nil
}

func GetTagList(userId uint64, accessLevel int) (model.TagListResponse, error) {
	var tagResponse = model.TagListResponse{}
	var tagList = make([]model.TagInfo, 0)
	var tags = make([]model.Tag, 0)
	var err error

	if accessLevel == model.ADMIN {
		tags, err = repo.GetAllTags(component.DB)
	} else {
		tags, err = repo.GetTagsByUserId(component.DB, userId)
	}

	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.TagListResponse{}, errors.New(model.SqlInsertionError)
	}

	for _, tag := range tags {
		tagInfo := convertTagToTagInfo(tag)
		if accessLevel == model.ADMIN {
			user, _ := repo.GetUserByUserId(component.DB, tag.UserId)
			brand, _ := repo.GetBrandByBrandId(component.DB, user.BrandId)
			tagInfo.Tag = fmt.Sprintf("%s-%s-%s", brand.Brand, user.Nickname, tagInfo.Tag)
		}
		tagList = append(tagList, tagInfo)
	}

	tagResponse.Tags = tagList
	return tagResponse, nil
}

func convertTagToTagInfo(tag model.Tag) model.TagInfo {
	var tagInfo = model.TagInfo{}
	tagInfo.Tag = tag.Tag
	tagInfo.Id = fmt.Sprintf("%d", tag.Id)
	return tagInfo
}
