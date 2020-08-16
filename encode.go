package main

import (
	"fmt"
	"image/jpeg"
	"os"

	"lukechampine.com/jsteg"
)

func main() {
	// open a jpeg
	inputFilename := os.Args[1]
	f, _ := os.Open(inputFilename)
	img, _ := jpeg.Decode(f)

	//add hidden data to jpeg
	outputFilename := os.Args[2]
	out, _ := os.Create(outputFilename)
	encodedString := os.Args[3]
	encodedStringLen := len(encodedString)
	data := []byte(fmt.Sprint(encodedStringLen) + ":" + encodedString)
	err := jsteg.Hide(out, img, data, nil)
	if err != nil {
		fmt.Println("hide: ", err)
	}
}
