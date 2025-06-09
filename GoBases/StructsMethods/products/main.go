package main

import (
	"fmt"
	"strings"
)

type Product interface {
	Price() float64
}

type SmallProduct struct {
	price float64
}

func (p SmallProduct) Price() float64 {
	return p.price
}

type MediumProduct struct {
	price float64
}

func (p MediumProduct) Price() float64 {
	return p.price + p.price*0.03 + p.price*0.03
}

type LargeProduct struct {
	price float64
}

func (p LargeProduct) Price() float64 {
	return p.price + p.price*0.06 + 2500
}

func NewProduct(t string, price float64) Product {
	switch strings.ToLower(t) {
	case "pequeno":
		return SmallProduct{price}
	case "médio", "medio":
		return MediumProduct{price}
	case "grande":
		return LargeProduct{price}
	default:
		return nil
	}
}

func main() {
	produtos := []Product{
		NewProduct("pequeno", 1000),
		NewProduct("médio", 1000),
		NewProduct("grande", 1000),
	}

	for _, p := range produtos {
		fmt.Printf("Preço final: %.2f\n", p.Price())
	}
}
