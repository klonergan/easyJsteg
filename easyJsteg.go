package main

import (
	"os"

	"./steg"
)

func main() {
	steg.Encode(os.Args[1], os.Args[2], os.Args[3])
}
