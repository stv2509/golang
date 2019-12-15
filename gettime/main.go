package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func ctime() {
	t := time.Now()
	fmt.Println("\nCurrent time:  ", t.Format("15:04:05"))

}

func ntptime() {
	t2, err := ntp.Time("0.pool.ntp.org")
	if err != nil {
		fmt.Println("\nNTP time:  ", err)
		return
	} else {
		fmt.Println("\nNTP time:  ", t2.Format("15:04:05"))
		return
	}
}
func main() {
	ctime()
	ntptime()
	fmt.Println("\n\n\nPress any key...")
	fmt.Scanf(" ")
	return
}
