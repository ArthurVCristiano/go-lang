package main

import "fmt"

func init(){
	fmt.Println("Executando função init - sempre executada primeiro independente de tudo")
	//pode ser uxada para setar algum valor e coisas do tipo
}
func main(){
	fmt.Println("Funcao main sendo executada")
}