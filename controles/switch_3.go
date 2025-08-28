package main

import "fmt"

func notaParaConceito2(n float64) string {
	switch {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 5 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	case n >= 10 || n < 0:
		return "Nota inválida"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(notaParaConceito2(10))
}
