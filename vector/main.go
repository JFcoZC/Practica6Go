/*
*Programa que crea un vector de longitud N de integer 64 con 0s
*/
package main

import "fmt"

//FUNCIONES
func crearVector(size int) []int {
	array:= make([]int, size)

	for i := 0; i < size; i++{
		array[i] = 0
	}//End for 

	return array
}

//Inicio del programa
func main() {
	resultado := crearVector(2)
	//Print vector
	fmt.Printf("%v",resultado)
}