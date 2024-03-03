package main

import (
	"AC3/geometria"
	"fmt"
	"math/rand"
)

func vetor10() {
	vetor := make([]int, 10)

	for i := 0; i < 10; i++ {
		vetor[i] = rand.Intn(100)
	}

	fmt.Println("Elementos do vetor:")
	for _, elemento := range vetor {
		fmt.Println(elemento)
	}
}

func inverterString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func inverte() {
	var inputString string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&inputString)

	fmt.Println("String Invertida:", inverterString(inputString))
}

func main() {
	vetor10()
	fmt.Println("-----------------------------------")
	inverte()
	fmt.Println("-----------------------------------")
	area := geometria.AreaRet(2, 3)
	fmt.Printf("Área do retângulo: %f\n", area)
	fmt.Println("-----------------------------------")
	peri := geometria.PeriRet(2, 3)
	fmt.Printf("Perímetro do retângulo: %f\n", peri)
}
