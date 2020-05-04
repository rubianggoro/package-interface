package main

import (
	"fmt"

	"github.com/rubianggoro/package-interface/motor"
	"github.com/rubianggoro/package-interface/sepeda"
)

func main() {

	fmt.Println("---------------MOTOR------------------")
	motor1 := motor.CreateMotor("Merah", 12, 12)

	fmt.Println(motor1)
	fmt.Println("Lebih cepat:", motor1.Cepat(12), "km/h")
	fmt.Println("Lebih lambat:", motor1.Lambat(12), "km/h")

	fmt.Println("---------------SEPEDA------------------")
	sepeda1 := sepeda.CreateSepeda("BMX", 2)

	fmt.Println(sepeda1)
	fmt.Println("Lebih cepat:", sepeda1.Cepat(5), "km/h")
	fmt.Println("Lebih lambat:", sepeda1.Lambat(5), "km/h")
}
