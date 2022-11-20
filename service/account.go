package service

import (
	"rackrock/repository"
	"rackrock/repository/domain"
)

type Account interface {
	GetAccount(account *domain.Account, cols ...string) error
}

func GetAccountService() Account {
	return &AccountService{Account: new(repository.AccountRepository)}
}

type AccountService struct {
	Account repository.AccountRepoInterface
}

func (self *AccountService) GetAccount(entity *domain.Account, cols ...string) error {
	return self.Account.GetAccount(entity, cols...)
}
