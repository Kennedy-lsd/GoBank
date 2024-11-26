package repos

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/Kennedy-lsd/GoBank/data"
)

type AccountRepo struct {
	DB *sql.DB
}

func NewAccoutRepo(db *sql.DB) *AccountRepo {
	return &AccountRepo{
		DB: db,
	}
}

func (r *AccountRepo) GetAll() ([]data.Account, error) {
	query := "SELECT * FROM accounts s LEFT JOIN balances c ON s.id = c.account_id"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accountsMap := make(map[uint]*data.Account)
	for rows.Next() {
		var account data.Account
		var balance data.Balance

		err = rows.Scan(
			&account.Id, &account.Name, &account.Age, &balance.Id, &balance.Amount, &balance.CreatedAt, &balance.AccountId,
		)
		if err != nil {
			return nil, err
		}

		// Initialize stock comments if not already done
		if _, exists := accountsMap[account.Id]; !exists {
			account.Balances = []data.Balance{} // Initialize to an empty slice
			accountsMap[account.Id] = &account
		}

		if balance.Id != nil {
			accountsMap[account.Id].Balances = append(accountsMap[account.Id].Balances, balance)
		}
	}

	// Convert map to slice
	var accounts []data.Account
	for _, account := range accountsMap {
		accounts = append(accounts, *account)
	}
	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].Id < accounts[j].Id
	})

	return accounts, nil
}

func (r *AccountRepo) GetOne(id int64) (*data.Account, error) {
	query := "SELECT * FROM accounts WHERE id = $1"

	var account data.Account

	err := r.DB.QueryRow(query, id).Scan(&account.Id, &account.Name, &account.Age)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) GetByName(name string) (*data.Account, error) {
	query := "SELECT * FROM accounts WHERE name = $1"

	var account data.Account

	err := r.DB.QueryRow(query, name).Scan(&account.Id, &account.Name, &account.Age)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *AccountRepo) Create(account *data.Account) error {
	query := "INSERT INTO accounts (name, age) VALUES ($1, $2) RETURNING id"

	err := r.DB.QueryRow(query, &account.Name, &account.Age).Scan(&account.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepo) Delete(id int64) error {
	query := "DELETE FROM accounts WHERE id = $1"

	result, err := r.DB.Exec(query, id)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no stocks found with the given ID")
	}

	return nil

}
