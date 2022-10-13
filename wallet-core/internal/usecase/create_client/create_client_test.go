package create_client_test

import (
	"testing"

	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/ruancaetano/digital-wallet-core/internal/usecase/create_client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (m *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (m *ClientGatewayMock) Save(client *entity.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func TestCreateClientUseCase_Execute(t *testing.T) {
	m := &ClientGatewayMock{}

	m.On("Save", mock.Anything).Return(nil)

	uc := create_client.NewCreateClientUseCase(m)

	output, err := uc.Execute(create_client.CreateClientInputDTO{
		Name:  "John",
		Email: "john@example.com",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)
	assert.Equal(t, output.Name, "John")
	assert.Equal(t, output.Email, "john@example.com")
	assert.NotEmpty(t, output.CreatedAt)
	assert.NotEmpty(t, output.UpdatedAt)

	m.AssertExpectations(t)
	m.AssertNumberOfCalls(t, "Save", 1)
}
