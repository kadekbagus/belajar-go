package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// map
	myMap := make(map[string]string)
	myMap["dog"] = "Brandy"
	myMap["cat"] = "koceng"

	log.Println(myMap["cat"])

	myMap2 := make(map[string]int)

	myMap2["id"] = 122
	myMap2["satu"] = 10

	log.Println(myMap2["id"], myMap2["satu"])

	myMap3 := make(map[string]User)

	kadek := User{
		FirstName: "Kadek",
		LastName:  "Bagus",
	}

	myMap3["admin"] = kadek

	log.Println(myMap3["admin"])
	log.Println(myMap3["admin"].LastName)

	//Slices
	var mySlice []string
	mySlice = append(mySlice, "Samurai")
	mySlice = append(mySlice, "Champloo")

	log.Println(mySlice[0])
}
