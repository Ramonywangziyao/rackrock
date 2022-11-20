package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
)

func CreateTag(tagRequest model.CreateTagRequest) (int64, error) {
	var tag = model.Tag{}
	tag.Tag = tagRequest.Tag

	id, err := repo.InsertTag(setting.DB, tag)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return -1, errors.New(model.SqlInsertionError)
	}

	return id, nil
}
