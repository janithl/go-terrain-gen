# A basic terrain generator in Go (Part 3)

Better map visualisation: Added colours and shading to output
of map to terminal

---

* To build: `go build -o bin/terrain .`
* To run: `bin/terrain`

[Blog post can be found here](https://janithl.github.io/2019/09/go-terrain-gen-part-3/)

## Usage:

```
$ bin/terrain --help
Usage of bin/terrain:
    -elev int
        levels of elevation on the map (default 9)
    -height int
        height of the map (default 16)
    -peaks int
        percentage probability a peak will randomly appear (default 5)
    -width int
        width of the map (default 16)
```
