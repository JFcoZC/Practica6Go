/*
*Programa que crea una matriz de floats aleatorios
*/
package main

import "fmt"
import "math/rand"

//FUNCIONES
func createMatrix(columns int, rows int, maxRandom float64) [][]float64 {

	matrix := make([][]float64, rows )

	for i := 0; i < rows; i++{
		matrix[i] = make([]float64, columns)
	}//End for 

	for i:= 0; i < rows; i++{

		for j:= 0; j < columns; j++{
			matrix[i][j] = rand.Float64()*maxRandom
		}//End for 2

	}//End for 1

	return matrix
}

//Inicio del programa
func main() {
	resultado := createMatrix(4,4,9)
	//Print vector
	fmt.Printf("%v",resultado)
}