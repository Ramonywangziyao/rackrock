package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
	"rackrock/utils"
)

func CreateEvent(eventRequest model.CreateEventRequest) (int64, error) {
	var event = model.Event{}
	event.EventName = eventRequest.EventName
	event.Type = eventRequest.EventType
	event.City = eventRequest.City
	event.TagId, _ = utils.ConvertStringToInt64(eventRequest.TagId)
	event.UserId, _ = utils.ConvertStringToInt64(eventRequest.UserId)
	event.StartTime, _ = utils.ConvertStringToTime(eventRequest.StartTime)
	event.EndTime, _ = utils.ConvertStringToTime(eventRequest.EndTime)
	event.LastDays = int(event.EndTime.Sub(event.StartTime).Hours() / 24)

	id, err := repo.InsertEvent(setting.DB, event)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return -1, errors.New(model.SqlInsertionError)
	}

	return id, nil
}
