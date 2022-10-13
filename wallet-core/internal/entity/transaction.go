package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := transaction.Validate()
	if err != nil {
		return nil, err
	}

	transaction.Commit()

	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil || t.AccountTo == nil {
		return errors.New("invalid accounts")
	}

	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient balance")
	}

	return nil
}

func (t *Transaction) Commit() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}
