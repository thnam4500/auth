package main

import "time"

type AccountStatus int32

var (
	AccountActive   AccountStatus = 1
	AccountDeactive AccountStatus = 0
)

type Account struct {
	UserID       int64  `json:"user_id"`
	Username     string `json:"username"`
	HashPassword string `json:"hash_password"`
	Point        int64  `json:"point"`
	Status       int32  `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}

type GetAccount struct {
	UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	Point    int64  `json:"point"`
	Status   int32  `json:"status"`
}

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetAccountRequest struct {
	Username string `json:"username"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAccount(username string, password string) *Account {
	return &Account{
		Username:     username,
		HashPassword: password,
		Point:        0,
		Status:       int32(AccountActive),
		CreatedAt:    time.Now().Unix(),
		UpdatedAt:    0,
	}
}
