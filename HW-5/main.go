package main

import (
	event "HW-5/LogOtusEvent"
	"os"
)

func main() {
	submitted := &event.HwSubmitted{Id: 3456, Code: "My Code", Comment: "please take a look at my homework"}
	accepted := &event.HwAccepted{Id: 3456, Grade: 4}

	event.LogOtusEvent(submitted, os.Stdout)
	event.LogOtusEvent(accepted, os.Stdout)
}
