package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	sortSli := make([]int, 3)

	intSli := make([]int, len(sortSli))

	for i := 0; ; i++ {
		fmt.Println("Please, input integer or X for Exit and press Enter:")
		in := bufio.NewReader(os.Stdin)
		userStr, err := in.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		userStr = strings.TrimSpace(userStr)

		if userStr == "X" {
			fmt.Println("Quit")
			return
		} else {
			x, err := strconv.Atoi(userStr)

			if err != nil {
				fmt.Println(err)
				continue
			}
			//fmt.Println("i=", i)
			if i < len(intSli) {
				intSli[i] = x
				_ = copy(sortSli, intSli)
			} else {
				sortSli = append(sortSli, x)
			}

			sort.Ints(sortSli)
			//fmt.Println("intSli", intSli)
			fmt.Println("After sort", sortSli)

		}
	}

}
