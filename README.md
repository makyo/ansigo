[![Build Status](https://travis-ci.org/makyo/ansigo.svg?branch=master)](https://travis-ci.org/makyo/ansigo)
[![GoDoc](https://godoc.org/github.com/makyo/ansigo?status.svg)](https://godoc.org/github.com/makyo/ansigo)

# ansigo

ANSI escape-code library for Golang.

[![Pretty colors](/docs/demo.png)](/docs/demo.png")

## Usage

```go
package main

import (
	"fmt"

	ansi "github.com/makyo/ansigo"
)

func main() {
	// You can get attributes...
	bold, err := ansi.Attributes.Find("bold")
	if err != nil {
		panic(err)
	}
	fmt.Println(bold.Apply("Some bold text"))
	fmt.Print("\n")

	// Or colors from three different spaces...
	red1, err := ansi.Colors8.Find("red")
	if err != nil {
		panic(err)
	}
	red2, err := ansi.Colors256.Find("DarkRed")
	if err != nil {
		panic(err)
	}
	red3, err := ansi.Colors24bit.Find("#661126")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s %s %s\n", red1.FG("Three"), red2.BG("different"), red3.BG(red1.FG("reds")))
	fmt.Print("\n")

	// Or you can let ansigo figure out the hard work.
	s, err := ansi.ApplyOne("underline", "Some text.")
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
	fmt.Print("\n")

	// But all this error checking is tiresome, and maybe you don't want to
	// have to keep nesting lookups.
	fmt.Println(ansi.MaybeApply("bold+brightyellow+rgb(145, 20, 31):bg+blink", "Warning!"))
}
```
