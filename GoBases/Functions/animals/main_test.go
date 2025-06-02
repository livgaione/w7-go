package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnimalCat(t *testing.T) {

	animal := "cat"

	quantity := 5

	expect := float64(25)

	result := Animals(animal)(quantity)

	require.Equal(t, expect, result)
}

func TestAnimalDog(t *testing.T) {

	animal := "dog"

	quantity := 5

	expect := float64(50)

	result := Animals(animal)(quantity)

	require.Equal(t, expect, result)
}

func TestAnimalTaran(t *testing.T) {

	animal := "taran"

	quantity := 2

	expect := float64(0.30)

	result := Animals(animal)(quantity)

	require.Equal(t, expect, result)
}

func TestAnimalHams(t *testing.T) {

	animal := "taran"

	quantity := 5

	expect := float64(0.75)

	result := Animals(animal)(quantity)

	require.Equal(t, expect, result)
}
