package main

import "log"

func main() {
	// if statement
	var isTrue bool
	isTrue = true
	num := 100

	if isTrue && num == 100 {
		log.Println("if")
	} else if !isTrue || num != 100 {
		log.Println("else if")
	} else {
		log.Println("else")
	}

	// switch statement
	myVar := "cat"

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
