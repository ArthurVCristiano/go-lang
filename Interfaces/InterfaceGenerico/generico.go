package main

import "fmt"

func generica(interf interface{}){
	fmt.Println(interf)
}

func main(){
	generica("String")
	generica(0)
	generica(true)
}