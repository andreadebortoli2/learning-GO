package main

import (
	"log"
	"packages-tutorial/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Andrea"

	log.Println(myVar.TypeName)
}
