package gateway

import "github.com/devfullcycle/fcutils/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
