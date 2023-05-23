package main

import "database/sql"

type Storage interface {
	CreateAccount(*Account) error
	GetAccountByID(string) (*Account, error)
	UpdateAccount(*Account) error
	DeleteAccount(string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	
}docker run --name some-postgres -e POSTGRES_PASSWORD=gobank -p 5432:5432 -d postgres