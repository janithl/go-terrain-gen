package main

import (
	"flag"

	"github.com/janithl/go-terrain-gen/part2/terrain"
)

func main() {
	height := flag.Int("height", 16, "height of the map")
	width := flag.Int("width", 16, "width of the map")
	elevation := flag.Int("elev", 9, "levels of elevation on the map")
	peakProbability := flag.Int("peaks", 5, "percentage probability a peak will randomly appear")
	flag.Parse()

	terrainMap := terrain.NewFullMap(*height, *width, *elevation, *peakProbability)
	terrainMap.Generate()
	terrainMap.Print()
}
