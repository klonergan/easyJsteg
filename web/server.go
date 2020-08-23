package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func testRoute(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var buf bytes.Buffer
		err := req.ParseMultipartForm(1 << 25)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		fmt.Println(req.Header.Get("Content-Type"))
		file, header, err := req.FormFile("img")
		if err != nil {
			w.WriteHeader(400)
			return
		}
		defer file.Close()
		name := header.Filename
		fmt.Println(name)
		io.Copy(&buf, file)
		fmt.Println("buf: ", buf.String())
		m := req.FormValue("m")
		fmt.Println(m)
		// body, err := ioutil.ReadAll(req.Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		// fmt.Println(string(body))
	} else {
		fmt.Fprintf(w, "test route\n")
		fmt.Println(*req)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	http.HandleFunc("/test", testRoute)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
