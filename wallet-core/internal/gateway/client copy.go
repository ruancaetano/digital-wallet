package gateway

import "github.com/ruancaetano/digital-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
