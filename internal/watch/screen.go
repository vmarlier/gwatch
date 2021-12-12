package watch

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
)

// DisplayOutput will take the a command output as a string and display it
func DisplayOutput(output []string, second int64, y int) {
	tm.Clear()

	// boxes will contain all the box which will contain the command outputs
	var boxes []*tm.Box
	// placements will contain all the x,y placement for the box
	var placements [][2]int
	// ---

	switch y {
	case 1:
		// 1 box, the box take all the screen
		// Placed top left
		boxes = append(boxes, tm.NewBox(96|tm.PCT, 100|tm.PCT, 0))
		placements = append(placements, [2]int{2, 2})
	case 2:
		// 2 boxes, the 2 boxes shares the width but take all the height
		// Placed top left and top middle
		boxes = append(boxes, tm.NewBox(48|tm.PCT, 100|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 100|tm.PCT, 0))
		placements = append(placements, [2]int{2, 2},
			[2]int{52, 2})
	case 3:
		// 3 boxes, 1 box take all the heigth, 2 boxes shares the width, 2 boxes shares the height
		// Placed top left, top middle and middle middle
		boxes = append(boxes, tm.NewBox(48|tm.PCT, 100|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 50|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 50|tm.PCT, 0))
		placements = append(placements, [2]int{2, 2},
			[2]int{52, 2},
			[2]int{52, 52})
	case 4:
		// 4 boxes, the 4 boxes shared the height and the width
		// Placed top left, top middle, middle left, middle middle
		boxes = append(boxes, tm.NewBox(48|tm.PCT, 50|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 50|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 50|tm.PCT, 0),
			tm.NewBox(48|tm.PCT, 50|tm.PCT, 0))
		placements = append(placements, [2]int{2, 2},
			[2]int{52, 2},
			[2]int{2, 52},
			[2]int{52, 52})
	}

	for i := 0; i < y; i++ {
		// Add some content to the box
		fmt.Fprintf(boxes[i], output[i])

		// Move Box(es)
		tm.Print(tm.MoveTo(boxes[i].String(), placements[i][0]|tm.PCT, placements[i][1]|tm.PCT))
	}

	// ---

	tm.Flush()
	time.Sleep(time.Second * time.Duration(second))
}

// Clear use flocking to clear the terminal
func Clear() {
	fmt.Print("\033[H\033[2J")
}
