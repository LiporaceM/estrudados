package main

import (
	"fmt"
	"sort"
)

func triangulo(a float64, b float64, c float64) {
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
	} else if a == b || b == c || a == c {
		fmt.Println("TRIANGULO ISOSCELES")
	}
}

func main() {
	triangulo(7.0, 5.0, 7.0)
	triangulo(6.0, 6.0, 10.0)
	triangulo(6.0, 6.0, 6.0)
	triangulo(5.0, 7.0, 2.0)
	triangulo(6.0, 8.0, 10.0)
}
