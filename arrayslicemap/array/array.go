package main

import "fmt"

func main() {
	var a [5]float64
	a[0], a[1], a[2], a[3] = 98.6, 30.0, 45.6, 10
	fmt.Println(a)

	total := 0.0
	for i := 0; i < len(a); i++ {
		total += a[i]
	}
	media := total / float64(len(a))
	fmt.Printf("Média: %.2f\n", media)

}
