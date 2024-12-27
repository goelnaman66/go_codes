package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page.")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addNumbers(2, 2)
	_, _ = fmt.Fprintf(w, "%s", fmt.Sprintf("The sum of 2+2 is %d", sum))
}

func addNumbers(x, y int) int {
	return x + y
}

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("The port is starting on port : %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
