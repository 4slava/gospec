/*
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func InString() (string, error) {

	inBuff := bufio.NewReader(os.Stdin)
	inStr, err := inBuff.ReadString('\n')
	if err != nil {
		return "", err
	}
	inStr = strings.TrimSpace(inStr)
	return inStr, err
}

func StringToDigitSli(inStr string) []int {
	digitSli := make([]int, 0)
	strSli := strings.Split(inStr, " ")
	for i := 0; i < len(strSli); i++ {
		terminalInt, err := strconv.Atoi(strSli[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		digitSli = append(digitSli, terminalInt)

	}
	return digitSli
}

func sortSli(text string, sli []int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s unsorted: %v\n", text, sli)
	sort.Ints(sli)
	fmt.Printf("%s sorted: %v\n", text, sli)

}

func main() {

	fmt.Println("Please, input integers with space-separator and press Enter for sorting:")

	inString, err := InString()
	if err != nil {
		fmt.Println(err, " Try again")
		return
	}

	sortingSli := StringToDigitSli(inString)
	if len(sortingSli) > 3 {
		var wg sync.WaitGroup
		partLen := int(math.Round(float64(len(sortingSli)) / 4))

		wg.Add(1)
		go sortSli("Part 1", sortingSli[0:partLen], &wg)
		wg.Add(1)
		go sortSli("Part 2", sortingSli[partLen:2*partLen], &wg)
		wg.Add(1)
		go sortSli("Part 3", sortingSli[2*partLen:3*partLen], &wg)
		wg.Add(1)
		go sortSli("Part 4", sortingSli[3*partLen:int(len(sortingSli))], &wg)

		wg.Wait()
		fmt.Println("Entire sorted list: ", sortingSli)
		return
	}
	sort.Ints(sortingSli)
	fmt.Println("Entire sorted list: ", sortingSli)
}
*/