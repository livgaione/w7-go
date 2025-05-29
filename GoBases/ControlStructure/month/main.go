package main

import "fmt"

func main() {

	var month int

	fmt.Printf("Entre com o mês desejado: ")
	fmt.Scan(&month)

	switch month {

	case 1:
		fmt.Printf("Janeiro\n")
	case 2:
		fmt.Printf("Fevereiro\n")
	case 3:
		fmt.Printf("Março\n")
	case 4:
		fmt.Printf("Abril\n")
	case 5:
		fmt.Printf("Maio\n")
	case 6:
		fmt.Printf("Junho\n")
	case 7:
		fmt.Printf("Julho\n")
	case 8:
		fmt.Printf("Agosto\n")
	case 9:
		fmt.Printf("Setembro\n")
	case 10:
		fmt.Printf("Outubro\n")
	case 11:
		fmt.Printf("Novembro\n")
	case 12:
		fmt.Printf("Dezembro\ns")
	}
}
