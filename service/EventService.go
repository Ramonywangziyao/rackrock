package service

import (
	"errors"
	"fmt"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/starter/component"
	"rackrock/utils"
	"strings"
)

func CreateEvent(eventRequest model.CreateEventRequest, creatorId uint64) (uint64, error) {
	var event = model.Event{}
	event.EventName = eventRequest.EventName
	event.Type = eventRequest.EventType
	event.City = eventRequest.City
	event.TagId, _ = utils.ConvertStringToUint64(eventRequest.TagId)
	event.UserId, _ = utils.ConvertStringToUint64(eventRequest.UserId)
	event.StartTime, _ = utils.ConvertStringToTime(eventRequest.StartTime)
	event.EndTime, _ = utils.ConvertStringToTime(eventRequest.EndTime)
	event.LastDays = int(event.EndTime.Sub(event.StartTime).Hours() / 24)
	event.CreatorId = creatorId

	id, err := repo.InsertEvent(component.DB, event)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return 0, errors.New(model.SqlInsertionError)
	}

	return id, nil
}

func GetEvent(eventId uint64) (model.Event, error) {
	event, err := repo.GetEventByEventId(component.DB, eventId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.Event{}, errors.New(model.SqlQueryError)
	}

	return event, nil
}

func GetEventList(userId, tagId uint64, startTime, endTime, sortBy, order, user string, eventType, page, pageSize int) (model.EventListResponse, error) {
	accessLevel, _ := GetUserAccessLevel(userId)
	tagIds := make([]uint64, 0)
	if accessLevel == model.ADMIN {
		tag, _ := repo.GetTagById(component.DB, tagId)
		tagIds, _ = repo.GetTagIdsByTag(component.DB, tag.Tag)
	} else {
		tagIds = append(tagIds, tagId)
	}

	whereClause := generateEventSearchWhereClause(userId, tagIds, startTime, endTime, user, eventType)
	sortOrder := getEventSortOrder(sortBy, order)
	offset := (page - 1) * pageSize
	events, err := repo.GetEvents(component.DB, whereClause, sortOrder, offset, pageSize)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.EventListResponse{}, errors.New(model.SqlQueryError)
	}

	eventListResponse, err := convertEventQueryResultToEventResponse(events)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
		return model.EventListResponse{}, errors.New(model.DataTypeConversionError)
	}

	eventListResponse.CurrentPage = page
	eventListResponse.PageSize = pageSize
	eventTotalCount, err := repo.GetEventsCountByUserId(component.DB, userId)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error: Get Page %s", err.Error()))
		eventListResponse.TotalPage = -1
	} else {
		eventListResponse.TotalPage = int(eventTotalCount) / pageSize
		if int(eventTotalCount)%pageSize > 0 {
			eventListResponse.TotalPage = eventListResponse.TotalPage + 1
		}
	}

	return eventListResponse, nil
}

func getEventSortOrder(sortBy, order string) string {
	if len(sortBy) == 0 {
		sortBy = "start_time"
	}

	if len(order) == 0 {
		order = "desc"
	}

	sortOrder := fmt.Sprintf("%s %s", sortBy, order)
	return sortOrder
}

func generateEventSearchWhereClause(userId uint64, tagIds []uint64, startTime, endTime, user string, eventType int) string {
	var where = ""
	var newClause = make([]string, 0)
	if userId > 0 {
		newClause = append(newClause, fmt.Sprintf("(creator_id = %d or user_id = %d)", userId, userId))
	}

	if len(tagIds) > 0 {
		tagIdsStrList := make([]string, 0)
		for _, tagId := range tagIds {
			tagIdsStrList = append(tagIdsStrList, fmt.Sprintf("%d", tagId))
		}
		tagIdsStr := strings.Join(tagIdsStrList, ",")
		newClause = append(newClause, fmt.Sprintf("tag_id in (%s)", tagIdsStr))
	}

	if len(startTime) > 0 {
		newClause = append(newClause, fmt.Sprintf("start_time >= '%s 00:00:00'", startTime))
	}

	if len(endTime) > 0 {
		newClause = append(newClause, fmt.Sprintf("end_time <= '%s 00:00:00'", endTime))
	}

	if len(user) > 0 {
		newClause = append(newClause, fmt.Sprintf("user_id in (%s)", user))
	}

	if eventType > 0 {
		newClause = append(newClause, fmt.Sprintf("type = %d", eventType))
	}

	where = strings.Join(newClause, " and ")

	return where
}

func convertEventQueryResultToEventResponse(events []model.Event) (model.EventListResponse, error) {
	var eventResponse = model.EventListResponse{}
	var res = make([]model.EventInfo, 0)
	for _, event := range events {
		eventInfo := model.EventInfo{}
		eventInfo.EventName = event.EventName
		eventInfo.Id = fmt.Sprintf("%d", event.Id)
		tag, err := repo.GetTagById(component.DB, event.TagId)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
			continue
		}
		eventInfo.Tag = model.TagInfo{Id: fmt.Sprintf("%d", tag.Id), Tag: tag.Tag}
		eventInfo.StartTime = event.StartTime.Format("2006-01-02")
		eventInfo.EndTime = event.EndTime.Format("2006-01-02")
		eventInfo.City = event.City
		switch event.Type {
		case model.CONSIGNMENT_EVENT_TYPE:
			eventInfo.Type = model.EventType{Value: model.CONSIGNMENT_EVENT_TYPE, Label: model.CONSIGNMENT_EVENT_TYPE_LABEL}
			break
		case model.PURCHASED_EVENT_TYPE:
			eventInfo.Type = model.EventType{Value: model.PURCHASED_EVENT_TYPE, Label: model.PURCHASED_EVENT_TYPE_LABEL}
			break
		}

		res = append(res, eventInfo)
	}
	eventResponse.Events = res
	return eventResponse, nil
}
