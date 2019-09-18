# ask

Ask the user a question, using text.

[![Build Status](https://travis-ci.org/xyproto/ask.svg?branch=master)](https://travis-ci.org/xyproto/ask) [![GoDoc](https://godoc.org/github.com/xyproto/ask?status.svg)](https://godoc.org/github.com/xyproto/ask) [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/ask/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/ask)](https://goreportcard.com/report/github.com/xyproto/ask)

### Example use

```go
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
```

### General info

* Version: 1.0.0
* Licence: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
