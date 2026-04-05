package main

import (
	"fmt"
	"math/rand"
)

type Role string

const (
	RoleAdmin    Role = "admin"
	RoleCustomer Role = "customer"
)

type User struct {
	Name   string
	Email  string
	Role   Role
	UserId int
}

type Order struct {
	UserId    int
	ProductId string
	Quantity  string
	Price     float64
	Total     float64
}

type NewApp struct {
	Orders map[int]Order
	Users  map[int]User
}

func (app *NewApp) CreateUser(name, email, role string) int {
	if name == "" || email == "" {
		fmt.Println("Empty name and email")
		return -1
	}

	if role != "admin" && role != "customer" {
		fmt.Println("bad role")
		return -1
	}

	userId := rand.Int()
	app.Users[userId] = User{
		Name:   name,
		Email:  email,
		Role:   Role(role),
		UserId: userId,
	}

	return userId
}
