package models

type Address struct {
	Country Country
	City string
	Street string
	HouseNumber string
	ApartmentNumber string
}