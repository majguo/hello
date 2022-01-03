package main

import (
	"fmt"
	"log"

	"github.com/majguo/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Jianguo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	names := []string{"Alan", "Yan", "Jianguo"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
