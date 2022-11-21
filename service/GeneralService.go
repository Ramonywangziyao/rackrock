package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
)

func CreateTag(tagRequest model.CreateTagRequest, userId uint64) (uint64, error) {
	var tag = model.Tag{}
	tag.Tag = tagRequest.Tag
	tag.UserId = userId

	id, err := repo.InsertTag(setting.DB, tag)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return -1, errors.New(model.SqlInsertionError)
	}

	return id, nil
}

func GetTagList(userId uint64) (model.TagListResponse, error) {
	var tagResponse = model.TagListResponse{}
	var tagList = make([]model.TagInfo, 0)
	tags, err := repo.GetTagsByUserId(setting.DB, userId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.TagListResponse{}, errors.New(model.SqlInsertionError)
	}

	for _, tag := range tags {
		tagInfo := convertTagToTagInfo(tag)
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

