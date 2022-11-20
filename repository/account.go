package repository

import (
	"rackrock/repository/domain"
	"rackrock/starter/component"
)

type AccountRepoInterface interface {
	GetAccount(entity *domain.Account, cols ...string) error

	AddAccount(entity *domain.Account) error
}

type AccountRepository struct {
}

func (self *AccountRepository) GetAccount(entity *domain.Account, cols ...string) error {
	return component.DB.Select(cols).Where(entity).First(entity).Error
}

func (self *AccountRepository) AddAccount(entity *domain.Account) error {
	return component.DB.Create(entity).Error
}
