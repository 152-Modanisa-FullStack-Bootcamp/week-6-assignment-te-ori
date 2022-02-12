package service

import (
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
		Return(&account.Account{Username: username, Balance: float32(customBalance)}).
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
			Return(&account.Account{Username: username, Balance: customBalance}).
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
