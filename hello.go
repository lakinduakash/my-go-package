package main

import (
	"fmt"
)

import "rsc.io/quote"

func Cool() {
	fmt.Print("cool")
}

func main() {
	fmt.Print("hello")
	quote.Hello()
}
