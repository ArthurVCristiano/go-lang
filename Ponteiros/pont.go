package main
import "fmt"

func main(){
	fmt.Println("Ponteiros")

	var v1 int = 10
	var v2 int = v1

	fmt.Println(v1, v2)

	v1++

	fmt.Println(v1, v2)

	// Ponteiro é uma referencia de memória

	var v3 int = 0
	var ponteiro *int

	v3 = 100
	ponteiro = &v3 // precisa do & para ela receber a variavel sem erro

	fmt.Println(v3, ponteiro) // sem ponteiro aparece o endereço de memoria

	fmt.Println(v3, *ponteiro) // com o * o valor aparece
}