package main

import "fmt"

const UserToken string = "JKA87LP"

func main() {
	fmt.Printf("Variables\n")

	var myName string = "Aditya"
	fmt.Println(myName)
	fmt.Printf("My name datatype is %T\n", myName)

	var age int = 21
	fmt.Println(age)
	fmt.Printf("My age datatype is %T\n", age)

	var cgpa float32 = 7.87
	fmt.Println(cgpa)
	fmt.Printf("My cgpa datatype is %T\n", cgpa)

	// Implicit type
	var website = "adityasunny.com"
	fmt.Println(website)
	fmt.Printf("My website datatype is %T\n", website)

	// no var keyword
	noOfFriends := 10
	fmt.Println(noOfFriends)
	fmt.Printf("My fiends count datatype is %T\n", noOfFriends)

	// const

	fmt.Println(UserToken)
	fmt.Printf("My user token datatype is %T\n", UserToken)

}
