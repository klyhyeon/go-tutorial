package main

import (
	"fmt"
	"log"

	"go-tutorial/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("yuhyeon")

	// Request a multiple greeting message.
	names := []string{"yuhyeon", "jiwoo"}
	messages, err := greetings.HelloMultiple(names)
	// If an error was returned, print it to the console and exit.
	if err != nil {
		log.Fatal(err)
	}

	// Get a greeting message and print it.
	//message := greetings.Hello("Yuhyeon")
	fmt.Println(message)

	fmt.Println(messages)
}
