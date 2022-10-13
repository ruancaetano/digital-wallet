package entity_test

import (
	"testing"

	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := entity.NewClient("John", "john@email.com")

	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.Name, "John")
	assert.Equal(t, client.Email, "john@email.com")
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := entity.NewClient("", "")

	assert.Nil(t, client)
	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}

func TestUpdateClient(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")

	assert.NotNil(t, client)

	lastUpdatedAt := client.UpdatedAt

	err := client.Update("John 2", "john2@email.com")

	assert.Nil(t, err)
	assert.Equal(t, client.Name, "John 2")
	assert.Equal(t, client.Email, "john2@email.com")
	assert.NotEqual(t, client.UpdatedAt, lastUpdatedAt)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := entity.NewClient("John", "john@email.com")

	assert.NotNil(t, client)

	err := client.Update("", "")

	assert.NotNil(t, err)
	assert.Error(t, err, "name is required")
}
