package main

import (
	"errors"
	"log"
)

func main() {

	result, err := devide(100.0, 10.0)
	if err != nil {
		log.Println(err)
	}

	log.Println("result of division is", result)
}

func devide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("cannot devide by 0")
	}

	result = x / y

	return result, nil
}
