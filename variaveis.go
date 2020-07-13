package main

import ("fmt"
		"reflect"
)

func main (){
	//Declarando variaveis (var <nome> <tipo> = <valor>)
	var nome string = "Miguel"
	//O Go consegue entender qual o tipo da variável sem precisar declarar
	var idade = 33
	// Podemos usar o decalrador de variáveis curtas para encurtar mais ainda nosso código(:=)
	versao := 1.1

	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Esse programa está na versão", versao)

	//imprimindo o tipo da variável (type() do python)
	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))

	var comando int
	// Lendo uma entrada do usuário.
	// Scanf(<tipo da variável>, <endereço da variáver (por isso o "&")>)
	fmt.Scanf("%d", &comando)

	//Podemos usar a função Scan()
	//Ela entende qual o tipo da variável e, portanto, não precisa colocar o "%d"
	fmt.Scan(&comando)

	// Essa é a forma de usar o if em GO
	// A condição sempre deve retornar um valor booleano
	if comando == 1 {
		fmt.Println("Monitorando...")
	}else if comando == 2{
		fmt.Println("Exibindo logs...")
	}else if comando == 0{
		fmt.Println("Saindo do programa")
	}else {
		fmt.Println("Não conheço este comando")
	}

	// Usando switch (in case of do pascal)
	switch comando{
	case 1: 
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo do programa")
	default:
		fmt.Println("Não conheço este comando")
	}
}