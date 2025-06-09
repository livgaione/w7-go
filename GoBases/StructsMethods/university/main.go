package main

import "fmt"

type Students struct {
	Name     string
	LastName string
	Id       int
	Date     string
}

func (s *Students) Register(list []Students) []Students {

	return append(list, *s)

}

func PrintStudents(list []Students) {
	fmt.Println("Lista de Alunos:")

	for _, m := range list {

		fmt.Printf("Nome do aluno: %s\n", m.Name)
		fmt.Printf("Sobrenome do aluno: %s\n", m.LastName)
		fmt.Printf("ID do aluno: %d\n", m.Id)
		fmt.Printf("Data de admissão do aluno: %s\n", m.Date)
		fmt.Println("-------------")
	}
}

func main() {

	var listStudents []Students // slice vazio

	student1 := Students{"Lucy", "Loro", 2102, "19/05/2025"} // adiciono infos do aluno
	student2 := Students{"Ester", "Vitória", 2103, "19/05/2025"}

	listStudents = student1.Register(listStudents) // a função register retorna um slice de students populado
	listStudents = student2.Register(listStudents) // o retorno do register será populado na variável listStudents

	PrintStudents(listStudents)

}
