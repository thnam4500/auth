package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAdd string
	store     Storage
}

func NewAPIServer(listenAdd string, store Storage) *APIServer {
	return &APIServer{
		listenAdd: listenAdd,
		store:     store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleGetAccount))
	router.HandleFunc("/register", makeHTTPHandleFunc(s.handleMethodRegister))
	router.HandleFunc("/login", makeHTTPHandleFunc(s.handleMethodLogin))
	log.Println("JSON API running on port:", s.listenAdd)
	http.ListenAndServe(s.listenAdd, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccount(w, r)
	}
	if r.Method == "PUT" {
	}
	if r.Method == "DELETE" {
	}
	return fmt.Errorf("method not allow")
}

func (s *APIServer) handleMethodLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleLogin(w, r)
	}
	return fmt.Errorf("method not allow")
}

func (s *APIServer) handleMethodRegister(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleRegister(w, r)
	}
	return fmt.Errorf("method not allow")
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	username := r.URL.Query().Get("username")
	account, err := s.store.GetAccount(username)
	if err != nil {
		return err
	}
	getAccount := GetAccount{
		UserID:   account.UserID,
		Username: account.Username,
		Point:    account.Point,
		Status:   account.Status,
	}
	return WriteJSON(w, http.StatusOK, getAccount)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccountRequest := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(&createAccountRequest); err != nil {
		return err
	}
	account := NewAccount(createAccountRequest.Username, createAccountRequest.Password)
	err := s.store.CreateAccount(account)
	if err != nil {
		return err
	}
	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	loginRequest := new(LoginRequest)
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		return err
	}
	account, err := s.store.GetAccount(loginRequest.Username)
	if err != nil {
		return err
	}
	if account.HashPassword != loginRequest.Password {
		return fmt.Errorf("password not correct")
	}
	return WriteJSON(w, http.StatusOK, map[string]interface{}{"status": "success"})
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	return s.handleCreateAccount(w, r)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
