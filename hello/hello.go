package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a falg to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names
	names := []string{"Jeg", "Beter", "Potato"}

	messages, err := greetings.Hellos(names)
	// If error, print error and exit
	if err != nil {
		log.Fatal(err)
	}
	// If no error, print message
	fmt.Println(messages)
}
