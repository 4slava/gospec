package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func makearray(s string) []int {
	ss := strings.Split(s, ",")
	var cs = []int{}
	sum := 0
	for sum < len(ss) {
		c, err := strconv.Atoi(ss[sum])
		if err != nil {
			log.Fatal(err)
		}
		cs = append(cs, c)
		sum += 1
	}
	return cs
}

func sorting(cs []int, c chan []int, wg sync.WaitGroup) {
	sort.Ints(cs)
	c <- cs
}

func slice(cs []int) ([]int, []int, []int, []int) {
	var cs1 = []int{}
	var cs2 = []int{}
	var cs3 = []int{}
	var cs4 = []int{}
	i := 0
	sel := 1
	for i < len(cs) {
		if sel == 1 {
			cs1 = append(cs1, cs[i])
		}
		if sel == 2 {
			cs2 = append(cs2, cs[i])
		}
		if sel == 3 {
			cs3 = append(cs3, cs[i])
		}
		if sel == 4 {
			cs4 = append(cs4, cs[i])
			sel = 0
		}
		sel += 1
		i += 1
	}
	fmt.Println("SubArray 1: ", cs1)
	fmt.Println("SubArray 2: ", cs2)
	fmt.Println("SubArray 3: ", cs3)
	fmt.Println("SubArray 4: ", cs4)
	return cs1, cs2, cs3, cs4
}

func main() {
	var wg sync.WaitGroup
	c := make(chan []int, 4)
	fmt.Println("Enter intager array devided by comas: ")
	var arr string
	fmt.Scanln(&arr)
	cs := makearray(arr)
	fmt.Println(cs)
	cs1, cs2, cs3, cs4 := slice(cs)
	go sorting(cs1, c, wg)
	go sorting(cs2, c, wg)
	go sorting(cs3, c, wg)
	go sorting(cs4, c, wg)
	c1 := <-c
	c2 := <-c
	c3 := <-c
	c4 := <-c
	nc1 := append(c1, c2...)
	nc2 := append(c3, c4...)
	nc := append(nc1, nc2...)
	fmt.Println("Sorted array: ", nc)

}
