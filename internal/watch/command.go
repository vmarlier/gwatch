package watch

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// GetCommandOutput will take cmd (string) and args ([]string) and execute the command to return the output
func GetCommandOutput(cmd string, args []string) (string, error) {
	command := exec.Command(cmd, args...)

	var output bytes.Buffer
	command.Stdout = &output

	err := command.Run()
	if err != nil {
		fmt.Println("command run error")
		log.Fatalln(err)
	}

	return output.String(), nil
}

// ParseCommand will take a string to return a command with the args
func ParseCommand(str []string) (string, []string) {
	cmd := strings.Split(str[0], " ")
	return cmd[0], cmd[1:]
}
