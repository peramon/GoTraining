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
	var positivoUint uint = 12
	var numeroComplejo complex128 = 10 + 8i
	fmt.Println("\nDATOS PRIMITIVOS \nuint -> ", positivoUint, "\nComplejo -> ", numeroComplejo)

	// FMT
	fmt.Println("\nFMT")
	helloMessage := "Hello"
	worldMessage := "World"

	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	name := "Paul"
	age := 23

	fmt.Printf("%s tiene la edad de %d anios\n", name, age)
	fmt.Printf("%v tiene la edad de %v anios\n", name, age)

	message := fmt.Sprintf("%s tiene la edad de %d anios", name, age)
	fmt.Println(message)

	// Tipo de datos
	fmt.Println("\nVER EL TIPO DE DATOS")
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("age: %T\n", age)

	// Uso de funciones
	fmt.Println("\nFUNCIONES")
	normalFunction("Hola mundo")
	tripleArgument(1, 5, "Nuevo")
	valor := returnValue(7)
	fmt.Println("Retorno funcion -> ", valor)
	valor1, valor2 := returnTwoMoreValues(3)
	fmt.Println("Funciones con 2 retornos\nValor 1 -> ", valor1, "\nValor 2 -> ", valor2)
	valor3, _ := returnTwoMoreValues(5)
	fmt.Println("Funciones con 2 retornos, trayendo solo 1\nValor 3 -> ", valor3)

	// Practica: Pasar el area de la fuguras calculadas a Funciones
	valorAreaRectangulo := calcularAreaRectanfulo(12, 10)
	fmt.Println("El resultado de la funcion para calcular el area del rectangulo es -> ", valorAreaRectangulo)
	valorAreaTrapecio := calcularAreaTrapecio(3.5, 9.5, 4)
	fmt.Println("El resultado de la funcion para calcular el area del trapecio es -> ", valorAreaTrapecio)
	valorAreaCirculo := calcularAreaCirculo(4)
	fmt.Println("El resultado de la funcion para calcular el area del circulo es -> ", valorAreaCirculo)

	// Ciclo For
	fmt.Println("\nCICLO FOR")
	// For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// For while
	fmt.Println("For while")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter += 2
	}

	// For forever
	/*counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever
	}*/

	// Practica: For invertido
	fmt.Println("For while invertido")
	counterDrecemental := 10
	for counterDrecemental > 0 {
		fmt.Println(counterDrecemental)
		counterDrecemental--
	}

	// CONDICIONALES
	fmt.Println("\nCONDICIONAL IF")
	valorIf := 2
	valorIfTwo := 3

	if valorIf == 1 {

		fmt.Println("Es el numero 1")
	} else {
		fmt.Println("No es el numero 1")
	}

	// With and
	if valorIfTwo == 3 && valorIf == 2 {
		fmt.Println("Es verdad 3 y 2")
	}

	// With or
	if valorIfTwo == 3 || valorIf == 0 {
		fmt.Println("Uno de los 2 si cumple la condicion")
	}

	// Practica: Determinar si un numero es par o impar
	numeroParImpar(8)
	numeroParImpar(3)

	// Practica: Comprobar usuario y password
	revisarOne := checkUserPass("peramon", "123456787")
	revisarTwo := checkUserPass("peramon", "123456")

	fmt.Println("Intento 1 -> ", revisarOne, "\nIntento 2 -> ", revisarTwo)

}

// FUNCIONES

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println("a -> ", a, "\nb -> ", b, "\nc -> ", c)
}

func returnValue(a int) int {
	return a * 2
}

func returnTwoMoreValues(a int) (c, d int) {
	return a * 2, a * 4
}

// Practica
// Area del rectangulo
func calcularAreaRectanfulo(base, altura int) int {
	return base * altura
}

// Area del Trapecio
func calcularAreaTrapecio(baseMayor, baseMenor, altura float64) float64 {
	return ((baseMayor + baseMenor) / 2) * float64(altura)
}

// Area del circulo
func calcularAreaCirculo(radioCirculo int) float64 {
	return 3.14 * float64(radioCirculo)
}

// Par o impar
func numeroParImpar(numero int) {
	if numero != 0 && numero%2 == 0 {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}
}

// Usuario y pass
func checkUserPass(user, pass string) bool {
	if user == "peramon" && pass == "123456" {
		return true
	} else {
		return false
	}
}
