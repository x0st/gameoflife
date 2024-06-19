package solution1

import (
	"fmt"
	"os"
	"time"
)

// I have hardcoded the 25 for the sake of speed.
// Could have used a variable to denote grid size.

// This solution allocates a new grid on every tick.

type Solution struct {
	grid [25][25]uint8
}

// Runtime: O(1)
func (r *Solution) seed() {
	r.grid = [25][25]uint8{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
}

// Runtime: O(1)
func (r *Solution) computeNextGeneration() {
	// Copy the current grid. I understand, this can be memory-heavy in a certain environment.
	nextGenGrid := *(&r.grid)

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			cellState := r.grid[i][j]
			liveNeighbors := r.countLiveNeighbors(i, j)

			switch cellState {
			case 1:
				// Any live cell with fewer than two live neighbours dies
				if liveNeighbors < 2 {
					nextGenGrid[i][j] = 0
					continue
				}

				// Any live cell with two or three live neighbours lives
				if liveNeighbors == 2 || liveNeighbors == 3 {
					nextGenGrid[i][j] = 1
					continue
				}

				// Any live cell with more than three live neighbours dies
				if liveNeighbors > 3 {
					nextGenGrid[i][j] = 0
					continue
				}
			case 0:
				// Any dead cell with exactly three live neighbours becomes a live cell
				if liveNeighbors == 3 {
					nextGenGrid[i][j] = 1
				}
			}
		}
	}

	r.grid = nextGenGrid
}

// Runtime: O(1)
func (r *Solution) print() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if r.grid[i][j] == 1 {
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
		liveNeighbors += r.grid[i-1][j]

		if j > 0 {
			liveNeighbors += r.grid[i-1][j-1]
		}

		if j < 24 {
			liveNeighbors += r.grid[i-1][j+1]
		}
	}

	if i < 24 {
		liveNeighbors += r.grid[i+1][j]

		if j > 0 {
			liveNeighbors += r.grid[i+1][j-1]
		}

		if j < 24 {
			liveNeighbors += r.grid[i+1][j+1]
		}
	}

	if j > 0 {
		liveNeighbors += r.grid[i][j-1]
	}

	if j < 24 {
		liveNeighbors += r.grid[i][j+1]
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
