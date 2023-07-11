package photo

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"

	// start module import
	"github.com/kevfo/photomosaic_generator/start"
)

func GenerateMosaic(output string, original image.Image, tileSize int, data map[string][3]float64) {
	oriBounds := original.Bounds()
	mosaic, height, width := image.NewRGBA(image.Rect(0, 0, oriBounds.Dx(), oriBounds.Dy())), oriBounds.Dy(), oriBounds.Dx()

	for y := height - tileSize; y >= 0; y -= tileSize {
		for x := 0; x < width; x += tileSize {
			crop := image.Rect(x, y, x+tileSize, y+tileSize)
			subImg := original.(interface {
				SubImage(r image.Rectangle) image.Image
			}).SubImage(crop).(image.Image)

			tileName := findClosest(start.ProcessImage(subImg), data)
			openTile, err := os.Open(tileName)
			if err != nil {
				fmt.Printf("Could not open the following image: '%s'\n", tileName)
				os.Exit(1)
			}
			oriTile, err := png.Decode(openTile)
			if err != nil {
				fmt.Printf("Could not decode the image '%s'\n", tileName)
				os.Exit(1)
			}
			tile := resizeImage(oriTile, tileSize, tileSize)

			// Replacing the image here:
			for x := crop.Min.X; x < crop.Max.X; x++ {
				for y := crop.Min.Y; y < crop.Min.Y; y++ {
					mosaic.Set(x, y, tile.At(x-crop.Min.X, y-crop.Min.Y))
				}
			}
		}
	}
	outputImage(mosaic, output)
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

func outputImage(img image.Image, fileName string) {
	filePath, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating '%s': %s\n", fileName, err)
		os.Exit(1)
	}
	err = png.Encode(filePath, img)
	if err != nil {
		fmt.Printf("An error occurred while encoding: ")
	}
}
