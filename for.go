package main

import "fmt"

func main()  {
	// for condicional
	for i := 0; i < 10; i++{
		fmt.Println(i);
	}
	//for while
	count := 0
	for count <+10{
		println(count)
		count++
	}
	// for forever, ciclo de for que itera sin final
	/* countForever :=0
	for{
		println(countForever)
		countForever++
	} */
	
	//for range que es cuando se tiene una colección de un objeto

	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}
	//
}