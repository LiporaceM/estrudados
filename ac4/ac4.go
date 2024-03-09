package main

import "fmt"

func hanoi(n int, origem, trabalho, destino string) {
	if n > 0 {
		hanoi(n-1, origem, destino, trabalho)

		fmt.Printf("Movendo o disco %d da origem %s para o destino %s\n", n, origem, destino)

		hanoi(n-1, trabalho, origem, destino)
	}
}
func findPair(arr []int, target int) (int, int, bool) {
	left, right := 0, len(arr)-1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			// Encontrou um par que soma para o alvo
			return arr[left], arr[right], true
		} else if sum < target {
			// Se a soma for menor que o alvo, mova o ponteiro esquerdo para a direita
			left++
		} else {
			// Se a soma for maior que o alvo, mova o ponteiro direito para a esquerda
			right--
		}
	}

	// Nenhum par encontrado, retorna (-1, -1) e false
	return -1, -1, false
}

func main() {
	numDisks := 3
	hanoi(numDisks, "A", "B", "C")

	fmt.Println(("----------------------------------"))

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 10

	num1, num2, found := findPair(array, target)

	if found {
		fmt.Printf("Par encontrado: (%d,%d)\n", num1, num2)
	} else {
		fmt.Println("(-1-1)")
	}
}
