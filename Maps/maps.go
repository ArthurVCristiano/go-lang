package main

import "fmt"

func main() {

	fmt.Println("Maps")

	usuario := map[string]string{
		//dentro do [] tem o tipo da chave, e depois do tipo da chave tem o tipo do valor
		"nome" : "Arthur",
		"sobrenome" : "Velho",
		"curso" : "Engenharia de Software",
	}

	mapMap := map[string]map[string]string{
		"pessoa": {
			"nome" : "Arthur",
			"idade" : "21",
		},
		"curso": {
			"nome" : "Engenharia de Software",
		"campus" : "udesc",
		},
	}
	fmt.Println(usuario)
	fmt.Println(mapMap)
	delete(mapMap, "campus") // delete é uma função para deletar um valor do map, precisa passar o nome do map e a chave que quer deletar
	fmt.Println(mapMap)
}
