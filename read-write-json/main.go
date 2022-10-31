package main

import "log"

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
}
