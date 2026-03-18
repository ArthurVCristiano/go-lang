//FUNCAO Q N TEM NOME, INVOCA DE JEITO DIFERENTE
package main

import "fmt"

func main(){
	func (texto string){
		fmt.Println(texto)
	}("passando parametro para variavel texto") // declara ela e já executa

		/* retorno :=  func(texto string) string {
			return fmt.Sprintf("Recebido -> %s", texto) // Sprintf funciona como um titulo
		}("passando Parametro")
		fmt.Println(retorno)

		*/
}