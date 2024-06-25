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
	grid1 [25]uint64
	grid2 [25]uint64
}

// Runtime: O(1)
func (r *Solution) seed() {
	r.grid1[11] |= 1 << 11
	r.grid1[12] |= 1 << 12
	r.grid1[13] |= 1 << 10
	r.grid1[13] |= 1 << 11
	r.grid1[13] |= 1 << 12
}

// Runtime: O(1)
func (r *Solution) computeNextGeneration() {
	for i := 1; i < 24; i++ {
		if r.grid1[i] == 0 &&
			r.grid1[i-1] == 0 &&
			r.grid1[i+1] == 0 {
			continue
		}

		for j := 1; j < 63; j += 1 {
			// <count live neighbors>
			var top, mid, bot uint64

			//     --- triplet of interest
			//     |
			//     v
			// 00011100 -> 00000111
			mid = r.grid1[i] >> (j - 1)
			top = r.grid1[i-1] >> (j - 1)
			bot = r.grid1[i+1] >> (j - 1)

			//      --- bit of interest
			//      |             |
			//      |             |
			// 00000111 -> 00000001 & 00000001
			topLeft := (top >> 2) & 1
			//       --- bit of interest
			//       |            |
			//       |            |
			// 00000111 -> 00000011 & 00000001
			topVal := (top >> 1) & 1
			//        --- bit of interest
			//        |
			//        |
			// 00000111 & 00000001
			topRight := top & 1

			botLeft := (bot >> 2) & 1
			botVal := (bot >> 1) & 1
			botRight := bot & 1

			midRight := mid & 1
			midLeft := (mid >> 2) & 1

			aliveNeighbors := uint8(topLeft + topVal + topRight + midLeft + midRight + botLeft + botVal + botRight)
			// </count live neighbors>

			if aliveNeighbors == 3 || (aliveNeighbors == 2 && (r.grid1[i]&(1<<j)) != 0) {
				r.grid2[i] |= 1 << j
			} else {
				r.grid2[i] &= ^(1 << j)
			}
		}
	}

	r.grid1 = r.grid2
}

// Runtime: O(1)
func (r *Solution) print() {
	// I am not sure if this works on all OS, but this is "move cursor to home position (0, 0)".
	_, _ = os.Stdout.Write([]byte{0x1B, 0x5B, 0x48})

	for i := 0; i < 25; i++ {
		for j := 0; j < 64; j++ {
			if r.grid1[i]&(1<<j) != 0 {
				_, _ = fmt.Fprintf(os.Stdout, "%2s", "*")
			} else {
				if i == 0 || i == 24 || j == 0 || j == 63 {
					_, _ = fmt.Fprintf(os.Stdout, "%2s", "#")
				} else {
					_, _ = fmt.Fprintf(os.Stdout, "%2s", " ")
				}
			}
		}

		_, _ = fmt.Fprint(os.Stdout, "\n")
	}
}

// Runtime: O(1)
// I understand, this method looks messy. I could have used hashmap to define grid and leverage it in this method.
func (r *Solution) countLiveNeighbors(i, j uint8) uint8 {
	var top, mid, bot uint64

	mid = r.grid1[i] >> (j - 1)
	top = r.grid1[i-1] >> (j - 1)
	bot = r.grid1[i+1] >> (j - 1)

	topLeft := (top >> 2) & 1
	topVal := (top >> 1) & 1
	topRight := top & 1

	botLeft := (bot >> 2) & 1
	botVal := (bot >> 1) & 1
	botRight := bot & 1

	midRight := mid & 1
	midLeft := (mid >> 2) & 1

	sum := uint8(topLeft + topVal + topRight + midLeft + midRight + botLeft + botVal + botRight)

	return sum
}

func (r *Solution) Run() {
	r.seed()

	for {
		r.print()
		r.computeNextGeneration()
		time.Sleep(time.Millisecond * 200)
	}
}

// 000[000][110][100][010][001][011][111]

// num >> (j-1)
//
// 00000000_00000010 | 00000000_00000100 => 00000000_00000110
// 00000000_00000000
// 00000000_00000000

// (00000000_00000111 & 00000000_00000001) +
// ((00000000_00000111 >> 1) & 00000000_00000001) +
// ((00000000_00000111 >> 2) & 00000000_00000001) +

// 0000_0000_0000_0000_0000_0000_0000_0111_0000_0000_0000_0000_0000_0000_0000_0000 abov
// 0000_0000_0000_0000_0000_0000_0000_0111_0000_0000_0000_0000_0000_0000_0000_0000 curr
// 0000_0000_0000_0000_0000_0000_0000_0111_0000_0000_0000_0000_0000_0000_0000_0000 belo
//                                    v
// 0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000_0000

/*
11,11
12,12
13,10
13,11
13,12



0000000000000000
0000000000000000
0000000110000000
0000000110000000
0000000000000000

11111
*/
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_00000***00000000_00000000_00000000_00000000
// 00000000_00000000_00000000_00000*1**0000000_00000000_00000000_00000000 abov
// 00000000_00000000_00000000_0000***1*0000000_00000000_00000000_00000000 curr
// 00000000_00000000_00000000_0000*111*0000000_00000000_00000000_00000000 belo
// 00000000_00000000_00000000_0000*****0000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000
// 00000000_00000000_00000000_0000000000000000_00000000_00000000_00000000

// x/2

// 1111_0000
// 0010_0000
