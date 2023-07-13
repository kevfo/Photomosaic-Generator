package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	// Custom modules for this project
	"github.com/kevfo/photomosaic_generator/photo"
	"github.com/kevfo/photomosaic_generator/start"
)

func main() {
	inputImage := flag.String("img", "image.png", "The name of the image that you want to convert into a photomosaic.")
	outputName := flag.String("out", "result.png", "The name of the resulting photomosaic.")
	library := flag.String("lib", "pictures", "The name of a folder of photos that you want to use to construct the photomosaic.")
	tileLength := flag.Int("length", 20, "The length of tiles per row that you want the photomosaic to have.")
	tileHeight := flag.Int("height", 40, "The height of tiles per row that you want the photomosaic to have.")
	flag.Parse()

	database := start.InitDatabase(*library)
	img, err := os.Open(*inputImage)
	if err != nil {
		fmt.Printf("An error happened while reading in %s: %s\n", *inputImage, err)
		os.Exit(1)
	}
	defer img.Close()

	toUse, err := png.Decode(img)
	if err != nil {
		fmt.Printf("Unable to decode %s - are you sure %s is a PNG image?\n", img.Name(), img.Name())
		os.Exit(1)
	}
	photo.GenerateMosaic(*outputName, toUse, *tileLength, *tileHeight, database)

	// Delete temp files:
	err = os.Remove("temp.png")
	if err != nil {
		fmt.Printf("An error occurred while removing 'temp.png': '%s'\n", err)
		os.Exit(1)
	}
	fmt.Println("All done!")
}
