package main

import (
	event "HW-5/otusevent"
	"os"
)

func main() {
	submitted := &event.HwSubmitted{ID: 3456, Code: "My Code", Comment: "please take a look at my homework"}
	accepted := &event.HwAccepted{ID: 3456, Grade: 4}

	event.LogOtusEvent(submitted, os.Stdout)
	event.LogOtusEvent(accepted, os.Stdout)
}
