package main

import (
	"fmt"
	"os"

	"./steg"
)

func main() {
	// err := steg.Encode(os.Args[1], os.Args[2], os.Args[3])
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	data, err := steg.Decode(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
