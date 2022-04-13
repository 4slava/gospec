package main

import (
	"fmt"
	"time"
)

func PrintABC() {
	for abc := 'a'; abc < 'a'+25; abc++ {
		fmt.Printf(" %c ", abc)
		time.Sleep(9 * time.Millisecond)
	}
}

func PrintInt() {
	for i := 0; i < 25; i++ {
		fmt.Print(i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	//	runtime.GOMAXPROCS(2)
	go PrintABC()
	go PrintInt()
	fmt.Println("\nGo Routines end: ")
	time.Sleep(2 * time.Second)

}
