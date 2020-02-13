package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func ctime() {

	t := time.Now()

	fmt.Println("\nCurrent time:  ", t.Format("15:04:05"))

}

func ntptime() {

	t2, err := ntp.Time("0.pool.ntp.org")

	if err != nil {
		log.Printf("\n\nError NTP message:  %s", err)
		return

	} else {
		fmt.Println("\nNTP time:  ", t2.Format("15:04:05"))
		return
	}
}

func main() {

	ctime()

	ntptime()

	fmt.Println("\n\nPress any key...")
	fmt.Scanf(" ")

	return
}
