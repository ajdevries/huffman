// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println("Start")
	c := make(chan string, 1) // Make channel with buffer size of 1
	c <- "Hello"              // Writing to a buffered channel won't blocks // HL
}

// STOP OMIT
