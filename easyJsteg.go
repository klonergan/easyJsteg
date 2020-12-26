package main

import (
	"os"

	"easyJsteg/steg"

	"flag"
	"fmt"
)

func main() {
	// define flags
	var d, e, o, m, f string
	flag.StringVar(&d, "d", "", "decode mode. supply a jpg filename to reveal information in.")
	flag.StringVar(&e, "e", "", "encode mode. supply a jpg filename to hide information in")
	flag.StringVar(&o, "o", "output.jpg", "define output filename for encode")
	flag.StringVar(&m, "m", "", "use to include a message as a string")
	flag.StringVar(&f, "f", "", "declare a filename of a file to be hidden")
	flag.Parse()
	if (d == "" && e == "") || (d != "" && e != "") {
		fmt.Println("use -d to decode OR -e to encode a message")
		return
	}
	// encode mode
	if e != "" {
		if f == "" {
			err := steg.Encode(e, o, m)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		} else {
			err := steg.EncodeFile(e, o, f)
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}
	}
	// decode mode
	if d != "" {
		err := steg.Decode(d, "./output")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
