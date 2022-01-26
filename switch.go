package main

import "fmt"

func main() {
	modulo := 5%5

	switch modulo {
	case 0:
		fmt.Println("es par")
	default:
		fmt.Println("es impar")
	}
}