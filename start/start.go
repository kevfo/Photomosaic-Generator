package start

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func InitDatabase(photoDir string) map[string][3]float64 {
	imgDatabase := make(map[string][3]float64)

	photos, err := os.ReadDir(photoDir)
	if err != nil {
		fmt.Printf("Unable to locate the directory %s - are you sure it hasn't been moved?\n", photoDir)
		os.Exit(1)
	}

	for _, photo := range photos {
		openPhoto, err := os.Open(photoDir + "/" + photo.Name())
		if err != nil {
			fmt.Printf("Unable to open '%s'\nError: '%s'", photo.Name(), err)
			os.Exit(1)
		}
		defer openPhoto.Close()

		pic, err := png.Decode(openPhoto)
		if err != nil {
			fmt.Printf("Unable to decode the file '%s' \nError: '%s'", photoDir+"/"+photo.Name(), err)
			os.Exit(1)
		}
		imgDatabase[photoDir+"/"+photo.Name()] = ProcessImage(pic)
		fmt.Printf("Done processing: '%s'\n", photoDir+"/"+photo.Name())
	}
	return imgDatabase
}

func ProcessImage(img image.Image) [3]float64 {
	averageRGBA, bounds := [3]float64{0.0, 0.0, 0.0}, img.Bounds()
	width, height := bounds.Dx(), bounds.Dy()
	for y := height - 1; y >= 0; y-- {
		for x := 0; x < width; x++ {
			point := img.At(x, y)
			r, g, b, _ := point.RGBA()
			averageRGBA[0] += float64(r >> 8)
			averageRGBA[1] += float64(g >> 8)
			averageRGBA[2] += float64(b >> 8)
		}
	}
	return [3]float64{
		averageRGBA[0] / float64(width*height),
		averageRGBA[1] / float64(width*height),
		averageRGBA[2] / float64(width*height),
	}
}
