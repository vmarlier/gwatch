package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func watch(args []string) {
	for {
		// get the command path
		_, err := exec.LookPath(args[0])
		if err != nil {
			log.Fatalf("Command %s not found\n", args[0])
		}

		// create the command
		cmd := exec.Command(args[0], args[1:]...)

		// read the return
		var out bytes.Buffer
		cmd.Stdout = &out

		// run the command
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("\r%s", out.String())

		// wait 2 seconds
		time.Sleep(2 + time.Second)

	}

}

func main() {
	watch(os.Args[1:])
}

/* TODO
1- Change the Time Interval
2- Highlighting the Difference Between Updates
3- Commands with Pipes
*/
