package service

import (
	"errors"
	"rackrock/model"
	"rackrock/repo"
	"rackrock/setting"
)

func GetDashboardInfo(userId uint64) (model.DashboardBasicResponse, error) {
	nickname, err := repo.GetUserNickNameById(setting.DB, userId)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	eventCount, err := repo.GetTotalEventCountById(setting.DB, userId)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	events, err := repo.GetEventIdsByUserId(setting.DB, userId)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	amountSoldCount, err := repo.GetTotalAmountSoldByEventIds(setting.DB, events)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	itemSoldCount, err := repo.GetTotalItemSoldByEventIds(setting.DB, events)
	if err != nil {
		return model.DashboardBasicResponse{}, errors.New(model.SqlQueryError)
	}

	return model.DashboardBasicResponse{
		nickname, eventCount, amountSoldCount, itemSoldCount,
	}, nil
}
