package main

import (
	"os"

	"./steg"

	"flag"
	"fmt"
)

func main() {
	// define flags
	var d, e bool
	var i, o, m, f string
	flag.BoolVar(&d, "d", false, "decode mode")
	flag.BoolVar(&e, "e", false, "encode mode")
	flag.StringVar(&i, "i", "", "define input jpg filename for encode or decode")
	flag.StringVar(&o, "o", "output.jpg", "define output filename for encode")
	flag.StringVar(&m, "m", "", "use to include a message as a string")
	flag.StringVar(&f, "f", "", "declare a filename of a file to be hidden")
	flag.Parse()
	if (d == true && e == true) || (d == false && e == false) {
		fmt.Println("use -d to decode OR -e to encode a message")
		return
	}
	// encode mode
	if e == true {
		if f == "" {
			err := steg.Encode(i, o, m)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		} else {
			err := steg.EncodeFile(i, o, f)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	}
	// decode mode
	if d == true {
		if i == "" {
			fmt.Println("use -i to define the file to be read")
			return
		}
		data, err := steg.Decode(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(data)
	}
}
