package main

import (
	"log"
	"os"

	"gwatch/internal/watch"

	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Interval int64      `short:"i" long:"interval" description:"Interval in seconds, default 2s." required:"false"`
	Commands [][]string `short:"c" long:"commands" description:"2 or more commands." required:"false"`
}

func main() {
	options, err := flags.Parse(&opts)

	// If help message was printed, don't go further
	if e, ok := err.(*flags.Error); ok && e.Type == flags.ErrHelp {
		os.Exit(0)
	}

	interval := int64(2)
	if opts.Interval > 2 {
		interval = opts.Interval
	}

	cmd, args := watch.ParseCommand(options)
	for {
		output, _ := watch.GetCommandOutput(cmd, args)
		watch.DisplayOutput(output, interval)
	}
	watch.Clear()
}

func checkIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
