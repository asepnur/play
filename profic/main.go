package main

import (
	"embed"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/golang/freetype"
)

// Embed the font file into the binary.
//go:embed Roboto-Regular.ttf
var fontFile embed.FS

// GenerateProfilePicture generates a profile picture with the first letter of the name.
func GenerateProfilePicture(name string, outputPath string) error {
	const (
		imgWidth  = 200
		imgHeight = 200
		fontSize  = 100
		dpi       = 72
	)

	// Create a new blank image with a random background color.
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	bgColor := randomColor()
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// Load the font.
	fontBytes, err := fontFile.ReadFile("Roboto-Regular.ttf")
	if err != nil {
		return err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}

	// Draw the first letter of the name.
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White) // Set text color to white for better contrast

	// Calculate the position to center the text.
	firstLetter := string(name[0])
	pt := freetype.Pt((imgWidth/2)-fontSize/4, (imgHeight/2)+fontSize/3)

	_, err = c.DrawString(firstLetter, pt)
	if err != nil {
		return err
	}

	// Save the image to the output path.
	outFile, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	return jpeg.Encode(outFile, img, nil)
}

// randomColor generates a random color.
func randomColor() color.RGBA {
	rand.Seed(time.Now().UnixNano())
	return color.RGBA{
		R: uint8(rand.Intn(256)),
		G: uint8(rand.Intn(256)),
		B: uint8(rand.Intn(256)),
		A: 255,
	}
}

func main() {
	name := flag.String("name", "A", "The name to generate the profile picture for.")
	outputPath := flag.String("output", "profile.jpg", "The output path for the generated profile picture.")
	flag.Parse()

	err := GenerateProfilePicture(*name, *outputPath)
	if err != nil {
		log.Fatalf("Error generating profile picture: %v", err)
	}

	log.Printf("Profile picture generated and saved to %s", *outputPath)
}
