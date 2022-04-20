package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(index, times int, wg *sync.WaitGroup, permit, finish chan int) {
	defer wg.Done()

	var wg2 sync.WaitGroup

	for i := 0; i < times; i++ {
		c := <-permit
		wg2.Add(2)
		go func() {
			p.leftCS.Lock()
			//fmt.Println("lock leftCS ", index)
			wg2.Done()
		}()

		go func() {
			p.rightCS.Lock()
			//fmt.Println("lock rightCS ", index)
			wg2.Done()
		}()
		wg2.Wait()

		fmt.Println("starting to eat ", index)
		time.Sleep(10 * time.Millisecond)

		wg2.Add(2)
		go func() {
			p.leftCS.Unlock()
			//fmt.Println("unlock leftCS ", index)
			wg2.Done()
		}()

		go func() {
			p.rightCS.Unlock()
			//fmt.Println("unlock rightCS ", index)
			wg2.Done()
		}()
		wg2.Wait()

		fmt.Println("finishing eating ", index)

		finish <- c

	}
}
func Host(n int, wg *sync.WaitGroup, permit, finish chan int) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		c := <-finish
		permit <- c

	}
}

func main() {
	people := 5
	times := 3
	concurrency := 2

	CSticks := make([]*ChopS, people)
	for i := 0; i < people; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, people)
	for i := 0; i < people; i++ {
		philos[i] = &Philo{
			CSticks[i], CSticks[(i+1)%people],
		}
	}

	var wg sync.WaitGroup
	permit := make(chan int, concurrency)
	finish := make(chan int, 0)
	for i := 0; i < concurrency; i++ {
		permit <- i
	}

	wg.Add(1)
	go Host(times*people, &wg, permit, finish)

	for i := 0; i < people; i++ {
		wg.Add(1)
		go philos[i].eat(i, times, &wg, permit, finish)

	}

	wg.Wait()

}
