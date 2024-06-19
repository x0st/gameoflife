package solution3

import (
	"fmt"
	"os"
	"time"
)

// I have hardcoded the 25 for the sake of speed.
// Could have used a variable to denote grid size.

// This is sort of a "double buffering" approach.

type Solution struct {
	grid map[string][2]uint8
}

func (r *Solution) seed() {
	r.grid = map[string][2]uint8{}

	r.grid["1112"] = [2]uint8{1, 0}
	r.grid["1213"] = [2]uint8{1, 0}
	r.grid["1313"] = [2]uint8{1, 0}
	r.grid["1312"] = [2]uint8{1, 0}
	r.grid["1311"] = [2]uint8{1, 0}
}

func (r *Solution) computeNextGeneration() {
	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			cellAddress := fmt.Sprintf("%d%d", i, j)
			cellLiveNeighbors := r.countLiveNeighbors(i, j)

			switch r.grid[cellAddress][0] {
			case 1:
				// Any live cell with fewer than two live neighbours dies
				if cellLiveNeighbors < 2 {
					r.grid[cellAddress] = [2]uint8{r.grid[cellAddress][0], 0}
					continue
				}

				// Any live cell with two or three live neighbours lives
				if cellLiveNeighbors == 2 || cellLiveNeighbors == 3 {
					r.grid[cellAddress] = [2]uint8{r.grid[cellAddress][0], 1}
					continue
				}

				// Any live cell with more than three live neighbours dies
				if cellLiveNeighbors > 3 {
					r.grid[cellAddress] = [2]uint8{r.grid[cellAddress][0], 0}
					continue
				}
			case 0:
				// Any dead cell with exactly three live neighbours becomes a live cell
				if cellLiveNeighbors == 3 {
					r.grid[cellAddress] = [2]uint8{r.grid[cellAddress][0], 1}
				}
			}
		}
	}

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			cellAddress := fmt.Sprintf("%d%d", i, j)
			r.grid[cellAddress] = [2]uint8{r.grid[cellAddress][1], 0}
		}
	}
}

// Runtime: O(1)
func (r *Solution) print() {
	// I am not sure if this works on all OS, but this is "clear screen".
	_, _ = os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})

	for i := 0; i < 25; i++ {
		for j := 0; j < 25; j++ {
			if r.grid[fmt.Sprintf("%d%d", i, j)][0] == 1 {
				_, _ = fmt.Fprintf(os.Stdout, "%3s", "*")
			} else {
				_, _ = fmt.Fprintf(os.Stdout, "%3s", "_")
			}
		}

		_, _ = fmt.Fprint(os.Stdout, "\n")
	}
}

// Runtime: O(1)
func (r *Solution) countLiveNeighbors(i, j int) uint8 {
	return r.grid[fmt.Sprintf("%d%d", i-1, j)][0] +
		r.grid[fmt.Sprintf("%d%d", i-1, j-1)][0] +
		r.grid[fmt.Sprintf("%d%d", i-1, j+1)][0] +
		r.grid[fmt.Sprintf("%d%d", i+1, j)][0] +
		r.grid[fmt.Sprintf("%d%d", i+1, j-1)][0] +
		r.grid[fmt.Sprintf("%d%d", i+1, j+1)][0] +
		r.grid[fmt.Sprintf("%d%d", i, j-1)][0] +
		r.grid[fmt.Sprintf("%d%d", i, j+1)][0]
}

func (r *Solution) Run() {
	r.seed()

	for {
		r.print()
		r.computeNextGeneration()
		time.Sleep(time.Millisecond * 300)
	}
}
