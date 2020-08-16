package steg

import (
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"os"

	"lukechampine.com/jsteg"
)

// Encode a message in an image
func Encode(inputFilename, outputFilename, encodedString string) error {
	// open a jpeg
	f, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}

	//add hidden data to jpeg
	out, err := os.Create(outputFilename)
	if err != nil {
		return err
	}
	encodedStringLen := len(encodedString)
	// the hidden data is formated length:type/data. the type is a file name or an m for message
	data := []byte(fmt.Sprint(encodedStringLen) + ":" + "m/" + encodedString)
	hideErr := jsteg.Hide(out, img, data, nil)
	if hideErr != nil {
		return hideErr
	}
	return nil
}

// EncodeFile encodes a file in an image
func EncodeFile(inputFilename, outputFilename, messageFilename string) error {
	// open a jpeg
	f, err := os.Open(inputFilename)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(f)
	if err != nil {
		return err
	}
	// open the file to be hidden
	h, err := ioutil.ReadFile(messageFilename)
	if err != nil {
		return err
	}
	encodedData := string(h)
	encodedDataLen := len(encodedData)
	// add hidden data to jpeg
	out, err := os.Create(outputFilename)
	if err != nil {
		return err
	}

	data := []byte(fmt.Sprint(encodedDataLen) + ":" + messageFilename + "/" + encodedData)
	hideErr := jsteg.Hide(out, img, data, nil)
	if hideErr != nil {
		return hideErr
	}
	return nil
}
