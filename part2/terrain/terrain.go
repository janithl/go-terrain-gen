package terrain

import (
	"fmt"
	"math/rand"
	"time"
)

type fullMap struct {
	width, height, elevation, peakProbability int
	elements                                  [][]int
}

// NewFullMap returns a new terrain map
func NewFullMap(width, height, elevation, peakProbability int) fullMap {
	elements := make([][]int, height)
	for i := 0; i < height; i++ {
		elements[i] = make([]int, width)
	}

	return fullMap{
		width:           width,
		height:          height,
		elevation:       elevation,
		peakProbability: peakProbability,
		elements:        elements,
	}
}

// Generate generates the terrain map
func (f *fullMap) Generate() {
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

				// if the random value meets our criteria, it's a peak
				if rand.Intn(100) < f.peakProbability {
					f.elements[h][w] = e
				}
			}
		}
	}
}

// Print prints the terrain map out
func (f *fullMap) Print() {
	// print out map
	for h := 0; h < f.height; h++ {
		fmt.Println(f.elements[h])
	}
}
