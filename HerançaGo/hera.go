package main 

import "fmt"

type pessoa struct {
	nome string 
	sobrenome string
	idade int
	altura int
}

type estudante struct{
	pessoa // aqui está passando todos os campos de pessoa e jogando dentro de estudante - herança no Go
	matricula int
	curso string
}

func main(){

	estudante1 := estudante{
		pessoa: pessoa{
			nome:      "Arthur",
			sobrenome: "Velho",
			idade:     21,
			altura:    181,
		},
		matricula: 151123,
		curso:     "Engenharia",
	}
	fmt.Println(estudante1)

	p1 := pessoa{"Thiaguinho", "Pagodeiro", 35, 167}
	fmt.Println(p1)

	e1 := estudante{p1, 1423, "Eng. de Software"}
	fmt.Println(e1.nome)
	
}