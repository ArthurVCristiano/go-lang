package main 
import "fmt"

func inverteSinal(numero int)int{
	return numero * -1
}

func inverteSinalComPonteiro(numero *int){
	// * indica q é ponteiro
	*numero = *numero * -1
}

func main(){
	numero := 20
	numeroInvertido := inverteSinal(numero)

	fmt.Println(numeroInvertido)
	// aqui passamos somente uma cópia da variavel numero

	novoNumero := 40
	inverteSinal(40)
	inverteSinalComPonteiro(&novoNumero)
	// & obrigatório para passar como parametro
	fmt.Println(novoNumero)
	//com ponteiro, a variavel muda seu valor no endereço de memória, otimização

}