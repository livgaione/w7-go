package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatisticsMinimum(t *testing.T) {

	op, _ := operation(minimum)

	result := op(2, 3, 3, 4, 10, 2, 4, 5)

	expected := float64(2)

	require.Equal(t, expected, result)

}

func TestStatisticsAverage(t *testing.T) {

	op, _ := operation(average)

	result := op(2, 3, 3, 4, 10, 2, 4, 5)

	expected := float64(4.125)

	require.Equal(t, expected, result)

}

func TestStatisticsMaximum(t *testing.T) {

	op, _ := operation(maximum)

	result := op(2, 3, 3, 4, 10, 2, 4, 5)

	expected := float64(10)

	require.Equal(t, expected, result)
}
