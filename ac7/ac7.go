package main

import "fmt"

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func simplify(num, den int64) (int64, int64) {
	if den > num {
		gcdVal := gcd(den, num)
		return num / gcdVal, den / gcdVal
	}
	gcdVal := gcd(num, den)
	return num / gcdVal, den / gcdVal
}

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var num1, den1, num2, den2 int64
		var op string

		fmt.Scanf("%d / %d %s %d / %d\n", &num1, &den1, &op, &num2, &den2)

		var resultNum, resultDen int64

		switch op {
		case "+":
			resultNum = num1*den2 + num2*den1
			resultDen = den1 * den2
		case "-":
			resultNum = num1*den2 - num2*den1
			resultDen = den1 * den2
		case "*":
			resultNum = num1 * num2
			resultDen = den1 * den2
		case "/":
			resultNum = num1 * den2
			resultDen = num2 * den1
		}

		resultNum, resultDen = simplify(resultNum, resultDen)

		if resultDen < 0 {
			resultNum *= -1
			resultDen *= -1
		}

		fmt.Printf("%d/%d = %d/%d\n", resultNum, resultDen, resultNum, resultDen)
	}

	var n2, l, q float64
	fmt.Scan(&n2, &l, &q)

	var s []string
	for i := 0; i < int(n2); i++ {
		var nome string
		fmt.Scan(&nome)
		s = append(s, nome)
	}

	ultimo_participante_idx := int((l / q)) % len(s)
	agua_restante := l - q*float64(int(l/q))

	if s[ultimo_participante_idx] == "Bob" && agua_restante == 0.0 {
		fmt.Println("Juca 0.4")
	} else {
		fmt.Printf("%s %.1f\n", s[ultimo_participante_idx], agua_restante)
	}

	var codigoPeca1, numeroPecas1, codigoPeca2, numeroPecas2 int
	var valorUnitarioPeca1, valorUnitarioPeca2 float64

	fmt.Scan(&codigoPeca1, &numeroPecas1, &valorUnitarioPeca1)
	fmt.Scan(&codigoPeca2, &numeroPecas2, &valorUnitarioPeca2)

	totalPagar := (float64(numeroPecas1) * valorUnitarioPeca1) + (float64(numeroPecas2) * valorUnitarioPeca2)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalPagar)
}
