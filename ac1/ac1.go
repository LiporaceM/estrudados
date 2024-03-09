package ac1

import (
	"fmt"
	"math"
)

//C=5/9. (F-32)

func calcMedia(valores ...float64) float64 {
	total := 0.0

	for _, valor := range valores {
		//For in range (em Go) sendo o range os valores que serao recebidos na função
		total += valor
	}

	return total / float64(len(valores))
	//A função ira retornar os valores somados (total) e dividilos pelo len de valores (quantidade de valores que foram recebidos)
}

func paridade(numero int) string { //lembrar de mudar o tipo de retorno de acordo com que for precisar
	if numero%2 == 0 {
		return "Par"
	}
	return "Ímpar"
	//Similar a um programa em python nesse caso, se o resto da divisao por 2 for 0, numero é par senao, numero impar.
}

func potencia(base, expoente float64) float64 {
	return math.Pow(base, expoente) //É uma função que faz a potencia entre 2 numeros
}

func ctof(celsius float64) float64 {
	return (celsius * 9 / 5) + 32

}

func main() {
	fmt.Println(calcMedia(10.0, 20.0, 30.0)) // 20
	fmt.Println(calcMedia(5.0, 15.0))        // 10
	fmt.Println(calcMedia(0.0, 0.0))         //erro
	fmt.Println("-------------------------------------------")

	fmt.Println(paridade(128)) //par
	fmt.Println(paridade(51))  //impar
	fmt.Println(paridade(0))   //par
	fmt.Println("-------------------------------------------")

	fmt.Println(potencia(2, 2))  //4
	fmt.Println(potencia(10, 5)) //100000
	fmt.Println(potencia(4, -2)) //0,0625
	fmt.Println("-------------------------------------------")

	fmt.Println(ctof(35))  //95
	fmt.Println(ctof(180)) //356
	fmt.Println(ctof(-10)) //14
}
