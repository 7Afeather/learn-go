package Greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// // Hello returns a greeting for the named person.
// func Hello(name string) (string, error) {
// 	// If no name was given, return an error with a message.
// 	if name == "" {
// 		return "", errors.New("empty name")
// 	}

// 	// If a name was received, return a value that embeds the name
// 	// in a greeting message.
// 	message := fmt.Sprintf("Hi, %v. Welcome!", name)
// 	return message, nil
// }

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	// rand.Intn(len(formats))生成不大于len(formats)的整数[0,len(formats))
	return formats[rand.Intn(len(formats))]
}
