package utils

import "my_account/account"

func CreateTestAccount(username string, balance float32) *account.Account {
	return &account.Account{Username: username, Balance: balance}
}
