package solution2

import (
	"fmt"
	"os"
	"time"
)

// I have hardcoded the 25 for the sake of speed.
// Could have used a variable to denote grid size.

// This solution needs a fixed amount of memory. It does not allocate new memory.
// This is sort of a "double buffering" approach.

type cellState struct {
	curr uint8
	next uint8
}

type Solution struct {
	grid [25][25]cellState
}

func (r *Solution) seed() {
	r.grid = [25][25]cellState{}

	r.grid[11][12].curr = 1
	r.grid[12][13].curr = 1
	r.grid[13][13].curr = 1
	r.grid[13][12].curr = 1
	r.grid[13][11].curr = 1
}

func (r *Solution) computeNextGeneration() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			liveNeighbors := r.countLiveNeighbors(i, j)

			switch r.grid[i][j].curr {
			case 1:
				// Any live cell with fewer than two live neighbours dies
				if liveNeighbors < 2 {
					r.grid[i][j].next = 0
					continue
				}

				// Any live cell with two or three live neighbours lives
				if liveNeighbors == 2 || liveNeighbors == 3 {
					r.grid[i][j].next = 1
					continue
				}

				// Any live cell with more than three live neighbours dies
				if liveNeighbors > 3 {
					r.grid[i][j].next = 0
					continue
				}
			case 0:
				// Any dead cell with exactly three live neighbours becomes a live cell
				if liveNeighbors == 3 {
					r.grid[i][j].next = 1
				}
			}
		}
	}

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			r.grid[i][j].curr = r.grid[i][j].next
		}
	}
}

// Runtime: O(1)
func (r *Solution) print() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if r.grid[i][j].curr == 1 {
				_, _ = fmt.Fprintf(os.Stdout, "%3s", "*")
			} else {
				_, _ = fmt.Fprintf(os.Stdout, "%3s", "_")
			}
		}

		_, _ = fmt.Fprint(os.Stdout, "\n")
	}
}

// Runtime: O(1)
// I understand, this method looks messy. I could have used hashmap to define grid and leverage it in this method.
func (r *Solution) countLiveNeighbors(i, j int) uint8 {
	liveNeighbors := uint8(0)

	if i > 0 {
		liveNeighbors += r.grid[i-1][j].curr

		if j > 0 {
			liveNeighbors += r.grid[i-1][j-1].curr
		}

		if j < 24 {
			liveNeighbors += r.grid[i-1][j+1].curr
		}
	}

	if i < 24 {
		liveNeighbors += r.grid[i+1][j].curr

		if j > 0 {
			liveNeighbors += r.grid[i+1][j-1].curr
		}

		if j < 24 {
			liveNeighbors += r.grid[i+1][j+1].curr
		}
	}

	if j > 0 {
		liveNeighbors += r.grid[i][j-1].curr
	}

	if j < 24 {
		liveNeighbors += r.grid[i][j+1].curr
	}

	return liveNeighbors
}

func (r *Solution) Run() {
	r.seed()

	for {
		r.print()
		r.computeNextGeneration()
		time.Sleep(time.Millisecond * 300)
	}
}
