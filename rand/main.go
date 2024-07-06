package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator to ensure different results each time
	rand.Seed(time.Now().UnixNano())

	// Generate a random integer (0 or 1)
	randomNumber := rand.Intn(2)

	// Print the generated random number
	fmt.Println("Random number:", randomNumber)
}
