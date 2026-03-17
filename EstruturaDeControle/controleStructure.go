package main

import "fmt"

func main(){
	fmt.Println("Estruturas de controle")

	numero := 10

	if numero % 2 == 0{
		fmt.Printf("O número digitado ( %d ) é par\n", numero)
	} else {
		fmt.Printf("O número digitado ( %d ) é ímpar\n", numero)
	}

	if outroNumero := numero; outroNumero > 0 {
		// aqui a variavel outroNumero só existe dentro do if, fora dele não existe
		fmt.Printf("O número digitado ( %d ) é maior que 0\n", outroNumero)
	}

	// if aninhado só pdoe ser usado dentro do if, fora dele a variavel não existe

	fmt.Println("Switch")
	dia := diaDaSemana(3)
	fmt.Println(dia)
	dia2 := diaDaSemana(20)
	fmt.Println(dia2)

	dia3 := diaDaSemana2(100)
	fmt.Println(dia3)

	dia4 := diaDaSemana3(1)
	fmt.Println(dia4)
	diaInvalido4 := diaDaSemana3(100)
	fmt.Println(diaInvalido4)

}

func diaDaSemana(dia int) string{
	switch dia{
	case 1: return "Domingo"
	case 2: return "Segunda-feira"
	case 3: return "Terça-feira"
	case 4: return "Quarta-feira"
	case 5: return "Quinta-feira"
	case 6: return "Sexta-feira"
	case 7: return "Sabado"

	default: return "Dia inválido"
	}

}

func diaDaSemana2(dia int) string{
	switch {
	case dia == 100 : return "dia 100 do mês"
	case dia == 200 : return "dia 200 do mês"
	default: return "Dia inválido"
		
	}
}

func diaDaSemana3(dia int) string{
	var diaSemana string

	switch {
	case dia == 1 : diaSemana = "Domingo"
	fallthrough // o fallthrough faz com que o próximo case seja executado mesmo que a condição do próximo case seja falsa,
				//  ele só funciona dentro do switch, não pode ser usado fora dele
				//não é muito usado
	case dia == 2 : diaSemana = "Segunda-feira"
	case dia == 3 : diaSemana = "Terça-feira"
	case dia == 4 : diaSemana = "Quarta-feira"
	case dia == 5 : diaSemana = "Quinta-feira"
	case dia == 6 : diaSemana = "Sexta-feira"
	case dia == 7 : diaSemana = "Sabado"

	default: diaSemana = "Dia inválido"
	}	
	return diaSemana
}