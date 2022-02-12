package service

import (
	"my_account/account"
)

type AccountService struct {
	initialBalanceAmount float32
	minumumBalanceAmount float32
	repository           account.IAccountRepository
}

func (srv *AccountService) AccountOf(username string) *account.Account {
	return srv.repository.AccountOf(username)
}

func (srv *AccountService) Put(username string) (float32, error) {
	if acc := srv.AccountOf(username); acc != nil {
		return acc.Balance, nil
	}

	return srv.repository.Put(username, srv.initialBalanceAmount)
}

func NewAccountService(initBalance, minBalance float32, repository account.IAccountRepository) *AccountService {
	return &AccountService{
		initialBalanceAmount: initBalance,
		minumumBalanceAmount: minBalance,
		repository:           repository,
	}
}
