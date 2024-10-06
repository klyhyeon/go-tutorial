package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.

	/* var message string // long way to initialize var
	/message = fmt.Sprintf("Hi, %v. Welcome!", name)
	*/
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func Again(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome again!", name)
	return message, nil
}
