package main

import "fmt"

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "é um inteiro"
	case float64:
		return "é um float"
	case string:
		return "é uma string"
	case bool:
		return "é um booleano"
	case func():
		return "é uma função"
	default:
		return "não sei que tipo é esse"
	}
}

func main() {
	fmt.Println(tipo(3.14))
	fmt.Println(tipo("Olá"))
	fmt.Println(tipo(42))
	fmt.Println(tipo(true))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo([]int{1, 2, 3}))
}
