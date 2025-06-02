package main

import (
	"errors"
	"fmt"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFunc, _ := operation(minimum)
	avgFunc, _ := operation(average)
	maxFunc, _ := operation(maximum)

	minVal := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	avgVal := avgFunc(2, 3, 3, 4, 1, 2, 4, 5)
	maxVal := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Mínimo:", minVal)
	fmt.Println("Média:", avgVal)
	fmt.Println("Máximo:", maxVal)
}

func operation(op string) (func(...int) float64, error) {
	switch op {
	case minimum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			min := nums[0]
			for _, n := range nums {
				if n < min {
					min = n
				}
			}
			return float64(min)
		}, nil

	case average:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			sum := 0
			for _, n := range nums {
				sum += n
			}
			return float64(sum) / float64(len(nums))
		}, nil

	case maximum:
		return func(nums ...int) float64 {
			if len(nums) == 0 {
				return 0
			}
			max := nums[0]
			for _, n := range nums {
				if n > max {
					max = n
				}
			}
			return float64(max)
		}, nil

	default:
		return nil, errors.New("operação inválida")
	}
}
