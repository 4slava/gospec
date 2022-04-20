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
}

func mergeSortSli(unsortSli, sli1, sli2, sli3, sli4 []int) []int {
	i1 := 0
	i2 := 0
	i3 := 0
	i4 := 0
	sortedSli := make([]int, 0)
	for range unsortSli {
		switch {
		case i1 < len(sli1) && sli1[i1] <= sli2[i2] && sli1[i1] <= sli3[i3] && sli1[i1] <= sli4[i4]:

			sortedSli = append(sortedSli, sli1[i1])
			if i1 == (len(sli1) - 1) {
				sli1[i1] = math.MaxInt
				continue
			}
			i1++
			continue

		case i2 < len(sli2) && sli2[i2] <= sli1[i1] && sli2[i2] <= sli3[i3] && sli2[i2] <= sli4[i4]:

			sortedSli = append(sortedSli, sli2[i2])
			if i2 == (len(sli2) - 1) {
				sli2[i2] = math.MaxInt
				continue
			}
			i2++
			continue

		case i3 < len(sli3) && sli3[i3] <= sli1[i1] && sli3[i3] <= sli2[i2] && sli3[i3] <= sli4[i4]:

			sortedSli = append(sortedSli, sli3[i3])
			if i3 == (len(sli3) - 1) {
				sli3[i3] = math.MaxInt
				continue
			}
			i3++
			continue

		case i4 < len(sli4) && sli4[i4] <= sli1[i1] && sli4[i4] <= sli2[i2] && sli4[i4] <= sli3[i3]:

			sortedSli = append(sortedSli, sli4[i4])
			if i4 == (len(sli4) - 1) {
				sli4[i4] = math.MaxInt
				continue
			}
			i4++
			continue

		default:
			fmt.Println("Something is wrong with sorting and merging")
			sortedSli = nil
			break

		}

	}
	return sortedSli
}

func main() {

	fmt.Println("Please, input integers with space-separator and press Enter for sorting:")

	inString, err := InString()
	if err != nil {
		fmt.Println(err, " Try again")
		return
	}

	unsortSli := StringToDigitSli(inString)

	if len(unsortSli) < 8 {
		sort.Ints(unsortSli)
		fmt.Println("Entire sorted list: ", unsortSli)
		return
	}

	var wg sync.WaitGroup
	partLen := int(math.Round(float64(len(unsortSli)) / 4))
	sli1 := unsortSli[0:partLen]
	sli2 := unsortSli[partLen : 2*partLen]
	sli3 := unsortSli[2*partLen : 3*partLen]
	sli4 := unsortSli[3*partLen : len(unsortSli)]
	fmt.Println(unsortSli)
	wg.Add(4)
	go sortSli("Part 1", sli1, &wg)
	go sortSli("Part 2", sli2, &wg)
	go sortSli("Part 3", sli3, &wg)
	go sortSli("Part 4", sli4, &wg)
	wg.Wait()
	fmt.Printf("Part 1 sorted: %v\n", sli1)
	fmt.Printf("Part 2 sorted: %v\n", sli2)
	fmt.Printf("Part 3 sorted: %v\n", sli3)
	fmt.Printf("Part 4 sorted: %v\n", sli4)

	sortedSli := mergeSortSli(unsortSli, sli1, sli2, sli3, sli4)

	fmt.Println("Entire sorted list: ", sortedSli)

}
