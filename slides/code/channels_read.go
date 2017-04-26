// +build OMIT

package main

import (
	"fmt"
)

// START OMIT
func main() {
	fmt.Println("Start")
	c := make(chan string) // Make channel

	v := <-c // Reading from a channel blocks until someone writes to it from it // HL
	fmt.Println(v)
}

// STOP OMIT
