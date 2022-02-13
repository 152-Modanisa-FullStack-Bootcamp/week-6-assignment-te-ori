package repository

import (
	"fmt"
	"my_account/account"
)

type repository map[string]float32

func (rep repository) Put(username string, amount float32) (float32, error) {
	rep[username] = amount

	return amount, nil
}

func (rep repository) AccountOf(username string) *account.Account {
	if balance, ok := rep[username]; ok {
		return &account.Account{Username: username, Balance: balance}
	} else {
		return nil
	}
}

func (rep repository) Accounts() []*account.Account {
	result := make([]*account.Account, len(rep))

	i := 0
	for key, value := range rep {
		result[i] = &account.Account{Username: key, Balance: value}
		i++
	}

	return result
}

func (rep repository) UpdateBalance(username string, amount float32) (float32, error) {
	if _, ok := rep[username]; !ok {
		return 0, fmt.Errorf("'%s' not found", username)
	}

	rep[username] = amount

	return amount, nil
}

func (rep repository) Count() int {
	return len(rep)
}

func NewAccountRepositoryMap() account.IAccountRepository {
	return make(repository)
}
