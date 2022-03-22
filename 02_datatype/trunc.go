package main

import "fmt"

func main() {
	var floatInput float32

	fmt.Printf("Please, input a floating point number and press Enter. \n")
	_, err := fmt.Scan(&floatInput)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Truncated number is %d", int(floatInput))

}
