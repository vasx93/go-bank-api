package repository

//! READ THIS https://github.com/amitshekhariitbhu/go-backend-clean-architecture/blob/main/repository/user_repository.go
// import (
// 	"database/sql"
// 	"fmt"

// 	db "github.com/vasx93/go-bank-api/internal/db"
// )

// type AccountRepository struct {
// 	storage *sql.DB
// }

// func New(s *sql.DB) *AccountRepository {
// 	return &AccountRepository{storage: s}
// }

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
