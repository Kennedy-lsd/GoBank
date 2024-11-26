package data

import "time"

type Account struct {
	Id       uint      `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Balances []Balance `json:"balances"`
}

type Balance struct {
	Id        *uint      `json:"id"`
	Amount    *int64     `json:"amount"`
	CreatedAt *time.Time `json:"created_at"`
	AccountId *uint      `json:"account_id"`
}

type AccountGetter interface {
	GetAll() ([]Account, error)
	GetOne(id int64) (*Account, error)
	GetByName(name string) (*Account, error)
}

type AccountSetter interface {
	Create(*Account) error
}

type AccountDeletter interface {
	Delete(id int64) error
}

type AccountRepository interface {
	AccountGetter
	AccountSetter
	AccountDeletter
}

// Balance

type BalanceGetter interface {
	GetAll() ([]Balance, error)
	// GetOne(id int64) (*Balance, error)
}

type BalanceSetter interface {
	Create(*Balance) (*Balance, error)
}

// type BalanceDeletter interface {
// 	Delete(id int64) error
// }

// type BalanceUpdater interface {
// 	Update(id int64, b *Balance) (*Balance, error)
// }

type BalanceRepository interface {
	BalanceGetter
	BalanceSetter
	// BalanceDeletter
	// BalanceUpdater
}
