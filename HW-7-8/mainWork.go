package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk(wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("%s is taking a walk\n", d.name)
	if d.walkDuration > 20 {
		return fmt.Errorf("%s is going for city\n", d.name)
	}

	time.Sleep(d.walkDuration * time.Second)
	fmt.Printf("%s is going home\n", d.name)

	return nil
}

var M = new(int)
var allH int
var mu sync.Mutex

func walkTheDogs(dogs []Dog, N int) {
	allW := len(dogs)

	if N < allW {
		fmt.Println("if slice > rootines:")
		j := 0
		n := N
		for j < allW {

			for _, d := range dogs[j:N] {
				fmt.Printf("\n%v : %v gorutiens\n", j, N)
				wg.Add(1)
				go func(d Dog, M *int) {
					err := d.Walk(&wg)
					if err != nil {
						mu.Lock()
						*M = *M + 1
						mu.Unlock()
						fmt.Print(err)
					}

				}(d, M)

			}
			wg.Wait()
			//ограничение количества потоков в N
			j = N
			if n+j < allW {
				N = n + j
			} else {
				N = allW
			}

		}
		mu.Lock()
		allH = allW - *M
		mu.Unlock()
		fmt.Printf("\n%v of %v dogs comeback home", allH, allW)

	} else {
		fmt.Println("if rootines > slice:")
		for _, d := range dogs {
			wg.Add(1)
			go func(d Dog, M *int) {
				err := d.Walk(&wg)
				if err != nil {
					*M = *M + 1
					fmt.Print(err)
				}

			}(d, M)

		}
		wg.Wait()
		allH := allW - *M
		fmt.Printf("\n%v of %v dogs comeback home", allH, allW)

	}
}

func main() {
	dogs := []Dog{
		{name: "vasya", walkDuration: 5}, {name: "john", walkDuration: 6}, {name: "max", walkDuration: 25},
		{name: "tuzik", walkDuration: 3}, {name: "belka", walkDuration: 4}, {name: "strelka", walkDuration: 35},
		{name: "vasya2", walkDuration: 5}, {name: "john2", walkDuration: 6}, {name: "max2", walkDuration: 25},
		{name: "tuzik2", walkDuration: 3}, {name: "belka2", walkDuration: 10}, {name: "strelka2", walkDuration: 35},
	}

	walkTheDogs(dogs, 5)

}
