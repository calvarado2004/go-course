package main

import (
	"log"

	"github.com/calvarado2004/myniceprogram/helpers"
)

func main() {

	log.Println("Program using custom modules")

	var myVar helpers.TypeUser
	myVar.TypeName = "Carlos Alvarado"
	myVar.TypeAge = 33
	myVar.TypeEmail = "carlos@email.com"
	log.Println(myVar.TypeName)
	log.Println(myVar.TypeEmail)

}
