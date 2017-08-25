package main

import (
	"bytes"
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	f, err := os.Open("turtle.jpeg")
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal("Cannot decode image:", err)
	}

	// the original turtle image is 2256 x 1504
	// crop to just the upper left corner

	cImg, err := cutter.Crop(img, cutter.Config{
		Height:  750,               // height in pixel or Y ratio(see Ratio Option below)
		Width:   1100,              // width in pixel or X ratio
		Mode:    cutter.TopLeft,    // Accepted Mode: TopLeft, Centered
		Anchor:  image.Point{0, 0}, // Position of the top left point
		Options: 0,                 // Accepted Option: Ratio
	})

	if err != nil {
		log.Fatal("Cannot crop image:", err)
	}
	fmt.Println("cImg dimension:", cImg.Bounds())
	// Output: cImg dimension: (10,10)-(510,510)

	// convert from image.Image to []byte
	buf := &bytes.Buffer{}
	if err := jpeg.Encode(buf, cImg, nil); err != nil {
		log.Fatalf("Error converting: %s\n", err)
	}

	// write []byte to file
	cropped_name := "cropped_file.jpg"
	out_file, out_err := os.Create(cropped_name)
	if out_err != nil {
		log.Fatalf("Error creating output file: %s\n", out_err)
	}
	if _, err := out_file.Write(buf.Bytes()); err != nil {
		log.Fatalf("Error writing to output file: %s\n", err)
	}
	fmt.Printf("Wrote %s\n", cropped_name)
}
