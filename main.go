package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// get the command path
	_, err := exec.LookPath("kubectl")
	if err != nil {
		log.Fatal("Command not found")
	}

	// create the command
	cmd := exec.Command("kubectl", "get", "pods", "-n", "gitlab")

	// read the return
	var out bytes.Buffer
	cmd.Stdout = &out

	// run the command
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
	sleep(2)

	fmt.Print("\033[H\033[2J")
}
