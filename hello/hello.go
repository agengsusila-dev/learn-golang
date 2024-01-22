package main

import (
	"fmt"

	"learn-golang/greetings"
)

func main() {
	message := greetings.Hello("Priti")
	fmt.Println(message)
}
