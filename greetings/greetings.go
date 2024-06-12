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

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string) // 创建一个key和value都是字符串的map
	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}
	return messages, nil
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
