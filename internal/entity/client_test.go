package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.Nil(t, client)
	assert.NotNil(t, err)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	err := client.Update("John Vargas", "jv@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Vargas", client.Name)
	assert.Equal(t, "jv@j.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	err := client.Update("", "j@j.com")
	assert.Error(t, err, "name is required")
}

func TestAddAcountToClient(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	account := NewAccount(client)
	err := client.AddAcount(*account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
