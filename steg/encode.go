package steg

import (
	"fmt"
	"image/jpeg"
	"os"

	"lukechampine.com/jsteg"
)

// Encode a message in an image
func Encode(inputFilename, outputFilename, encodedString string) error {
	// open a jpeg
	// inputFilename := os.Args[1]
	f, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	//add hidden data to jpeg
	// outputFilename := os.Args[2]
	out, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	// encodedString := os.Args[3]
	encodedStringLen := len(encodedString)
	data := []byte(fmt.Sprint(encodedStringLen) + ":" + encodedString)
	hideErr := jsteg.Hide(out, img, data, nil)
	if hideErr != nil {
		return hideErr
	}
	return nil
}
