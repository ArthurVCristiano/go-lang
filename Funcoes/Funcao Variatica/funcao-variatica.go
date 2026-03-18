package main 

import "fmt"

// função vai resolver de 1 a N númeriso inteiros
func soma (numeros ... int)int{
	total := 0 
	for _, numero := range numeros{
		total += numero
	}
	return total
}

func escrever( texto string, numeros ... int){
	for _, numero := range numeros {
		fmt.Println(texto,numero)
	}
}

func main(){
	totalSoma := soma(1,2,3,4,5,50,60,62,78,100)
	fmt.Println(totalSoma)

	escrever("Escreve e adiciona os números: ", 1,2,3,4,5)
	//cria um slice

}