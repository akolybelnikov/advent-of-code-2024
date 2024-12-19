package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

type Robot struct {
	P, V image.Point
}

func main() {
	// read form file
	input, err := utils.ReadFile("resources/day14.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input, 101, 103))
	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string, w, h int) int {
	ans := 1
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	positions := make(map[int]int)
	for _, line := range lines {
		line = strings.Trim(line, " ")
		var x, y, vx, vy int
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		pos := calc(w, h, x, y, vx*100, vy*100)
		assign(pos, &positions, w, h)
	}
	for _, v := range positions {
		ans *= v
	}

	return ans
}

// part2 calculates the number of steps required for all robots to occupy unique positions in a bounded grid area.
func part2(input string) int {
	area := image.Rectangle{Min: image.Point{}, Max: image.Point{X: 101, Y: 103}}

	var robots []Robot
	quads := map[image.Point]int{}
	for _, s := range strings.Split(strings.TrimSpace(input), "\n") {
		var r Robot
		_, err := fmt.Sscanf(s, "p=%d,%d v=%d,%d", &r.P.X, &r.P.Y, &r.V.X, &r.V.Y)
		if err != nil {
			utils.HandleErr(err)
		}
		robots = append(robots, r)
		r.P = r.P.Add(r.V.Mul(100)).Mod(area)
		quads[image.Point{X: sgn(r.P.X - area.Dx()/2), Y: sgn(r.P.Y - area.Dy()/2)}]++
	}
	fmt.Println(quads[image.Point{X: -1, Y: -1}] * quads[image.Point{X: 1, Y: -1}] *
		quads[image.Point{X: 1, Y: 1}] * quads[image.Point{X: -1, Y: 1}])

	// calculates the smallest time `t` at which all robots occupy distinct positions in a 2D wrapping grid
	for t := 1; ; t++ {
		seen := map[image.Point]struct{}{}
		for i := range robots {
			robots[i].P = robots[i].P.Add(robots[i].V).Mod(area)
			seen[robots[i].P] = struct{}{}
		}
		if len(seen) == len(robots) {
			return t
		}
	}
}

// calc returns the new position of a point
// handling the wrap around with negative modulo logic
func calc(w, h, x, y, vx, vy int) [2]int {
	nx, ny := ((x+vx)%w+w)%w, ((y+vy)%h+h)%h
	return [2]int{nx, ny}
}

// assign places the robot into one of the quadrants based on its position
func assign(pos [2]int, positions *map[int]int, w, h int) {
	switch {
	case pos[0] < w/2 && pos[1] < h/2:
		(*positions)[0] += 1
	case pos[0] > w/2 && pos[1] > h/2:
		(*positions)[2] += 1
	case pos[0] > w/2 && pos[1] < h/2:
		(*positions)[1] += 1
	case pos[0] < w/2 && pos[1] > h/2:
		(*positions)[3] += 1
	default:
	}
}

// generateImage creates a png representation of the grid at any given second
func generateImage(grid *map[image.Point]image.Point, out int) string {
	// Image dimensions
	width, height := 110, 110

	pointColor := color.RGBA{G: 255, A: 255} // Green color for points

	// Create a blank RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := image.Point{X: x, Y: y}
			if _, ok := (*grid)[p]; ok {
				drawPoint(img, p, 1, pointColor)
			}
		}
	}

	// Save the image to a file in the /images directory
	fileName := fmt.Sprintf("images/grid_%d.png", out)
	outFile, err := os.Create(fileName)
	utils.HandleErr(err)
	defer func(outFile *os.File) {
		closeErr := outFile.Close()
		if closeErr != nil {
			fmt.Println(closeErr)
		}
	}(outFile)

	// Encode the image as PNG
	encodeErr := png.Encode(outFile, img)
	utils.HandleErr(encodeErr)

	return fileName
}

// Helper function to draw a single point as a rectangle
func drawPoint(img *image.RGBA, pt image.Point, size int, col color.Color) {
	halfSize := size / 2
	for y := pt.Y - halfSize; y <= pt.Y+halfSize; y++ {
		for x := pt.X - halfSize; x <= pt.X+halfSize; x++ {
			// Ensure the point stays within image bounds
			if x >= 0 && x < img.Bounds().Dx() && y >= 0 && y < img.Bounds().Dy() {
				img.Set(x, y, col)
			}
		}
	}
}

func sgn(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}
