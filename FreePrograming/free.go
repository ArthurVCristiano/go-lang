package main

import "fmt"

private var cpfLogin = login()
private	var password = password()

func main() {
	fmt.Println("Estrutura de Controle")
	
	vanguardLogin(cpfLogin)
	vanguardPassword(password)
	loginAndPassword(cpfLogin, password)
}

func login() string {

	fmt.Println("Digite seu CPF:")
	var cpfReplied string
	fmt.Scanln(&cpfReplied)


	return cpf
}

func password() string {
	fmt.Println("Digite sua senha:")
	var passwordReplied string
	fmt.Scanln(&passwordReplied)

	return passwordReplied
}

func vanguardLogin() bool{
	var cpf = this.cpfLogin()
	if cpf != "" {
		return true
	} else {
		return false
	}
}

func vanguardPassword() bool{
	var password = this.password()

	if password != "" {
		return true
	} else {
		return false
	}
}

func loginAndPassword(cpfFunc string, passwordFunc string) string {

	cpfFunc = login()	
	passwordFunc = password()

	switch{
	case vanguardLogin() && vanguardPassword(): 
		cpfFunc = login() && passwordFunc = password() : return "Login e senha válidos"

	case vanguardLogin() && !vanguardPassword():
		cpfFunc = login() && passwordFunc = password() : return "Login válido, senha inválida"

	case !vanguardLogin() && vanguardPassword():
		cpfFunc = login() && passwordFunc = password() : return "Login inválido, senha válida"

	default: return "Login e senha inválidos"
	}
}

