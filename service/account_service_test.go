package service

import (
	"fmt"
	"my_account/account"
	"my_account/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	MIN_AMOUNT   = float32(-100)
	INIT_BALANCE = float32(0)
)

func TestGetAccount(t *testing.T) {
	username := "orhan"
	customBalance := float32(200)

	repository := mock.NewMockIAccountRepository(gomock.NewController(t))

	repository.EXPECT().
		AccountOf(username).
		Return(testAccount(username, customBalance)).
		Times(1)

	service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

	result := service.AccountOf(username)

	assert.NotNil(t, result)
	assert.Equal(t, username, result.Username)
	assert.Equal(t, customBalance, result.Balance)
}

func TestPutAccount(t *testing.T) {
	username := "orhan"
	customBalance := float32(200)

	t.Run("Put user not exist", func(t *testing.T) {
		repository := mock.NewMockIAccountRepository(gomock.NewController(t))

		repository.EXPECT().
			AccountOf(username).
			Return(nil).
			Times(1)

		repository.EXPECT().
			Put(username, INIT_BALANCE).
			Return(INIT_BALANCE, nil).
			Times(1)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		initBalance, err := service.Put(username)

		assert.Nil(t, err)
		assert.Equal(t, INIT_BALANCE, initBalance)
	})
	t.Run("Put existing user", func(t *testing.T) {
		repository := mock.NewMockIAccountRepository(gomock.NewController(t))
		repository.EXPECT().
			AccountOf(username).
			Return(testAccount(username, customBalance)).
			Times(1)
		repository.EXPECT().
			Put(username, INIT_BALANCE).
			Return(customBalance, nil).
			Times(0)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		initBalance, err := service.Put(username)

		assert.Nil(t, err)
		assert.Equal(t, customBalance, initBalance)
	})
}

func TestGetAccounts(t *testing.T) {
	mockReturnAccounts := []*account.Account{
		testAccount("user1", 100),
		testAccount("user2", -50),
		testAccount("user3", 70),
		testAccount("user4", 23.05),
	}

	repository := mock.NewMockIAccountRepository(gomock.NewController(t))
	repository.EXPECT().
		Accounts().
		Return(mockReturnAccounts).
		Times(1)

	service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

	result := service.Accounts()

	assert.Equal(t, mockReturnAccounts, result)
}

func TestUpdateBalance(t *testing.T) {
	t.Run("Update balance of unexisting user", func(t *testing.T) {
		username := "orhan"

		repository := mock.NewMockIAccountRepository(gomock.NewController(t))
		repository.EXPECT().
			AccountOf(username).
			Return(nil).
			Times(1)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		newBalance, availableLimit, err := service.UpdateBalance(username, 50)

		assert.Equal(t, fmt.Errorf("'%s' not found", username), err)
		assert.Equal(t, float32(0), newBalance)
		assert.Equal(t, float32(0), availableLimit)

	})

	t.Run("Update balance with positive amount", func(t *testing.T) {
		username := "orhan"
		currentBalance := float32(20)
		newAmont := float32(50)

		repository := mock.NewMockIAccountRepository(gomock.NewController(t))
		repository.EXPECT().
			AccountOf(username).
			Return(testAccount(username, currentBalance)).
			Times(1)

		repository.EXPECT().
			UpdateBalance(username, currentBalance+newAmont).
			Return(currentBalance+newAmont, nil).
			Times(1)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		newBalance, availableLimit, err := service.UpdateBalance(username, newAmont)

		assert.Nil(t, err)
		assert.Equal(t, float32(70), newBalance)
		assert.Equal(t, float32(170), availableLimit)

	})

	t.Run("Update balance with negative amount which does not exceed min amount", func(t *testing.T) {
		username := "orhan"
		currentBalance := float32(20)
		newAmont := float32(-50)

		repository := mock.NewMockIAccountRepository(gomock.NewController(t))
		repository.EXPECT().
			AccountOf(username).
			Return(testAccount(username, currentBalance)).
			Times(1)

		repository.EXPECT().
			UpdateBalance(username, currentBalance+newAmont).
			Return(currentBalance+newAmont, nil).
			Times(1)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		newBalance, availableLimit, err := service.UpdateBalance(username, newAmont)

		assert.Nil(t, err)
		assert.Equal(t, float32(-30), newBalance)
		assert.Equal(t, float32(70), availableLimit)

	})

	t.Run("Update balance with negative amount which exceeds min amount", func(t *testing.T) {
		username := "orhan"
		currentBalance := float32(20)
		newAmont := float32(-130)

		repository := mock.NewMockIAccountRepository(gomock.NewController(t))
		repository.EXPECT().
			AccountOf(username).
			Return(testAccount(username, currentBalance)).
			Times(1)

		repository.EXPECT().
			UpdateBalance(username, currentBalance+newAmont).
			Return(currentBalance+newAmont, nil).
			Times(0)

		service := NewAccountService(INIT_BALANCE, MIN_AMOUNT, repository)

		newBalance, availableLimit, err := service.UpdateBalance(username, newAmont)

		assert.Equal(t, fmt.Errorf("no sufficient balance"), err)
		assert.Equal(t, float32(20), newBalance)
		assert.Equal(t, float32(120), availableLimit)

	})
}

func testAccount(username string, balance float32) *account.Account {
	return &account.Account{Username: username, Balance: balance}
}
