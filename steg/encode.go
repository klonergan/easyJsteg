package steg

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"

	"lukechampine.com/jsteg"
)

// Encode a message in an image
func Encode(inputFilename, outputFilename, messageString string) error {
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
	messageStringLen := len(messageString)
	// the hidden data is formated length:type/data. the type is a file name or an m for message
	data := []byte(fmt.Sprint(messageStringLen) + ":" + "m/" + messageString)
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

// EncodeFromFile modifies the input file and returns error
func EncodeFromFile(imgFile io.Reader, messageString string, messageFileBytes []byte, messageFileName string) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	img, err := jpeg.Decode(imgFile)
	if err != nil {
		return out, err
	}
	messageStringLen := len(messageString)
	var data []byte
	if messageStringLen > 0 && len(messageFileBytes) == 0 {
		data = []byte(fmt.Sprint(messageStringLen) + ":" + "m/" + messageString)
	}
	if len(messageFileBytes) > 0 && messageStringLen == 0 {
		data = []byte(fmt.Sprint(len(messageFileBytes)) + ":" + messageFileName + "/" + string(messageFileBytes))
	}
	if len(messageFileBytes) > 0 && messageStringLen > 0 {
		data = []byte(fmt.Sprint(len(messageFileBytes)) + ":" + messageFileName + "/" + string(messageFileBytes) + fmt.Sprint(messageStringLen) + ":" + "m/" + messageString)
	}
	err = jsteg.Hide(out, img, data, nil)
	if err != nil {
		return out, err
	}
	return out, nil
}
