package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {

	grade := 2

	sum := 10

	expected := 10

	result := Average(grade, sum)

	require.Equal(t, expected, result)

}
