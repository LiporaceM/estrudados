package main

import "fmt"

func hanoi(n int, origem, trabalho, destino string) {
	if n > 0 {
		hanoi(n-1, origem, destino, trabalho)

		fmt.Printf("Movendo o disco %d da origem %s para o destino %s\n", n, origem, destino)

		hanoi(n-1, trabalho, origem, destino)
	}
}
func findPair(arr []int, target int) (int, int) {
	seen := make(map[int]bool)
	for _, num := range arr {
		complement := target - num

		if seen[complement] {
			return complement, num
		}

		seen[num] = true
	}
	return -1, -1
}

func main() {
	numDisks := 3
	hanoi(numDisks, "A", "B", "C")

	fmt.Println(("----------------------------------"))

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	target := 5

	num1, num2 := findPair(array, target)

	if num1 == -1 && num2 == -1 {
		fmt.Println("(-1,-1)")
	} else {
		fmt.Printf("Par encontrado: (%d,%d)\n", num1, num2)
	}
}
