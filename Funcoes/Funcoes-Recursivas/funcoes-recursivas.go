package main

import "fmt"

func fibonacci(posicao uint)uint{
	//uint - número positovs e somente eles
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao - 2 ) + fibonacci(posicao -1)
}

func main(){
	fmt.Println("Funções recursivas")

	posicao := uint(50)
	fmt.Println(fibonacci(posicao))
}

//funcao recursiva não é muito utilziada para performance, te muitas chamadas que acabam deixando o sistema lento