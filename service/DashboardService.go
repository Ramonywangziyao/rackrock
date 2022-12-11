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
		events, err := repo.GetEventIdsByUserId(component.DB, userId)
		if err != nil {
			return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
		}

		amountSoldCount, err := repo.GetTotalAmountSoldByEventIds(component.DB, events)
		if err != nil {
			return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
		}

		itemSoldCount, err := repo.GetTotalItemSoldByEventIds(component.DB, events)
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
