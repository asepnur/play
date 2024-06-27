package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

const (
	imgWidth  = 461
	imgHeight = 461
	fontPath  = "./Inter-ExtraBold.ttf"
	outputDir = "./output"
)

var (
	fontColorModel1_2Letters = color.RGBA{177, 78, 198, 255}  // #B14EC6
	bgColorModel1_2Letters   = color.RGBA{249, 227, 255, 255} // #F9E3FF
	fontColorModel2_2Letters = color.RGBA{80, 78, 198, 255}   // #504EC6
	bgColorModel2_2Letters   = color.RGBA{228, 227, 255, 255} // #E4E3FF

	fontColorModel1_1Letter = color.RGBA{177, 78, 198, 255}  // #B14EC6
	bgColorModel1_1Letter   = color.RGBA{249, 227, 255, 255} // #F9E3FF
	fontColorModel2_1Letter = color.RGBA{80, 78, 198, 255}   // #504EC6
	bgColorModel2_1Letter   = color.RGBA{228, 227, 255, 255} // #E4E3FF
)

func main() {
	// Load the font
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		log.Fatalf("Failed to load font: %v", err)
	}

	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Fatalf("Failed to parse font: %v", err)
	}

	alphanumericSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	// Create output directories
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Generate images for 2-letter combinations
	for _, c1 := range alphanumericSet {
		for _, c2 := range alphanumericSet {
			text := string(c1) + string(c2)
			generateImage(f, text, fontColorModel1_2Letters, bgColorModel1_2Letters, filepath.Join(outputDir, fmt.Sprintf("%s_model_1_2_letters.png", text)))
			generateImage(f, text, fontColorModel2_2Letters, bgColorModel2_2Letters, filepath.Join(outputDir, fmt.Sprintf("%s_model_2_2_letters.png", text)))
		}
	}

	// Generate images for 1-letter combinations
	for _, c := range alphanumericSet {
		text := string(c)
		generateImage(f, text, fontColorModel1_1Letter, bgColorModel1_1Letter, filepath.Join(outputDir, fmt.Sprintf("%s_model_1_1_letter.png", text)))
		generateImage(f, text, fontColorModel2_1Letter, bgColorModel2_1Letter, filepath.Join(outputDir, fmt.Sprintf("%s_model_2_1_letter.png", text)))
	}
}

func generateImage(f *truetype.Font, text string, fontColor, bgColor color.Color, outputPath string) {
	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(f)
	c.SetFontSize(220)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(&image.Uniform{fontColor})

	pt := freetype.Pt(imgWidth/4, imgHeight/2+110) // Adjust as necessary to center text
	_, err := c.DrawString(text, pt)
	if err != nil {
		log.Fatalf("Failed to draw string: %v", err)
	}

	// Save the image
	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create image file: %v", err)
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		log.Fatalf("Failed to encode image to PNG: %v", err)
	}

	fmt.Printf("Generated %s\n", outputPath)
}
