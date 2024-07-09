package main

import (
	"fmt"

	"example.com/hello/greetings"
)

func main() {
	message := greetings.Hello("Gladys")

	fmt.Println(message)
}
