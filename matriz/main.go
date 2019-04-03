/*
*Programa que crea una matriz de 0s
*/
package main

import "fmt"

//FUNCIONES
func createMatrix(columns int, rows int) [][]float64 {

	matrix := make([][]float64, rows )

	for i := 0; i < rows; i++{
		matrix[i] = make([]float64, columns)
	}//End for 

	return matrix
}

//Inicio del programa
func main() {
	resultado := createMatrix(2,2)
	//Print vector
	fmt.Printf("%v",resultado)
}