package main

import "fmt"

type usuario struct{
	nome string
	idade uint8

}

func (u usuario) salvar(){
	fmt.Printf("Salvando os dados do usuário %s no banco de dados \n", u.nome)
}

func (u usuario) maiorIdade()bool{
	return u.idade >=18
}

func (u *usuario) fazerAniversario(){
	u.idade++
}

func main(){	
	usuario1 := usuario{"Arthur Velho Cristiano", 21}
	usuario2 := usuario{ "Sofia Helena Bisson", 21}

	usuario1.salvar()
	usuario2.salvar()

	maiorIdade := usuario2.maiorIdade()
	fmt.Println(maiorIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}