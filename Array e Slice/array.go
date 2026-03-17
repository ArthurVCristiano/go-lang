package main

import "fmt"

func main(){

	var array [5]string

	array [0] = "Array1"

	fmt.Println(array)

	
	// jeito mais fácil
	array2 := [3]int {1,2,3}
	fmt.Println(array2)

	//numero de posições flexiveis
	array3 := [...]int {1,2,3,78,67}
	fmt.Println(array3)

	//array n é muito usado no go, por isso foi criado o slice, que é mais flexivel.

	slice23 := []int{1,2,3,4,5,6,7,98}
	fmt.Println(slice23)

	// fmt.Println(reflect.TypeOf(slice)) -> devolve o tipo da array - (TypeOf)

	slice23 = append(slice23, 18) // append adiciona um novo valor, mas como slice não possui posição direta, o append adiciona uma depois da última
	fmt.Println(slice23)

	slice2 := array3[4:5] // slice receve um pedaço do array
	fmt.Println(slice2)	

	nomesArray := [5]string{"Arthur","Sofia","Laura","Lena","Almir"}
	sliceCasal := nomesArray[0:2]
	fmt.Println(sliceCasal)

	// =================================ARRAYS INTERNOS =================================

	slice3 := make([]float32 10,15)


	fmt.Println(slice3)
	fmt.Println(len(slice3))  // lenght
	fmt.Println(cap(slice3))  // capacidade
}