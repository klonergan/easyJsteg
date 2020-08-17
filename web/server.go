package main

import (
	"fmt"
	"log"
	"net/http"
)

func testRoute(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "test route\n")
}

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/test", testRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
