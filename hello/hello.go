package main

import (
	"fmt"
	"log"

	G "github.com/7Afeather/learn-go/greetings"
)

// func main() {
// 	// 获取问候消息并打印。
// 	message := G.Hello("Gladys")
// 	fmt.Println(message)
// }

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := G.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}
