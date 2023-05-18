package models

type Order struct {
	Id string
	User User
	Address Address
	Status OrderStatus
	Products []Product
}
