package main

import (
	"fmt"
	"math"
)

func main() {
	const PI = 3.1415
	var raio = 3.2
	area := PI * math.Pow(raio, 2)

	a, b, c := 1, 2, 3

	fmt.Println(a, b, c)

	fmt.Println("A área da circunferência é", area)
}
