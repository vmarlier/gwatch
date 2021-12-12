package main

import (
	"log"
	"os"

	"gwatch/internal/watch"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Interval int64    `short:"i" long:"interval" description:"Interval in seconds, default 2s." required:"false"`
	Commands []string `short:"c" long:"commands" description:"2 or more commands." required:"false"`
}

func main() {
	options, err := flags.Parse(&opts)

	// If help message was printed, don't go further
	if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
		os.Exit(0)
	}

	// Interval is 2 seconds minimum
	interval := int64(2)
	if opts.Interval > 2 {
		interval = opts.Interval
	}

	// Check which variable is containing the commands and how many of them there is to execute
	commands, nCommands := checkOptions(options, opts.Commands)
	if nCommands > 4 {
		log.Fatalln("gwatch cannot handle more than 4 commands at once.")
	}

	// Parse the command(s) and separate main command from args/options
	// $ gwatch "kubectl get pods" -> [kubectl] + [[get pods]]
	// $ gwatch "ps ax" "kubectl get pods" -> [ps kubectl] + [[ax] [get pods]]
	cmd, args := watch.ParseCommand(commands, nCommands)

	for {
		outputs, _ := watch.GetCommandOutputs(cmd, args, nCommands)
		watch.DisplayOutput(outputs, interval, nCommands)
	}
	watch.Clear()
}

// checkOptions will check that options or &opts.Commands has the commands to be executed
func checkOptions(options []string, commands []string) ([]string, int) {
	if len(options) > 0 {
		return options, len(options)
	}

	return commands, len(commands)
}

func checkIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
