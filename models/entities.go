package models

import (
	"time"
)

type Address struct {
	ID      string
	Zip     string
	City    string
	Address string
}

type Recipient struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
	Email     string
}

type Order struct {
	ID        string
	Number    string
	Manager   *User
	Date      time.Time
	Delivery  Address
	Recipient Recipient
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}
