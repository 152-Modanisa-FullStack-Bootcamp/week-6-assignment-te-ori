package account

type Account struct {
	Username string
	Balance  float32
}

type IAccountService interface {
	AccountOf(username string) *Account
	Accounts() []*Account
	Put(username string) (float32, error)
	UpdateBalance(username string, amount float32) error
}

type IAccountRepository interface {
	AccountOf(username string) *Account
	Accounts() []*Account
	Put(username string, initBalance float32) (float32, error)
	UpdateBalance(username string, amount float32) error
	Count() int32
}
