package photo

import (
	"image"
	"math"

	// start module import
	_ "../start"
)

/*
1) Divide image up into however many tiles
2) Get RGBA value of tile
3) Replace tile with closest value
4) Repeat 1) to 3) until all tiles done.
*/
func GenerateMosaic(output string, original image.Image, tileSize int) image.Image {
	oriBounds := original.Bounds()
	mosaic, height, width := image.NewRGBA(image.Rect(0, 0, oriBounds.Dx(), oriBounds.Dy())), oriBounds.Dy(), oriBounds.Dx()

	for y := height - 1; y >= 0; y -= tileSize {
		for x := 0; x < width; x += tileSize {

		}
	}
	return mosaic
}

func distance(a1, a2 [3]float64) float64 {
	return math.Sqrt(math.Pow(a2[0]-a1[0], 2) + math.Pow(a2[1]-a1[1], 2) + math.Pow(a2[2]-a1[2], 2))
}

func findClosest(colors [3]float64, data map[string][3]float64) string {
	name, closest := "", math.Inf(1)
	for image, rgbaVal := range data {
		dist := distance(colors, rgbaVal)
		if dist < closest {
			closest, name = dist, image
		}
	}
	return name
}

func resizeImage(img image.Image, newWidth, newHeight int) image.Image {
	bounds := img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()

	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	for x := 0; x < newWidth; x++ {
		for y := 0; x < newHeight; y++ {
			oriX := x * width / newWidth
			oriY := y * height / newHeight
			color := img.At(oriX, oriY)
			resizedImg.Set(x, y, color)
		}
	}
	return resizedImg
}
