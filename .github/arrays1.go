package main

import "fmt"

func main() {
	var numeros [3]int
	var x int
	var soma int

	fmt.Print("digite um numero")
	fmt.Scan(&x)

	numeros[0] = x
	numeros[1] = x
	numeros[2] = x

	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}
	fmt.Print(soma)
}
