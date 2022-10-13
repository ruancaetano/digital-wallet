package entity_test

import (
	"testing"

	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaciton(t *testing.T) {
	client1, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client1)
	account1, _ := entity.NewAccount(client1)
	assert.NotNil(t, account1)

	client2, _ := entity.NewClient("John2", "john2@email.com")
	assert.NotNil(t, client2)
	account2, _ := entity.NewAccount(client2)
	assert.NotNil(t, account2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := entity.NewTransaction(account1, account2, float64(100))

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1.Balance, float64(900))
	assert.Equal(t, account2.Balance, float64(1100))
}

func TestCreateNewTransacitonWithInvalidAmount(t *testing.T) {
	client1, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client1)
	account1, _ := entity.NewAccount(client1)
	assert.NotNil(t, account1)

	client2, _ := entity.NewClient("John2", "john2@email.com")
	assert.NotNil(t, client2)
	account2, _ := entity.NewAccount(client2)
	assert.NotNil(t, account2)

	account1.Credit(1000)
	account2.Credit(1000)

	_, err := entity.NewTransaction(account1, account2, float64(0))

	assert.Error(t, err, "amount must be greater than zero")

	_, err = entity.NewTransaction(account1, account2, float64(2000))
	assert.Error(t, err, "insufficient balance")
}

func TestCreateNewTransacitonWithAccounts(t *testing.T) {
	client1, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client1)
	account1, _ := entity.NewAccount(client1)
	assert.NotNil(t, account1)

	account1.Credit(1000)

	_, err := entity.NewTransaction(account1, nil, float64(100))
	assert.Error(t, err, "invalid accounts")
}
