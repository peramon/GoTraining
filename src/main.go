package main

import (
	"fmt"
)

func main() {
	// Declaracion de constantes
	fmt.Println("DECLARANDO VARIABLES")
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi1 -> ", pi, "\npi2 -> ", pi2)

	// Declaracion de variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(" int -> ", a, "\nfloat64 -> ", b, "\nstring -> ", c, "\nbool -> ", d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("Area cuadrado = ", areaCuadrado)

	// Operadores aritméticos

	fmt.Println("\nOPERACIONES ARITMETICAS")
	x := 10
	y := 50

	// Suma
	result := x + y

	fmt.Println("suma -> ", result)

	// Resta
	result = y - x
	fmt.Println("resta ->", result)

	// Multiplicacion
	result = x * y
	fmt.Println("multiplicacion -> ", result)

	// Division
	result = y / x
	fmt.Println("division -> ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo -> ", result)

	// Incrementar
	x++
	fmt.Println("Incrementa -> ", x)

	// Decremental
	x--
	fmt.Println("Decremental", x)

	// Area del rectangulo
	baseRectangulo := 15
	alturaRectangulo := 12
	fmt.Println("El area del rectangulo con base 15 y altura 12 es -> ", baseRectangulo*alturaRectangulo)

	// Area del trapecio
	var baseMenorTrapecio float64 = 3.5
	var baseMayorTrapecio float64 = 9.5
	var alturaTRapecio int = 4
	var areaTrapecio float64 = ((baseMayorTrapecio + baseMenorTrapecio) / 2) * float64(alturaTRapecio)
	fmt.Println("Base de un trapecio", " con base menor ", baseMenorTrapecio, ", base mayor", baseMayorTrapecio, "y altura", alturaTRapecio, "-> ", areaTrapecio)

	// Area de un circulo
	var radioCirculo int = 12
	areaCirculo := pi * float64(radioCirculo)

	fmt.Println("EL area del circulo con radio ", radioCirculo, "es -> ", areaCirculo)

	// Datos primitivos
	var positivoUint uint = 12;
	var numeroComplejo complex128 = 10 + 8i
	fmt.Println("\nDATOS PRIMITIVOS \nuint -> ", positivoUint,"\nComplejo -> ",numeroComplejo)




}
