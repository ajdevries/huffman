// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println("Start")
	c := make(chan string) // Make channel

	c <- "Hello" // Writing to a channel blocks until someone reads from it // HL
	fmt.Println("Never get here")
}

// STOP OMIT
