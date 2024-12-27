package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "jack"
	log.Println(myMap["dog"])

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	log.Println(myMap2["First"])

	myMap3 := make(map[string]User)
	me := User{
		FirstName: "Naman",
		LastName:  "Goel",
	}
	myMap3["Naman"] = me
	log.Println(myMap3["Naman"])

	// slices in go
	var myString []string
	myString = append(myString, "Naman")
	myString = append(myString, "Shubh")
	myString = append(myString, "Abhay")
	sort.Strings(myString)
	log.Println(myString)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sort.Ints(numbers)
	log.Println(numbers)
	log.Println(numbers[6:9])

	names := []string{"one", "two", "three", "four"}
	log.Println(names)
}
