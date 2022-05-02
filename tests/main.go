package main

import (
	"errors"
	"log"
)

func main() {

	result, err := divide(100, 12)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Result of division is: ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32

	if y == 0 {
		return result, errors.New("you can't divide by zero")
	}

	result = x / y

	return result, nil
}
