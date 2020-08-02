package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//wg.Add(2) // в группе две горутины
	work := func(s time.Duration, id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(s * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	// вызываем горутины
	for i := 1; i < 100; i++ {
		wg.Add(1)
		go work(1, i)
	}

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
