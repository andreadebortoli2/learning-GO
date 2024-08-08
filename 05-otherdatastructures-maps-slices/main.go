package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// MAPS
	myMap := make(map[string]string)
	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap2["me"] = me

	log.Println(myMap2["me"].FirstName, myMap2["me"].LastName)

	// SLICES
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Mary")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5}

	log.Println(numbers)
	log.Println(numbers[0:2])
}
