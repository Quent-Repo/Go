package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	var v string
	var sep string
	for i := 1; i < 10; i++ {
		s += sep + strconv.Itoa(i)
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(v)
}
