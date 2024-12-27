package main

import "fmt"

func main() {
	fmt.Println("hello world!!")

	var whatToSay string
	whatToSay = "Hi, there!!"
	fmt.Println(whatToSay)

	var i int
	i = 7
	fmt.Println("i is set to ", i)

	whatWasSaid := say_something()
	fmt.Println((whatWasSaid))

	whatWasSaid_1, whatWasSaid_2 := say_something_2()
	fmt.Println(whatWasSaid_1, whatWasSaid_2)

}

func say_something() string {
	return "something"
}

func say_something_2() (string, string) {
	return "Hii", "Naman"
}
