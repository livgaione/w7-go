package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTax(t *testing.T) {

	salary := float64(155000)

	expected := float64(41850)

	result := tax(float64(salary))

	require.Equal(t, expected, result)

}
