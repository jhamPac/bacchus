package main

import (
	"fmt"
	"os"
)

func main() {
	firstArg := os.Args[1]
	fmt.Printf("The first argument is %s", firstArg)
}
