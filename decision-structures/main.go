package main

import "log"

func main() {

	// if else
	var isTrue bool

	isTrue = false

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}

	myNum := 10
	myBool := false
	if (myNum > 99) && (!myBool) {
		log.Println("nice face")
	} else if myBool {
		log.Println("cat face")
	} else {
		log.Println("dog face")
	}

	// switch case
	myVar := "dogs"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")
	default:
		log.Println("cat is something else")
	}
}
