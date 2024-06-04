package createacount

import (
	"github.com/devfullcycle/fcutils/internal/entity"
	"github.com/devfullcycle/fcutils/internal/gateway"
)

type CreateAccountInputDTO struct {
	ClientID string `json:"client_id"`
}

type CreateAccountOutputDTO struct {
	ID string
}

type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

func NewCreateAccountUseCase(a gateway.AccountGateway, c gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: a,
		ClientGateway:  c,
	}
}

func (c *CreateAccountUseCase) Execute(input CreateAccountInputDTO) (*CreateAccountOutputDTO, error) {
	client, err := c.ClientGateway.Get(input.ClientID)
	if err != nil {
		return nil, err
	}
	account := entity.NewAccount(client)
	err = c.AccountGateway.Save(account)
	if err != nil {
		return nil, err
	}
	return &CreateAccountOutputDTO{
		ID: account.ID,
	}, nil
}
