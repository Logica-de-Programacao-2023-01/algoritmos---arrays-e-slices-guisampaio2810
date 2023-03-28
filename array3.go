package main

import "fmt"

func main() {
	lista := [4]int{5, 7, 9, 12}
	produto := 1

	for i := 0; i < len(lista); i++ {
		produto = produto * lista[i]
	}
	fmt.Println(produto)
}
