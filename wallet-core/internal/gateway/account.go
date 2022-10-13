package gateway

import "github.com/ruancaetano/digital-wallet-core/internal/entity"

type AccountGateway interface {
	Get(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
