package main

import "fmt"

type Marketplace struct {
	Id          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var Products []Marketplace

func (m *Marketplace) Save() {
	Products = append(Products, *m)
}

func (m *Marketplace) GetAll() {
	fmt.Printf("Id: %d\n", m.Id)
	fmt.Printf("Producto: %s\n", m.Name)
	fmt.Printf("Precio: %.2f\n", m.Price)
	fmt.Printf("Descripción: %s\n", m.Description)
	fmt.Printf("Categoría: %s\n", m.Category)
}

func GetAll() {
	fmt.Println("Lista de Produtos:")
	for _, m := range Products {
		fmt.Printf("Id: %d\n", m.Id)
		fmt.Printf("Produto: %s\n", m.Name)
		fmt.Printf("Preço: %.2f\n", m.Price)
		fmt.Printf("Descrição: %s\n", m.Description)
		fmt.Printf("Categoria: %s\n", m.Category)
		fmt.Println("-------------")
	}
}

func GetByID(id int) *Marketplace {
	for _, p := range Products {
		if p.Id == id {
			return &p
		}
	}
	return nil
}

func main() {
	prod1 := Marketplace{Id: 1, Name: "Produto A", Price: 99.90, Description: "Descrição A", Category: "Categoria A"}
	prod2 := Marketplace{Id: 2, Name: "Produto B", Price: 150.00, Description: "Descrição B", Category: "Categoria B"}

	prod1.Save()
	prod2.Save()

	GetAll()

	p := GetByID(2)
	if p != nil {
		fmt.Printf("Encontrado: %s - %.2f\n", p.Name, p.Price)
	} else {
		fmt.Println("Produto não encontrado.")
	}
}
