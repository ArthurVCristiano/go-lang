package main
import (
	"fmt"
	//"time"

)



func main(){
	//i := 8
	/*
	for i < 10 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second) // código dorme por um segundo
	}

	fmt.Println(i)

	for j :=  0; j < 10; j+=3 {
		fmt.Println("Incrementando J", j )
		time.Sleep(time.Second)
	}
*/
	nomes := [3]string{"Arthur","Renato","Rafael"}

	for indice, nome := range nomes{
		// valor do indice - posição do vetor, nome que é o elemento do vetor  := range vetor
		fmt.Println(indice, nome)
		// primeiro valor sempre vai ser o indice, para deixar ele vazio e colocar só o nome, os parametros devem ser _, nome
	}

	for _, nome := range nomes{
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA"{
		//percorre letras da string
		fmt.Println(indice,string(letra))
		// para retornar as letras transformar para string
	}

	usuario := map[string]string{
		"nome" : "Arthur",
		"sobrenome" : "Velho",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
	//só funciona em map, não em struct


	for{
		fmt.Println("Executando infinitamente ..... ")
	}
}