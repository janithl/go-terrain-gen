package main

import (
	"flag"

	"github.com/janithl/go-terrain-gen/part4/terrain"
)

func main() {
	height := flag.Int("height", 16, "height of the map")
	width := flag.Int("width", 16, "width of the map")
	elevation := flag.Int("elev", 9, "levels of elevation on the map")
	peakProbability := flag.Int("peaks", 5, "percentage probability a peak will randomly appear")
	cliffProbability := flag.Int("cliffs", 5, "percentage probability a cliff will randomly appear")
	steps := flag.Bool("steps", false, "show the generation steps")
	flag.Parse()

	terrainMap := terrain.NewFullMap(*width, *height, *elevation, *peakProbability, *cliffProbability)
	terrainMap.Generate(*steps)
	terrainMap.Print()
}
