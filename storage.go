package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	GetAccount(username string) (*Account, error)
	UpdateAccount(*Account) error
}

type PostgresStore struct {
	db *sql.DB
}

func (s *PostgresStore) CreateAccount(account *Account) (_ error) {
	query := `insert into account
		(username,hash_password,status,created_at)
		values 
		($1,$2,$3,$4)`
	_, err := s.db.Query(query, account.Username, account.HashPassword, 1, time.Now().Unix())
	if err != nil {
		return err
	}

	return nil
}
func (s *PostgresStore) GetAccount(username string) (*Account, error) {
	query := `select id,username,hash_password,point,status from account where username = $1`
	row := s.db.QueryRow(query, username)
	account := Account{}
	switch err := row.Scan(&account.UserID, &account.Username, &account.HashPassword, &account.Point, &account.Status); err {
	case nil:
		return &account, nil
	case sql.ErrNoRows:
		return nil, err
	}
	return &account, nil
}
func (s *PostgresStore) UpdateAccount(_ *Account) (_ error) {
	return nil
}

func NewPostgresStorage() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account (
		id serial primary key, 
		username varchar(50), 
		hash_password varchar(50),
		point int, 
		created_at int8,
		updated_at int8,
		status int2
	)`

	_, err := s.db.Exec(query)
	return err
}
