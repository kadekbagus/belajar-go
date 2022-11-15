package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func Devide(w http.ResponseWriter, r *http.Request) {
	var x, y float32
	x = 100.0
	y = 0.0
	f, err := devideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func devideValues(x, y float32) (float32, error) {
	result := x / y
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/devide", Devide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

	// go mod init myapp
}
