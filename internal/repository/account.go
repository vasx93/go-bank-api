package repository

import (
	db "github.com/vasx93/go-bank-api/internal/db"
	"github.com/vasx93/go-bank-api/internal/db/postgres"
)

//! READ THIS https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/repository/user_repository.go

// type AccountRepository interface {
// 	CreateAccount(account Account) error
// 	GetAccount(id string) (Account, error)
// 	UpdateAccount(account Account) error
// 	DeleteAccount(id string) error
// 	ListAccounts() ([]Account, error)
// }

type AccountRepository struct {
*postgres.PostgresDB
}

func New(p db.DbConnection) *AccountRepository {
	return &PostgresAccountRepository{p}
}

// func (a *AccountRepository) Find() ([]db.Account, error) {
// 	query := `SELECT * FROM accounts`

// 	rows, err := a.storage.Query(query)
// 	if err != nil {
// 		return nil, fmt.Errorf("could not find accounts: %v", err)
// 	}

// 	defer rows.Close()

// 	var accounts []db.Account
// 	for rows.Next() {
// 		var account db.Account
// 		if err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.EncryptedPassword, &account.Balance, &account.CreatedAt); err != nil {
// 			return nil, fmt.Errorf("could not scan account: %v", err)
// 		}
// 		accounts = append(accounts, account)
// 	}

// 	return accounts, nil
// }
