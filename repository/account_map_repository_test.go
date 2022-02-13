package repository

import (
	"my_account/account"
	"my_account/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	MIN_AMOUNT   = float32(-100)
	INIT_BALANCE = float32(0)
)

func TestPutAccount(t *testing.T) {
	t.Run("Put account which does not exist", func(t *testing.T) {
		username := "orhan"
		repository := NewAccountRepositoryMap()

		assert.Equal(t, 0, repository.Count())
		amount, err := repository.Put(username, INIT_BALANCE)

		assert.Nil(t, err)
		assert.Equal(t, INIT_BALANCE, amount)
		assert.Equal(t, 1, repository.Count())
	})
}

func TestGetAccount(t *testing.T) {
	username := "orhan"
	repository := NewAccountRepositoryMap()

	assert.Equal(t, 0, repository.Count())
	repository.Put(username, INIT_BALANCE)

	user := repository.AccountOf(username)

	assert.Equal(t, username, user.Username)
	assert.Equal(t, INIT_BALANCE, user.Balance)
}
func TestGetAccounts(t *testing.T) {
	mockReturnAccounts := []*account.Account{
		utils.CreateTestAccount("user1", 100),
		utils.CreateTestAccount("user2", -50),
		utils.CreateTestAccount("user3", 70),
		utils.CreateTestAccount("user4", 23.05),
	}

	repository := NewAccountRepositoryMap()

	assert.Equal(t, 0, repository.Count())

	for _, e := range mockReturnAccounts {
		repository.Put(e.Username, e.Balance)
	}

	accounts := repository.Accounts()

	assert.Equal(t, len(mockReturnAccounts), len(accounts))
}

func TestUpdateBalance(t *testing.T) {
	username := "orhan"
	amount := float32(20)

	t.Run("Update Balance of existing user", func(t *testing.T) {
		repository := NewAccountRepositoryMap()

		repository.Put(username, 50)

		newBalance, err := repository.UpdateBalance(username, amount)

		assert.Nil(t, err)
		assert.Equal(t, float32(20), newBalance)
	})
	t.Run("Update Balance of nonexisting user", func(t *testing.T) {
		repository := NewAccountRepositoryMap()

		newBalance, err := repository.UpdateBalance(username, amount)

		assert.NotNil(t, err)
		assert.Equal(t, float32(0), newBalance)
	})
}
