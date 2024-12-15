package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"image/png"
	"os"
	"strings"

	"github.com/akolybelnikov/advent-of-code-2024/internal/utils"
)

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

	frames := make([]string, 101*103)
	for i := 0; i < 101*103; i++ {
		img := fmt.Sprintf("images/grid_%d.png", i)
		frames[i] = img
	}
	generateGIF(frames)

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

// part two
func part2(input string) int {
	ans := 1
	w, h := 101, 103
	lines, err := utils.ParseLines(input)
	utils.HandleErr(err)
	points := make(map[image.Point]image.Point)
	frames := make([]string, 0)
	for _, line := range lines {
		line = strings.Trim(line, " ")
		var p, v image.Point
		_, _ = fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p.X, &p.Y, &v.X, &v.Y)
		points[v] = p
		imageName := generateImage(&points, 0)
		frames = append(frames, imageName)
	}

	for i := 1; i < (101 * 103); i++ {
		newPoints := make(map[image.Point]image.Point)
		for point, vector := range points {
			np := calc(w, h, point.X, point.Y, vector.X, vector.Y)
			newPoints[image.Point{X: np[0], Y: np[1]}] = vector
		}
		imageName := generateImage(&newPoints, i)
		frames = append(frames, imageName)
		points = newPoints
	}

	generateGIF(frames)

	return ans
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

func generateGIF(frames []string) {
	// Output GIF
	outputFile, err := os.Create("output.gif")
	if err != nil {
		panic(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		utils.HandleErr(err)
	}(outputFile)

	// Create the GIF structure
	var g gif.GIF

	// Set desired delay per frame
	frameDelay := 50 // Delay in hundredths of a second (e.g., 5 = 50ms per frame)

	// Define the transparent color for the background
	transparentColor := color.Transparent
	// Create a custom palette with green as the main color (adjust as necessary)
	greenColor := color.RGBA{G: 255, A: 255} // Pure green
	// Create a palette containing green and transparent colors
	palette := []color.Color{
		greenColor,
		transparentColor,
	}

	// Loop through each PNG file
	for _, file := range frames {
		// Open the PNG file
		f, err := os.Open(file)
		utils.HandleErr(err)

		// Decode the PNG image
		img, err := png.Decode(f)
		utils.HandleErr(err)

		// Create a paletted image with a transparent background
		palettedImage := image.NewPaletted(img.Bounds(), palette)

		// Manually draw the image pixels into the palettedImage, preserving transparency
		for y := 0; y < img.Bounds().Dy(); y++ {
			for x := 0; x < img.Bounds().Dx(); x++ {
				// Get the color of the pixel at (x, y)
				at := img.At(x, y)
				// If the pixel is transparent, we leave it as transparent (no change)
				if at == greenColor {
					palettedImage.Set(x, y, at)
				} else {
					palettedImage.Set(x, y, transparentColor)
				}
			}
		}

		// Append the frame to the GIF structure
		g.Image = append(g.Image, palettedImage)
		g.Delay = append(g.Delay, frameDelay) // Set delay for this frame

		_ = f.Close()
	}

	// Encode and save the GIF to the output file
	err = gif.EncodeAll(outputFile, &g)
	utils.HandleErr(err)
}
