package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TerminalInputToStr() (string, error) {

	terminalBuff := bufio.NewReader(os.Stdin)
	terminalStr, err := terminalBuff.ReadString('\n')
	if err != nil {
		return "", err
	}

	terminalStr = strings.TrimSpace(terminalStr)
	return terminalStr, nil
}

func StrToIntSli(terminalStr string) []int {
	splitStrSli := make([]int, 0)
	terminalStrSli := strings.Split(terminalStr, " ")
	for i := 0; i < len(terminalStrSli); i++ {
		terminalInt, err := strconv.Atoi(terminalStrSli[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		splitStrSli = append(splitStrSli, terminalInt)

	}
	return splitStrSli
}

func Swap(slix []int, i int) {
	temp := slix[i]
	slix[i] = slix[i+1]
	slix[i+1] = temp
}

func BubbleSort(sli []int) {
	swaped := false
	for i := 0; i < len(sli)-1; i++ {
		if sli[i] > sli[i+1] {
			Swap(sli, i)
			swaped = true
		}
	}
	if swaped {
		BubbleSort(sli)
	}

}

func main() {
	fmt.Println("Please, input integers with space-separator and press Enter for Bubble Sorting:")
	terminalStr, err := TerminalInputToStr()
	if err != nil {
		fmt.Println(err, " Try again")
		return
	}

	userIntSli := StrToIntSli(terminalStr)

	BubbleSort(userIntSli)

	fmt.Println(userIntSli)

}
