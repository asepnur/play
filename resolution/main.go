package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"net/http"
)

func getImageResolution(url string) (int, int, error) {
	// Fetch the image from the URL
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	// Decode the image
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	// Get the dimensions
	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	return width, height, nil
}

func main() {
	url := "https://lh3.googleusercontent.com/a/ACg8ocKeC7EfKMIbPW-S_oJ5CcADpdz8Ncxw1yqxXbD_EhYONK_UsA=s461-c-no"

	// Get the image resolution
	width, height, err := getImageResolution(url)
	if err != nil {
		log.Fatalf("failed to get image resolution: %v", err)
	}

	fmt.Printf("The resolution of the image is %dx%d pixels\n", width, height)
}
