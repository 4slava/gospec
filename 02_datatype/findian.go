package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Please, write something word and press Enter.")

	//	var userStr string
	//	_, err := fmt.Scanln(&userStr) //Scan with spase-separator - read only one word

	in := bufio.NewReader(os.Stdin)
	userStr, err := in.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return
	}

	userStr = strings.ReplaceAll(userStr, "\n", "") //delete CRLF \n
	userStr = strings.ReplaceAll(userStr, "\r", "") //delete CRLF \r

	//userStr = strings.TrimSpace(userStr) //better then delete CRLF

	lowerStr := strings.ToLower(userStr)

	if strings.HasPrefix(lowerStr, "i") && strings.HasSuffix(lowerStr, "n") && strings.Contains(lowerStr, "a") {
		fmt.Print("Found!")
		return
	}

	fmt.Print("Not Found!")

}
