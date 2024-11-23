package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve the value of the FLAG environment variable
	flag := os.Getenv("FLAG")

	// Check if FLAG is set
	if flag == "" {
		fmt.Println("FLAG environment variable not set")
	} else {
		fmt.Printf("The value of FLAG is: %s\n", flag)
	}
}
