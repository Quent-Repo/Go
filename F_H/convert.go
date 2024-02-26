package main

import (
	"fmt"
)

func main() {
	for true {
		F := 0
		fmt.Println("Enter temp in °F to get back °C")
		fmt.Scan(&F)
		fmt.Printf("%v°F is ", F)
		fmt.Println((F-32)*5/9, `°C`)
	}
}
