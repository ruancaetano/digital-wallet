package create_transaction_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/ruancaetano/digital-wallet-core/internal/entity"
	"github.com/ruancaetano/digital-wallet-core/internal/usecase/create_transaction"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Transaction gateway mock
type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
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
	mockedAccountFromId := uuid.New().String()
	mockedAccountToId := uuid.New().String()

	mockedAccountGateway := &AccountGatewayMock{}
	mockedAccountGateway.On("Get", mockedAccountFromId).Return(&entity.Account{
		Balance: 1000,
		Client: &entity.Client{
			ID: uuid.NewString(),
		},
		ID: mockedAccountFromId,
	}, nil)
	mockedAccountGateway.On("Get", mockedAccountToId).Return(&entity.Account{
		ID: mockedAccountToId,
		Client: &entity.Client{
			ID: uuid.NewString(),
		},
		Balance: 1000,
	}, nil)

	mockedTransactionGateway := &TransactionGatewayMock{}
	mockedTransactionGateway.On("Create", mock.Anything).Return(nil)

	uc := create_transaction.NewCreateTransactionUseCase(mockedTransactionGateway, mockedAccountGateway)

	output, err := uc.Execute(create_transaction.CreateTransactionInputDTO{
		AccountIDFrom: mockedAccountFromId,
		AccountIDTo:   mockedAccountToId,
		Amount:        500,
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, output.ID)

	mockedTransactionGateway.AssertExpectations(t)
	mockedTransactionGateway.AssertNumberOfCalls(t, "Create", 1)

	mockedAccountGateway.AssertExpectations(t)
	mockedAccountGateway.AssertNumberOfCalls(t, "Get", 2)
}
