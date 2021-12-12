package watch

import (
	"bytes"
	"os/exec"
	"strings"
)

// GetCommandOutput will take cmd (string) and args ([]string) and execute the command to return the output
func GetCommandOutputs(cmd []string, args [][]string, y int) (outputs []string, err error) {
	for i := 0; i < y; i++ {
		command := exec.Command(cmd[i], args[i]...)

		var output bytes.Buffer
		command.Stdout = &output

		err = command.Run()
		if err != nil {
			return []string{}, err
		}

		outputs = append(outputs, output.String())
	}

	return outputs, nil
}

// ParseCommand will take a string to return a command with the args
func ParseCommand(str []string, y int) (cmd []string, args [][]string) {
	for i := 0; i < y; i++ {
		pass := strings.Split(str[i], " ")
		cmd = append(cmd, pass[0])
		args = append(args, pass[1:])
	}

	return cmd, args
}
