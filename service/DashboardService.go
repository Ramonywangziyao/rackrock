package service

import (
	"errors"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/starter/component"
)

func GetDashboardInfo(userId uint64) (model.DashboardBasicResponse, error) {
	nickname, err := repo.GetUserNickNameById(component.DB, userId)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	eventCount, err := repo.GetTotalEventCountById(component.DB, userId)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	if eventCount > 0 {
		events, err := repo.GetEventsByUserId(component.DB, userId)
		if err != nil {
			return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
		}

		eventIds := make([]uint64, 0)
		for _, event := range events {
			if event.ReportStatus == model.READY {
				eventIds = append(eventIds, event.Id)
			}
		}

		if len(eventIds) == 0 {
			return model.DashboardBasicResponse{
				nickname, eventCount, 0, 0,
			}, nil
		}

		amountSoldCount, err := repo.GetTotalAmountSoldByEventIds(component.DB, eventIds)
		if err != nil {
			return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
		}

		itemSoldCount, err := repo.GetTotalItemSoldByEventIds(component.DB, eventIds)
		if err != nil {
			return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
		}

		return model.DashboardBasicResponse{
			nickname, eventCount, amountSoldCount, itemSoldCount,
		}, nil
	}

	return model.DashboardBasicResponse{
		nickname, eventCount, 0, 0,
	}, nil
}
