package main

import (
	"fmt"
	"github.com/xyproto/ask"
)

func main() {
	var (
		yes  bool
		name string
	)
	for !yes {
		name = ask.Ask("What is your name? ")
		yes = ask.YesNo("Your name is "+name+"?", false)
	}
	fmt.Printf("Greetings, %s!\n", name)
}
