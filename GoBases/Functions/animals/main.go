package main

import "fmt"

func main() {

	var animal string

	var quantity int

	fmt.Printf("Insira o animal: ")

	fmt.Scanf("%s", &animal)

	fmt.Printf("Insira a quantidade do animal que seja adicionar: ")
	fmt.Scanln(&quantity)

	total := Animals(animal)(quantity)

	fmt.Printf("Total de comida para %d %s(s): %.2f kg\n", quantity, animal, total)

}

func Animals(animal string) func(int) float64 {
	switch animal {
	case "cat":
		return func(quantity int) float64 {
			return float64(quantity) * 5
		}
	case "dog":
		return func(quantity int) float64 {
			return float64(quantity) * 10
		}
	case "hams":
		return func(quantity int) float64 {
			return float64(quantity) * 0.25
		}
	case "taran":
		return func(quantity int) float64 {
			return float64(quantity) * 0.15
		}
	default:
		fmt.Println("Animal não encontrado")
	}
	return nil
}
