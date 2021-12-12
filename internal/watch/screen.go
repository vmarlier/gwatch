package watch

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

// DisplayOutput will take the a command output as a string and display it
func DisplayOutput(output string, second int64) {
	Clear()

	// By moving cursor to top-left position we ensure that console output
	// will be overwritten each time, instead of adding new.
	tm.MoveCursor(0, 0)

	tm.Println(output)

	tm.Flush() // Call it every time at the end of rendering

	time.Sleep(time.Second * time.Duration(second))
}

func DisplayOutputBox(output string, second int64) {
	tm.Clear()

	// Create Box with 30% width of current screen, and height of 20 lines
	box := tm.NewBox(40|tm.PCT, 90|tm.PCT, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box, output)

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box.String(), 2|tm.PCT, 5|tm.PCT))

	// Create Box with 30% width of current screen, and height of 20 lines
	box2 := tm.NewBox(40|tm.PCT, 90|tm.PCT, 0)

	// Add some content to the box
	// Note that you can add ANY content, even tables
	fmt.Fprint(box2, output)

	// Move Box to approx center of the screen
	tm.Print(tm.MoveTo(box2.String(), 52|tm.PCT, 5|tm.PCT))

	tm.Flush()
	time.Sleep(time.Second * time.Duration(second))
}

// DisplayOutput will take several commands outputs as strings and display them
func DisplayMultiplesOutput(outputs ...string) {
	tm.Clear() // Clear current screen

	// By moving cursor to top-left position we ensure that console output
	// will be overwritten each time, instead of adding new.
	tm.MoveCursor(1, 1)

	tm.Println("Current Time:", time.Now().Format(time.RFC1123))

	width := tm.Width() / 2
	tm.MoveCursor(width, 1)

	tm.Println("Current Time:", time.Now().Format(time.RFC1123))

	tm.Flush() // Call it every time at the end of rendering

	time.Sleep(time.Second)
}

// Clear use flocking to clear the terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}
