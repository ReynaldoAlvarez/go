package main

import "fmt"

func main() {
	normalFunction("hola vidal")
	tripleArg(1,"vidal",true)
	value := returnValue(10)
	fmt.Println(value)
	value1, value2 := doubleReturn(2)
	fmt.Println("valor1: ", value1)
	fmt.Println("valor2: ", value2)
	// cuando queremos solo un valor de los dos que nos retorna la funcion
	valor1, _ :=doubleReturn(5)
	fmt.Println("valor: ", valor1)
	_, valor2 :=doubleReturn(5)
	fmt.Println("valor: ", valor2)
}

func tripleArg(id int, name string, status bool){
	fmt.Println("id", id)
	fmt.Println("name: ", name)
	fmt.Println(" status: ", status)
	
}
func normalFunction(message string){
	fmt.Println(message)
}

// cuando una funcion nos retorna un valor

func returnValue(a int) int{
	return a*2
}

// cuando requerimos que nos retorne 2 o mas valores

func doubleReturn(a int) (c, d int){
	return a, a+2
}