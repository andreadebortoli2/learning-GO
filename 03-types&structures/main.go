package main

import (
	"log"
	"time"
)

var s = "seven"

// in GO ther aren't OOP classes, as a turn around i can declaire a type with its struct, if type an var start with the captal letter they are visible from all packages
/* var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time */
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 555 555-1212",
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber)
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s3, "World"
}
