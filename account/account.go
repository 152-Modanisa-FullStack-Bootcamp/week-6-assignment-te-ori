package account

type Account struct {
	Username string  `json:"username"`
	Balance  float32 `json:"balance"`
}

type IAccountService interface {
	AccountOf(username string) *Account
	Accounts() []*Account
	Put(username string) (float32, error)
	UpdateBalance(username string, amount float32) (float32, float32, error)
}

type IAccountRepository interface {
	AccountOf(username string) *Account
	Accounts() []*Account
	Put(username string, initBalance float32) (float32, error)
	UpdateBalance(username string, amount float32) (newBalance float32, err error)
	Count() int
}
