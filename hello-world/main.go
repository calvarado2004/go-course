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

//	About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 5)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 5 is: %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {

	var x float32
	var y float32

	x = 100
	y = 23

	f, err := divideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by zero")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
}

func divideValues(x, y float32) (float32, error) {

	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {

	http.HandleFunc("/hello-world", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
