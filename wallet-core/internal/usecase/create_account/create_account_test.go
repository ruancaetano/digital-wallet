package create_account_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/ruancaetano/digital-wallet-core/internal/usecase/create_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Client gateway mock
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

// Account gateway mock
type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) Get(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	mockedId := uuid.New().String()

	mockedAccountGateway := &AccountGatewayMock{}
	mockedAccountGateway.On("Save", mock.Anything).Return(nil)

	mockedClientGateway := &ClientGatewayMock{}
	mockedClientGateway.On("Get", mock.Anything).Return(&entity.Client{
		ID: mockedId,
	}, nil)

	uc := create_account.NewCreateAccountUseCase(mockedAccountGateway, mockedClientGateway)

	output, err := uc.Execute(create_account.CreateAccountInputDTO{
		ClientID: mockedId,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)

	mockedAccountGateway.AssertExpectations(t)
	mockedAccountGateway.AssertNumberOfCalls(t, "Save", 1)

	mockedClientGateway.AssertExpectations(t)
	mockedClientGateway.AssertNumberOfCalls(t, "Get", 1)
}
