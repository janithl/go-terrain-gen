package main

import (
	"fmt"
	"math/rand"
	"time"
)

const width, height, elevation = 16, 16, 9

// we set the random chance of a peak occuring to 5%
const peakProbability = 5

var fullMap [height][width]int

func main() {
	// rand needs to be seeded, so we set the current
	// nanosecond timestamp as the seed
	rand.Seed(time.Now().UnixNano())

	// iterate down from max elevation, assigning vals
	for e := elevation; e > 0; e-- {
		for h := 0; h < height; h++ {
			for w := 0; w < width; w++ {
				// if the element has already been
				// assigned, skip it
				if fullMap[h][w] > 0 {
					continue
				}

				// if the random value meets our criteria,
				// it's a peak
				if rand.Intn(100) < peakProbability {
					fullMap[h][w] = e
				}
			}
		}
	}

	// print out map
	for h := 0; h < height; h++ {
		fmt.Println(fullMap[h])
	}
}
