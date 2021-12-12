package watch

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

// DisplayOutput will take the a command output as a string and display it
func DisplayOutput(output string, second int64) {
	tm.Clear()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(96|tm.PCT, 100|tm.PCT, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, output)

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 2|tm.PCT, 2|tm.PCT))

	tm.Flush()
	time.Sleep(time.Second * time.Duration(second))
}

// DisplayOutput will take several commands outputs as strings and display them
func DisplayMultiplesOutput(output string, second int64) {
	tm.Clear()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(48|tm.PCT, 100|tm.PCT, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, output)

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 2|tm.PCT, 2|tm.PCT))

	// Create Box with 30% width of current screen, and height of 20 lines
	box2 := tm.NewBox(48|tm.PCT, 100|tm.PCT, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box2, output)

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box2.String(), 52|tm.PCT, 2|tm.PCT))

	tm.Flush()
	time.Sleep(time.Second * time.Duration(second))
}

// Clear use flocking to clear the terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}
