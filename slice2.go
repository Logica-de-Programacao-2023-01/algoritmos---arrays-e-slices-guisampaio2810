package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Print(numeros)

	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Println(numeros)

	numeros = numeros[:len(numeros)-1]
	fmt.Println(numeros)

	numeros = numeros[1:]
	fmt.Println(numeros)
}
