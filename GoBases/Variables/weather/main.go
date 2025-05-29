package main

import "fmt"

func main() {

	var (
		temperature int
		humidity    int
		pressure    int
	)

	temperature = 25
	humidity = 70
	pressure = 1013

	fmt.Println("The current weather is", temperature, "degrees Celsius,", humidity, "percent humidity, and", pressure, "hPa pressure.")

}
