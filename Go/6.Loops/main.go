package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}
	for i, animal := range animals {
		log.Println(i, animal)
	}

	myMap := make(map[string]string)
	myMap["horse"] = "Charlie"
	myMap["dog"] = "Jack"
	myMap["cat"] = "Cherry"
	for animalType, animal := range myMap {
		log.Println(animalType, animal)
	}

	firstLine := "My name is Naman"
	for i, w := range firstLine {
		log.Println(i, w)
	}

	type User struct {
		FirstName string
		LastName  string
		Phone     string
		Age       int
	}
	var users []User
	users = append(users, User{"Naman", "Goel", "9045946258", 23})
	users = append(users, User{"Shubh", "Goel", "9087654768", 21})

	for i, user := range users {
		log.Println(i, user.FirstName)
	}

}
