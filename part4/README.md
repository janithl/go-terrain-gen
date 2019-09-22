# A basic terrain generator in Go (Part 4)

Explore how to assign values to adjacent map elements, introduce
randomness in the form of cliff probability

---

- To build: `go build -o bin/terrain .`
- To run: `bin/terrain`

[Blog post can be found here](https://janithl.github.io/2019/09/go-terrain-gen-part-4/)

## Usage:

```
$ bin/terrain --help
Usage of bin/terrain:
  -cliffs int
        percentage probability a cliff will randomly appear (default 15)
  -elev int
        levels of elevation on the map (default 9)
  -height int
        height of the map (default 16)
  -peaks int
        percentage probability a peak will randomly appear (default 5)
  -steps
        show the generation steps
  -width int
        width of the map (default 16)
```
