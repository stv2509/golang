package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//ReadDir(dir string) (map[string]string, error)
//RunCmd(cmd []string, env map[string]string) int
//Command(name string, arg ...string)

func RunCmd(cmd []string, env map[string]string) int {

	for k, v := range env {
		fmt.Printf("key: %s, value: %s,\n", k, v)
	}

	cmdRun := exec.Command(cmd[0], cmd[1:]...)

	for k, v := range env {
		os.Setenv(k, v)
	}

	cmdOut, err := cmdRun.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return 1
	}
	fmt.Println(string(cmdOut))
	return 0
}
