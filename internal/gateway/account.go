package gateway

import "github.com/devfullcycle/fcutils/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	UpdateBalance(acount *entity.Account) error
}
