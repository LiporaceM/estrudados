package main

import "fmt"

type Contato struct {
	Nome  string
	Email string
}

func (c *Contato) alterarEmail(novoEmail string) {
	c.Email = novoEmail
}
func adicionarContato(contato Contato, arrayContatos []Contato) {
	for i := range arrayContatos {
		if arrayContatos[i].Nome == "" {
			arrayContatos[i] = contato
			break
		}
	}
	fmt.Println(arrayContatos)
}
func excluirContato(arrayContatos []Contato) {
	for i := range arrayContatos {
		if arrayContatos[i].Nome != "" {
			arrayContatos[i].Nome = ""
			arrayContatos[i].Email = ""
		}
	}
	fmt.Println(arrayContatos)
}
func main() {
	var x = 10
	var nome, email string
	listaContatos := make([]Contato, 5)
	Contato1 := Contato{
		Nome:  nome,
		Email: email,
	}
	for x > 0 {
		fmt.Print("Insira o numero para realizar a função desejada:\n1. Adicionar Contato \n2. Remover Contato\n3. Parar programa\n")
		fmt.Scanln(&x)
		if x == 1 {
			fmt.Print("Digite o nome: ")
			fmt.Scanln(&nome)
			Contato1.Nome = nome
			fmt.Print("Digite o email: ")
			fmt.Scanln(&email)
			Contato1.Email = nome
			adicionarContato(Contato1, listaContatos)
			x = 10
		}
		if x == 2 {
			excluirContato(listaContatos)
			x = 10
		}
		if x == 3 {
			break
		}
	}

}
