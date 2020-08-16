package steg

import (
	"os"
	"strconv"
	"strings"

	"lukechampine.com/jsteg"
)

// Decode a message in an image and print it to console
func Decode(filename string) (string, error) {
	// open a jpeg
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	hidden, revealErr := jsteg.Reveal(f)
	if revealErr != nil {
		return "", revealErr
	}
	str := string(hidden)
	firstIndex := strings.Index(str, ":")
	dataLength, err := strconv.ParseUint(str[0:firstIndex], 10, 64)
	if err != nil {
		return "", err
	}
	data := str[firstIndex+1 : uint64(firstIndex)+dataLength+1]
	return data, nil
}
