package database

import (
	"database/sql"

	"github.com/devfullcycle/fcutils/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupTestSuite() {
	db, err := sql.Open("sqlite3", ":memory")
	s.Nil(err)
	s.db = db
	db.Exec("Create table clients (id varchar(255), name varchar(255), email varchar(255), created_At date)")
	db.Exec("Create table accounts (id varchar(255), client_id varchar(255), balance int, craeted_At date)")
	s.client, _ = entity.NewClient("John", "j@j")
}
