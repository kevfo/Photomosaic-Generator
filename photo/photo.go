package photo

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
	"os"

	// start module import
	"github.com/kevfo/photomosaic_generator/start"
)

func GenerateMosaic(output string, original image.Image, tileWidth, tileHeight int, data map[string][3]float64) {
	fmt.Println("Building the photomosaic now...")
	oriBounds := original.Bounds()
	mosaic, height, width := image.NewRGBA(image.Rect(0, 0, oriBounds.Dx(), oriBounds.Dy())), oriBounds.Dy(), oriBounds.Dx()

	for y := height; y > tileHeight; y -= tileHeight {
		for x := 0; x < width-tileWidth; x += tileWidth {
			crop := image.Rect(x, y, x+tileWidth, y-tileHeight)
			subImg := original.(*image.NRGBA).SubImage(crop)
			outputImage(subImg, "temp.png")

			// Open the temp image file here:
			temp, err := os.Open("temp.png")
			if err != nil {
				fmt.Printf("An error happened: %s\n", err)
				os.Exit(1)
			}
			defer temp.Close()

			tempImg, err := png.Decode(temp)
			if err != nil {
				fmt.Printf("An error happened: %s\n", err)
				os.Exit(1)
			}
			tileName := findClosest(start.ProcessImage(tempImg), data)
			openTile, err := os.Open(tileName)
			if err != nil {
				fmt.Printf("Could not open the following image: '%s'\n", tileName)
				os.Exit(1)
			}
			defer openTile.Close()

			oriTile, err := png.Decode(openTile)
			if err != nil {
				fmt.Printf("Could not decode the image '%s'\n", tileName)
				os.Exit(1)
			}
			tile := resizeImage(oriTile, tileWidth, tileHeight)

			// Replacing the image here:
			for yCrop := y; yCrop < y+tileHeight; yCrop++ {
				for xCrop := x; xCrop < x+tileWidth; xCrop++ {
					mosaic.Set(xCrop, yCrop-tileHeight, tile.At(xCrop-x, yCrop-y))
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
	name, closest := "", math.Pow(100, 3)
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
		for y := 0; y < newHeight; y++ {
			oriX := x * width / newWidth
			oriY := y * height / newHeight
			resizedImg.Set(x, y, img.At(oriX, oriY))
		}
	}
	return resizedImg
}

func convertToNRGBA(original image.Image) image.NRGBA {
	bounds := original.Bounds()
	newImg := image.NewNRGBA(bounds)
	draw.Draw(newImg, bounds, original, bounds.Min, draw.Src)
	return *newImg
}

func outputImage(img image.Image, fileName string) {
	filePath, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating '%s': %s\n", fileName, err)
		os.Exit(1)
	}
	defer filePath.Close()

	err = png.Encode(filePath, img)
	if err != nil {
		fmt.Printf("An error occurred while encoding: %s\n", err)
		os.Exit(1)
	}
}
