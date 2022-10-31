package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Kadek",
			"last_name": "Gadget",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Nunuk",
			"last_name": "Gadget",
			"hair_color": "red",
			"has_dog": false
		}
	]
	`
	log.Println(myJson)

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)
	log.Printf("unmarshalled: %v", unmarshalled[0].FirstName)

	// write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "west"
	m1.HairColor = "blue"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Brandy"
	m2.LastName = "Brows"
	m2.HairColor = "brown"
	m2.HasDog = true

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "      ")

	if err != nil {
		log.Println("error marshalling: ", err)
	}

	fmt.Println(string(newJson))
}
