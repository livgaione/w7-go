package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSalary(t *testing.T) {

	category := "A"
	hours := 40
	expected := 180000

	result := Salary(hours, category)

	require.Equal(t, expected, result)

}
