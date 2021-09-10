package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the user input module"
	fmt.Println(welcome)

	// Taking user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the pizza rating: ")

	// comma ok || err syntax
	/*
		In go there is no try catch block
		so there is a ok err syntax which
		takes the input or err respectively.

		syntax:
				input, err := reader.ReadString('\n')
	*/
	rating, _ := reader.ReadString('\n')
	fmt.Println("User Rating: ", rating)

}
