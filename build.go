package main

import (
	"flag"
	"fmt"

	// Custom modules for this project
	"github.com/kevfo/photomosaic_generator/start"
)

func main() {
	inputImage := flag.String("img", "image.png", "The name of the image that you want to convert into a photomosaic.")
	outputName := flag.String("out", "result.png", "The name of the resulting photomosaic.")
	library := flag.String("lib", "pictures", "The name of a folder of photos that you want to use to construct the photomosaic.")
	numTiles := flag.Int("tiles", 10, "The size of tiles per row that you want the photomosaic to have.")
	flag.Parse()

	database := start.InitDatabase(*library)
	fmt.Println(database)

	fmt.Println(*inputImage, *outputName, *numTiles)
}
