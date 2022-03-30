package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	nameMap := make(map[string]string)
	var nameStr string
	var addrStr string
	var err error
	in := bufio.NewReader(os.Stdin)

	//	for nameStr != "X" || addrStr != "X" {

	fmt.Println("Enter name")

	nameStr, err = in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Enter adress")

	addrStr, err = in.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	nameStr = strings.TrimSpace(nameStr)
	addrStr = strings.TrimSpace(addrStr)
	nameMap[nameStr] = addrStr
	//	}

	barr, err := json.Marshal(nameMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("JSON object: \n %s", barr)

}
