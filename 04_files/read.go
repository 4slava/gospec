package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	var fileName string
	p := Person{}
	slicePersons := []Person{}

	fmt.Print("Enter a file name: ")
	_, err := fmt.Scanln(&fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	for {

		// Read line from file
		bline, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return
		}

		// Use scanner for reading string line and split words
		scanner := bufio.NewScanner(strings.NewReader(string(bline)))
		scanner.Split(bufio.ScanWords)

		//reading and adding to struct first word - fname
		scanner.Scan()
		p.fname = scanner.Text()

		//reading and adding a second word - lname
		scanner.Scan()
		p.lname = scanner.Text()

		//append struct to slice of structs
		slicePersons = append(slicePersons, p)

	}

	i := 0
	for range slicePersons {
		fmt.Printf("%s %s \n", slicePersons[i].fname, slicePersons[i].lname)
		i++
	}

}
