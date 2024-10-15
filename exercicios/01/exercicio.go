package main

import "fmt"

func main() {
	var myvar int = 0
	myvar = 10

	fmt.Println("Myvar", myvar)

	soma := 15 + 8
	fmt.Println("Primeira soma: ", soma)

	soma += 1
	fmt.Println("Soma mais soma e mais 1: ", soma)

	soma *= 3

	fmt.Println("Soma total * 3: ", soma)

	souninja := true

	fmt.Println("sou ninja booleano", souninja)

	var comidas [3]string = [3]string{"arroz", "feijão", "ovo"}

	fmt.Println(comidas[1])

	fmt.Println("verificar se a variável `soma' é igual a variável `myvar`: ", myvar == soma)

	fmt.Println("verificar se a variável `myvar` é menor ou igual à variável `soma`.: ", myvar <= soma)

	var divisao2 = divisao(10, 2)
	fmt.Println("Invoque a função divisão, passando os parâmetros 10 e 2.", divisao2)
}

// func myVar() int{
// 	var myvar int = 0
// 	myvar = 10

// 	return myvar
// }

func divisao(firstNumber int32, secondNumber int32) int32 {
	return firstNumber / secondNumber
}
