package steg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lukechampine.com/jsteg"
)

// Decode a message in an image and print it to console
func Decode(filename string) error {
	// open a jpeg
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	hidden, revealErr := jsteg.Reveal(f)
	if revealErr != nil {
		return revealErr
	}
	str := string(hidden)
	firstIndex := strings.Index(str, ":")
	// if current headers aren't available, just return all hidden data
	if firstIndex == -1 {
		fmt.Println(str)
		return nil
	}
	secondIndex := strings.Index(str, "/")
	dataLength, err := strconv.ParseUint(str[0:firstIndex], 10, 64)
	if err != nil {
		return err
	}
	// if message or filetype wasn't hidden, return only a message
	if secondIndex == -1 {
		data := str[firstIndex+1 : uint64(firstIndex)+dataLength+1]
		fmt.Println(data)
		return nil
	}
	messageType := str[firstIndex+1 : secondIndex]
	// if the messageType is a text message, print it to the console
	if messageType == "m" {
		data := str[secondIndex+1 : uint64(secondIndex)+1+dataLength]
		fmt.Println(data)
		return nil
	}
	newf, err := os.Create("./output/" + messageType)
	if err != nil {
		return err
	}
	fmt.Println("Decoded file saved in output/" + messageType)
	data := []byte(str[secondIndex+1 : uint64(secondIndex)+1+dataLength])
	_, err = newf.Write(data)
	if err != nil {
		return err
	}
	return nil
}
