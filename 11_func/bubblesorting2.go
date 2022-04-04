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

func Swap(slix []int) {
	temp := slix[0]
	slix[0] = slix[1]
	slix[1] = temp
}

func BubbleSort(sli []int) {
	swaped := false
	for i := 1; i < len(sli); i++ {
		if sli[i-1] > sli[i] {
			Swap(sli[i-1 : i+1])
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
