package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Create a structure to unmarshal the json input. Slice of Person
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	//Json input to unmarshal
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true

		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "brown",
			"has_dog": false
		}

	]`

	//Create a slice var of type Person
	var unMarshalled []Person

	//Asign the json to the slice of Person variable
	err := json.Unmarshal([]byte(myJson), &unMarshalled)

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled content: %v", unMarshalled)

	//write json from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false
	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Peter"
	m2.LastName = "Parker"
	m2.HairColor = "brown"
	m2.HasDog = false
	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	fmt.Println(string(newJson))

}
