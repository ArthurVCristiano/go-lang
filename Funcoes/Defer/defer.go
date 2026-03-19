package main

import "fmt"

func escreverNaTela(){
	fmt.Println("Escrevendo função 1")
}

func escreverNaTela2(){
	fmt.Println("Escrevendo função 2")
}

func alunoAprovado (n1,n2 float32)bool{
	defer fmt.Println("Média calculada. Resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovao")
	// serve para adiar o codigo de uma função
	media := (n1 + n2)/2
	
	if media >= 6 {
		return true
	}

	return false
}

func main(){
	defer escreverNaTela()
	escreverNaTela2()

	fmt.Println(alunoAprovado(7,9))

	// DEFER -> "Adiar"então ele vai adiar a função1 e executar a segunda função
}