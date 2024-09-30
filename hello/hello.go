package main

import (
	"fmt"

	"go-tutorial/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Yuhyeon")
	fmt.Println(message)
}
