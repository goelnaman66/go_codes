package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)

	changeValue(&myString)
	log.Println("myString is set to", myString)
}

func changeValue(s *string) {
	newValue := "Red"
	*s = newValue
}
