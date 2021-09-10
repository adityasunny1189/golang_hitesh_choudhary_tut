package main

import (
	// "math/rand"
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// random numbers
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1)

	// random number from crypto
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("Random number: ", myRandomNum)
}
