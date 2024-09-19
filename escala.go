package main

import "fmt"

const kelvinT = 373

func main() {

	celsius := kelvinT - 273
	fmt.Printf("A conversão da temperatura de da água de Kelvin (%dK) para °C é de: %d°C", kelvinT, celsius)

}
