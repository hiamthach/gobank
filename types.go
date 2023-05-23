package main

import (
	"math/rand"
	"strconv"
)

type Account struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Number    int64  `json:"number"`
	Balance   int64  `json:"balance"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID:        strconv.Itoa(rand.Intn(100000)),
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Int63n(10000000000)),
	}
}
