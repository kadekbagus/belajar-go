package main

import (
	"fmt"
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

func main() {
	fmt.Println(s)
	log.Println("tes")

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	fmt.Println(user.FirstName)
}
