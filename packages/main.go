package main

import (
	"log"

	"github.com/kadekbagus/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType

	myVar.TypeName = "kadek"
	myVar.TypeNumber = 100

	log.Println(myVar.TypeName)
}
