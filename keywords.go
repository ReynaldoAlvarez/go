package main

import "fmt"

func main() {
	//deffer; permite ejecutar en ultima instancia donde se encuentre este codigo

	defer fmt.Println("hola")
	fmt.Println("mundo")

	// continue y breack

	for i := 0;i < 10; i++ {
		fmt.Println(i)
		if i == 4{
			fmt.Println("es",i)
		}	

		if i==7{
			fmt.Println("es",i)
			break
		}
	}
}