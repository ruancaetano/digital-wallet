package entity_test

import (
	"testing"

	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client)

	account, err := entity.NewAccount(client)

	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, account.Client, client)
}

func TestCreateNewAccountWithEmptyClient(t *testing.T) {
	account, err := entity.NewAccount(nil)

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.Error(t, err, "client is required")
}

func TestCreditAccount(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client)

	account, _ := entity.NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100)
	assert.Equal(t, account.Balance, float64(100))

	account.Credit(50)
	assert.Equal(t, account.Balance, float64(150))
}

func TestDebitAccount(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client)

	account, _ := entity.NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100)
	assert.Equal(t, account.Balance, float64(100))

	err := account.Debit(50)
	assert.Equal(t, account.Balance, float64(50))
	assert.Nil(t, err)

	err = account.Debit(50)
	assert.Equal(t, account.Balance, float64(0))
	assert.Nil(t, err)
}

func TestDebitAccountWhenInvalidAmount(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")
	assert.NotNil(t, client)

	account, _ := entity.NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100)
	assert.Equal(t, account.Balance, float64(100))

	err := account.Debit(1000)

	assert.Error(t, err, "insufficient balance")
}
