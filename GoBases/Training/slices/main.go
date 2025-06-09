package main

import (
	"fmt"
)

func main() {

	var names []string

	names = append(names, "Livia", "Maria", "João", "Lucas", "Gabiela", "Daniela")

	names = append(names, "Julia")

	var newName []string

	for _, n := range names {
		if len(n) > 5 {
			newName = append(newName, n)
		}
	}

	fmt.Println(newName)
}
