package main

import (
	"fmt"
	"time"
)

/*
Race condition is a problem that can occur in the concurrency execution of code that is interrelated.
For example functions that are executed concurrently and have access to the same variable or to the same output.
The outcome of the program depends on the interleaving and sometimes it can get one answer, sometimes gets another.
And interleaving in turn is non-deterministic and depends on the OS and runtime.
*/

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
