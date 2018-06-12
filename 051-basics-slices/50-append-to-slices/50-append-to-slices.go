package main

import "fmt"

func main() {
	cars := []string{"Ford Mustang", "Fiat 126p"}
	fmt.Println(cars)

	cars = append(cars, "BMW M5", "Tesla Roadster")
	fmt.Println(cars)

	motorbikes := []string{"Yamaha MT", "Hoda Goldwing"}

	allVehicles := append(cars, motorbikes...)
	fmt.Println(allVehicles)
}
