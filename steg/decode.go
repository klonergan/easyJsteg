package steg

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"lukechampine.com/jsteg"
)

// Decode a message in an image
func Decode() {
	// open a jpeg
	f, _ := os.Open(os.Args[1])

	hidden, err := jsteg.Reveal(f)
	if err != nil {
		fmt.Println("hidden: ", err)
	}
	str := string(hidden)
	firstIndex := strings.Index(str, ":")
	dataLength, err := strconv.ParseUint(str[0:firstIndex], 10, 64)
	if err != nil {
		fmt.Println("Error parsing data length")
		return
	}
	fmt.Println(str[firstIndex+1 : uint64(firstIndex)+dataLength+1])
}
