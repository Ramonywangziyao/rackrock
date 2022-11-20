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

func GetEventList(userId, tagId int64, startTime, endTime, sortBy, orderBy, brand string, eventType, page int) (model.EventListResponse, error) {
	whereClause := generateEventSearchWhereClause(userId, tagId, startTime, endTime, sortBy, orderBy, brand, eventType)
	offset := (page - 1) * model.EventPageSize
	events, err := repo.GetEvents(setting.DB, whereClause, offset, model.EventPageSize)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.EventListResponse{}, errors.New(model.SqlQueryError)
	}

	eventListResponse, err := convertEventQueryResultToEventResponse(events)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.EventListResponse{}, errors.New(model.DataTypeConversionError)
	}

	return eventListResponse, nil
}

func generateEventSearchWhereClause(userId, tagId int64, startTime, endTime, sortBy, orderBy, brand string, eventType int) string {
	//TODO
	return ""
}

func convertEventQueryResultToEventResponse(events []model.Event) (model.EventListResponse, error) {
	//TODO
	return model.EventListResponse{}, nil
}
