/* valor1 == valor2: Retorna TRUE si valor1 y valor2 son exactamente iguales.
valor1 != valor2: Retorna TRUE si valor1 es diferente de valor2.
valor1 < valor2: Retorna TRUE si valor1 es menor que valor2
valor1 > valor2: Retorna TRUE si valor1 es mayor que valor2
valor1 >= valor2: Retorna TRUE si valor1 es igual o mayor que valor2
valor1 <= valor2: Retorna TRUE si valor1 es menor o igual que valor2.

Operadores lógicos
Son aquellos que retorna TRUE o FALSE si cumplen o no una condición utilizando puertas lógicas.

Operador AND:
Este operador indica que todas las condiciones declaradas deben cumplirse para poderse marcar como TRUE. En Go, se utiliza este símbolo &&.

Ejemplo1: 1>0 && 2 >0 Esto retornará TRUE porque tanto la primera como la segunda condición son verdaderas.

Ejemplo2: 2<0 && 1 > 0 Esto retornará FALSE porque una de las condiciones no es verdadera.

Operador OR:
Este operador indica que al menos una de las condiciones debe cumplirse para marcarse como TRUE. En Go, se representa con el símbolo ||.

Ejemplo: 2<0 || 1 > 0 Esto retornará TRUE porque la segunda condición se cumple, a pesar que la primera no.

Operador NOT:
Este operador retornará el opuesto al boleano que está dentro de la variable. Ejemplo:

myBool :=  true
fmt.Println(!myBool) // Esto retornará false
Una vez ya estudiado la teoría, en la siguiente clase vamos a ver cómo utilizarlo con más detalles en Go.å */
package main

import "fmt"

func main(){
	valor1 := 1
	valor2 :=1
	if valor1 == 0{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}

	if valor1 == 1 && valor2==2{
		fmt.Println("valor", valor1, valor2)
	}
}