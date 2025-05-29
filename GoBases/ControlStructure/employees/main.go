package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	for i, v := range employees {

		if v > 21 {
			fmt.Printf("%s\n", i)
		}
	}

	employees["Frederico"] = 25

	fmt.Println(employees)

	delete(employees, "Pedro")

	fmt.Println(employees)

}
