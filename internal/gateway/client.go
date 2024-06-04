package gateway

import "github.com/devfullcycle/fcutils/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
