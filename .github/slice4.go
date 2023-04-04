package main

import (
	"crypto/x509"
	"fmt"
)

func main() {
	var slice {}int
	var , tamanho, x int

	fmt.Print("digite o tamanho do slice")
	fmt.Scan(&tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Print("digite um numero:")
		fmt.Scan(%x)
		slice = append(slice, x)
		soma = soma + x

	}

	fmt.Println(slice)
	fmt.Println(soma)

}
