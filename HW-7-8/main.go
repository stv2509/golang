package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
	// "github.com/mattn/go-colorable"
)


func () error


func Run(task []func() error, N int, M int) error {
	wg.Add(len(N))
	return
}

func main() {

}


func fetchAll() error {
	var N = 50
	quit := make(chan bool)
	errc := make(chan error)
	done := make(chan error)
	for i := 0; i < N; i++ {
		go func(i int) {
			// dummy fetch
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			err := error(nil)
			if rand.Intn(4) == 0 {
				err = fmt.Errorf("goroutine %d's error returned", i)
			}
			ch := done // we'll send to done if nil error and to errc otherwise
			if err != nil {
				ch = errc
			}
			select {
			case ch <- err:
				return
			case <-quit:
				return
			}
		}(i)
	}
	count := 0
	for {
		select {
		case err := <-errc:
			close(quit)
			return err
		case <-done:
			count++
			if count == N {
				return nil // got all N signals, so there was no error
			}
		}
	}
}
