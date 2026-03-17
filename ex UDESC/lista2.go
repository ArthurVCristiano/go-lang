package main

import "fmt"

func main(){

	var nota1, nota2, nota3, media float64
	var nome string

	fmt.Println("Digite o nome do aluno")
	 fmt.Scan(&nome)

	fmt.Println("Digite a primeira nota")
	 fmt.Scan(&nota1)

	fmt.Println("Digite a segunda nota")
	 fmt.Scan(&nota2)

	fmt.Println("Digite a terceira nota")
	 fmt.Scan(&nota3)

	 media = ((nota1 + nota2 + nota3) / 3)

	fmt.Println("O aluno %s teve a média de:%f ", nome, media )

}