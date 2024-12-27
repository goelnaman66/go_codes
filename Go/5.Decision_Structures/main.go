package main

import "log"

func main() {
	var isTrue bool
	isTrue = true

	if isTrue {
		log.Println("isTrue is True")
	} else {
		log.Println("is True is False")
	}

	///////////////////////////////////////////////////////
	num := 100
	if isTrue && num > 99 {
		log.Println("Number is greter than 99")
	} else {
		log.Println("Not True")
	}
	////////////////////////////////////////////////////////////

	
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")
	
	default:
		log.Println("myVar is set to something else")
	}

}
