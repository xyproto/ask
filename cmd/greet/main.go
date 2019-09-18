package main

import (
	"fmt"
	"github.com/xyproto/ask"
)

func main() {
	var (
		correct bool
		name    string
	)
	for !correct {
		name = ask.Ask("What is your name? ")
		correct = ask.YesNo("Your name is "+name+"?", false)
	}
	fmt.Printf("Greetings, %s!\n", name)
}
