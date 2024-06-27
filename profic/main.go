package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fogleman/gg"
)

const (
	width    = 461
	height   = 461
	fontPath = "./Inter-ExtraBold.ttf"
	fontSize = 150
)

// generateImage generates an image with given text, font color, and background color
func generateImage(text, fontColor, bgColor, model, outputDir string) error {
	dc := gg.NewContext(width, height)

	// Set background color
	bgR, bgG, bgB, bgA := hexToRGBA(bgColor)
	dc.SetRGBA(bgR, bgG, bgB, bgA)
	dc.Clear()

	// Load font
	if err := dc.LoadFontFace(fontPath, fontSize); err != nil {
		return fmt.Errorf("could not load font: %v", err)
	}

	// Calculate position to center the text
	dc.SetHexColor(fontColor)
	dc.DrawStringAnchored(text, width/2, height/2, 0.5, 0.5)

	// Save the image
	outputPath := fmt.Sprintf("%s/%s_%s.png", outputDir, text, model)
	return dc.SavePNG(outputPath)
}

// hexToRGBA converts hex color to RGBA components
func hexToRGBA(hex string) (float64, float64, float64, float64) {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return float64(r) / 255, float64(g) / 255, float64(b) / 255, 1.0
}

func main() {
	upperAlphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	models := []struct {
		fontColor string
		bgColor   string
		modelName string
	}{
		{"#B14EC6", "#F9E3FF", "model_1"},
		{"#504EC6", "#E4E3FF", "model_2"},
	}

	outputDir := "./output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("could not create output directory: %v", err)
	}

	for _, model := range models {
		for _, char1 := range upperAlphanumeric {
			// Generate for single character
			text := string(char1)
			if err := generateImage(text, model.fontColor, model.bgColor, model.modelName, outputDir); err != nil {
				log.Printf("could not generate image for %s: %v", text, err)
			}
			for _, char2 := range upperAlphanumeric {
				// Generate for two characters
				text = string(char1) + string(char2)
				if err := generateImage(text, model.fontColor, model.bgColor, model.modelName, outputDir); err != nil {
					log.Printf("could not generate image for %s: %v", text, err)
				}
			}
		}
	}

	log.Println("Profile picture generation completed.")
}
