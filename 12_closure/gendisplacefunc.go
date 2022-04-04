package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, s float64) func(time float64) float64 {
	return func(time float64) float64 {
		return 0.5*a*float64(math.Pow(time, 2)) + (v * time) + s
	}
}

func UserInput(text string) float64 {
	var in float64
	fmt.Println(text)
	_, err := fmt.Scanln(&in)
	if err != nil {
		fmt.Println(err, " Try again")
		UserInput(text)
	}
	return in
}

func main() {
	a := UserInput("Enter the acceleration, m/s^2:")
	v := UserInput("Enter the initial velocity, m/s:")
	s := UserInput("Enter the initial displacement, m:")
	fn := GenDisplaceFn(a, v, s)
	time := UserInput("Enter a time, s:")

	fmt.Printf("Displacement after %vs = %vm", time, fn(time))
}
