package main

import (
	"fmt"
	"sort"
)

func anosParaUltrapassar(PA, PB int, G1, G2 float64) int {
	anos := 0
	for PA <= PB {
		// Calcula o crescimento populacional para ambas as cidades
		PA += int(float64(PA) * (G1 / 100))
		PB += int(float64(PB) * (G2 / 100))
		anos++

		// Verifica se passou de 100 anos
		if anos > 100 {
			return anos
		}
	}
	return anos
}

func main() {
	fmt.Println("Hello World!")

	fmt.Println("----------------------------------------------------")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	sides := []float64{a, b, c}
	sort.Sort(sort.Reverse(sort.Float64Slice(sides)))
	a, b, c = sides[0], sides[1], sides[2]

	if a >= b+c {
		fmt.Println("NAO FORMA TRIANGULO")
		return
	}

	if a*a == b*b+c*c {
		fmt.Println("TRIANGULO RETANGULO")
	}

	if a*a > b*b+c*c {
		fmt.Println("TRIANGULO OBTUSANGULO")
	}

	if a*a < b*b+c*c {
		fmt.Println("TRIANGULO ACUTANGULO")
	}

	if a == b && b == c {
		fmt.Println("TRIANGULO EQUILATERO")
		return
	}

	if a == b || b == c || a == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
	fmt.Println("----------------------------------------------------")
	var T, PA, PB int
	var G1, G2 float64

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&PA, &PB, &G1, &G2)
		anos := anosParaUltrapassar(PA, PB, G1, G2)

		if anos > 100 {
			fmt.Println("Mais de 1 seculo.")
		} else {
			fmt.Printf("%d anos.\n", anos)
		}
	}
}
