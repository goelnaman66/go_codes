package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

// using this it will be tied to User struct
func (m *User) printFirstName() string {
	return m.FirstName
}

func main() {
	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	log.Println(say_something("Hello"))

	// using User struct
	user := User{
		FirstName:   "Naman",
		LastName:    "Goel",
		PhoneNumber: "1234567890",
		Age:         23,
	}
	log.Println(user)
	log.Println(user.printFirstName())
}

func say_something(s string) (string, string) {
	return s, "World"
}
