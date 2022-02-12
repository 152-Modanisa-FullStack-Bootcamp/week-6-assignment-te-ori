package service

import (
	"fmt"
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

func (srv *AccountService) Accounts() []*account.Account {
	return srv.repository.Accounts()
}

func (srv *AccountService) UpdateBalance(username string, amount float32) (float32, float32, error) {
	user := srv.AccountOf(username)

	if user == nil {
		return 0, 0, fmt.Errorf("'%s' not found", username)
	}

	if user.Balance+amount < srv.minumumBalanceAmount {
		return user.Balance, user.Balance - srv.minumumBalanceAmount, fmt.Errorf("no sufficient balance")
	}

	newBalance, err := srv.repository.UpdateBalance(username, user.Balance+amount)

	return newBalance, newBalance - srv.minumumBalanceAmount, err
}

func NewAccountService(initBalance, minBalance float32, repository account.IAccountRepository) *AccountService {
	return &AccountService{
		initialBalanceAmount: initBalance,
		minumumBalanceAmount: minBalance,
		repository:           repository,
	}
}
