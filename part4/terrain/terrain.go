package terrain

import (
	"fmt"
	"math"
	"math/rand"
	"os/exec"
	"time"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type fullMap struct {
	width, height, elevation, peakProbability, cliffProbability int
	elements                                                    [][]int
}

// NewFullMap returns a new terrain map
func NewFullMap(width, height, elevation, peakProbability, cliffProbability int) fullMap {
	elements := make([][]int, height)
	for i := 0; i < height; i++ {
		elements[i] = make([]int, width)
	}

	return fullMap{
		width:            width,
		height:           height,
		elevation:        elevation,
		peakProbability:  peakProbability,
		cliffProbability: cliffProbability,
		elements:         elements,
	}
}

// adjacentElevation checks if an adjacent element
// to the given element (h, w) is at a given elevation
func (f *fullMap) adjacentElevation(h, w, elevation int) bool {
	for y := max(0, h-1); y <= min(f.height-1, h+1); y++ {
		for x := max(0, w-1); x <= min(f.width-1, w+1); x++ {
			if f.elements[y][x] == elevation+1 {
				// if this element is *not* randomly a cliff, return true
				return rand.Intn(100) > f.cliffProbability
			}
		}
	}

	return false
}

// Generate generates the terrain map
func (f *fullMap) Generate(viewSteps bool) {
	// rand needs to be seeded, so we set the current
	// nanosecond timestamp as the seed
	rand.Seed(time.Now().UnixNano())

	// iterate down from max elevation, assigning vals
	for e := f.elevation; e > 0; e-- {
		for h := 0; h < f.height; h++ {
			for w := 0; w < f.width; w++ {
				// if the element has already been assigned, skip it
				if f.elements[h][w] > 0 {
					continue
				}

				// if the element is next to a element with elevation x, it
				// should get elevation x - 1
				// alternately, if the random value meets our criteria, it's a peak
				if f.adjacentElevation(h, w, e) || rand.Intn(100) < f.peakProbability {
					f.elements[h][w] = e
				}
			}
		}

		// if viewSteps is true, we print every step of the generation process,
		// sleep for a second, and clear the screen for the next step
		if viewSteps {
			f.Print()
			time.Sleep(time.Second)
			exec.Command("clear")
			fmt.Println()
		}
	}
}

// Print prints the terrain map out
func (f *fullMap) Print() {
	mapColours := [4]int{36, 34, 32, 33} // blue, cyan, green, yellow
	mapShades := [4]string{"░", "▒", "▓", "█"}

	// print out map
	for h := 0; h < f.height; h++ {
		for w := 0; w < f.width; w++ {
			// print a space (black) if elevation is zero
			if f.elements[h][w] == 0 {
				print(" ")
				continue
			}

			// get the approximate colour nearest to the elevation number
			elementColour := float64(f.elements[h][w]) / float64(f.elevation) * float64(len(mapColours)-1)

			// get the colour index
			colourIndex := int(math.Round(elementColour))

			// get the approximate shade within that colour
			elementShade := (elementColour - math.Floor(elementColour)) * float64(len(mapShades)-1)

			// get its index
			shadeIndex := int(math.Round(elementShade))

			// print out the corresponding ANSI code and unicode char
			fmt.Printf("\033[%dm%s\033[0m", mapColours[colourIndex], mapShades[shadeIndex])
		}

		// print a newline
		fmt.Println()
	}
}
