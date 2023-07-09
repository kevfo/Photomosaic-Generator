package photo

import (
	"image"
	"math"
)

func GenerateMosaic(output string, original image.Image, numTiles int) image.Image {
	oriBounds := original.Bounds()
	mosaic, height, width := image.NewRGBA(image.Rect(0, 0, oriBounds.Dx(), oriBounds.Dy())), oriBounds.Dy(), oriBounds.Dx()

	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {

		}
	}
	return mosaic
}

func distance(a1, a2 [3]float64) float64 {
	return math.Sqrt(math.Pow(a2[0]-a1[0], 2) + math.Pow(a2[1]-a1[1], 2) + math.Pow(a2[2]-a1[2], 2))
}

func findClosest()
