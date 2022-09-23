package greetings

import (
	"errors"
	"fmt"
)

// This is a comment

func Hello(name string) (string, error) {
	//If no name was given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
