package main

import (
	"fmt"
	"os"

	"./steg"
)

func main() {
	err := steg.Encode(os.Args[1], os.Args[2], os.Args[3])
	if err != nil {
		fmt.Println(err)
		return
	}
}
