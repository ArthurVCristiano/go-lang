package main 

import "fmt"

func calculoMatematico(n1,n2 int) (soma int, subtracao int){
	soma = n1 + n2
	subtracao = n1 - n2

	return
} 

func main(){
    soma, subtracao := calculoMatematico(20,10)
	fmt.Println(soma, subtracao)
}