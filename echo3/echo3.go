package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//sun := "gun"
	wake := " sight"
	//var life string
	fmt.Println(strings.Join(os.Args[1:], wake))
	fmt.Println(os.Args[1:])
}
