package main

import "log"

func main() {
	// loop
	for i := 0; i <= 10; i++ {
		log.Println("oke", i)
	}

	// loop through slice
	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	// loop through map
	animals2 := make(map[string]string)
	animals2["dog"] = "Brandy"
	animals2["cat"] = "hiyo"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	// loop through string
	var firstLine = "Once upon a midnight dreary"

	for letter, quote := range firstLine {
		log.Println(letter, quote)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Marry", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 40})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 10})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age, l.Email)
	}

}
