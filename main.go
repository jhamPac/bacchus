package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	nameArg := os.Args[2]
	typeArg := os.Args[4]
	fmt.Printf("The file name is %s.js\nthe logic type is %s", nameArg, typeArg)
	err := ioutil.WriteFile(nameArg+".js", []byte("import React from 'react'\n\nexport default function Marketing(props) {}"), 0755)

	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
