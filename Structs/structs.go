package main 

import "fmt"

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct {
	endereco string
	numero int
}

func main(){

	var u usuario
	u.nome = "Arthur"
	u.idade = 21
	fmt.Println(u)

	enderecoUsuario := endereco{"Servidao Joao Pedro Cristiano", 160}

	usuario2 := usuario{"Arthur2", 23, enderecoUsuario}
	fmt.Println(usuario2)
}