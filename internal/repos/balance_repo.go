package repos

import (
	"database/sql"

	"github.com/Kennedy-lsd/GoBank/data"
)

type BalanceRepo struct {
	DB *sql.DB
}

func NewBalanceRepo(db *sql.DB) *BalanceRepo {
	return &BalanceRepo{
		DB: db,
	}
}

func (r *BalanceRepo) GetAll() ([]data.Balance, error) {
	var balances []data.Balance
	query := "SELECT * FROM balances"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var balance data.Balance

		err := rows.Scan(&balance.Id, &balance.Amount, &balance.CreatedAt, &balance.AccountId)

		if err != nil {
			return nil, err
		}

		balances = append(balances, balance)
	}

	return balances, nil
}

func (r *BalanceRepo) Create(balance *data.Balance) (*data.Balance, error) {
	query := "INSERT INTO balances (amount, account_id) VALUES ($1, $2) RETURNING id, created_at"

	err := r.DB.QueryRow(query, &balance.Amount, &balance.AccountId).Scan(&balance.Id, &balance.CreatedAt)

	if err != nil {
		return nil, err
	}

	return balance, nil
}
