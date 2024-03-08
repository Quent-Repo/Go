package main

import (
	"fmt"
	"strings"
	"math"
)

func main() {
	var sei = [20]string{"run", "ore", "ovd"}
	var works = []string{"window plane", "coffee black"}
	var sun string = "bass"
	const work int = 50
	fmt.Println("car", work, sun)
	sun = "change"
	// fmt.Println(sun)
	// fmt.Scan(&sun)
	// fmt.Println(sun)
	g := ""
	fmt.Scan(&g)
	works = append(works, g)

	fmt.Println(works)
	fmt.Println(sei)

	// if sun == "gun" {
	// 	fmt.Println("that right")
	// } else {
	// 	fmt.Println("This is not right " + sun)
	// }
	firstSun := []string{}
	for _, works := range works {
		var names = strings.Fields(works)
		firstSun = append(firstSun, names[0])
	}
	fmt.Print(firstSun)
	for index, element := range works {
		fmt.Println(index, element)
	}

	p := new(int)
	q := new(int)
	fmt.Println(*p)
	fmt.Print(p == q)
	*p = 2
	fmt.Println(*p)

	for t := 0.0;t<cycles*2*math.Pi; t+= res{
		x:= math.Sin(t)
		y:=math.Sin(t*freq + phase)

	}
}
