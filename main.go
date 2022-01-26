package main

import (
	"fmt"
)

//import ("math")
func main()  {
	x,e := fmt.Println("hello world")
	fmt.Println(x)
	fmt.Println(e)
	// variable
	// a los nombres de las variables no se les puede asignar como caracteres especiales como $ @, rtc
	var p = "que tal"
	
	fmt.Println(p)
	// si se crea variables sin valor, es necesario declarar el tipo de variable que se almacenara
	var y string
	y= "vidal "
	fmt.Println(y)
	// abreviacion de declaracion de variable
	pais := "bolivia"
	fmt.Println(pais)
	// existe una forma para saver el tipo de valiable de una variable
	fmt.Printf("%T",pais)

	
}