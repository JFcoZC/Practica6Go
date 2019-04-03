/*
*Programa que realiza la multiplicacion de vector(n,m) por matriz(m,l)
*/
package main

import "fmt"
import "math/rand"

//FUNCIONES
func createMatrixAleatoria(columns int, rows int, maxRandom float64) [][]float64 {

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
}//End func create matrix
//-----------------------------------------------------------
func crearVectorAleatorio(size int, maxRandom float64) []float64 {
	array:= make([]float64, size)

	for i := 0; i < size; i++{
		array[i] = rand.Float64()*maxRandom
	}//End for 

	return array
}//End func createVectorAleatorio
//-----------------------------------------------------------
func multiplicacion(size int, maxRandomMatrix float64, maxRandomVector float64)[]float64{

	vector := crearVectorAleatorio(size,maxRandomVector)
	matrix := createMatrixAleatoria(size,size,maxRandomMatrix)

	resultado := make([]float64, size)

	for i:=0; i < size; i++{
		temp := 0.0

		for j:= 0; j < size; j++{
			temp = temp + (vector[j])*(matrix[i][j])
		}//End for 2

		resultado[i] = temp

	}//End for 1

	return resultado

}//End func multiplicacion
//-----------------------------------------------------------
//Inicio del programa
func main() {

	resultado := multiplicacion(3,1,10)
	//Print resultado multiplicacion
	fmt.Printf("%v",resultado)
}